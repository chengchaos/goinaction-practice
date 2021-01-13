package cgiserv

import (
    "log"
    "os"
)

func SetOut() {

    curr, err := os.Getwd()
    if err != nil {
        log.Println("Get Error: ", err)
        return
    }
    log.Println("Current Dir =>", curr)



}
