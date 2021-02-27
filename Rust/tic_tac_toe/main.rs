// Gananath R
// Rust N-Sized tic-tac-toe game
// 2021

//#![allow(dead_code)]
//#![allow(warnings, unused)]

mod game;

use game::utils::*;
use std::error::Error;
use std::io::stdin;

fn main() -> Result<(), Box<dyn Error>> {
    println!("\x1B[2J\x1B[1;1H");
    let size = input("size");
    let players = input("player");
    let size: usize = size.trim().parse().expect("Failed to read user input");
    let players = players.trim();
    let players = get_player(players)?; //
    let mut game = Board::new(players, size);
    game.view();
    loop {
        let mut player_moves = String::new();
        println!(
            "\n\n\t[You]: {}  [Computer]: {} \n",
            game.player.to_uppercase(),
            game.ai.to_uppercase()
        );
        eprint!(
            "Please enter the position number between 0 and {} : ",
            size * size - 1
        );
        let _ = stdin()
            .read_line(&mut player_moves)
            .expect("\nFailed to read player_moves");
        let player_moves: usize = player_moves
            .trim_end()
            .parse()
            .expect("Failed to read user input");

        let moves_available = game.available(&player_moves);

        if (player_moves < 0_usize) || (player_moves > size * size - 1 || !moves_available) {
            game.view();
            continue;
        }
        game.player(&player_moves);

        if game.end() {
            let mut usr_input = String::new();
            println!("\n\t\t\tDo you want to continue(Y/N): ");
            let _ = stdin()
                .read_line(&mut usr_input)
                .expect("Failed to read user input");
            let usr_input = usr_input.trim_end().to_uppercase();
            let usr_input = usr_input.as_str();
            if usr_input == "Y" {
                main();
            } else {
                break;
            }
        }

        game.ai();

        if game.end() {
            let mut usr_input = String::new();
            println!("\n\t\t\tDo you want to continue(Y/N): ");
            let _ = stdin()
                .read_line(&mut usr_input)
                .expect("Failed to read user input");
            let usr_input = usr_input.trim_end().to_uppercase();
            let usr_input = usr_input.as_str();
            if usr_input == "Y" {
                main();
            } else {
                break;
            }
        }
    }
    Ok(())
}
