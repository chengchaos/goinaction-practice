package main

import (
	"log"
	"net/http"
	"net/http/cgi"
)

func init() {
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate |
		log.Lmicroseconds |
		//log.Llongfile)
		log.Lshortfile,
	)
}
func index(w http.ResponseWriter, r *http.Request) {

	//files := []string{
	//	"templates/layout.html",
	//	"templates/navbor.html",
	//	"templates/index.html",
	//}
	//templates := template.Must(template.ParseFiles(files...))
	//threads, err := data.Threads()
	//if err == nil {
	//	templates.ExecuteTemplate(w, "layout", threads)
	//}
}

var pythonHandler = new(cgi.Handler)
var pythonPath string = "/usr/bin/python3"

func handlePython(w http.ResponseWriter, r *http.Request) {

	scriptFile := "/Users/chengchao/tmp" + r.URL.Path
	log.Printf("request: %s => script: %s",
		r.URL.Path,
		scriptFile)

	pythonHandler.Path = pythonPath
	pythonHandler.Args = append(pythonHandler.Args, scriptFile)

	pythonHandler.ServeHTTP(w, r)
}

func main() {

	// 一个默认多路复用器
	mux := http.NewServeMux()

	// 指定目录的静态文件服务器
	files := http.FileServer(http.Dir("/Users/chengchao/tmp"))

	mux.Handle("/static/", http.StripPrefix("/static/", files))
	mux.HandleFunc("/cgi-bin/", handlePython)
	mux.HandleFunc("/", index)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}
