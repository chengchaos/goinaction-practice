package p8

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "reflect"
)

type (
    GResult struct {
        Title   string `json:"title"`
        Content string `json:"content"`
    }
    GResponse struct {
        ResponseData struct {
            Results []GResult `json:"results"`
        } `json:"responseData""`
    }
)

func PracticeJson() (gr *GResponse) {

    uri := "http://localhost:8888/get-data"
    resp, err := http.Get(uri)
    if err != nil {
        log.Println("ERROR: ", err)
        return
    }
    defer resp.Body.Close()

    // 将 JSON 解码到结构类型
    //var gr *GResponse
    // NewDecoder 传入一个 Reader
    err = json.NewDecoder(resp.Body).Decode(&gr)
    if err != nil {
        log.Println("ERROR: ", err)
        return
    }

    fmt.Println("gr =>", gr)
    return
}


type Contact struct {
    Name string `json:"name"`
    Title string `json:"title"`
    Contact struct{
        Home string `json:"home"`
        Cell string `json:"cell"`
    } `json:"contact"`
}

var JSON = `{
    "name" : "程超",
    "title" : "先生",
    "contact" : {
        "home" : "415.333.3333",
        "cell" : "215.555.5555"
    }
}`
func PracticeJsonString() {
    var c Contact
    // 字符串要先转换成 byte 的切片。
    err := json.Unmarshal([]byte(JSON), &c)
    if err != nil {
        log.Println("ERROR: ", err)
        return
    }

    fmt.Println("c => ", c)
}


func Json2Map() {
    // 将 JSON 放到 Map 中。
    var c map[string]interface{}
    err := json.Unmarshal([]byte(JSON), &c)
    if err != nil {
        log.Println("ERROR: ", err)
        return
    }

    fmt.Println("Name: ", c["name"])
    fmt.Println("Title: ", c["title"])
    fmt.Println("Home: ", c["contact"].(map[string]interface{})["home"])
    fmt.Println("Cell: ", c["contact"].(map[string]interface{})["cell"])

}

func Map2Json() {
    c := make(map[string]interface{})
    c["name"] = "程超"
    c["title"] = "Mr."
    c["contact"] = map[string]interface{}{
        "home" : "412.3434.3333",
        "cell" : "152.3333.4444",
    }

    data, err := json.MarshalIndent(c, "", "    ")
    if err != nil {
        log.Println("ERROR: ", err)
        return
    }

    // []uint8
    fmt.Println("type =>", reflect.TypeOf(data))
    fmt.Println(string(data))




}