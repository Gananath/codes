/*
 * Author: Gananath R
 * A go lang library for dynamically creating n dimensional tic-tac-toe game
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
            str := strings.ToUpper(board[r])
            if str=="X"{
                fmt.Print("\033[32;1m ",str,"\033[0m ","|")
            }else if str=="O"{
                fmt.Print("\033[31;1m ",str,"\033[0m ","|")
            }else{
                fmt.Print("\033[37;1m ",str,"\033[0m ","|")
            }
            
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
    // Returns the winning move or blocks the winning move of the opponent
    possibleMoves := PossibleMoves(board)
    copyBoard := make([]string, len(board))
    letters := []string{"o","x"}
    for _,char := range letters{
        for _,e := range possibleMoves{
            copy(copyBoard,board)
            copyBoard[e] = char
            PrintBoard(copyBoard,size)
            if HasWon(copyBoard,char, size){
                return e
            }
        }
        
    }
    // Initialize global pseudo random generator
    rand.Seed(time.Now().Unix())
    // Play for corners
    top_right := size-1
    bottom_left := size*top_right
    bottom_right := bottom_left+top_right
    corners := []int{0,top_right,bottom_left,bottom_right}
    RandomizeSlice(corners)
    for _,e := range corners {
        if Contains(possibleMoves,e){
            return e
        }
    }
    
    // Play for the edges
    edges := GetEdges(size)
    RandomizeSlice(edges)
    for _,e := range edges {
        if Contains(possibleMoves,e){
            return e
        }
    }
    // Else take random positions
    return possibleMoves[rand.Intn(len(possibleMoves))]
}


func GameStatus(board []string,char string,size int) int{
    if HasWon(board,char,size){
        // Won the game
        return 1
    }else if IsFilled(board){
        // Game is a draw
        return -1
    }else{
        // Still playing
        return 0
    }
}

func Playing(board []string, char string,size int) bool{
    status := GameStatus(board,char,size)
    if status==1{
        fmt.Println("\t",strings.ToUpper(char),"has won the game")
        return true
    }else if status==-1 {
        fmt.Println("\tGame is a draw")
        return true
    } 
    return false
}


func PossibleMoves(board []string) []int{
    var possibleMoves[] int
    for i,e := range board{
        if e=="-"{
            possibleMoves = append(possibleMoves,i)
        }
    }
    return possibleMoves
}

func GetEdges(size int)[]int{
    var top,left,right,bottom [] int
    for i:=1;i<size-1;i++{
        top = append(top,i)
        left = append(left,i*size)
        right = append(right,size-1+size*i)
        bottom = append(bottom,(size*size-i)-1)
    }
    // Appending all slices to one
    x := append(top,left...)
    x = append(x,right...)
    x = append(x,bottom...)
    return x
}


func Contains(s []int, e int) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

func RandomizeSlice(x []int){
    rand.Shuffle(len(x), func(i, j int) { x[i], x[j] = x[j], x[i] })
}


