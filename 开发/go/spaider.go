package main

import (
    "fmt"
    "regexp"
    "time"
    "sync"
    "net/http"
    "ioutil"
)

func DownloadFile(url string, filename string) (ok bool) {
    resp, err := http.Get(url)
    HandleError(err, "http.Get url")
    defer resp.Body.Close()
    bytes, err := ioutil.ReadAll(resp.Body)
    HandleError(err, "ioutil.ReadALL urls")
    filename = "/opt/img/" + filename
    err := ioutil.WriteFile(filename, bytes, 0644)
    if err != nil {
        return false
    } else {
        return true
    }
}

func HandleError(err error, why string) {
    if err != nil {
        fmt.Println(why, err)
    }
}
