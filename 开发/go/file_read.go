package main

import (
    "os"
    "fmt"
    "io"
)

func main() {
    //读取数据
    file, err := os.Open("/opt/file.txt")
    if err != nil {
        fmt.Println("err:", err)
        return
    }
    
    defer file.Close()
    
    bs := make([]byte, 4, 4)
    n := -1
    for {
        n, err = file.Read(bs)
        if n == 0 || err == io.EOF {
            fmt.Println("文件已读完，结束读取操作！")
            break
        }
        fmt.Println(string(bs[:n]))
    }
}
