package main

import (
    "fmt"
    "net/http"
    "log"
    "html/template"
)

type Film struct {
    Title string
    Director string
}

func main() {
    fmt.Println("abc")

    home := func(w http.ResponseWriter, r *http.Request){
        films := map[string][]Film{
            "Films": {
                {Title: "Film 01",Director: "Joe"},
                {Title: "Film 02",Director: "Yogesh"},
                {Title: "Film 03",Director: "James"},
            },
        }
        templ := template.Must(template.ParseFiles("index.html"))
        templ.Execute(w, films)
    }
    http.HandleFunc("/", home)
    
    log.Fatal(http.ListenAndServe(":8000", nil))
}