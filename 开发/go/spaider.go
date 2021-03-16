package main

import (
    "fmt"
    "regexp"
    "time"
    "sync"
    "net/http"
    "io/ioutil"
    "strconv"
    "strings"
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

var (
    chanImgUrls chan string
    chanTask chan string
    wg sync.WaitGroup
    reImg = `https?://[^"]+?(\.((jpg)|(jpeg)|(png)|(gif)|(bmp)))`
)

func main() {
    //DownloadFile("https://xxx", "1.jpg")
    chanImgUrls = make(chan string, 1000000)
    chanTask = make(chan string, N) //N以预计任务数填写
    for i := 0; i < N; i++ {
        wg.Add(1)
        go getImgUrls("https://xxx" + strconv.Itoa(i) + "xx.html")
    }
    wg.Add(1)
    go checkOK()
    for i :=0; i < N; i++ { //N代表开几个协程
        wg.Add(1)
        go DownloadImg()
    }
    wg.Wait()
}

func getImgUrls(url string) {
    urls := getImgs(url)
    for _, url := range urls {
        chanImgUrls <- url
    }
    chanTask <- url
    wg.Done()
}

func getImgs(url string) (urls []string) {
    pageStr := GetPageStr(url)
    re := regexp.MustCompile(reImg)
    results := re.FindAllStringSubMatch(pageStr, -1)
    fmt.Printf("%d条需爬取~", len(results))
    for _, result := range results {
        url = result[0]
        urls = append(urls, url)
    }
    return
}

func GetPageStr(url string) (pageStr string) {
    resp, err := http.Get(url)
    HandleError(err, "http.Get.url")
    defer resp.Body.Close()
    pageBytes, err := ioutil.ReadAll(resp.Body)
    HandleError(err, "ioutil.ReadAll resp.Body")
    pageStr = string(pageBytes)
    return pageStr
}

func DownloadImg() {
    for url := range chanImgUrls {
        filename := GetFilenameFromUrl(url)
        ok := DownloadFile(url, filename)
        if !ok {
            fmt.Printf("%s 下载失败\n", filename)
        } else {
            fmt.Printf("%s 下载成功\n", filename)
        }
    }
    wg.Done()
}

func GetFilenameFromUrl(url string) (filename string) {
    lastIdx := strings.LastIndex(url, "/")
    filename = url[lastIdx+1:]
    timePrefix := strconv.Itoa(int(time.Now().UnixNano()))
    filename = timePrefix + "_" + filename
    return
}

func checkOK() {
    var count int
    for {
        url := <- chanTask
        fmt.Printf("%s 已完成爬取任务\n", url)
        count++
        if count == N { //N是之前预计爬取多少条
            close(chanImgUrls)
            break
        }
    }
    wg.Done()
}
