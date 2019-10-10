package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"swaggerdoc/lib"
	"text/template"
	"time"
)

type global struct {
	Url    string
	Prefix string
	Token  string
	R      int64
}

func joinComment(source, newLine string) string {
	if source != "" {
		return source + "\n" + newLine
	} else {
		return newLine
	}
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		servDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			log.Fatalf("err filepath: %s\n", servDir+"/index.html")
		}
		t, _ := template.ParseFiles(servDir + "/index.html")
		t.Execute(w, global{Url: "crm_base_service.test", Prefix: "/api/v1", Token: "chenlei6", R: rand.New(rand.NewSource(time.Now().UnixNano())).Int63()})

	})

	http.HandleFunc("/make", func(w http.ResponseWriter, r *http.Request) {
		var commentString string

		r.ParseForm()
		d := r.Form["jsonData"][0]

		lib.FindRequest(d, "")
		comment := lib.MakeComment(lib.AllRequest[0])
		lib.AllRequest = lib.AllRequest[1:]

		commentString = joinComment(commentString, "/**")
		for _, c := range comment {
			commentString = joinComment(commentString, " *"+c)
		}

		commentString = joinComment(commentString, " */")
		commentString = joinComment(commentString, "\n\n")

		w.Write([]byte(commentString))

	})

	http.ListenAndServe(":9999", nil)
}
