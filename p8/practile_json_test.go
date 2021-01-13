package p8

import (
    "fmt"
    "testing"
)

func TestPracticeJson(t *testing.T) {
    gr := PracticeJson()
    rd := gr.ResponseData

    rs := rd.Results

    for i, r := range rs {
        fmt.Printf("%d : %s => %s\n",
            i,
            r.Title,
            r.Content)
        fmt.Printf("2 > ${i} : #{r.Title} => #{r.Content}\n")
    }
}


func TestPracticeJsonString(t *testing.T) {
    PracticeJsonString()
}

func TestJson2Map(t *testing.T) {
    Json2Map()
}

func TestMap2Json(t *testing.T) {
    Map2Json()
}