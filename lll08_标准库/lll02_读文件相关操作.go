package main

import (
	"fmt"
	"os"
)

func main() {

	//打开文件

	// OpenFile is the generalized open call; most users will use Open
	// or Create instead. It opens the named file with specified flag
	// (O_RDONLY etc.). If the file does not exist, and the O_CREATE flag
	// is passed, it is created with mode perm (before umask). If successful,
	// methods on the returned File can be used for I/O.
	// If there is an error, it will be of type *PathError.
	/*
		OpenFile是一个通用的“打开”调用方式，但是大多数用户会使用Open或者Create来代替它
		它以特殊的标识打开文件（就是文件的打开方式）
		在文件不存在的时候，如果传递了O_CREATE标识，它会创建文件并赋予文件传入的perm权限（就是说第三个参数只有在文件不存在且传入了O_CREATE标识的情况下才有作用）

	*/

	/*
		OpenFile(name string, flag int, perm FileMode)
			name：文件名
			flag：文件的打开方式，有如下可选的方式
				// Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.
				O_RDONLY int = syscall.O_RDONLY // open the file read-only. 只读
				O_WRONLY int = syscall.O_WRONLY // open the file write-only. 只写
				O_RDWR   int = syscall.O_RDWR   // open the file read-write. 可读可写
				// The remaining values may be or'ed in to control behavior.
				O_APPEND int = syscall.O_APPEND // append data to the file when writing. 写的时候末尾追加
				O_CREATE int = syscall.O_CREAT  // create a new file if none exists. 没有就新建
				O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist.
				O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O.
				O_TRUNC  int = syscall.O_TRUNC  // truncate regular writable file when opened.
			perm：赋予文件的权限


	*/
	f, err := os.OpenFile("a.txt", os.O_RDONLY|os.O_CREATE, 0777) //ugo 从第二位开始（0不能省掉！），三个数字分别对应user，group，other的权限
	if err == nil {
		fmt.Printf("f.Name(): %v\n", f.Name())
		for {
			buf := make([]byte, 10) //创建一个缓冲区来接收读取的内容
			_, err2 := f.Read(buf)
			if err2 != nil {
				fmt.Printf("err2: %v\n", err2) //直到读到文件末尾就结束循环
				break
			}
			fmt.Printf("buf: %v\n", string(buf))

		}

	} else {
		fmt.Printf("err: %v\n", err)
	}
	f.Close() //关闭文件

	//打开一个文件 调用Open方法其实就是以只读的方式打开文件,底层调用的还是OpenFile
	// 等价于：OpenFile(name, O_RDONLY, 0)
	f2, err2 := os.Open("a.txt")
	if err2 == nil {
		fmt.Printf("f2.Name(): %v\n", f2.Name())
	}

	//创建一个文件 等价于：OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666)
	f3, err3 := os.Create("aa.txt")
	if err3 == nil {
		fmt.Printf("f3.Name(): %v\n", f3.Name())
	}
	f3.Close()
	// 读文件
	f4, err4 := os.Open("aaa.txt")
	if err4 == nil {

		//Read
		for {
			//每次读35个字节的内容
			con := make([]byte, 35)
			_, err5 := f4.Read(con)
			if err5 == nil {
				fmt.Printf("con: %v\n", string(con))
			} else {
				fmt.Printf("err5: %v\n", err5)
				break
			}
		}

		//ReadAt //从指定的位置开始读
		con := make([]byte, 10)
		_, err5 := f4.ReadAt(con, 3) //从第三个字节的位置开始
		if err5 == nil {
			fmt.Printf("con: %v\n", string(con))

		}
	} else {
		fmt.Printf("err4: %v\n", err4)
	}
	f4.Close()

}
