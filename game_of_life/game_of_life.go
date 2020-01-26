/*
 * Author: Gananath R
 * A dynamic size text based Conway's Game of Life in golang
 * 
 * 
 */

package main

import (
    "fmt"
    "math/rand"
    "time"
)
var m,n, r_num int// = 6,6,1
 

const (
        DebugColor   = "\033[1;31m%s\033[0m"
)


// initializing
// /*
func init(){    
    fmt.Println("Enter the row value of the matrix: ")
    fmt.Scan(&m)
    fmt.Println("Enter the column value of the matrix: ")
    fmt.Scan(&n)
    fmt.Println("Enter number of initial lifes: ")
    fmt.Scan(&r_num)
    
}
//*/


func main(){
    //fmt.Println("m value=",m," random lifes=",r);
    var cell, count int
    episodes := 3
    state := random_state(m,n)
    gen := 0
    for e :=0; e<episodes; {
        fmt.Println("\033[H\033[2J")  // clear screen for nix systems
        visualize(state)
        fmt.Println(gen)
        gen++
        for i :=0; i< m; i++{
            for j := 0; j < n; j++{
                if state[i][j]== 1 {
                    cell = 1
                }else{
                    cell = 0
                }                
                count = neighbours_count(state,i,j)
                
                if (cell==1) && (count < 2){
                // Rule 1: Any live cell with fewer than two live neighbours dies (referred to as underpopulation or exposur
                //fmt.Print(1)
                    state[i][j] = 0
                }else if ((cell==1) && (count > 3)){
                // Rule 2: Any live cell with more than three live neighbours dies (referred to as overpopulation or overcrowding).
                //fmt.Print(2)
                    state[i][j] = 0
                }else if ((cell==1) && ((count==2) || (count==3))){
                // Rule 3: Any live cell with two or three live neighbours lives, unchanged, to the next generation.
                //fmt.Print(3)
                    continue
                }else if ((cell==0)&&(count==3)){
                //Rule 4: Any dead cell with exactly three live neighbours will come to life.
                //fmt.Print(4)
                    state[i][j] = 1
                }
                //fmt.Print(count)
        } // j for loop
        //fmt.Println()
        
        } // i for loop 
    //break
    e = 1 //+ e// infinite loop
    time.Sleep(100 * time.Millisecond) 
    } // e for loop
}

func random_state(m int,n int) [][]int {
    // Creating a 2d zero matrix
    matrix := make([][]int, m) // row
    for i := range matrix {
        matrix[i] = make([]int, n) // column
    }
    // Random generating seeds for GoL
    for i :=0; i < r_num;{ 
        row := random(2,m-1)
        col := random(2,n-1)

        matrix[row][col] = 1
        matrix[row-1][col] =1
        matrix[row-2][col] =1
        matrix[row-1][col+1] =1
        matrix[row][col-1] =1
        matrix[row-1][col-2] =1
        i++
    }
    return matrix
}

// Visualization of GoL states
func visualize(state [][]int){
    for i:=0; i< m; i++ {
        for j:=0; j < n; j++ {
            if state[i][j] == 0{
                fmt.Print(".")
            }else{
                fmt.Print("*")
            }
        }
    fmt.Println()
    }
}

// Counting surrounding living neighbours 
func neighbours_count(matrix [][]int,i int, j int) int {
    
    count := 0
    if (i>0) { 
        if (matrix[i-1][j] == 1) {
        // Top
        count ++
      }
      }
    if (i< m-1) { 
    if matrix[i+1][j] == 1 {
        // Bottom
        count ++
    }
    }
    if (j>0) {
    if matrix[i][j-1] == 1 {
        // Left
        count ++
    }
    }
    if (j< n-1) {
    if matrix[i][j+1] == 1 {
        // Right
        count ++
    }
    }
    if (i>0) && (j>0){
    if matrix[i-1][j-1] == 1 {
        // Top Left
        count ++
    }
    }
    if (i>0) && (j< n-1) {
    if matrix[i-1][j+1] == 1 {
        // Top Right
        count ++
    }
    }
    if  (i< m-1) && (j>0){
    if matrix[i+1][j-1] == 1 {
        // Bottom Left
        count ++
    }
    }
    if  (i<m-1) && (j<n-1){
    if matrix[i+1][j+1] == 1 {
        // Bottom right
        count ++
    }
    }
    return count
}


// Custom random function for generating random number inside a range
func random(min int, max int) int {
    return rand.Intn(max-min) + min
}
