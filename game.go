/*
 * Author: Gananath R
 * A go lang library for creating n dimensional tic-tac-toe games
 * 
 */

package main
import (
    "fmt"
    "strings"
    "./libtac"
)

var size int 
var player, ai string

func init(){    
    fmt.Print("\n\tEnter the size of tic tac toe: ")
    fmt.Scan(&size)
    fmt.Print("\n\tDo you want to be 'x' or 'o': ")
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
    var u_input string
    // creating vacant board
    board := libtac.CreateBoard(size) 
    fmt.Println(size)
    libtac.PrintBoard(board,size)

    fmt.Println("\t[Board Map]")

    r := 0
    for i:=0;i<size;i++{
        fmt.Printf("\n\t\t")
        for j:=0;j<size;j++{
            fmt.Print(r,"|")
            r++
        }
    }

    for true {
        fmt.Printf("\n")
        fmt.Println("\n\t[You]:",strings.ToUpper(player),"  [Computer]:",strings.ToUpper(ai),"\n")
        fmt.Print(" Please enter the position number between 0 and ",size*size-1,": ")
        fmt.Scan(&p_input)
        // User moves
        possibleMoves := libtac.PossibleMoves(board)
        if (p_input<0)||(p_input>size*size-1)||!libtac.Contains(possibleMoves,p_input){
            libtac.PrintBoard(board,size)
            continue
        }
        board[p_input] = player
        libtac.PrintBoard(board,size)
        if libtac.Playing(board,player,size)==true{
            fmt.Print("\n\tDo you want to continue(Y/N): ")
            fmt.Scan(&u_input)
            if strings.ToUpper(u_input)=="Y"{
                main()
            }else{
                break
            }
            
        }
        // Computer moves
        move := libtac.SimpleAI(board,size)
        board[move] = ai
        libtac.PrintBoard(board,size)
        if libtac.Playing(board,ai,size)==true{
            fmt.Print("\n\tDo you want to continue(Y/N): ")
            fmt.Scan(&u_input)
            if strings.ToUpper(u_input)=="Y"{
                main()
            }else{
                break
            }
        }
    }
}



