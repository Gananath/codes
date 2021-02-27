use super::state::*;
use rand::seq::SliceRandom;
use rand::thread_rng;
use std::error::Error;
use std::fmt;
use std::io::stdin;

#[derive(Debug)]
pub struct Board<'a> {
    pub board: Vec<&'a str>,
    pub player: &'a str,
    pub ai: &'a str,
    pub size: usize,
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
