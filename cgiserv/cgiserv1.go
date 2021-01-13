package cgiserv

import (
    "log"
    "net/http"
    "net/http/cgi"
    "os"
)

/*
https://www.cnblogs.com/yjf512/archive/2012/12/25/2831891.html
 */
func SetOut() {

    curr, err := os.Getwd()
    if err != nil {
        log.Println("ERROR: ", err)
        return
    }

    log.Println("current dir =>", curr)

    http.HandleFunc("/cgi-bin/", func(w http.ResponseWriter, r *http.Request) {
        handler := new(cgi.Handler)
        handler.Path = curr + "/cgiserv/cgiserv1_runner" + r.URL.Path
        log.Println("handler.Path =>", handler.Path)
        handler.Dir = curr + "/"
        handler.ServeHTTP(w, r)
    })

    log.Fatal(http.ListenAndServe(":7777", nil))
}
