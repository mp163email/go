package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	//连接路径
	join := filepath.Join("dir", "subdir", "file.txt")
	fmt.Println("Joined path : ", join)

	//分割路径 路径和文件名分割开来 filepath.Split
	dir, file := filepath.Split(join)
	fmt.Println(dir, file)

	//获取绝对路径 filepath.Abs
	abs, err := filepath.Abs("file.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(abs)

	//遍历目录及目录下的子目录和文件 filepath.Walk
	err1 := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			panic(err)
		}
		fmt.Println(path, info.IsDir())
		return nil
	})
	if err1 != nil {
		panic(err1)
	}

}
