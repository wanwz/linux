package main

import (
    "fmt"
    "sort"
)

func main() {
    map1 := make(map[int]string, 5)
    map1[1] = "aaa"
    map1[2] = "bbb"
    map1[3] = "ccc"
    map1[4] = "ddd"
    map1[5] = "eee"
    map1[6] = "eab"
    sli := []int{}
    for k, _ := range map1 {
        sli = append(sli, k)
    }
    sort.Ints(sli)
    for i := 0; i < len(map1); i++ {
        fmt.Println(map1[sli[i]])
    }
}
