package main

import (
    "fmt"
    "sort"
)

func main() {
    map1 := make(map[int]int, 5)
    map1[1] = 1
    map1[2] = 2
    map1[5] = 5
    map1[3] = 3
    map1[4] = 4
    sli := []int{}
    for k, _ := range map1 {
        sli = append(sli, k)
    }
    sort.Ints(sli)
    for i := 0; i < len(map1); i++ {
        fmt.Println(map1[sli[i]])
    }
}
