/*
 * Author: Gananath R
 * A go lang library for dynamically creating n dimensional tic-tac-toe game
 * 
 * 
 * 
 * 
 */

package libtac

import (
	"fmt"
    "strings"
    "math/rand"
    "time"
)



func PrintBoard(board []string,size int){
    // Prints the board
    fmt.Println("\033[H\033[2J")  // clear screen for nix systems
    var r int = 0
    for i:=0;i<size;i++{
        fmt.Printf("\n\t\t")
        for j:=0;j<size;j++{
            fmt.Print(strings.ToUpper(board[r]),"|")
            r++
        }
    }
    fmt.Printf("\n\n")

}

func IsFilled(board []string) bool{
    // Check if the board is full or not
    for _, s :=range board{
        if s=="-"{
            return false
        }
    }
    return true
}

func CreateBoard(n int) []string{
    // Creates a vacant board
    board := make([]string,n*n)
    for i,_ := range board{
        board[i] = "-"
    }
    return board
}

func HasWon(board []string,char string,size int) bool{
    // Check if anyone won the game
    // row, column and diagonal validation
    var r,c int = 0,0 // row,column
    var d1,d2 int = 0,size-1 // diagonal start and end element
    for i:=0;i<size;i++{
        // Row variables
        iter_r := 0
        prev_r := board[r]
        // Column variables
        iter_c := 0
        prev_c := board[i]
        
        // Diagonal variables
        prev_d1 := board[d1]
        prev_d2 := board[d2]
        iter_d1 := 0
        iter_d2 := 0
        for j:=0;j<size;j++{
            // Row same element validation
            if (prev_r == board[r])&&(board[r]==char){
                iter_r++
            }
            r++
            // Column same element validation
            c = i+j*size
            if (prev_c == board[c])&&(board[c]==char){
                iter_c++
            }
            // Diagonal same element validation
            if(i==d1){
                d1 := (d2+2)*j
                if (prev_d1 == board[d1])&&(board[d1]==char){
                    iter_d1++
                }
            
            }
            if(i==d2){
                d2 := (d2*(j+1))
                if (prev_d2 == board[d2])&&(board[d2]==char){
                    iter_d2++
                }
            }
            
        }
        if (iter_r==size)||(iter_c==size)||(iter_d1==size)||(iter_d2==size){
            return true
        }
    }
    return false
}

func SimpleAI(board []string,size int)int{
    // returns the winning move or blocks the winning move of the opponent
    copyBoard := make([]string, len(board))
    letters := []string{"o","x"}
    var possibleMoves[] int
    for i,e := range board{
        if e=="-"{
            possibleMoves = append(possibleMoves,i)
        }
    }
    for _,char := range letters{
        for _,e := range possibleMoves{
            copy(copyBoard,board)
            copyBoard[e] = char
            if HasWon(copyBoard,char, size){
                return e
            }
        }
        
    }
    // initialize global pseudo random generator
    rand.Seed(time.Now().Unix())
    // takes corner positions if vacant
    top_right := size-1
    bottom_left := size*top_right
    bottom_right := bottom_left+top_right
    corners := []int{0,top_right,bottom_left,bottom_right}
    rand.Shuffle(len(corners), func(i, j int) { corners[i], corners[j] = corners[j], corners[i] }) // shuffling the arrays order
    for _,c := range corners {
        if Contains(possibleMoves,c){
            return c
        }
    }
    return possibleMoves[rand.Intn(len(possibleMoves))]
}


func GameStatus(board []string,char string,size int) int{
    if HasWon(board,char,size){
        // won the game
        return 1
    }else if IsFilled(board){
        // game is a draw
        return -1
    }else{
        // still playing
        return 0
    }
}

func Contains(s []int, e int) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

func Playing(board []string, char string,size int) bool{
    status := GameStatus(board,char,size)
    if status==1{
        fmt.Println(strings.ToUpper(char),"has won the game")
        return true
    }else if status==-1 {
        fmt.Println("Game is a draw")
        return true
    } 
    return false
}


