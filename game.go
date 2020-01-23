/*
 * Author: Gananath R
 * A go lang library for creating n dimensional tic-tac-toe game
 * 
 * 
 * 
 * 
 */

package main
import (
    "fmt"
    "./libtac"
    "strings"

)



var size int = 2

var player, ai string

func init(){    
    fmt.Println("Enter the size of tic tac toe: ")
    fmt.Scan(&size)
    fmt.Println("Do you want to be 'x' or 'o': ")
    fmt.Scan(&player)

    if (strings.ToUpper(player) == "X"){
        player = "x"
        ai = "o"
    }else if (strings.ToUpper(player) == "O"){
        player = "o"
        ai = "x"
    }else{
        fmt.Println("Something went wrong..resetting")
        player = " "
        ai = " "
    }

}

func main(){
    var p_input int
    // creating vacant board
    board := libtac.CreateBoard(size) 
    fmt.Println(size)
    libtac.PrintBoard(board,size)
    // user playing
    fmt.Println("Board Map")
    r := 0
    for i:=0;i<size;i++{
        fmt.Printf("\n\t\t")
        for j:=0;j<size;j++{
            fmt.Print(r,"|")
            r++
        }
    }
    fmt.Printf("\n")

    for true {
        fmt.Println(" You:",strings.ToUpper(player),"  Computer:",strings.ToUpper(ai))
        fmt.Println(" ")
        fmt.Println(" Please enter the position number between 1 and ",size*size)
        fmt.Scan(&p_input)
        if (p_input<1)||(p_input>size*size){
            continue
        }
        board[p_input-1] = player
        libtac.PrintBoard(board,size)
        if libtac.Playing(board,player,size)==true{
            break
        }
        move := libtac.SimpleAI(board,size)
        board[move] = ai
        libtac.PrintBoard(board,size)
        if libtac.Playing(board,ai,size)==true{
            break
        }
    }


}

