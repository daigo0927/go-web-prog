package main

import (
	"net/http"
	"html/template"
	"math/rand"
	"time"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl.html")
	t.Execute(w, "Hello World!")
}

func compare(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("compare.html")
	rand.Seed(time.Now().Unix())
	t.Execute(w, rand.Intn(10) > 5)
}

func iterate(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("iterate.html")
	daysOfWeek := []string{"mon", "tue", "wed", "thu", "fri", "sat", "sun"}
	t.Execute(w, daysOfWeek)
}

func assign(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("assign.html")
	t.Execute(w, "hello")
}

func include(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("t1.html", "t2.html")
	t.Execute(w, "Hello World")
}

func formatDate(t time.Time) string {
	layout := "2006-01-02"
	return t.Format(layout)
}

func date(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{"fdate": formatDate}
	t := template.New("date.html").Funcs(funcMap)
	t, _ = t.ParseFiles("date.html")
	t.Execute(w, time.Now())
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	http.HandleFunc("/compare", compare)
	http.HandleFunc("/iterate", iterate)
	http.HandleFunc("/assign", assign)
	http.HandleFunc("/include", include)
	http.HandleFunc("/date", date)
	server.ListenAndServe()
}
