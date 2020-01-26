extern crate rand;

use rand::Rng;
use std::thread;
use std::time::Duration;

fn main() {
    
    // parameter row, column for the enviroment and r_num for number of species
    let (row,col,r_num)=(4,4,1);
    let mut state = random_state(row,col,r_num);
    let mut episode = 0;
    loop{
        break;
        visualize(state.clone());
        println!("Generation:{}",episode);
        for i in 0..state.len()-1 {
            for j in 0..state[0].len()-1 {
                let cell =if state[i][j]==1{1}else{0};
                let count = neighbours_count(state.clone(),i,j);
                //println!("cell:{} count:{}",state.len(), state[0].len());
                if cell==1 && count < 2 {
                // Rule 1: Any live cell with fewer than two live neighbours dies (referred to as underpopulation or exposur
                    state[i][j] = 0;
                }else if cell==1 && count > 3 {
                // Rule 2: Any live cell with more than three live neighbours dies (referred to as overpopulation or overcrowding).
                    state[i][j] = 0;
                }else if (cell==1 && count==2) || count==3{
                // Rule 3: Any live cell with two or three live neighbours lives, unchanged, to the next generation.
                //fmt.Print(3)
                    state[i][j] = 1;
                }else if cell==0 && count==3 {
                //Rule 4: Any dead cell with exactly three live neighbours will come to life.
                //fmt.Print(4)
                    state[i][j] = 1;
                }
                
            } // j for loop
        
        }
    thread::sleep(Duration::from_millis(100));
    episode += 1;
    }
}


fn random_state(m:usize,n:usize,r_num:usize)-> Vec<Vec<i32>> {
    let mut matrix = vec![vec![0i32;n];m];
    for _i in 0..r_num{
        let row = rand::thread_rng().gen_range(2, m-1);
        let col = rand::thread_rng().gen_range(2, n-1);
        matrix[row][col] = 1;
        matrix[row-1][col] = 1;
        matrix[row-2][col] = 1;
        matrix[row-1][col+1] = 1;
        matrix[row][col-1] = 1;
        matrix[row-1][col-2] = 1;

    }
    
    matrix
}

fn visualize(matrix:Vec<Vec<i32>>){
    print!("\x1B[2J");
    print!("\n\t\t");
    for i in 0..matrix.len()-1{
        for j in 0..matrix[0].len()-1{
            if matrix[i][j]==1{
                print!("*");
            }else{
                print!(".")
            }
            
        }
        print!("\n\t\t");
    }
}

fn neighbours_count(matrix:Vec<Vec<i32>>,i:usize, j:usize)-> usize {
    
    let mut count =0;
    let m = matrix[0].len();
    let n = matrix.len();
    
    if i>0 { 
        if matrix[i-1][j] == 1 {
        // Top
        count += 1;
      }
      }
    if i< m-1{ 
    if matrix[i+1][j] == 1 {
        // Bottom
        count += 1;
    }
    }
    if j>0 {
    if matrix[i][j-1] == 1 {
        // Left
        count += 1;
    }
    }
    if j< n-1 {
    if matrix[i][j+1] == 1 {
        // Right
        count += 1;
    }
    }
    if i>0 && j>0 {
    if matrix[i-1][j-1] == 1 {
        // Top Left
        count += 1;
    }
    }
    if i>0 && j< n-1 {
    if matrix[i-1][j+1] == 1 {
        // Top Right
        count += 1;
    }
    }
    if  i< m-1 && j>0 {
    if matrix[i+1][j-1] == 1 {
        // Bottom Left
        count += 1;
    }
    }
    if  i<m-1 && j<n-1 {
    if matrix[i+1][j+1] == 1 {
        // Bottom right
        count += 1;
    }
    }
    count
}

