use super::utils::*;


///// game state
impl<'a> Board<'a> {
    pub fn available(&self, moves: &usize) -> bool {
        let possible_moves = self.possible_moves();
        let moves_available = self.contains(&possible_moves, moves);
        moves_available
    }

    pub fn contains<T>(&self, possible_moves: &Vec<T>, player_moves: &T) -> bool
    where
        T: std::fmt::Debug + PartialEq,
    {
        if possible_moves.contains(player_moves) {
            return true;
        }

        false
    }

    pub fn has_won(&self) -> bool {
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
