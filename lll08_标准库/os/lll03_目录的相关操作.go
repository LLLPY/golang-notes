package main

import (
	"fmt"
	"os"
)

func dir(dir_name string) {
	f, err := os.Open(dir_name)
	if err != nil {
		return
	}
	de, err2 := f.ReadDir(-1)
	if err2 != nil {
		return
	}

	for _, v := range de {
		fmt.Println(dir_name, ":", v.Name())
		if v.IsDir() {
			cur_dir_path := dir_name + "/" + v.Name()
			dir(cur_dir_path) //如果是目录，就继续递归调用
		}
	}

}

func main() {

	// If n > 0, ReadDir returns at most n DirEntry records.
	// In this case, if ReadDir returns an empty slice, it will return an error explaining why.
	// At the end of a directory, the error is io.EOF.
	//
	// If n <= 0, ReadDir returns all the DirEntry records remaining in the directory.
	//打开一个目录 n大于0，返回n个文件（目录）对象，如果没有n个，返回所有；n小于，返回所有文件（目录）对象

	f, err := os.Open("bbb")
	defer f.Close()
	if err == nil {
		de, err2 := f.ReadDir(-1) //-1<0 ===>获取目录下的所有内容
		if err2 == nil {
			for _, v := range de {
				fmt.Printf("v.IsDir(): %v\n", v.IsDir()) //判断是不是一个目录
				fmt.Printf("v.Name(): %v\n", v.Name())   //返回文件[目录]名
			}
		}
	}

	//获取当前目录下的所有文件和目录
	// dir(".")
	var a string
	fmt.Scan(a)

}
