package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

var (
	cleanPath string
	numSave   int
	daySave   int
	vrbs      bool
)

func init() {
	flag.StringVar(&cleanPath, "c", "", "Directories to be cleaned [Absolute path]")
	flag.IntVar(&numSave, "n", 10, "Number of copies to be kept")
	flag.IntVar(&daySave, "t", 10, "Retention days")
	flag.BoolVar(&vrbs, "v", false, "Print details[-v true]")
}

func main() {
	flag.Usage = func() {
		fmt.Printf("Usage of %s: pw-cleanup [OPTIONS]...\n", os.Args[0])
		fmt.Printf("Perfect World Game Management Tool.\n")
		flag.PrintDefaults()
		fmt.Printf("Examples\n")
		fmt.Printf("  pw-cleanup -c /var/test -n 5 -t 3\n")
	}
	flag.Parse()

	//cleanup verbose
	rst := cleanup(cleanPath, daySave, numSave, vrbs)
	if rst == true {
		fmt.Println("[INFO] Done.")
	} else {
		os.Exit(10)
	}

}

func cleanup(cleanPath string, daySave, numSave int, vrbs bool) (result bool) {
	if cleanPath == "" {
		fmt.Printf("[ERROR] Args [cleanPath] can't be empty\n")
		return false
	}

	nameDir, err := filepath.Glob(filepath.Join(cleanPath, "*"))
	if err != nil {
		fmt.Printf("[ERROR] %s\n", err)
		return false
	}
	//fmt.Println(nameDir)

	curtNum := len(nameDir)
	// curTime := time.Now().Unix()
	curTime := time.Now()

	if curtNum > numSave {
		//fmt.Println(curtNum, numSave)
		for _, files := range nameDir {
			fileInfo, err := os.Stat(files)
			if err != nil {
				fmt.Printf("[ERROR] %s\n", err)
				return false
			}
			fileTime := fileInfo.ModTime()
			subD := int(curTime.Sub(fileTime).Hours() / 24)
			//fmt.Println(subD)
			if subD > daySave {
				//以保留天数作为基准来删除文件
				//fmt.Printf("Absolute path: %s\n", filepath.Join(cleanPath, fileInfo.Name()))
				delDir := filepath.Join(cleanPath, fileInfo.Name())
				if vrbs == true {
					fmt.Printf("[INFO] The file has been deleted: %s\n", delDir)
				}
				err := os.RemoveAll(delDir)
				if err != nil {
					fmt.Printf("[ERROR] File deletion exception: %s\n", err)
					return false
				}
			}
			//fmt.Printf("当前时间戳：%v; 文件时间: %s; 文件时间戳：%d; 时间相差天数: %d\n", curTime, fileTime, fileTime.Unix(), int(subD.Hours()/24))
		}

	} else {
		fmt.Printf("[ERROR] The number of files in the current directory is less than the number of reserved copies, and no delete operation is performed\n")
		return false
	}
	return true
}
