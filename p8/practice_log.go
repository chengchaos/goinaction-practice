package p8

import (
    "log"
)

func init() {
    log.SetPrefix("TRACE: ")
    log.SetFlags(log.Ldate |
        log.Lmicroseconds |
        //log.Llongfile)
        log.Lshortfile,
    )
}

func PracticeLog() {
    // Println 写到标准日志记录器
    log.Println("message")

    // Fatalln 在调用 Println 后会调用 os.Exit(1)
    //log.Fatalln("fatal message")

    // Panicln 在调用 Println 后会调用 panic()
    log.Panicln("panic message")
}
