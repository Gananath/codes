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
var m,n, r_num int //= 10,10,1
 
func init(){
    fmt.Print("Enter the row value: ")
    fmt.Scan(&m)
    fmt.Print("Enter the column value: ")
    fmt.Scan(&n)
    fmt.Print("Enter number of initial lifes: ")
    fmt.Scan(&r_num)
    
}


func main(){
    //fmt.Println("m value=",m," random lifes=",r);
    var cell, count int    
    state := RandomState(m,n)    
    gen := 0
    for true {
        fmt.Println("\033[H\033[2J")  // clear screen for nix systems
        Visualize(state)
        count_map := NeigbourhoodMap(state)
        //BoardMap(state)
        fmt.Println("\t\tGeneration: ",gen)
        gen++
        for i :=0; i< m; i++{
            for j := 0; j < n; j++{
                if state[i][j]== 1 {
                    cell = 1
                }else{
                    cell = 0
                }
                //fmt.Println("here",i,j)                
                count = count_map[i][j]
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
                    state[i][j] = 1
                }else if ((cell==0)&&(count==3)){
                //Rule 4: Any dead cell with exactly three live neighbours will come to life.
                //fmt.Print(4)
                    state[i][j] = 1
                }
                //fmt.Print(count)
        } // j for loop
        //fmt.Println()
        } // i for loop 

    time.Sleep(100 * time.Millisecond) 
    } // e for loop
    
}

func RandomState(m int,n int) [][]int {
    // Creating a 2d zero matrix
    matrix := make([][]int, m) // row
    for i := range matrix {
        matrix[i] = make([]int, n) // column
    }
    // Random generating seeds for GoL
    for i :=0; i < r_num;{ 
        row := Random(2,m-1)
        col := Random(2,n-1)
        
        matrix[row][col] = 1
        matrix[row-1][col] =1
        matrix[row-2][col] =1
        matrix[row-1][col+1] =1
        matrix[row][col-1] =1
        matrix[row-1][col-2] =1
        /*
         *  dash-dash-dash 
         *  dash-dash-dash pattern
        matrix[row][col] = 1
        matrix[row][col-1] = 1
        matrix[row][col+1] = 1
        matrix[row-1][col-1] = 1
        matrix[row-1][col] = 1
        matrix[row-1][col+1] = 1
        */
        i++
    }
    return matrix
}

// Visualization of GoL states
func Visualize(state [][]int){
    for i:=0; i< m; i++ {
        fmt.Print("\n\t\t")
        for j:=0; j < n; j++ {
            if state[i][j] == 0{
                fmt.Print(".")
            }else{
                    fmt.Print("\033[32;1m*\033[0m")
            }
        }
    fmt.Print("\n")
    }
}

// Visualizating position of  states
func BoardMap(state [][]int){
    for i:=0; i< m; i++ {
        for j:=0; j < n; j++ {
            fmt.Print(state[i][j])
        }
    fmt.Println()
    }
}

// Neigbourhood maps
func NeigbourhoodMap(state [][]int)[][]int{
    //Gets the neigbourhood map of the state
    // Creating a 2d zero matrix
    matrix := make([][]int, m) // row
    for i := range matrix {
        matrix[i] = make([]int, n) // column
    }    
    for i:=0; i< m; i++ {
        for j:=0; j < n; j++ {
            count := NeighboursCount(state,i,j)
            matrix[i][j] = count
        }
    }
    return matrix
}

// Counting surrounding living neighbours 
func NeighboursCount(matrix [][]int,i int, j int) int {
    count := 0
    if (i>0) &&(matrix[i-1][j] == 1) {
        // Top
        count ++
    }
    if (i< m-1)&&matrix[i+1][j] == 1 {
        // Bottom
        //BoardMap(matrix)
        count ++
    }
    if (j>0) && matrix[i][j-1] == 1 {
        // Left
        count ++
    }
    if (j< n-1)&&matrix[i][j+1] == 1 {
        // Right
        count ++
    }
    if  (i>0) && (j>0)&&matrix[i-1][j-1] == 1 {
        // Top Left
        count ++
    }
    if (i>0) && (j< n-1)&& matrix[i-1][j+1] == 1 {
        // Top Right
        count ++
    }
    if (i< m-1) && (j>0)&&matrix[i+1][j-1] == 1 {
        // Bottom Left
        count ++
    }
    if (i<m-1) && (j<n-1)&&matrix[i+1][j+1] == 1 {
        // Bottom right
        count ++
    }
    return count
}


// Custom random function for generating random number inside a range
func Random(min int, max int) int {
    return rand.Intn(max-min) + min
}
