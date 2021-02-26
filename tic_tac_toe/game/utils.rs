use rand::seq::SliceRandom;
use rand::thread_rng;
use std::error::Error;
use std::fmt;
use std::io::stdin;

#[derive(Debug)]
pub struct Board<'a> {
    board: Vec<&'a str>,
    pub player: &'a str,
    pub ai: &'a str,
    size: usize,
}

// game core logic
pub fn input(val: &str) -> String {
    let mut inp = String::new();
    match val {
        "size" => eprint!("\nPlease enter the size: "),
        "player" => eprint!("\nPlease enter the player (X/O): "),
        _ => eprint!("Wrong size or player"),
    }

    let _ = stdin()
        .read_line(&mut inp)
        .expect("\nFailed to read player_moves");
    inp
}

pub fn get_player(val: &str) -> Result<(&str, &str), Box<dyn Error>> {
    let mut player = val;
    let mut ai = "";
    match player.to_uppercase().as_str() {
        "X" => {
            player = "x";
            ai = "o"
        }
        "O" => {
            player = "o";
            ai = "x";
        }
        _ => return Err("Not allowed".into()),
    }
    Ok((player, ai))
}

impl<'a> Board<'a> {
    pub fn new((player, ai): (&'a str, &'a str), size: usize) -> Self {
        Self {
            board: vec!["-"; size * size],
            player: player,
            ai: ai,
            size: size,
        }
    }

    pub fn view(&mut self) {
        print!("\x1B[2J\x1B[1;1H"); //clear screens
        let mut r = 0;
        for _ in 0..self.size {
            print!("\n\t\t\t\t |");
            for _ in 0..self.size {
                print!(" {} |", self.board[r]);

                r += 1;
            }
        }
    }

    pub fn possible_moves(&self) -> Vec<usize> {
        let mut vec: Vec<usize> = Vec::new();
        for (i, j) in self.board.iter().enumerate() {
            if j == &"-" {
                vec.push(i);
            } else {
                continue;
            }
        }
        vec
    }

    fn get_edges(&self) -> Vec<usize> {
        let mut vec: Vec<usize> = Vec::new();
        for i in 0..self.size {
            vec.push(i); //top
            vec.push(i * self.size); //left
            vec.push(self.size - 1 + self.size * i); //right
            vec.push((self.size * self.size - i) - 1); //bottom
        }
        vec
    }

    fn simple_ai(&self) -> usize {
        let mut corners: Vec<usize> = vec![0];
        let possible_moves = self.possible_moves();
        for i in &["o", "x"] {
            for j in &possible_moves {
                let mut copy_board = self.board.clone();
                copy_board[*j] = i;
                //print_board(&copy_board,&size);
                if self.has_won() {
                    return *j;
                }
            }
        }

        // finding corners
        let top_right: usize = self.size - 1;
        corners.push(top_right);
        let bottom_left: usize = self.size * top_right;
        corners.push(bottom_left);
        let bottom_right: usize = bottom_left + top_right;
        corners.push(bottom_right);
        corners.shuffle(&mut thread_rng()); //shuffle the vec
        for i in corners {
            if self.contains(&possible_moves, &i) {
                return i;
            }
        }
        // Play for the edges
        let mut edges = self.get_edges();
        edges.shuffle(&mut thread_rng());

        for i in edges {
            if self.contains(&possible_moves, &i) {
                return i;
            }
        }
        //println!("{:?}",corners);
        // Else take random positions
        *possible_moves.choose(&mut thread_rng()).unwrap()
    }

    pub fn player(&mut self, player: &usize) {
        self.board[*player] = self.player;
        self.view();
    }

    pub fn ai(&mut self) {
        // Computer play
        let c_move = self.simple_ai();
        self.board[c_move] = self.ai;
        self.view();
    }
}

///// game state
impl<'a> Board<'a> {
    pub fn available(&self, moves: &usize) -> bool {
        let possible_moves = self.possible_moves();
        let moves_available = self.contains(&possible_moves, moves);
        moves_available
    }

    fn contains<T>(&self, possible_moves: &Vec<T>, player_moves: &T) -> bool
    where
        T: std::fmt::Debug + PartialEq,
    {
        if possible_moves.contains(player_moves) {
            return true;
        }

        false
    }

    fn has_won(&self) -> bool {
        // Check if anyone won the game
        // row, column and diagonal validation
        let mut r: usize = 0;
        let mut c: usize = 0;
        //diagonal start and end element
        let mut d1: usize = 0;
        let mut d2: usize = self.size - 1;

        for i in 0..self.size {
            // Row variables
            let mut iter_r = 0;
            let prev_r = self.board[r];
            // Column variables
            let mut iter_c = 0;
            let prev_c = self.board[i];

            // Diagonal variables
            let prev_d1 = self.board[d1];
            let prev_d2 = self.board[d2];
            let mut iter_d1 = 0;
            let mut iter_d2 = 0;
            for j in 0..self.size {
                // Row same element validation
                if (prev_r == self.board[r]) && (self.board[r] == self.player) {
                    iter_r += 1;
                }
                r += 1;
                //Column same element validation
                c = i + j * self.size;
                if (prev_c == self.board[c]) && (self.board[c] == self.player) {
                    iter_c += 1;
                }
                // Diagonal same element validation
                if i == d1 {
                    d1 = (d2 + 2) * j;
                    if (prev_d1 == self.board[d1]) && (self.board[d1] == self.player) {
                        iter_d1 += 1
                    }
                }
                if i == d2 {
                    d2 = d2 * (j + 1);
                    if (prev_d2 == self.board[d2]) && (self.board[d2] == self.player) {
                        iter_d2 += 1;
                    }
                }
            }
            if (iter_r == self.size)
                || (iter_c == self.size)
                || (iter_d1 == self.size)
                || (iter_d2 == self.size)
            {
                return true;
            }
        }
        return false;
    }

    fn is_filled(&self) -> bool {
        for i in &self.board {
            if i == &"-" {
                return false;
            }
        }
        true
    }

    fn game_status(&self) -> u8 {
        if self.has_won() {
            //won the game
            return 1;
        } else if self.is_filled() {
            //game draw
            return 2;
        } else {
            //still playing
            return 0;
        }
    }

    pub fn end(&self) -> bool {
        let status = self.game_status();
        match status {
            1 => {
                println!("\n\t\t\t {} has won the game", self.player);
                return true;
            }
            2 => {
                println!("\n\t\t\tGame is a draw");
                return true;
            }
            _ => return false,
        }
    }

    fn _check(&self) -> bool {
        // Orphan function

        return true;
    }
}
