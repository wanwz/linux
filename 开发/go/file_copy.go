package main

import (
    "os"
    "fmt"
    "io"
    "io/ioutil"
)

func main() {
    srcFile := "/opt/aa.txt"
    dstFile := "/opt/ab.txt"
    
    total, err := CopyFile1(srcFile, dstFile)
    fmt.Println(total, err)
}

func CopyFile1(srcFile,dstFile string) (int64, error) {
    file1, err := os.Open(srcFile)
    if err != nil {
        return 0, err
    }
    
    file2, err := os.OpenFile(dstFile, os.O_WRONLY|os.O_CREATE, os.ModePerm)
    if err != nil {
        return 0, err
    }
    
    defer file1.Close()
    defer file2.Close()
    
    bs := make([]byte, 1024, 1024)
    n := -1
    total := 0
    for {
        n, err = file1.Read(bs)
        if err == io.EOF || n == 0 {
            fmt.Println("the copy is over!")
            break
        }else if err != nil {
            fmt.Println("something is wrong!")
            return total, err
        }
        total += n
        file2.Write(bs[:n])
    }
    return total, nil
}

func CopyFile2(srcFile, dstFile string) (int64, error) {
    file1, err := os.Open(srcFile)
    if err != nil {
        return 0, err
    }
    file2, err := os.OpenFile(dstFile, os.O_WRONLY|os.O_CREATE, os.ModePerm)
    if err != nil {
        return 0, err
    }
    
    defer file1.Close()
    defer file2.Close()
    return io.Copy(file2, file1)
}

func CopyFile3(srcFile, dstFile string) (int, error) {
    bs, err := ioutil.ReadFile(srcFile)
    if err != nil {
        return 0, err
    }
    
    err = ioutil.WriteFile(dstFile, bs, 0777)
    if err != nil {
        return 0, err
    }
    return len(bs), nil
}
