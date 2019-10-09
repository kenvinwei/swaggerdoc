package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

type Person struct {
	Name string
	age string
}

func main()  {
	http.HandleFunc("/", func(w http.ResponseWriter,r *http.Request) {
		dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
		t, err := template.ParseFiles(dir+"/index.html")
		if err != nil {
			fmt.Println("parse file err:", err)
			return
		}
		p := Person{Name: "Mary", age: "31"}
		if err := t.Execute(w, p); err != nil {
			fmt.Println("There was an error:", err.Error())
		}
	})
	http.ListenAndServe(":9999", nil)
}
