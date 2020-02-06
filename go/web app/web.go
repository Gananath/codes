/*
 * Author: Gananath R (2020)
 * 
 * An example web app project in go with sqlite3 database
 * 
 * 
 */
package main

import (
    "net/http"
    "html/template"
    "database/sql"
    "fmt"
    _"github.com/mattn/go-sqlite3"
    //"strconv"
)

// initializing db, err, testvalues as global variables
var db *sql.DB
var err error
var tpl *template.Template


type ShowdbValues struct{
    Id int
    Firstname,Lastname,Date string
}

func init(){
    tpl = template.Must(template.ParseGlob("templates/*.html"))

}

func connect_db(filename string)(*sql.DB, error){
    db, err = sql.Open("sqlite3", filename)
    if err != nil {
        fmt.Println("Error")
    }else{
        fmt.Println("Connection Established")
    }
    err =db.Ping()
    if err!=nil{
        fmt.Println("Ping failed")
    }
    return db,err
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
    //fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path)

    tpl.ExecuteTemplate(w,"index.html" ,nil)
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
    //var about_page = template.Must(template.ParseFiles("template/about.html"))
    var id int
    var firstname, lastname, date string
    if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
	}
    rows, _ := db.Query("SELECT * FROM test")
    re := ShowdbValues{}
    var results []ShowdbValues
    for rows.Next(){
        rows.Scan(&id,&firstname,&lastname,&date)
        re.Id = id
        re.Firstname = firstname
        re.Lastname = lastname
        re.Date = date
        //fmt.Println(strconv.Itoa(id) + ": " + firstname + " " + lastname)
        results = append(results, re)
    }
    tpl.ExecuteTemplate(w,"about.html",results)
}
func FormHandler(w http.ResponseWriter, r *http.Request){
    if r.Method != "POST"{
        http.Redirect(w,r,"/about",http.StatusSeeOther)
    }
    fname := r.FormValue("firstname")
    lname := r.FormValue("lastname")
    if fname =="" || lname ==""{
        http.Redirect(w,r,"/about",http.StatusSeeOther)
    }else{
        stmt, err := db.Prepare("INSERT INTO test (firstname, lastname) VALUES (?, ?)")
        checkErr(err)
        stmt.Exec(fname, lname)
        http.Redirect(w,r,"/about",http.StatusSeeOther)
    }
    
}

func main() {
    db,err = connect_db("./database/example.db")
    defer db.Close()
    checkErr(err)
    http.HandleFunc("/", IndexHandler)
    http.HandleFunc("/about", AboutHandler)
    http.HandleFunc("/action", FormHandler)
    http.ListenAndServe(":8080", nil)
}


