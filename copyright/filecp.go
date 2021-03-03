package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	srcPath := flag.String("path", "", "输入要生成的路径")
	flag.Parse()
	if *srcPath == "" {
		fmt.Println("路径为空请重试")
		return
	}

	err := filepath.Walk(*srcPath, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}

		ok := strings.HasSuffix(path, ".aspx") || strings.HasSuffix(path, ".js") || strings.HasSuffix(path, ".cs")
		if ok {
			if err != nil {
				fmt.Println("读取文件错误", err)
			}
			f, err := os.OpenFile("d://code.txt", os.O_WRONLY|os.O_CREATE, 0644)
			if err != nil {
				fmt.Println("打开文件错误", err)

			} else {
				b, err := ioutil.ReadFile(path) // just pass the file name
				if err != nil {
					fmt.Print(err)
				}

				// 查找文件末尾的偏移量
				n, _ := f.Seek(0, 2)

				// 从末尾的偏移量开始写入内容
				_, err = f.WriteAt(b, n)
			}

		}
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}
