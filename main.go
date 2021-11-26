package main

import (
	"net/http"
	"text/template"
)

type Index struct {
	Content string
}

//入口文件
func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		t, _ := template.ParseFiles("assets/html/index.html")
		t.Execute(rw, &Index{
			Content: "中国真美好",
		})
	})
	err := http.ListenAndServe(":8080", nil)
	if nil != err {
		panic(nil)
	}

}
