package main

import (
	"fmt"
	"os"
)

//创建文件
func create_file() {
	// Create creates or truncates the named file. If the file already exists,
	// it is truncated. If the file does not exist, it is created with mode 0666
	// (before umask). If successful, methods on the returned File can
	// be used for I/O; the associated file descriptor has mode O_RDWR.
	// If there is an error, it will be of type *PathError.
	/*
		新建一个文件，如果文件已经存在，它会被覆盖掉；如果不存在，会创建一个权限为666（可读可写）的文件

	*/

	f, err := os.Create("a.txt")
	if err == nil {
		fmt.Printf("f.Name(): %v\n", f.Name())
	} else {
		fmt.Printf("err: %v\n", err)
	}
}

// 创建目录
func create_dir() {
	// Mkdir creates a new directory with the specified name and permission
	// bits (before umask).
	// If there is an error, it will be of type *PathError.
	//创建单个目录 os.ModePerm:权限777 可读可写可执行
	err := os.Mkdir("aaa", os.ModePerm)
	if err == nil {
		fmt.Printf("目录创建成功！\n")
	} else {
		fmt.Printf("err: %v\n", err)

	}

	//创建递归目录
	err2 := os.MkdirAll("aaa/bbb/ccc", os.ModePerm)
	if err2 == nil {
		fmt.Printf("目录创建成功！\n")

	} else {
		fmt.Printf("err2: %v\n", err2)
	}

	//创建临时目录
	s, err3 := os.MkdirTemp(".", "temp") //在当前目录下创建一个临时文件
	if err3 == nil {
		fmt.Printf("s: %v\n", s)
	} else {
		fmt.Printf("err3: %v\n", err3)
	}

}

//删除目录和文件
func delete_dir_file() {

	//删除文件
	err := os.Remove("a.txt")
	if err == nil {
		println("文件删除成功！")
	} else {
		fmt.Printf("err: %v\n", err)
	}

	//删除(空)目录，如果目录非空就会报错
	err2 := os.Remove("aaa")
	if err2 == nil {
		println("目录删除成功！")
	} else {
		fmt.Printf("err2: %v\n", err2)
	}

	//删除递归目录 removes path and any children it contains.
	err3 := os.RemoveAll("aaa")
	if err3 == nil {
		println("目录删除成功！")
	} else {
		fmt.Printf("err3: %v\n", err3)
	}

}

//获取工作目录
func pwd() {
	dir, err := os.Getwd()
	if err == nil {
		fmt.Printf("dir: %v\n", dir)
	} else {
		fmt.Printf("err: %v\n", err)
	}
}

//修改工作目录
func chdir() {

	// Chdir changes the current working directory to the named directory.
	err := os.Chdir("/home/lll/Desktop")
	if err == nil {
		fmt.Println(os.Getwd())
	} else {
		fmt.Printf("err: %v\n", err)
	}

}

//获取临时目录
func temp_dir() {
	s := os.TempDir()
	fmt.Printf("s: %v\n", s)
}

//重命名文件
func rename() {
	// Rename renames (moves) oldpath to newpath.
	// If newpath already exists and is not a directory, Rename replaces it.
	// OS-specific restrictions may apply when oldpath and newpath are in different directories.
	// If there is an error, it will be of type *LinkErr
	/*

		移动旧路径的东西到新路径中，如果新路径中有同名的，会被覆盖掉
	*/
	err := os.Rename("aaa", "bbb")
	if err == nil {
		println("重命名成功！")
	} else {
		fmt.Printf("err: %v\n", err)
	}
}

// 读文件
func read_file() {

	//读取指定路径的文件，读取成功，返回对应的字节数组
	b, err := os.ReadFile("b.txt")
	if err == nil {
		fmt.Printf("b: %c\n", b)
	} else {
		fmt.Printf("err: %v\n", err)
	}
}

// 写文件
func write_file() {
	// WriteFile writes data to the named file, creating it if necessary.
	// If the file does not exist, WriteFile creates it with permissions perm (before umask);
	// otherwise WriteFile truncates it before writing, without changing permissions.
	/*
		如果文件已经存在，就会被覆盖
		如果文件不存在，创建一个新的文件，并赋予文件传入的权限
	*/
	data := []byte("hello world!")

	err := os.WriteFile("c.txt", data, os.ModePerm)
	if err == nil {
		println("文件写入成功！")
	} else {
		fmt.Printf("err: %v\n", err)
	}
}

func main() {

	create_file()
	create_dir()
	delete_dir_file()
	pwd()
	// chdir()
	temp_dir()
	rename()
	read_file()
	write_file()
}
