# 《go语言修炼手册》





## 简介



​				Go is **expressive**, **concise**, **clean**, and **efficient**. Its **concurrency mechanisms** make it easy to write programs that get the most out of multicore and networked machines, while its novel type system enables flexible and modular program construction. Go **compiles quickly** to machine code yet has the convenience of garbage collection and the power of run-time reflection. It's a **fast**, **statically typed**, **compiled language** that feels like a dynamically typed, interpreted language.

> 根据官方文档的简介，可以得出如下如下几点：
>
> - go语言是一个高效，简洁和富有表现力的语言
> - go语言天生具有并发机制
> - go语言编译快
> - go语言是一门快速的，静态的，编译性语言





## 安装



### 1.下载



下载地址：[golang](https://golang.google.cn/dl/)

如下图所示，找到golang的下载地址后，下载对应的linux版的最新版

![img](https://img-blog.csdnimg.cn/de69d45f1c174e64a5009871286a6d64.png)![点击并拖拽以移动](data:image/gif;base64,R0lGODlhAQABAPABAP///wAAACH5BAEKAAAALAAAAAABAAEAAAICRAEAOw==)编辑



### 2.安装



我们按照官方给的提示安装：

![img](https://img-blog.csdnimg.cn/e3cc07dee84745b0a405e8115cf53420.png)![点击并拖拽以移动](data:image/gif;base64,R0lGODlhAQABAPABAP///wAAACH5BAEKAAAALAAAAAABAAEAAAICRAEAOw==)编辑

#### 2.1移除之前的安装残留，同时安装当前版本

执行命令(普通用户可能权限不够，所以加了sudo)：

```bash
sudo rm -rf /usr/local/go && tar -C /usr/local -xzf go1.18.4.linux-amd64.tar.gz
```

![点击并拖拽以移动](data:image/gif;base64,R0lGODlhAQABAPABAP///wAAACH5BAEKAAAALAAAAAABAAEAAAICRAEAOw==)

#### 2.2 添加环境变量

执行命令：

```bash
export PATH=$PATH:/usr/local/go/bin
```

![点击并拖拽以移动](data:image/gif;base64,R0lGODlhAQABAPABAP///wAAACH5BAEKAAAALAAAAAABAAEAAAICRAEAOw==)

#### 2.3检查是否安装成功

执行命令：

```bash
go version
```

![点击并拖拽以移动](data:image/gif;base64,R0lGODlhAQABAPABAP///wAAACH5BAEKAAAALAAAAAABAAEAAAICRAEAOw==)

 如果提示了go语言的版本，就说明安装成功啦！

![img](https://img-blog.csdnimg.cn/28ba454d83824328a6f615ca2228563f.png)![点击并拖拽以移动](data:image/gif;base64,R0lGODlhAQABAPABAP///wAAACH5BAEKAAAALAAAAAABAAEAAAICRAEAOw==)编辑



### 3.编写第一个go文件 ！



打开任一编辑器，输入如下代码(以hello.go命名文件)：

```Go
package main
import "fmt"

func main()  {
	
	fmt.Println("hello,go!")
}
```

![点击并拖拽以移动](data:image/gif;base64,R0lGODlhAQABAPABAP///wAAACH5BAEKAAAALAAAAAABAAEAAAICRAEAOw==)

同目录下在终端中输入：

```bash
go run hello.go
```

![点击并拖拽以移动](data:image/gif;base64,R0lGODlhAQABAPABAP///wAAACH5BAEKAAAALAAAAAABAAEAAAICRAEAOw==)

![img](https://img-blog.csdnimg.cn/ce848144e51a43c8a4f31e4f99cdfe1e.png)![点击并拖拽以移动](data:image/gif;base64,R0lGODlhAQABAPABAP///wAAACH5BAEKAAAALAAAAAABAAEAAAICRAEAOw==)编辑

 如上，执行成功！

同步更新于个人博客系统：***\*[golang学习笔记系列之go语言的环境搭建(linux系统下)](http://www.lll.plus/learningPlanet/750)\****



## 代码的组织

go应用使用**包**和**模块**来组织代码，包对应到文件系统就是文件夹，模块就是go的源文件。一个包中可以有很多模块，或者多个子包。

- 包<----->文件夹
- 模块<----->go文件



### 项目管理工具

早期的go项目使用gopath来管理项目，不方便而且容易出错，从golang1.11开始使用gomod管理项目，除此之外，还有第三方项目管理工具，例如govendor。



#### go mod包管理步骤

1. 创建项目     =======>       mkdir go_pro
2. 初始化项目 =======>       进入到项目的根目录，go mod init go_pro   
3. 创建包         =======>       在项目中创建新的包
4. 创建模块     =======>       在项目中直接创建模块或者在新创的包中创建模块
5. 相互调用     =======>       项目中各模块之间可以相互调用



#### go mod使用方法

- 初始化模块：

  ```
  go mod init <项目模块名称>
  ```

- 依赖关系处理，根据go.mod文件

  ```
  go mod tidy
  ```

- 将依赖包复制到项目下的vendor目录

  ```
  go mod vendor
  ```

- 显示依赖关系

  ```
  go list -m all
  ```

- 显示详细依赖关系

  ```
  go list -m -json all
  ```

- 下载依赖

  ```
  go mod download [path@version]
  ```

  



## 标识符，关键字以及命名规则



### 标识符

标识符的英文是**identifier**，通俗地讲，就是给变量，常量，函数，结构体，数组，切片，接口起名字。



#### 标识符的规范要求

- 由数字，字母，下划线组成
- 不能以数字开头
- 区分大小写
- 尽量做到见名知意



```go
//正确的标识符
var abc string
var a12 int
var _123 int[]

//错误的标识符
var 123abc int //不能以数字开头
var abc&afa string //出现了未知的字符
```



### 关键字

Go 共有 25 个保留关键字，各有其作用，不能用作[标识符](https://so.csdn.net/so/search?q=标识符&spm=1001.2101.3001.7020)。Go 的 25 个关键字按照作用可以分为 3 类，分别为包管理、程序实体声明与定义与程序流程控制。

```shell
包管理（2个）：
	import	package

程序实体声明与定义（8个）：
	chan	const	func	interface	map	struct	type	var

程序流程控制（15个）：
	break	case	continue	default	defer	else	fallthrough	
	for		go		goto		if		range	return	select		switch
```

除了上面的保留关键字，go语言还有36个预定义标识符，其中包含了基本类型名称和一些基本的内置函数，如下表：

| append    | bool       | byte   | cap   | close   | complex |
| --------- | ---------- | ------ | ----- | ------- | ------- |
| complex64 | complex128 | uint16 | copy  | false   | float32 |
| float64   | imag       | int    | int8  | int16   | uint32  |
| int32     | int64      | iota   | len   | make    | new     |
| nil       | panic      | unit64 | print | println | real    |
| recover   | string     | true   | uint  | uint8   | uintprt |



### 命名规则



#### 区分大小写

命名规则涉及变量，常量，全局函数，结构，接口，方法等的命名。go语言从语法层面进行了以下限定：任何需要对外暴露的名字必须以大写字母开头，不需要对外暴露的则应该以小写字母开头。



当命名以一个大写字母开头，如：Hello，那么使用这种形式的标识符的对象就**可以被外部包的代码所使用**（前提是导入了这个包），这被称为导出（类似于java中的public）；命名如果以小写字母开头，则**对包外是不可见的**，但是它们在整个包的内部是可见并可用的（类似于java中的private）。

![image-20220823112855304](/home/lll/snap/typora/57/.config/Typora/typora-user-images/image-20220823112855304.png)





#### 包名称

保持package的名称和目录一致，尽量采取有意义的包名，简洁明了，尽量不要和标准库冲突。包名应该为**小写**单词，不要使用下划线或者混合大小写。



#### 文件名

尽量采取简洁明了的文件名，简短，有意义，应该为**小写**单词，使用**下划线**分割各个单词。



#### 结构体命名

一般采用驼峰命名法，首字母根据访问控制来规定大小写。

```go
//客户订单
type CustomerOrder struct {
    Name string
    Address string 
}

order := CustomerOrder{"tom", "上海"}
```



#### 接口命名

命名规则同结构体。

单个函数的结构名以“er”作为后缀，例如：Reader，Writer。

```go
type Reader interface {
    Read(p []byte) (n int, err error) 
    
}
```



#### 变量命名

基本命名规则同结构体，但遇到特有名词时，需要遵循以下规则：

- 如果变量为私有，且特有名词为首个单词，则使用小写，如appService若变量类型为bool，则名称应以Has，Is，Can或Allow开头

- ```go
  var isExist bool
  var hasConflict bool
  var canManage bool
  var allowGitHook bool
  ```



#### 常量命名

常量均使用大写字母组成，并且使用下划线分割单词。

```go
const APP_URL = “http://www.lll.plus”
```

如果是枚举类型的常量，需要首先创建相应的类型

```go
type Scheme string

const (
	HTTP Scheme = "http"
    HTTPS Scheme = "https"
)
```

 

### 错误处理

错误处理的原则就是不能丢弃任何有返回err的调用，不要使用丢弃，必须马上处理。**接收到错误，要么返回err，或者使用log记录下来尽早return**；一旦有错误发生，马上返回，尽量不要使用panic，除非你知道你在做什么，错误描述如果英文必须为小写，不需要标点结尾，采用独立的错误流进行处理。

```go
//错误的写法
if err != nil {
    //错误处理
}else {
    //正常代码
}


//正确写法
if err != nil {
    //错误处理
    return //或者继续
}
```



### 单元测试

单元测试文件名命名规范为**example_test.go**，测试用例的函数名称必须以Test开头，例如：==TestExample==，每个重要的函数都要首先编写测试用例，测试用例和正规代码一起提交方便进行回归测试。



## 变量和常量



### 变量

变量是计算机语言中能存储计算结果或能表示值的抽象概念，不同的的变量保存的数据类型可能不一样。



#### 声明变量

go语言中的变量需要声明后才能使用，同一作用域内不支持重复声明。并且go语言的变量声明后**必须使用**，否者会报错。



#### 默认语法

```go
var indertifer type

//例：
var age int    //int类型的变量不赋值，默认值是0
var price float64 //默认值是0
var flag bool  //默认值是false 
```



#### 类型推断

不需要为变量指定类型，根据变量的值来自动推断变量的类型。

```go
var name = “TOm”
var num = 10
```

 

#### 同时声明多个变量

```go
var name,age,gender = "Tom",18,true
```



#### 短变量声明

短变量只能在**函数内部**进行声明，使用**:=**运算符对变量进行声明和初始化。

```go
identifer := ""

//例：注意，一定是在函数内部声明，否则会报错！

func fun(){
    age:=18
	name:=tom
}

```



#### 匿名变量

如果我们接收到多个变量，有一些变量使用不到，可以使用下划线，便是变量名，这种变量叫做匿名变量。

```go
func Fun() (string, int) {
	return "Tom", 18
}
//匿名变量 匿名变量不使用也不会报错
tom_name, _ := Fun()
fmt.Printf("tom_name: %v\n", tom_name)
```



### 常量

常量，就是在程序编译阶段就确定下来的值，二程序在运行时则无法改变该值。在go语言中，常量可以是数值类型，布尔类型，字符串类型等等。



#### 常量声明

定义一个常量使用==const==关键字，语法格式如下：

```go
const constantName [type] = value
```

- const：定义常量的关键字
- constantName：常量名
- type：常量的类型，可选
- value：常量的值



```go
package main

import "fmt"

func main() {

	//常规语法
	const PI float32 = 3.14
	const H = 100

	//定义多个
	const (
		NAME = "Tom"
		AGE  = 18
	)

	//注意：声明多个变量时，如果省略了赋值，则表示和上面的一行值相同
	const (
		a = 1
		b //b=a=1
		c = 100
		d //d=c=100
	)

	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("c: %v\n", c)
	fmt.Printf("d: %v\n", d)

	//同时定义
	const X, Y, Z = 2, 3, 4
	fmt.Printf("X: %v\n", X)
	fmt.Printf("Y: %v\n", Y)

}

```



#### iota

iota比较特殊，可以被认为是一个可被编译器修改的常量，它的默认值是0，每调用一次值就加一，遇到const关键字时被重置为0。

```go
//iota
	const (
		A = iota //0 
		B = iota //1
		C = iota //2

	)
	fmt.Printf("A: %v\n", A)
	fmt.Printf("B: %v\n", B)
	fmt.Printf("C: %v\n", C)

	//iota中间截断
	const (
		aa = iota //0
		_         //1
		cc = iota //2
	)

	fmt.Printf("aa: %v\n", aa)
	fmt.Printf("cc: %v\n", cc)


	const(
		aaa=iota  //0
		bbb=1000  //1000 [1]
		ccc=iota  //2
	)

	fmt.Printf("aaa: %v\n", aaa)
	fmt.Printf("ccc: %v\n", ccc)
```



## 基本数据类型

在go语言中，数据类型用于声明函数和变量。

数据类型的出现是为了把数据分成所需内存大小不同的数据，编程的时候需要用大数据的时候才去申请大内存，需要小数据的时候就去申请小的内存，就可以充分利用空间。

go语言按类别有以下几种数据：

| 类型       | 描述                                                         |
| ---------- | ------------------------------------------------------------ |
| 布尔类型   | 布尔类型的值只可以是常量true或false。一个简单的例子：var flag=true |
| 数字类型   | 数字类型包括整型int，浮点型float32和浮点型float64；同时支持复数。（注意：**不能用0或者非0表示条件判断**） |
| 字符串类型 | 字符串就是一串固定长度的字符连接起来的字符序列。go语言的字符串是由单个字节连接起来的，go语言的字符串的字节使用utf8编码标识unicode文本。 |
| 派生类型   | 派生类型包括：指针类型，数组类型，切片类型，结构体类型，channel类型，接口类型，map类型等等 |



### 整数类型

go语言也有基于架构的类型，例如：int，uint和uintptr。

| 类型   | 描述                                                        |
| ------ | ----------------------------------------------------------- |
| uint8  | 无符号8位整型（0到255）                                     |
| uint16 | 无符号16位整数（0到65535）                                  |
| uint32 | 无符号32位整数（0到4294967295）                             |
| uint64 | 无符号64位整数（0到18446744073709551616）                   |
| int8   | 有符号8位整数（-128到127）                                  |
| int16  | 有符号16位整数（-32768到32767）                             |
| int32  | 有符号32位整数（-2147483648到2147483647）                   |
| int64  | 有符号64位整数（-9223372036854775808到9223372036854775807） |



```go
package main

import (
	"fmt"
	"math"
	"unsafe"
)

func main() {

	var a uint8
	var b uint16
	var c uint32
	var d uint64

	fmt.Printf("a: %v %T %dB  %v~%v\n", a, a, unsafe.Sizeof(a), 0, math.MaxUint8)
	fmt.Printf("b: %v %T %dB  %v~%v\n", b, b, unsafe.Sizeof(b), 0, math.MaxUint16)
	fmt.Printf("c: %v %T %dB  %v~%v\n", c, c, unsafe.Sizeof(c), 0, math.MaxUint32)
	fmt.Printf("d: %v %T %dB  %v~%v\n", d, d, unsafe.Sizeof(d), 0, math.MaxInt64)

	var e int8
	var f int16
	var g int32
	var h int64
	fmt.Printf("e: %v %T %dB  %v~%v\n", e, e, unsafe.Sizeof(e), math.MinInt8, math.MaxInt8)
	fmt.Printf("f: %v %T %dB  %v~%v\n", f, f, unsafe.Sizeof(f), math.MinInt16, math.MaxInt16)
	fmt.Printf("g: %v %T %dB  %v~%v\n", g, g, unsafe.Sizeof(g), math.MinInt32, math.MaxInt32)
	fmt.Printf("h: %v %T %dB  %v~%v\n", h, h, unsafe.Sizeof(h), math.MinInt64, math.MaxInt64)

	var i float32
	var j float64
	fmt.Printf("i: %v %T %dB  %v~%v\n", i, i, unsafe.Sizeof(i), -math.MaxFloat32, math.MaxFloat32)
	fmt.Printf("j: %v %T %dB  %v~%v\n", j, j, unsafe.Sizeof(j), -math.MaxFloat64, math.MaxFloat64)

	var k complex64
	var l complex128
	fmt.Printf("k: %v %T %dB  \n", k, k, unsafe.Sizeof(k))
	fmt.Printf("l: %v %T %dB  \n", l, l, unsafe.Sizeof(l))

}

/*
a: 0 uint8 1B  0~255
b: 0 uint16 2B  0~65535
c: 0 uint32 4B  0~4294967295
d: 0 uint64 8B  0~9223372036854775807
e: 0 int8 1B  -128~127
f: 0 int16 2B  -32768~32767
g: 0 int32 4B  -2147483648~2147483647
h: 0 int64 8B  -9223372036854775808~9223372036854775807
i: 0 float32 4B  -3.4028234663852886e+38~3.4028234663852886e+38
j: 0 float64 8B  -1.7976931348623157e+308~1.7976931348623157e+308
k: (0+0i) complex64 8B  
l: (0+0i) complex128 16B 
*/

```





### 进制的相互转换

```go
package main

import "fmt"

func main() {

	//十进制的数
	a := 10

	fmt.Printf("a的十进制表示a: %d\n", a)
	fmt.Printf("a的二进制表示a: %b\n", a)
	fmt.Printf("a的八进制表示a: %o\n", a)
	fmt.Printf("a的十六进制表示a: %x\n", a)

	//定义一个八进制的数 以0开头
	b := 077

	fmt.Printf("b的十进制表示a: %d\n", b)
	fmt.Printf("b的二进制表示a: %b\n", b)
	fmt.Printf("b的八进制表示a: %o\n", b)
	fmt.Printf("b的十六进制表示a: %x\n", b)

	//定义一个十六进制的数 以0x开头
	c := 0x111

	fmt.Printf("c的十进制表示a: %d\n", c)
	fmt.Printf("c的二进制表示a: %b\n", c)
	fmt.Printf("c的八进制表示a: %o\n", c)
	fmt.Printf("c的十六进制表示a: %x\n", c)

}

```





### 浮点类型

| 类型       | 描述                  |
| ---------- | --------------------- |
| float32    | IEEE-754 32位浮点型数 |
| float64    | IEEE-754 64位浮点型数 |
| complex64  | 32位实数和虚数        |
| complex128 | 64位实数和虚数        |



### 其他数字类型

| 类型 | 描述      |
| ---- | --------- |
| byte | 类似uint8 |
| rune | 类似int32 |





### 字符串类型

在go语言中，字符串使用双引号""或者反引号来创建。双引号用来创建可解析的字符串，支持转义，但不能用来引用多行；反引号用来创建原生的字符串，可由多行组成，但不支持转义，并且可以包含除了反引号外其他所有字符。双引号创建可解析的字符串应用最广泛，反引号用来创建原生的字符串多用于书写多行消息，HTML以及正则表达式。



#### 字符串的拼接

```go
package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {

	//单行字符串 支持转义
	a := "hello world"

	//多行字符串 不支持转义
	b := `
			<div>
			<p>hello</p>
			</div>
		`
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)

	//字符串的拼接
	s1 := "hello"
	s2 := "world"
	//1.加号拼接
	res1 := s1 + " " + s2
	fmt.Printf("s: %v\n", res1)

	//2.字符串格式化 Sprintf
	res2 := fmt.Sprintf("%s %s", s1, s2)
	fmt.Printf("res2: %v\n", res2)

	//3.strings.join()
	/*
		join会根据字符串数组得到内容，计算出一个拼接之后的字符串的长度，然后申请对应大小的内存，一个一个字符填入，
		在已有一个数组的情况下，这种效率会很高，但是本来没有，去构造这个数组的代价也很高
	*/
	res3 := strings.Join([]string{s1, s2}, " ") //第一个参数是字符串数组，第二个参数是连接符
	fmt.Printf("res3: %v\n", res3)

	//4.buffer.WriteString()
	/*
		这个比较理想，可以当成可变字符串使用，对内存的增长也有优化，如果能预估字符串的长度，还可以使用buffer.Grow()接口来设置capacity
		同时这是直接写到缓冲区，因此效率比较高
	*/
	var buffer bytes.Buffer
	buffer.WriteString(s1)
	buffer.WriteString(s2)
	fmt.Printf("buffer.String(): %v\n", buffer.String())


}

```



#### 转义字符

go语言的字符串常见的转义字符包含回车，换行，单双引号，制表符等

| 转义符 | 含义                          |
| ------ | ----------------------------- |
| \r     | 回车符号                      |
| \n     | 换行符                        |
| \t     | 制表符（tab键，或者四个空格） |
| \\'    | 单引号                        |
| \\"    | 双引号                        |
| \\\    | 反斜杠                        |





#### 切片操作

```go
//切片操作
	s := "I am a student."
	m, n := 2, 4

	fmt.Printf("s[m:n]: %v\n", s[m:n]) //获取字符串s索引位置从m到n-1的值
	fmt.Printf("s[:n]: %v\n", s[:n])   //获取字符串s索引位置从0到n-1的值
	fmt.Printf("s[m:]: %v\n", s[m:])   //获取字符串s索引位置从m到len(s)-1的值
	fmt.Printf("s[:]: %v\n", s[:])     //获取字符串s
	fmt.Printf("s[m]: %v\n", s[m])     //获取字符串s索引位置m的字符的ascii值
```



#### 字符串的一些常用方法

| 方法                                                    | 描述                                                         |
| ------------------------------------------------------- | ------------------------------------------------------------ |
| len(s)                                                  | 获取字符串s的长度                                            |
| +或者fmt.Sprintf                                        | 拼接字符串                                                   |
| strings.Split(s,seq)                                    | 用分隔符seq分割字符s                                         |
| strings.Contains(s,subs)                                | 查询字符串s中是否包含子串subs                                |
| strings.HasPrefix(s,prefix)/strings.HasSuffix(s,suffix) | 判断前/后缀是否是指定的字符串                                |
| strings.Index(s,subs)/strings.LastIndex(s,subs)         | 查询子串subs在s中第一次(最后一次)出现的索引位置，若没有则返回-1 |
| strings.join(s_arr,seq)                                 | 将字符串数组用seq拼接成一个字符串                            |

```go
//分割字符串
	fmt.Printf("strings.Split(s, \" \"): %v\n", strings.Split(s, " "))

	//查询某个字符串是否包含指定的字符串
	fmt.Printf("strings.Contains(s, \"student\"): %v\n", strings.Contains(s, "student"))

	//判断前缀是否是指定的字符串
	fmt.Printf("strings.HasPrefix(s, \"hello\"): %v\n", strings.HasPrefix(s, "hello"))

	//判断后缀是否是指定的字符串
	fmt.Printf("strings.HasSuffix(s, \"student.\"): %v\n", strings.HasSuffix(s, "student."))

	//查找指定字符串第一次出现的索引位置
	fmt.Printf("strings.Index(s, \"a\"): %v\n", strings.Index(s, "aaa"))

	//查找指定字符串最后一次出现的索引位置
	fmt.Printf("strings.LastIndex(s, \"a\"): %v\n", strings.LastIndex(s, "a"))

	//拼接字符串数组
	fmt.Printf("strings.Join([]string{\"i\", \"am\", \"a\", \"student.\"}, \" \"): %v\n", strings.Join([]string{"i", "am", "a", "student."}, " "))
```

**运行结果**

```
strings.Split(s, " "): [I am a student.]
strings.Contains(s, "student"): true
strings.HasPrefix(s, "hello"): false
strings.HasSuffix(s, "student."): true
strings.Index(s, "a"): -1
strings.LastIndex(s, "a"): 5
strings.Join([]string{"i", "am", "a", "student."}, " "): i am a student.
```



### 格式化输出



#### 普通占位符

| 占位符 | 说明                     | 举例                                                         |
| ------ | ------------------------ | ------------------------------------------------------------ |
| %v     | 相应值的默认格式。       | fmt.Printf("a: %v\n", 100)   ----> a:100                     |
| %#v    | 会打印出数据的完整格式。 | tom := studnet{"Tom"}fmt.Printf("tom: %#v\n", tom) ---->tom: main.studnet{Name:"Tom"} |
| %T     | 相应值的类型。           | fmt.Printf("a: %T\n", 100)   ----> a:int                     |
| %%     | 输出%。                  | fmt.Printf("%%") ---->%                                      |

```go
//相应数据的默认格式
a := 10
fmt.Printf("a: %#v\n", a)

tom := studnet{"Tom"}
//数据的完整格式
fmt.Printf("tom: %#v\n", tom)

//%
fmt.Printf("%%\n")
```

**运行结果**

```
a: 10
tom: main.studnet{Name:"Tom"}
%
```



#### 布尔占位符

| 占位符 | 说明              | 举例                                      |
| ------ | ----------------- | ----------------------------------------- |
| %t     | 输出true或false。 | fmt.Printf("flag: %t\n", true)  ---->true |

```go
//布尔占位符
flag := true
fmt.Printf("flag: %t\n", flag)
```

**运行结果**

```
flag: true
```



#### 整数占位符

| 占位符 | 说明                                       | 举例                               |
| ------ | ------------------------------------------ | ---------------------------------- |
| %b     | 输出二进制表示                             | fmt.Printf("%b\n", 10)  ---->1010  |
| %c     | 相应的Unicode码值所表示的码符              | fmt.Printf("%c\n",97) ---->a       |
| %d     | 输出十进制表示                             | fmt.Printf("%d\n", 10)  ---->10    |
| %o     | 输出八进制表示                             | fmt.Printf("%o\n", 10)  ---->12    |
| %x     | 十六进制表示（小写字母）                   | fmt.Printf("%x\n", 10) ---->a      |
| %X     | 十六进制表示（大写字母）                   | fmt.Printf("%X\n", 10) ---->A      |
| %q     | 单引号围绕的字符字面值，由go语法安全的转义 | fmt.Printf("%q\n", 10) ---->'\n'   |
| %U     | Unicode格式：U+1234等同于“U+%04X”          | fmt.Printf("%U\n", 10) ---->U+000A |

```go
//二进制输出
fmt.Printf("%b\n", 10)

//相应的Unicode码值对应的码符
fmt.Printf("%c\n", 97)

//八进制
fmt.Printf("%o\n", 10)

//十六进制 小写字母
fmt.Printf("%x\n", 10)

//十六进制 大写字母
fmt.Printf("%X\n", 10)

fmt.Printf("%q\n", 10)
fmt.Printf("%U\n", 10)
```

**运行结果**

```
1010
a
12
a
A
'\n'
U+000A
```



#### 浮点数和复数占位符

| 占位符 | 说明                                                         | 举例                                                         |
| ------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| %b     | 无小数部分的，指数为二的幂的科学计数法，与strconv.FormatFloat的'b'转换格式一样 | fmt.Printf("%b\n", 10.1) ---->5685794529555251p-49           |
| %e     | 科学计数法（小写字母）                                       | fmt.Printf("%e\n", 10.1) ---->1.010000e+01                   |
| %E     | 科学计数法（大写字母）                                       | fmt.Printf("%E\n", 10.1) ---->1.010000E+01                   |
| %f     | 有小数点而无指数                                             | fmt.Printf("%5.2f\n", 10.1) ---->10.10（总长度为5，小数点两位） |
| %g     | 根据情况选择%e或%f以产生更紧凑的输出（无末尾的0）            | fmt.Printf("%g\n", 10.10000) ---->10.1                       |
| %G     | 根据情况选择%E或%f以产生更紧凑的输出（无末尾的0）            | fmt.Printf("%G\n", 10.000+2.0600i) ---->(10+2.06i)           |

```go
//浮点数和复数
fmt.Printf("%b\n", 10.1)
fmt.Printf("%e\n", 10.1)
fmt.Printf("%E\n", 10.1)
fmt.Printf("2%5.2f2\n", 10.1)
fmt.Printf("%g\n", 10.10000)
fmt.Printf("%G\n", 10.000+2.0600i)
```

**运行结果**

```
5685794529555251p-49
1.010000e+01
1.010000E+01
210.102
10.1
(10+2.06i)
```



#### 字符串与字节切片占位符

| 占位符 | 说明                                     | 举例                                                         |
| ------ | ---------------------------------------- | ------------------------------------------------------------ |
| %s     | 输出字符串表示（string类型或[]byte类型） | fmt.Printf("[]byte{\"hello\", \"world\"}: %s\n", []byte("hello")) ---->[]byte{"hello", "world"}: hello |
| %q     | 双引号围绕的字符串，由go语言安全地转义   | fmt.Printf("\"hello\": %q\n", "hello") ---->"hello"          |
| %x     | 十六进制，小写字母，每字节两个字符       | fmt.Printf("\"hello\": %x\n", "hello") ---->"hello": 68656c6c6f |
| %X     | 十六进制，大写字母，每字节两个字符       | fmt.Printf("\"hello\": %X\n", "hello") ---->"hello": 68656C6C6F |

```go
//字符串
fmt.Printf("[]byte{\"hello\", \"world\"}: %s\n", []byte("hello"))
fmt.Printf("\"hello\": %q\n", "hello")
fmt.Printf("\"hello\": %x\n", "hello")
fmt.Printf("\"hello\": %X\n", "hello")
```

**运行结果**

```
[]byte{"hello", "world"}: hello
"hello": "hello"
"hello": 68656c6c6f
"hello": 68656C6C6F
```



#### 指针占位符

| 占位符 | 说明                               | 举例                                                         |
| ------ | ---------------------------------- | ------------------------------------------------------------ |
| %p     | 十六进制表示，输出指针所指向的地址 | h := 100   p_h := &h     fmt.Printf("%p\n", &h)     fmt.Printf("%p\n", p_h)  ---->两个输出是一样的：0xc0000ba018，都是输出的h的地址 |

```go
h := 100
p_h := &h
fmt.Printf("h: %p\n", &h)
fmt.Printf("p_h: %p\n", p_h)
```

**运行结果**

```
h: 0xc0000140f8
p_h: 0xc0000140f8
```



### golang运算符

go语言内置的运算符有：

- 算术运算符
- 关系运算符
- 逻辑运算符
- 位运算符
- 赋值运算符



#### 算术运算符

| 运算符 | 描述         |
| ------ | ------------ |
| +      | 相加         |
| -      | 相减         |
| *      | 相乘         |
| /      | 相除         |
| %      | 取模（求余） |

> 注意：++（自增）和--（自减）在go语言中是单独的语句，并不是运算符。

```go
a := 100
	b := 200
	//算术运算符
	res := a + b
	fmt.Printf("res: %v\n", res)

	res = a - b
	fmt.Printf("res: %v\n", res)

	res = a * b
	fmt.Printf("res: %v\n", res)

	res = a / b
	fmt.Printf("res: %v\n", res)

	res = a % b
	fmt.Printf("res: %v\n", res)
	a++
```

**运行结果**

```
res: 300
res: -100
res: 20000
res: 0
res: 100
```



#### 关系运算符

| 运算符 | 描述                                              |
| ------ | ------------------------------------------------- |
| ==     | 检查两个值是否相等，返回true或false               |
| !=     | 检查两个值是否不相等，返回true或false             |
| >      | 检查左边的值是否大于右边的值，返回true或false     |
| >=     | 检查左边的值是否大于等于右边的值，返回true或false |
| <      | 检查左边的值是否小于右边的值，返回true或false     |
| <=     | 检查左边的值是否小于等于右边的值，返回true或false |

```go
// 关系运算符
	var res2 bool

	res2 = a == b
	fmt.Printf("res2: %v\n", res2)

	res2 = a > b
	fmt.Printf("res2: %v\n", res2)

	res2 = a >= b
	fmt.Printf("res2: %v\n", res2)

	res2 = a < b
	fmt.Printf("res2: %v\n", res2)

	res2 = a <= b
	fmt.Printf("res2: %v\n", res2)

	res2 = a != b
	fmt.Printf("res2: %v\n", res2)
```

**运行结果**

```
res2: false
res2: false
res2: false
res2: true
res2: true
res2: true
```



#### 逻辑运算符

| 运算符 | 描述                                                         |
| ------ | ------------------------------------------------------------ |
| &&     | 逻辑and运算符。只有两边的操作数都为true才返回true，否则返回false。 |
| \|\|   | 逻辑or运算符。只要有一个操作数为true就返回true，否则返回false。 |
| !      | 逻辑not运算符。取相反的结果。                                |

```go
//逻辑运算
	var res3 bool
	res3 = true && true
	fmt.Printf("res3: %v\n", res3)
	res3 = true || false
	fmt.Printf("res3: %v\n", res3)
	res3 = !res3
	fmt.Printf("res3: %v\n", res3)
```

```
res3: true
res3: true
res3: false
```



#### 位运算符

位运算是对整数在内存中的二进制位进行操作。

| 运算符 | 描述                                                         |
| ------ | ------------------------------------------------------------ |
| &      | 参与运算的两个数各对应的二进位相与。（两位均为1才为1）       |
| \|     | 参与运算的两个数各对应的二进位相或。（两位只要有一个为1就为1） |
| ^      | 参与运算的两个数各对应的二进位异或。（当相对应的二进位相异时，结果才为1，否则为0） |
| <<     | 左移n位就是乘以2的n次方。a<<b是把a的各二进位全部左移b位，高位丢弃，低位补零。 |
| <<     | 右移n位就是除以2的n次方。a>>b是把a的各二进位全部右移b位。    |

```go
//位运算
	aa := 10
	bb := 11
	fmt.Printf("aa: %b\n", aa)
	fmt.Printf("bb: %b\n", bb)

	res4 := aa & bb
	fmt.Printf("res4:%b  %v\n", res4, res4)
	res4 = aa | bb
	fmt.Printf("res4:%b  %v\n", res4, res4)
	res4 = aa ^ bb
	fmt.Printf("res4:%b  %v\n", res4, res4)

	res4 = aa << 2
	fmt.Printf("res4:%b  %v\n", res4, res4)
	res4 = aa >> 1
	fmt.Printf("res4:%b  %v\n", res4, res4)
```

**运行结果**

```
aa: 1010
bb: 1011
res4:1010  10
res4:1011  11
res4:1  1
res4:101000  40
res4:101  5
```



#### 赋值运算符

| 运算符 | 描述                                                 |
| ------ | ---------------------------------------------------- |
| =      | 简单的赋值运算符，将等号右边的结果赋值给左边的变量。 |
| +=     | 相加后再赋值                                         |
| -=     | 相减后再赋值                                         |
| *=     | 相乘后再赋值                                         |
| /=     | 相除后再赋值                                         |
| %=     | 取模后再赋值                                         |
| <<=    | 左移后再赋值                                         |
| >>=    | 右移后再赋值                                         |
| &=     | 按位与后赋值                                         |
| \|=    | 按位或后赋值                                         |
| ^=     | 异或后赋值                                           |





## 流程控制



### go语言中的条件

**条件语句**是用来判断给定的条件是否满足，并根据判断的结果决定执行的语句，go语言中的条件语句也是这样的。



### go语言中的条件语句

1. if语句：if语句由一个布尔表达式后紧跟一个或多个语句组成。
2. if...else语句：if语句后可以使用可选的else语句，else语句中的表达式在布尔表达式为false时执行。
3. if嵌套语句
4. switch语句：switch语句用于基于不同条件执行不同的动作。
5. select语句：select语句类似于switch语句，但是select会随机执行一个可运行的case，如果没有case可运行，它将阻塞，直到有case可运行。



#### if语句

**语法**

```go
if 布尔表达式{
    /*布尔表达式为true时要执行的语句*/
}
```

> 注意：在go语言中，布尔表达式不需要使用括号。

```go
//示例
a := 100
	if a > 10 {
		fmt.Println("hello")
	}
```

**运行结果**

```
hello
```



> **初始变量可以声明在布尔表达式里面，但是这个变量的作用于只能用在当前条件语句中！**
>
> ```go
> //示例
> 	if b := 100; b > 10 {
> 		fmt.Println("if条件成立！")
> 		fmt.Printf("b: %v\n", b)
> 	} else {
> 		fmt.Println("if条件不成立！")
> 		fmt.Printf("b: %v\n", b)
> 	}
> ```
>
> **运行结果**
>
> ```
> if条件成立！
> b: 100
> ```
>
> ```go
> if b := 100; b > 10 {
> 		fmt.Println("if条件成立！")
> 		fmt.Printf("b: %v\n", b)
> 	} else {
> 		fmt.Println("if条件不成立！")
> 		fmt.Printf("b: %v\n", b)
> 	}
> 	fmt.Printf("b: %v\n", b)
> ```
>
> **运行结果**
>
> ```
> 19:24: undefined: b
> ```
>
> 



golang中if语句使用提示：

1. 不能使用布尔类型以外的其他值作为判断条件
2. 不能使用括号将条件语句括起来
3. 大括号必须存在，即使只有一行语句
4. 左括号必须于if，else同一行
5. 在if之后，条件语句之前，可以添加变量初始化语句，使用分号;分割



#### if else语句

**语法**

```go
if 布尔表达式{
    /*布尔表达式为true时要执行的语句*/
}else{
    /*布尔表达式为false时要执行的语句*/
}
```

```go
//示例
if age := 18; age > 18 {
		fmt.Println("你已经成年了！")
	} else {
		fmt.Println("你还未成年！")
	}
```

**运行结果**

```
你还未成年！
```



#### if else if语句

**语法**

```go
if 布尔表达式1{
    /*布尔表达式1为true时要执行的语句*/
}else if 布尔表达式2{
    /*布尔表达式2为true时要执行的语句*/
}else{
    /*以上所有布尔表达式都不成立时要执行的语句*/
}
```

```go
//示例
var aa int
	fmt.Println("请输入一个数：")
	fmt.Scan(&aa)
	if aa == 100 {
		fmt.Printf("aa: %v\n", aa)
	} else if aa < 100 {
		fmt.Println("太小了！")
	} else {
		fmt.Println("太大了！")
	}

```

```
请输入一个数：
10
太小了！
```



#### switch语句

##### 单值匹配

**语法**

```go
//单值匹配
switch 条件表达式{
    case var1:
    	...
    case var2:
    	...
    case var3:
    	...
    default:
    	...
}

```



```go
//示例
	aaa := 1
	switch aaa {
	case 1:
		fmt.Println("唱歌！")
	case 2:
		fmt.Println("跳舞！")
	case 3:
		fmt.Println("唱歌+跳舞！")
	default:
		fmt.Println("才艺表演。")
	}
```

**运行结果**

```
唱歌！
```



##### 多值匹配

**语法**

```go
//多值匹配
switch 条件表达式{
    case var1,var2,var3...:
    	...
    case var4,var5...:
    	...
    default:
    	...
}
```

```go
//示例
var day int
	fmt.Println("请输入一个数字：")
	fmt.Scan(&day)
	switch day {
	case 1, 2, 3, 4, 5:
		fmt.Println("工作日！")
	case 6, 7:
		fmt.Println("假期！")
	default:
		fmt.Println("输入有误！")
	}
```

**运行结果**

```
请输入一个数字：
6
假期！
```



##### case后接条件表达式

**语法**

```go
//case后接条件表达式
switch{
    case 条件表达式1:
    	...
    case 条件表达式2:
    	...
    case 条件表达式3:
    	...
    default:
    	...
}
```

```go
//示例
	var n int
	fmt.Println("请输入一个数字：")
	fmt.Scan(&n)
	switch {
	case n == 10:
		fmt.Println("猜对了！")
	case n < 10:
		fmt.Println("太小了！")
	case n > 10:
		fmt.Println("太大了！")
	default:
		fmt.Println("不可能执行到这条语句，除非上面的case都不成立。")
	}
```

**运行结果**

```
请输入一个数字：
23
太大了！
```



##### fallthrough

**语法**

```go

//fallthrough 可以执行满足条件的下一个case
switch 条件表达式{
    case var1:
    	...
    	fallthrough
    case var2:
    	...
    	fallthrough
    case var3:
    	...
    default:
    	...
}

```

```go
bb := 100
	switch bb {
	case 100:
		fmt.Printf("bb: %v\n", bb)
		fallthrough
	case 200:
		fmt.Println("200")
		fallthrough
	case 300:
		fmt.Println("300")
	default:
		fmt.Println("case都不成立就执行这里。")
	}

```

**运行结果**

```
bb: 100
200
300
```

golang中switch使用注意事项：

1. 支持多条件匹配
2. 不同的case之间不使用break分隔，默认只会执行一个case
3. 如果想执行多个case，需要使用fallthrough关键字，也可以用break终止
4. case后还可以使用条件表达式



### go语言中的循环语句

go语言中的循环只有for循环，去除了while，do while循环，使用起来更加简洁。

1. for循环。
2. for range循环。（类似于python中的for in）



#### for语句

**语法**

```go
for 初始语句;条件表达式；结束语句{
    循环语句
}
```

```go
//for循环
	for i := 0; i < 10; i++ {
		fmt.Printf("i: %v\t", i)
	}
	fmt.Println("\n")
	//初始条件写在外面
	j := 0
	for ; j < 10; j++ {
		fmt.Printf("j: %v\t", j)
	}

	fmt.Println("\n")
	//结束语句写在循环体内
	for i := 0; i < 10; {
		fmt.Printf("i: %v\t", i)
		i++

	}
	fmt.Println("\n")

	//永真循环
	k := 0
	for {
		fmt.Printf("k: %v ", k)
		k++
		if k == 10 {
			break
		}
	}

```

**运行结果**

```
i: 0    i: 1    i: 2    i: 3    i: 4    i: 5    i: 6    i: 7    i: 8    i: 9

j: 0    j: 1    j: 2    j: 3    j: 4    j: 5    j: 6    j: 7    j: 8    j: 9

i: 0    i: 1    i: 2    i: 3    i: 4    i: 5    i: 6    i: 7    i: 8    i: 9

k: 0 k: 1 k: 2 k: 3 k: 4 k: 5 k: 6 k: 7 k: 8 k: 9 
```



#### for range语句

go语言中可以使用for range遍历数组，切片，字符串，map及通道（channel）。通过for range遍历的返回值有一下规律：

1. 数组，切片，字符串返回索引和值
2. map返回键和值
3. 通道只返回通道内的值

```go
//for range
	//遍历数组
	var arr = [...]int{1, 2, 3, 4, 5, 6, 7}
	for i, v := range arr {
		fmt.Printf("%v:%v\n", i, v)

	}

	//遍历map
	var m = make(map[string]string, 0)
	m["name"] = "Tom"
	m["age"] = "18"
	m["num"] = "1234"
	for k, v := range m {
		fmt.Printf("%v:%v\n", k, v)

	}
```

**运行结果**

```
name:Tom
age:18
num:1234
```



### break关键字

break关键字可以结束for，switch和select的代码块。

```go
//跳出for循环
k := 0
	for {
		fmt.Printf("k: %v ", k)
		k++
		if k == 10 {
			break
		}
	}
```

**运行结果**

```
k: 0 k: 1 k: 2 k: 3 k: 4 k: 5 k: 6 k: 7 k: 8 k: 9 
```



```go
//配合标签使用
label:
	for i := 0; i < 10; i++ {
		fmt.Printf("i: %v\n", i)
		if i == 5 {
			break label
		}

	}
	fmt.Println("END...")
```

**运行结果**

```
i: 0
i: 1
i: 2
i: 3
i: 4
i: 5
END...
```

go语言中break的注意事项：

1. 单独在select中使用break和不使用break没有啥区别。
2. 单独在switch中使用break，并且没有使用fallthrough，和不使用break没有啥区别。
3. 在switch中，配合fallthrough关键字，能够终止fallthrough后面case语句的执行。
4. 带标签的break，可以跳出多层select/select作用域。让break更加灵活，写法更加简单，不需要使用控制变量一层一层跳出循环，没有带break的只能跳出当前语句。



### continue关键字

continue只能用在循环中，在go中只能用在for循环中，它可以终止本次循环，进入下一轮循环。

continue也可以配合标签使用。

```go
//打印10以内的偶数
	for i := 0; i < 11; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Printf("i: %v\n", i)
	}
```

**运行结果**

```
i: 0
i: 2
i: 4
i: 6
i: 8
i: 10
```



**配合标签使用**

```go
label:

	for i := 0; i < 3; i++ {

		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				continue label //跳到label所在的地方
			}
			fmt.Printf("i: %v  j: %v\n", i, j)
		}

	}
	fmt.Println("i=1,j=1时跳出了循环。")
```

**运行结果**

```
i: 0  j: 0
i: 0  j: 1
i: 0  j: 2
i: 1  j: 0
i: 2  j: 0
i: 2  j: 1
i: 2  j: 2
i=1,j=1时跳出了循环。
```



### goto关键字

goto语句通过标签进行代码间的无条件跳转。goto语句可以在快速跳出循环，避免重复退出上有一定帮助。

#### 跳到指定标签

```go
for i := 0; i < 10; i++ {
		fmt.Printf("i: %v\n", i)
		if i == 5 {
			goto label
		}

	}

label:
	fmt.Println("循环结束！")
```

**运行结果**

```
i: 0
i: 1
i: 2
i: 3
i: 4
i: 5
循环结束！
```



#### 跳出多重循环

```go
//与break不一样的是，break只能跳出一层循环，而goto可以直接跳出所有循环
for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				fmt.Printf("i=%v j=%v k=%v\n", i, j, k)
				if i == 1 && j == 1 && k == 1 {
					goto mylabel
				}

			}

		}

	}
mylabel:
	fmt.Println("循环结束！")
```

**运行结果**

```
i=0 j=0 k=0
i=0 j=0 k=1
i=0 j=0 k=2
i=0 j=1 k=0
i=0 j=1 k=1
i=0 j=1 k=2
i=0 j=2 k=0
i=0 j=2 k=1
i=0 j=2 k=2
i=1 j=0 k=0
i=1 j=0 k=1
i=1 j=0 k=2
i=1 j=1 k=0
i=1 j=1 k=1
循环结束！
```





## 复杂数据类型



### 数组

数组是**相同数据类型**的一组数据的集合，数组一旦定义长度不能修改，数组	可以通过索引来访问元素。

#### **数组的定义**

```go
var array_name [SIZE]TYPE
```

- array_name：数组名
- SIZE：数组的大小
- TYPE：数组中数据的类型

```go
type Student struct {
	Name string
	num  int
}
// 数组的定义
var arr_int [10]int
fmt.Printf("arr_int: %v\n", arr_int)
var arr_str [10]Student
fmt.Printf("arr_str: %v\n", arr_str)
```

**运行结果**

```
arr_int: [0 0 0 0 0 0 0 0 0 0]
arr_str: [{ 0} { 0} { 0} { 0} { 0} { 0} { 0} { 0} { 0} { 0}]
```



#### 数组的初始化

初始化就是给数组的元素赋初值，没有初始化的数组，默认元素都是**零值**（数值型的默认值是0，布尔型的默认值所示false，字符串型的默认值是空字符）。

```go
//给数组赋初始值
	var arr_int_init = [10]int{1, 2, 3}
	fmt.Printf("arr_int_init: %v\n", arr_int_init)

	//给指定位置赋初始值
	var arr_float_init = [10]float64{0: 100.0, 3: 200.0}
	fmt.Printf("arr_float_init: %v\n", arr_float_init)

	//使用...不指定数组的大小,根据初始值来判断数组的大小
	var arr_string_init = [...]string{"hello", "world"}
	fmt.Printf("arr_string_init: %v\n", arr_string_init)
```

**运行结果**

```
arr_int_init: [1 2 3 0 0 0 0 0 0 0]
arr_float_init: [100 0 0 200 0 0 0 0 0 0]
arr_string_init: [hello world]
```



#### 数组的访问

可以通过索引的方式来访问数组。数组的最大下标为数组的长度减一，最小为0，大于这个值会发生数组越界。

```go
package main

import "fmt"

func main() {

	var arr_int = [10]int{1, 2, 3, 4, 5, 6}

	//访问第一个元素
	fmt.Printf("arr_int[0]: %v\n", arr_int[0])

	//访问第3个元素
	fmt.Printf("arr_int[2]: %v\n", arr_int[2])

	//访问最后一个元素
	fmt.Printf("arr_int[len(arr_int)-1]: %v\n", arr_int[len(arr_int)-1])

	//for遍历数组
	for i := 0; i < len(arr_int); i++ {

		fmt.Printf("a[%v]=%v \n", i, arr_int[i])

	}
	print("#########################################\n")
	//for range遍历数组
	for i, v := range arr_int {
		fmt.Printf("a[%v]=%v \n", i, v)

	}
}

```

**运行结果**

```
arr_int[0]: 1
arr_int[2]: 3
arr_int[len(arr_int)-1]: 0
a[0]=1 
a[1]=2 
a[2]=3 
a[3]=4 
a[4]=5 
a[5]=6 
a[6]=0 
a[7]=0 
a[8]=0 
a[9]=0 
#########################################
a[0]=1 
a[1]=2 
a[2]=3 
a[3]=4 
a[4]=5 
a[5]=6 
a[6]=0 
a[7]=0 
a[8]=0 
a[9]=0 
```



### 切片

和数组类似，切片也是一组相同数据类型的数据的集合，但与数组不一样的是，切片的长度是可变的。对于数组来说，当我们对要保存的元素的个数不确定时，如果申请太小的数组，可能就不够用；如果申请太大的数组，可能就造成了不必要的浪费。鉴于这个原因，就有了切片，我们可以把切片理解为可变长度的数组，其实它底层就是使用数组实现的，只不过增加了一个自动扩容功能。



#### 切片的定义



语法**1**

```go
var slice_name []TYPE
```

- slice_name：切片名
- TYPE：切片类型



语法**2**

```go
//使用make函数定义切片时，会同时将切片初始化
slice_name := make([]TYPE,SIZE)
```

- slice_name：切片名
- TYPE：切片中元素的类型
- SIZE：初始化切片的大小

```go
package main

import "fmt"

func main() {

	//切片的定义
	var s1 []int
	s1 = append(s1, 1)
	fmt.Printf("s1: %v\n", s1)

	//make定义切片的同时会将其初始化
	s2 := make([]string, 10)
	fmt.Printf("s2: %v\n", s2)

	//访问切片中的元素
	fmt.Printf("s1[0]: %v\n", s1[0])

	//在切片的末尾添加元素
	s1 = append(s1, 2)
	fmt.Printf("s1: %v\n", s1)
	//修改切片中的元素
	s1[0] = 100
	fmt.Printf("s1: %v\n", s1)

	//获取切片的大小
	fmt.Printf("len(s1): %v\n", len(s1))

	fmt.Printf("cap(s1): %v\n", cap(s1))

}

```

**运行结果**

```
s1: [1]
s2: [         ]
s1[0]: 1
s1: [1 2]
s1: [100 2]
len(s1): 2
cap(s1): 2
```



#### 切片的初始化

```go
package main

import "fmt"

func main() {

	//切片的初始化
	//方法一
	var s1 = []int{1, 2, 3}
	fmt.Printf("s1: %v\n", s1)

	//方法二：make
	s2 := make([]int, 10)
	fmt.Printf("s2: %v\n", s2)

	//方法三：借助数组
	arr := [3]int{1, 2, 3}
	s3 := arr[:]
	fmt.Printf("s3: %v\n", s3)

	//切片/数组/字符串的切片操作： s[a:b] 左闭右开 于python不一样的是，go语言中不能修改步长，步长只能是1
	s4 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("s4[1:9]: %v\n", s4[1:3])
    
    //访问操作和遍历操作同数组
}

```

**运行结果**

```
s1: [1 2 3]
s2: [0 0 0 0 0 0 0 0 0 0]
s3: [1 2 3]
s4[1:9]: [2 3]
```



#### 切片的crud操作

```go
package main

import (
	"fmt"
)

func main() {

	s := []int{1, 2, 3, 4, 5}

	//add
	s = append(s, 1)
	fmt.Printf("s: %v\n", s)

	//delete：删除索引为index的元素
	index := 2
	s = append(s[:index], s[index+1:]...)
	fmt.Printf("s: %v\n", s)

	//update
	s[4] = 6
	fmt.Printf("s: %v\n", s)

	//query
	target := 5
	for i, v := range s {
		if v == target {
			println("找到了！", i)
			break
		}

	}

}

```

**运行结果**

```
s: [1 2 3 4 5 1]
s: [1 2 4 5 1]
s: [1 2 4 5 6]
找到了！ 3
```



### map

map是一种key:value键值对的数据结构。map内部实现是hash表。map最重要的一点是通过key能够快速的检索出数据。



#### map的定义

```go
var m[K_TYPE]V_TYPE
```

- m：map名
- K_TYPE：key的类型
- V_TYPE：value的类型



```go
package main

import (
	"fmt"
)

func main() {

	//map的定义
	//方法一
	var m1 map[string]string
	fmt.Printf("m1: %v\n", m1)

	//方法二
	m2 := make(map[string]string)
	fmt.Printf("m2: %v\n", m2)

	//map的初始化
	var m3 = map[string]string{
		"name": "Tom",
		"age":  "18",
	}
	fmt.Printf("m3: %v\n", m3)

	//增加/修改map
	m3["num"] = "1234"
	m3["age"] = "20"
	fmt.Printf("m3: %v\n", m3)

	//根据k获取v
	fmt.Printf("m3[\"name\"]: %v\n", m3["name"])

	//判断某个k是否存在  v,ok=m[k]--->如果k存在，ok为true，否则为false
	v, ok := m3["name"]
	if ok {
		print(v)
	}

	//删除某个k
	delete(m3, "age")
	fmt.Printf("m3: %v\n", m3)
}

```

**运行结果**

```
m1: map[]
m2: map[]
m3: map[age:18 name:TOm]
m3: map[age:20 name:TOm num:1234]
m3["name"]: TOm
TOmm3: map[name:TOm num:1234]
```



#### map的遍历

通过for range对map进行遍历。

```go
package main

import "fmt"

func main() {

	var m = map[string]string{
		"name": "Tom",
		"age":  "18",
		"num":  "1234",
	}

	//for range遍历
	//1.只拿到key
	for k := range m {
		fmt.Printf("k: %v v: %v\n", k, m[k])
	}

	//2.同时拿到k和v
	for k, v := range m {
		fmt.Printf("k: %v v: %v\n", k, v)
	}

}

```

**运行结果**

```
k: name v: Tom
k: age v: 18
k: num v: 1234
k: num v: 1234
k: name v: Tom
k: age v: 18
```



## 函数



### golang函数简介

函数是go语言中的**一级公民**，我们把所有的功能单元都定义在函数中，可以重复使用。函数包含函数的名称，参数列表和返回值类型，这些构成了函数的签名。



### golang中函数的特性

- go语言中有3种函数：普通函数，匿名函数，方法（定义在struct上的函数）。
- go语言中不允许函数重载（overload），也就是说不允许函数同名。
- **go语言中的函数不能嵌套函数**，但可以嵌套匿名函数。
- 函数是一个值，可以将函数赋值给变量，使得这个变量也成为函数。
- 函数可以作为参数传递给另一个函数。
- 函数的返回值可以是一个函数。
- 函数调用的时候，如果有参数传递给函数，则先拷贝参数的副本，再将副本传递给函数。（值传递）
- 函数参数可以没有名称。



### golang中函数的定义和调用

函数在使用之前必须先定义，可以调用函数来完成某个任务。函数可以重复使用，从而达到代码重用。

**语法**

```go
func function_name([parameter list])(return_types){
    //函数体
}
```

- func：声明函数
- function_name：函数名
- [parameter list]：参数列表，参数就像一个占位符，当函数被调用时，你可以将值传递给函数，这个值被称为实际参数。参数列表指定的是参数类型，顺序及参数个数。参数是可选的，也就是说函数也可以不包含参数。
- return_types：函数返回值的类型。
- 函数体：函数的逻辑代码部分



```go
//定义一个两数之和
func sum(a int,b int)(res int){
    res=a+b
    return res
}
```



### golang函数的返回值

函数可以有0个或多个返回值，返回值需要指定数据类型，返回值通过return关键字来指定。

return可以有参数，也可以没有参数，这些返回值可以有名称，也可以没有名称。go中的函数可以有多个返回值。

- return关键字中指定了参数时，返回值可以不用名称。如果return省略参数，则返回值部分必须带名称。
- 当返回值有名称时，必须使用括号包围，逗号分隔，即使只有一个返回值。
- 但即使返回值命名了，return中也可以强制指定其他返回值名称，也就是说return的优先级更高。
- 命名的返回值是预先声明好的，在函数内部可以直接使用，无需再次声明。命名返回值的名称不能和函数参数名称相同，否则报错提示变量重复声明。
- return中可以有表达式，但不能出现赋值表达式，这其他语言可能有所不同。例如可以：return a+b，但是不能：return a=b+c



```go
package main

import "fmt"

//没有返回值
func foo1() {
	print("没有参数，也没有返回值！\n")
}

//有参数，没有返回值
func foo2(name string) {
	print("有一个参数：")
	fmt.Printf("name: %v\n", name)
}

//有参数，有返回值 且给返回值命名了
func foo3(name string, age int) (name2 string, age2 int) {
	name2 = name
	age2 = age
	return name2, age2
	// return    如果return不指定返回值，则默认返回上面定义的返回值
}

// 有参数，有返回值，但没有给返回值命名 这个时候就必须需要return来指定返回值
func foo4(name string, age int) (string, int) {
	return name, age

}

func main() {

	foo1()
	foo2("Tom")
	name2, age2 := foo3("Tom", 18)
	fmt.Printf("name2: %v\n", name2)
	fmt.Printf("age2: %v\n", age2)
	name, _ := foo4("Tom", 18)  //丢弃age
	fmt.Printf("name: %v\n", name)

}

```

**运行结果**

```
没有参数，也没有返回值！
有一个参数：name: Tom
name2: Tom
age2: 18
name: Tom
```

> tips：
>
> - go中经常会使用其中一个返回值作为函数是否执行成功，是否有错误信息的判断条件。例如return value,exists；return value,ok（map中就是这样使用的）；return value,err等
> - 当函数的返回值过多时，例如有4个以上的返回值，应该将这些返回值收集到容器中，然后以返回容器的方式去返回。
> - 当函数有多个返回值时，如果其中某个或几个返回值不想使用，可以通过下划线_来丢弃这些返回值。





### golang函数的参数

go语言函数可以有0个或多个参数，参数需要指定数据类型。

声明函数时的参数列表叫做形参，调用时传递的参数叫做实参。

go语言是通过**传值的方式传参**的，意味着传递给函数的是拷贝后的副本，所以函数内部访问，修改的也是这个副本。(但是map，slice，interface，channel这些数据类型本身就是指针类型，所以就算是拷贝传值也是拷贝的指针，拷贝后的参数任然指向底层数据结构，所以修改它们可能会影响外部数据结构的值。)

go语言可以使用变长参数，有时候不能确定参数的个数时，可以使用变长参数，可以在函数定义语句的参数部分使用**args...type**的方式。这时会将...代表的参数全部保存到一个名为args的slice中，并且这些参数的数据类型都是type。



```go
package main

import "fmt"

func test(a int) {
	a = 200
	fmt.Printf("里面的a: %v\n", a)
}

//两数之和
func sum(a int, b int) (c int) {
	c = a + b
	return
}

//...接收无数个参数
func foo(name string, age int, args ...string) {
	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)
	for _, v := range args {
		fmt.Printf("v: %v\n", v)
	}
}
func main() {

	a := 100
	test(a)
	fmt.Printf("外面的a=%v\n", a)

	res := sum(1, 2)
	fmt.Printf("res: %v\n", res)
	foo("Tom", 18, "Hello", "My", "name", "is", "Tom")

}

```

**运行结果**

```
里面的a: 200
外面的a=100
res: 3
name: Tom
age: 18
v: Hello
v: My
v: name
v: is
v: Tom
```



### golang函数类型与函数变量

可以使用type关键字来定义一个函数类型，语法如下：

```go
type fun func(par_type1,par_type2...) res_type1...
```

- fun：自己定义的函数类型名
- par_type1，par_type2...：表示各个参数的类型
- res_type1...：表示各个返回值的类型

```go
package main

import "fmt"

func sum(a int, b int) (c int) {
	c = a + b
	return
}

func main() {

	//自己声明函数的类型，然后再将函数赋值给一个变量
	type my_func func(int, int) int
	var my_sum my_func
	my_sum = sum
	res := my_sum(1, 2)
	fmt.Printf("res: %v\n", res)

	//直接通过短变量的形式将一个函数赋值给一个变量
	c := sum
	res2 := c(1, 2)
	fmt.Printf("res2: %v\n", res2)

}

```

**运行结果**

```
res: 3
res2: 3
```



### 高阶函数

go语言的函数，可以作为函数的参数，传递给另外一个函数，同时也可以作为函数的返回值返回。



#### 函数作为参数和返回值

```go
package main

import "fmt"

func sum(a int, b int) (c int) {
	c = a + b
	return
}

func sub(a int, b int) (c int) {
	c = a - b
	return
}

func test(a int, b int, f func(int, int) int) {
	res := f(a, b)
	fmt.Printf("%v+%v=res: %v\n", a, b, res)
}

func cal(flag string) func(int, int) int {
	switch flag {
	case "+":
		return sum
	case "-":
		return sub
	default:
		return sum
	}

}

func main() {

	//函数作为参数传给函数
	test(10, 20, sum)

	//函数作为返回值
	f_sum := cal("+")
	res := f_sum(1, 2)
	fmt.Printf("res: %v\n", res)

	f_sub := cal("-")
	res2 := f_sub(1, 2)
	fmt.Printf("res2: %v\n", res2)
}

```



**运行结果**

```
10+20=res: 30
res: 3
res2: -1
```





### 匿名函数

go语言函数不能嵌套，但是在函数内部可以定义匿名函数，实现一些简单的功能。所谓匿名函数，就是没有名称的函数。

**语法**

```go
func ([参数列表])([返回值列表]){
    //函数体
}
```

```go
package main

import "fmt"

func main() {

	//将匿名函数赋值给一个变量
	say_hello := func(name string, age int) {
		fmt.Printf("My name is  %v,And i',m %v old.\n", name, age)
	}
	say_hello("Tom", 18)

	//直接调用匿名函数
	func(words string) {
		fmt.Printf("words: %v\n", words)
	}("Hi,我是一个匿名函数！")

}

```



**运行结果**

```
My name is  Tom,And i',m 18 old.
words: Hi,我是一个匿名函数！
```

 

### 闭包

闭包就是延伸了函数作用于的函数，使得函数可以访问到函数体外的非全局变量。

```go
package main

import "fmt"

//求平均值的闭包
func make_avger() func(int) float64 {

	s := make([]int, 0) 
    //对于下面的匿名函数来说，切片s就相当于函数体外的非全局变量
    //并且调用make_avager方法后，在全局中是访问不到切片s的
  

	return func(num int) float64 {
		s = append(s, num)
		sum := 0
		for _, v := range s {
			sum += v
		}
		return float64(sum / len(s))
	}

}
func main() {

	avg := make_avger()
	for i := 0; i < 10; i++ {
		fmt.Printf("%.2f\n", avg(i))

	}
}

```



**运行结果**

```
0.00
0.00
1.00
1.00
2.00
2.00
3.00
3.00
4.00
4.00
```



### 递归

函数内部调用函数自己的函数称为递归函数（同闭包一样，递归不是go语言特有的，而是go语言具备实现递归的条件）。

```go
package main

//使用递归实现斐波那契数列
func fib(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}

func main() {
	res := fib(10)
	print(res)
}

```



**运行结果**

```
102334155
```



## defer语句



go语言中的defer语句会将其后面紧跟随的语句延迟处理。在defer归属的函数即将返回时，将延迟处理的语句按照defer定义的顺序逆序执行，也就是说，先defer的语句后执行，后defer的语句先执行。

**defer特性**

- 关键字defer用于注册延迟调用。
- 这些调用直到return前才执行。因此可以用来做资源清理。
- 多个defer语句，按先进后出的执行顺序。
- defer语句中的变量，在defer声明时就决定了。



```go
package main

func main() {
	print("start\n")
	defer print("step 1\n")
	defer print("step 2\n")
	defer print("step 3\n")
	print("end\n")
}

```

**运行结果**

```
start
end
step 3
step 2
step 1
```

> 如运行结果所示，defer后的语句都被**延迟**执行了，并且有多个defer语句存在时，后defer的语句先执行。



## init函数



golang有一个特殊的init函数，先于main函数执行，可以用于实现包级别的一些初始化工作。



**init函数的主要特点**

- init函数先于main函数执行，不能被主动调用，它是**自动执行**的
- init函数即没有参数，也没有返回值
- 每个包可以有多个init函数
- 包的每个源文件也可以有多个init函数，这点比较特殊，就是可以在同一个文件下定义多个init函数
- 同一个包的init执行顺序没有明确的定义，但是同一个文件下的init函数是按顺序执行的
- 不同包的init函数按照包导入的依赖关系决定执行顺序。



**golang初始化顺序**

顺序：变量初始化>init函数>main函数

```go
package main

func init_var() int {
	print("初始化变量...\n")
	return 10
}

func init() {
	print("init函数1...\n")
}

func init() {
	print("init函数2...\n")
}

var i int = init_var()

func main() {
	print("main函数被执行了...\n")

}

```

**运行结果**

```
初始化变量...
init函数1...
init函数2...
main函数被执行了...
```



## 指针



go语言中的函数传参都是值拷贝，当我们想修改某个变量的时候，我们可以创建一个指向该变量地址的指针变量。传递数据使用指针，而无需拷贝数据。

类型指针不能进行偏移和运算。

go语言中的指针操作非常简单，只需要记住两个符号：&（取地址）和*（根据地址取值）。



### 指针地址和指针类型

每个变量在运行时都拥有一个地址，这个地址代表变量在内存中的位置。go语言中使用&字符放在变量前面对变量进行取地址操作。go语言中的值类型（int，float，bool，string，array，struct）都有对应的指针类，如：*int，**string等。



### 指针语法

一个指针变量指向了一个值的内存地址。（也就是我们声明了一个指针之后，可以像变量赋值一样，把一个值的内存地址放入到指针当中。）

**语法**

```go
var var_name *var_type
```

```go
package main

import "fmt"

func main() {

	var p *int
	var pp **int
	var ppp ***int
	a := 10
	p = &a    //将指针p指向a的地址 即*p=a=10
	pp = &p   //将指针pp指向p的地址 即*pp=p-> *(*pp)=*p=a=10
	ppp = &pp //将指针ppp指向pp的地址 即*ppp=pp-> *(*ppp)=*pp -> *(*(*ppp))=*(*pp)=*p=a=10

	fmt.Printf("p: %v\n", p)
	fmt.Printf("p: %v\n", &p)
	fmt.Printf("pp: %v\n", pp)
	fmt.Printf("ppp: %v\n", ***ppp)
	fmt.Println(*p == a)
	fmt.Println(*(*pp) == a)
	fmt.Println(*(*(*ppp)) == a)

}

```

**运行结果**

```
p: 0xc0000ba000
p: 0xc0000b4018
pp: 0xc0000b4018
ppp: 10
true
true
true
```



### 指向数组的指针



```go
package main

import "fmt"

func main() {

	var arr = [5]int{1, 2, 3, 4, 5}
	var arr_p [5]*int //数组类型的指针
	fmt.Printf("arr: %v\n", arr)
	for i := 0; i < len(arr); i++ {
		arr_p[i] = &arr[i]
	}

	for i := 0; i < len(arr_p); i++ {
		fmt.Printf("arr_p[i]=%v *arr_p[i]=%v arr[i]=%v\n", arr_p[i], *arr_p[i], arr[i])

	}

}

```

**运行结果**

```
arr: [1 2 3 4 5]
arr_p[i]=0xc0000b2030 *arr_p[i]=1 arr[i]=1
arr_p[i]=0xc0000b2038 *arr_p[i]=2 arr[i]=2
arr_p[i]=0xc0000b2040 *arr_p[i]=3 arr[i]=3
arr_p[i]=0xc0000b2048 *arr_p[i]=4 arr[i]=4
arr_p[i]=0xc0000b2050 *arr_p[i]=5 arr[i]=5
```



## 结构体



### golang类型定义和类型别名



**类型定义语法**

```go
type my_type old_type
```

- my_type：自己新定义的类型
- old_type：已经存在的类型



**类型别名语法**

```go
type my_type = old_type //用已存在的类型赋值给新类型
```

**两者的区别**

- 类型定义相当于定义了一个全新的类型，与之前的类型不同；但是类型别名并没有定义新的类型，而是使用一个别名来代替之前的类型
- 类型别名只会在代码中存在，在编译完成之后并不会存在该别名
- 因为类型别名和原来的类型是一致的，所以原类型所拥有的方法，类型别名定义的变量也拥有；但是如果是重定义的一个类型，那么不可以调用之前的任何方法

```go
package main

import "fmt"

func main() {

	//类型定义
	type my_string string

	var name my_string = "Tom"
	fmt.Printf("name: %T %v\n", name, name) //类型为自己定义的新类型

	//类型别名
	type my_string2 = string
	var name2 my_string2 = "Tom"
	fmt.Printf("name2: %T %v\n", name2, name2) //类型仍旧是string

}

```

**运行结果**

```
name: main.my_string Tom
name2: string Tom
```



### golang结构体

go语言没有面向对象的概念，但是可以使用结构体来实现面向对象编程的一些特性，你如：继承，组合等特性。



**结构体的定义**

```go
type struct_name struct{
    member1 type1
    member2 type2
    member3 type3
    ...
    
}
```

- type：结构体定义关键字
- struct_name：结构体名
- struct：结构体定义关键字
- member type：成员定义



```go
package main

import "fmt"

func main() {

	//一个结构体就相当于一个新的类型
	type Student struct {
		name  string
		num   int
		age   int
		email string
	}

	tom := Student{"Tom", 1, 18, "110@qq.com"}
	fmt.Printf("tom: %v\n", tom)
	fmt.Printf("tom.name: %v\n", tom.name)
	fmt.Printf("tom.age: %v\n", tom.age)

}

```

**运行结果**

```
tom: {Tom 1 18 110@qq.com}
tom.name: Tom
tom.age: 18
```



**匿名结构体**

对于临时需要使用结构体时，可以定义匿名结构体。

```go
// 匿名结构体
	var dog struct {
		name string
		age  int
	}

	dog.name = "Jerry"
	dog.age = 3
	fmt.Printf("dog: %v\n", dog)
	fmt.Printf("dog.name: %v\n", dog.name)
```

**运行结果**

```
dog: {Jerry 3}
dog.name: Jerry
```

 

### 结构体的初始化

对于通过结构体定义的一个变量来说，未初始化的变量的每个成员属性都是零值；go语言中提供了两种结构体初始化的方法，分别是“k-v”式和“列表”式。

```go
package main

import "fmt"

func main() {
	type Student struct {
		name  string
		age   int
		email string
	}

	tom := Student{}
	fmt.Printf("tom: %v\n", tom) //未初始化的结构体每个成员属性都是零值

	//“k-v”式初始化 可以对部分值进行初始化，也可以对全部值初始化
	jerry := Student{name: "Jerry", age: 18}
	fmt.Printf("jerry: %v\n", jerry)

	//“列表”式初始化 必须将所有值初始化
	autumn := Student{"Autumn", 18, "110@qq.com"}
	fmt.Printf("autumn: %v\n", autumn)
}

```

**运行结果**

```
tom: { 0 }
jerry: {Jerry 18 }
autumn: {Autumn 18 110@qq.com}
```



### **结构体指针**

```go
package main

import "fmt"

func main() {

	type Student struct {
		name string
		age  int
	}

	tom := Student{name: "Tom", age: 18}

	//定义一个结构体指针
	var s_p *Student
	s_p = &tom
	fmt.Printf("s_p: %v\n", s_p)
	fmt.Printf("(*s_p): %v\n", (*s_p))
	fmt.Printf("(*s_p).name: %v\n", (*s_p).name)
	fmt.Printf("s_p.name: %v\n", s_p.name) //在取成员变量的值的时候可以将*省略

	//使用new关键字创建结构体指针
	var jerry = new(Student)
	jerry.name = "Jerry"
	fmt.Printf("jerry: %v\n", jerry)
	fmt.Printf("(*jerry): %v\n", (*jerry)) //取值
	fmt.Printf("jerry.name: %v\n", jerry.name)

}

```

**运行结果**

```
s_p: &{Tom 18}
(*s_p): {Tom 18}
(*s_p).name: Tom
s_p.name: Tom
jerry: &{Jerry 0}
(*jerry): {Jerry 0}
jerry.name: Jerry
```



### 结构体作为函数参数

go语言的结构体可以像普通变量一样，作为函数参数，参数的传递方式分为两种：

1. 直接传结构体，在函数体内对结构体的操作不会影响原结构体（因为操作的只是原结构体的副本）
2. 传递结构体的地址，在函数体内对结构的任何操作都会对原结构体生效

```go
package main

import "fmt"

type Student struct {
	name string
	age  int
}

//值传递
func show_student(s Student, name string, age int) {
	s.name = name
	s.age = age
	fmt.Printf("s: %v\n", s)

}

//地址传递（引用传递）
func show_student2(s *Student, name string, age int) {
	s.name = name
	s.age = age
	fmt.Printf("s: %v\n", s)
}

func main() {

	tom := Student{name: "Tom", age: 18}
	jerry := Student{name: "Jerry", age: 18}

	//tom直接传值
	fmt.Printf("tom: %v\n", tom)
	show_student(tom, "Tom2", 20)
	fmt.Printf("tom: %v\n", tom)
	fmt.Println("-------------------")
	//jerry传递地址
	fmt.Printf("jerry: %v\n", jerry)
	show_student2(&jerry, "Jerry2", 20)
	fmt.Printf("jerry: %v\n", jerry)

}

```

**运行结果**

```
tom: {Tom 18}
s: {Tom2 20}
tom: {Tom 18}
-------------------
jerry: {Jerry 18}
s: &{Jerry2 20}
jerry: {Jerry2 20}
```



### 结构体的嵌套

go语言的结构体中可以嵌套另一个结构体。

```go
package main

import "fmt"

type Person struct {
	name string
	age  int
	dog  Dog //宠物狗，另一个结构体
}

type Dog struct {
	name string
	age  int
}

func main() {

	var tom Person
	tom.name = "Tom"
	tom.age = 18

	var erha = Dog{name: "二哈", age: 3}
	tom.dog = erha

	fmt.Printf("tom: %v\n", tom)
	fmt.Printf("tom.name: %v\n", tom.name)
	fmt.Printf("tom.dog.name: %v\n", tom.dog.name)
	tom.dog.name = "二哈哈哈"
	fmt.Printf("tom.dog: %v\n", tom.dog)
}

```

**运行结果**

```
tom: {Tom 18 {二哈 3}}
tom.name: Tom
tom.dog.name: 二哈
tom.dog: {二哈哈哈 3}
```



## 方法



go语言没有面向对象的特性，也没有类的概念。但是，可以使用结构体来模拟这些特性；在面向对象里面类有方法的概念，在go语言中，我们也可以为每一个结构体定义属于自己的方法。



### **语法**

go语言中的方法是一种特殊的函数，定义于struct之上（与struct绑定），被称为struct的接受者（receiver）。

```go
func (recv recv_type) method_name (params_li)(return_type){
    
}
```

- func：定义函数的关键字
- recv recv_type：接受者和接受者的类型
- method_name：方法名
- params_li：[参数]和参数类型列表
- return_type：返回值类型

> 方法和函数的定义基本上一模一样，只不过在func和函数名之间添加了一个接受者和接受者的类型。



```go
package main

import "fmt"

type Dog struct {
	name string
	age  int
}

//为结构体Dog定义方法
func (dog Dog) say_hi() {
	fmt.Printf("我叫%v，我今年%v岁啦！\n", dog.name, dog.age) 
}

func (dog Dog) woof() {
	fmt.Printf("%v正在汪汪大叫~\n", dog.name)
}

func main() {

	wc := Dog{"旺财", 2}
	wc.say_hi()
	wc.woof()

}

```

**运行结果**

```
我叫旺财，我今年2岁啦！
旺财正在汪汪大叫~
```



**方法的注意事项**

- 方法的receiver type并非一定是struct类型，type定义的类型别名，slice，map，channel，func类型等都可以。
- struct结合它的方法就等价于面向对象中的类。只不过struct可以和它的方法分开，并非一定要属于同一个文件，但必须属于同一个包。
- 方法有两种接收类型：（T Type）和(T *Type)。
- 方法就是函数，go语言中没有重载的说法，也就是同一个结构体中的所有方法都必须是唯一的。
- 如果receiver是一个指针类型，则会自动解除引用。
- 方法和type是分开的，意味着实例的行为和数据存储是分开的，但是它们通过receiver建立起关联关系。（个人感觉：只不过是多了一个必须参数receiver）



### 方法的接受者类型

和函数的传值方法一样，方法的接受者类型同样支持值传递和地址传递，通过地址传递的，在方法内部对接受者进行操作会影响到接受者。

```go
package main

import "fmt"

type Person struct {
	name string
}

//传递值
func (p Person) fix_name(name string) {
	p.name = name

}

//传递指针
func (p *Person) fix_name2(name string) {
	p.name = name

}

func main() {

	tom := Person{name: "Tom"}
	jerry := &Person{name: "Jerry"}
	fmt.Printf("tom.name: %v\n", tom.name)
	tom.fix_name("Tooooom")
	fmt.Printf("tom.name: %v\n", tom.name)
	print("----------------------------\n")
	fmt.Printf("jerry.name: %v\n", jerry.name)
	jerry.fix_name2("Jeeeeery")
	fmt.Printf("jerry.name: %v\n", jerry.name)

}

```

**运行结果**

```
tom.name: Tom
tom.name: Tom
----------------------------
jerry.name: Jerry
jerry.name: Jeeeeery
```

> 可以看到通过传递地址的fix_name2在内部修改结构体的属性后，该结构体的属性值会随之变化，而通过值传递的fix_name在内部修改结构体的属性值并不会对原结构体产生任何影响。





## 接口

go语言的接口，是一种新的类型定义，它把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就实现了这个接口。

接口就像一个公司里面的领导，它会定义一些通用规范，只是设计，而不实现规范。



**语法**

```go
type interface_name interface{
    func_name(parma_li)(return_li)
}
```

- interface_name：接口名
- func_name：方法名
- param_li：参数列表
- return_li：返回值列表





```go
package main

import "fmt"

// 定义一个USB的读写接口
type USB interface {
	read()
	write(string)
}

// 计算机
type Computer struct {
	name string
}

//手机
type Mobile struct {
	name string
}

// 计算机实现USB接口
func (c Computer) read() {
	fmt.Printf("%v is reading...\n", c.name)
}

func (c Computer) write() {
	fmt.Printf("%v is writing...\n", c.name)
}

func main() {

	c := Computer{name: "Mac"}
	c.read()
	c.write()

}

```

**运行结果**

```
Mac is reading...
Mac is writing...
```



### 接口和类型的关系

- 一个类型可以实现多个接口
- 多个类型可以实现一个接口（多态：不同的类型调用同一个方法，具体实现细节不一样）

```go
package main

import "fmt"

//一个类型实现多个接口
type Music interface {
	PlayMusic()
}

type Vide interface {
	PlayVideo()
}

type Mobile struct {
	name string
}

func (m Mobile) PlayMusic() {
	fmt.Printf("%v播放音乐...\n", m.name)
}

func (m Mobile) PlayVideo() {
	fmt.Printf("%v播放视频...\n", m.name)

}

// 多个类型实现一个接口
type Eat interface {
	eat()
}

type Dog struct {
	name string
}

type Cat struct {
	name string
}

func main() {

	m := Mobile{name: "Apple"}
	m.PlayMusic()
	m.PlayVideo()

	erha := Dog{name: "二哈"}
	fmt.Printf("%v is eating...\n", erha.name)
	tom := Cat{name: "Tom"}
	fmt.Printf("%v is eating...\n", tom.name)

}

```

**运行结果**

```
Apple播放音乐...
Apple播放视频...
二哈 is eating...
Tom is eating...
```



### 接口嵌套

接口可以通过嵌套，创建新的接口。

```go
package main

type Fly interface {
	fly()
}

type Swim interface {
	swim()
}

type FlyFish interface {
	Fly
	Swim
}

type Fish struct {
}

func (f Fish) fly() {
	print("flying...\n")
}

func (f Fish) swim() {
	print("swimming...\n")
}

func main() {

	f := Fish{}
	f.fly()

	var ff FlyFish
	ff = Fish{}
	ff.fly()	

}

```

**运行结果**

```
flying...
flying...
```



### 通过接口实现开闭原则（OCP）

面向对象的可复用设计的第一块基石，便是所谓的“开闭”原则（Open-Close Principle，常缩写为OCP，意为对扩展开放，对修改关闭）。虽然go语言不是面向对象语言，但是也可以模拟实现这个原则。



```go
package main

//定义一个宠物接口
type Pet interface {
	eat()
	sleep()
}

//定义两个结构体，然后分别实现接口
type Dog struct{}
type Cat struct{}

func (d Dog) eat() {
	print("dog is eating...\n")
}

func (d Dog) sleep() {
	print("dog is sleeping...\n")
}

func (c Cat) eat() {
	print("cat is eating...\n")
}

func (cat Cat) sleep() {
	print("cat is sleeping...\n")
}

type Person struct {
	name string
}

func (p Person) care(pet Pet) {

	pet.eat()
	pet.sleep()

}

func main() {

	d := Dog{}
	c := Cat{}
	p := Person{name: "Tom"}
	p.care(d)
	p.care(c)

}

```

**运行结果**

```
dog is eating...
dog is sleeping...
cat is eating...
cat is sleeping...
```

> 个人理解：为了方便理解，我觉得可以把它和面向对象语言里面的概念对比来理解，go语言中的接口有点类似于面向对象语言中的抽象基类，而实现了接口中所有方法的结构体可以看成是这个抽象基类的实现类，只有实现了接口中所有方法的实现类，才能把它看成是这个接口的一个子类，在参数类型是这个接口类型的地方，传入这个子类也是合法的。
>
> 在上面的例子中，可以看到，Pet这个接口有两个方法eat和sleep，在结构体Dog和Cat中都分别将它们实现了，因此可以将Dog和Cat看做是Pet的一个子类；而在后面的调用中也显示出了这个特性：当Person的care方法的参数是Pet类型时，传入是Dog或者Cat类型的参数也是正确的！
>
> 同时也可以举一个反例来看：
>
> 我将Cat结构体的sleep方法给注释掉，同时将p.sleep()的调用给注释掉，但是这样编译都不会通过。所以，很显然，因为Cat没有实现Pet中所有的方法，所以在使用Pet的地方是不能使用Cat的，而报错的提示也确实如此！

```go
package main

//定义一个宠物接口
type Pet interface {
	eat()
	sleep()
}

//定义两个结构体，
type Dog struct{}
type Cat struct{}

func (d Dog) eat() {
	print("dog is eating...\n")
}

func (d Dog) sleep() {
	print("dog is sleeping...\n")
}

func (c Cat) eat() {
	print("cat is eating...\n")
}

// func (cat Cat) sleep() {
// 	print("cat is sleeping...\n")
// }

type Person struct {
	name string
}

func (p Person) care(pet Pet) {

	pet.eat()
	// pet.sleep()

}

func main() {

	d := Dog{}
	c := Cat{}
	p := Person{name: "Tom"}
	p.care(d)
	p.care(c)

}

```

**运行结果**

```
# command-line-arguments
./lll06_通过接口实现OCP原则.go:46:9: cannot use c (variable of type Cat) as type Pet in argument to p.care:
        Cat does not implement Pet (missing sleep method)
```





## 继承

通过结构体嵌套来实现继承。

```go
package main

import "fmt"

//"父类"
type Pet struct {
	name string
	age  int
}

func (p Pet) eat() {
	fmt.Printf("%v is eating\n", p.name)
}

func (p Pet) sleep() {
	fmt.Printf("%v is sleeping\n", p.name)
}

// “子类”
type Dog struct {
	Pet
}

func main() {

	erha := Dog{
		Pet{name: "二哈", age: 2},
	}

	// 子类可以直接调用父类中的方法
	erha.eat()
	erha.sleep()

}

```

**运行结果**

```
二哈 is eating
二哈 is sleeping
```



## 并发编程



golang中的并发，是函数相互独立运行的能力，goroutines是并发运行的函数。golang提供了goroutines作为并发处理的一种方式。

创建一个协程非常简单，就是在一个任务函数前面添加一个go关键字：

```go
go task()
```



```go
package main

import (
	"fmt"
	"time"
)

func show(name string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("name: %v\n", name)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {

	go show("Python") //开启一个go协程
	go show("Golang") //开启另外一个go协程
	time.Sleep(time.Millisecond * 2000)
	print("hello world!\n") //默认情况下：主线程结束后其他线程也会结束
}

```

**运行结果**

```
name: Golang
name: Python
name: Python
name: Golang
name: Golang
name: Python
name: Python
name: Golang
name: Python
name: Golang
hello world!
```

 

### 并发编程之channel通道

go提供了一种称为通道的机制，用于在goroutine之间共享数据。当goroutine执行并发活动时，需要在goroutine之间共享数据资源，channel通道充当goroutine之间的管道并提供一种机制来保证同步交换。

需要在声明通道时指定数据类型，我们可以共享内置，命名，结构和引用类型的值和指针。数据在通道上传递：**在任何给定的时间只有一个goroutine可以访问数据项，因此按照设计不会发生数据竞争。**

根据数据交换的行为，有两种通道类型：**无缓冲通道**和**缓冲通道**。无缓存通道用于执行goroutine之间的同步信号，而缓冲通道用于执行异步信号。无缓冲通道保证在发送和接收发生的瞬间执行两个goroutine之间的交换。缓冲通道没有这样的保证。



通道由make函数创建，该函数指定chan关键字和通道的元素类型。

**语法**

```go
chan_ubuffered:=make(chan int) //整型无缓冲通道
chan_buffered:=make(chan int, 10) //整型有缓冲通道
```



**将值发送给通道使用==<-==运算符**

```go
c:=make(chan string)
c <- "hello"
```



**从通道中取值**

```go
var name string
name <- c
```



**无缓冲通道**

在无缓冲通道中，在接收到任何值之前没有能力保存它。在这种类型的通道中，发送和接收的goroutine在任何发送或接收操作完成之前的同一时刻都是准备就绪的。如果两个goroutine没有在同一时刻准备好，则通道会让其中已准备好的一方等待另一方，直到两个都准备就绪。同步是通道上发送和接收之间交互的基础。没有另一个就不可能发生。



**缓冲通道**

在缓冲通道中，有能力在接收到一个或多个值之前保存它们。在这种类型的通道中，不强制goroutine在同一时刻准备好执行发送和接收。当发送或接收阻塞时也有不同的条件。**只有当通道中没有要接收的值时，接收才会阻塞**；**仅当没有可用缓冲区来放置正在发送的值时，发送才会阻塞。**



```go
package main

import (
	"fmt"
	"time"
)

// 将值传入通道
func send(c *chan string, val string) {
	*c <- val
}

// 将值从通道中取出
func get(c *chan string) *string {
	val := <-*c
	fmt.Printf("val: %v\n", val)
	return &val
}

func main() {

	c := make(chan string) //无缓冲通道
	defer close(c)

	go send(&c, "hello")
	go get(&c)
	go send(&c, "world")
	go get(&c)

	time.Sleep(time.Millisecond * 1000)
	fmt.Printf("main end...")
}

```

**运行结果**

```
val: hello
val: world
main end...
```



### WaitGroup实现同步



在golang中，如果除了主协程以外还有其他的协程，当主协程结束的时候其他协程也会一起结束，不管它们有没有执行完毕，即默认是协程守护的。使用WaitGroup可以使用协程之间的同步，（个人感觉类似于python中的join方法）。



```go
package main

import (
	"fmt"
)

func show_msg(msg int) {
	fmt.Printf("msg: %d\n", msg)
}

func main() {
	for i := 0; i < 10; i++ {
		go show_msg(i)
	}

	print("end main...")
}

```

**运行结果**

```
end main...msg: 3
msg: 0
msg: 1
msg: 2
msg: 6
msg: 4
msg: 9
```

> 可以看到上面的代码的每次运行结果都可能不一样，因为每一个show_msg都用一个协程来启动，它们和主协程一样，但是只要主协程执行结束，不管它们有没有被执行，都会被结束。



**使用WaitGroup后**

```go
package main

import (
	"fmt"
	"sync"
)

var mg sync.WaitGroup

func show_msg(msg int) {
	fmt.Printf("msg: %d\n", msg)
	defer mg.Add(-1) //任务完成就减一
}

func main() {

	for i := 0; i < 10; i++ {
		mg.Add(1) //添加任务就加一
		go show_msg(i)
	}
	// Counter is 0, no need to wait.
	mg.Wait() //如果WaitGroup的计数为0就停止等待
	print("end main...")
}

```

**运行结果**

```
msg: 9
msg: 0
msg: 1
msg: 2
msg: 3
msg: 4
msg: 5
msg: 6
msg: 7
msg: 8
end main...
```



> 在使用WaitGroup后，通过WaitGroup来进行计数，每添加一个协程，计数加一（Add(1)）；每完成一个协程，计数减一(Add(-1))，只有当计数为0时Wait才不会等待。



### 并发编程之runtime包

runtime包中包含了一些协程管理相关的API。



#### runtime.Gosched()

功能：让出处理机，给其他的协程去使用。

> // Gosched yields the processor, allowing other goroutines to run. It does not
>
> // suspend the current goroutine, so execution resumes automatically.

```go
package main

import (
	"fmt"
	"runtime"
)

func show_msg(msg int) {
	fmt.Printf("msg: %v\n", msg)

}

func main() {
	go show_msg(100)

	runtime.Gosched() //在主协程中使用的

	print("main end...")
}

```

**运行结果**

```
msg: 100
main end...
```

> 在上面的代码中，我们可以看到在主协程中调用了Gosched方法，因此当主协程和show_msg协程抢占处理机时，如果show_msg协程先抢到执行，则执行结果如上；如果主协程先抢到处理机，则它会主动让出处理机，此时show_msg会拿到处理机，因此执行结果还是如上。



#### runtime.Goexit()

功能：退出协程

```go
package main

import (
	"fmt"
	"runtime"
)

func show_msg(msg int) {
	for i := 0; i < 5; i++ {
		if i == 3 { //i等于3时直接结束协程
			runtime.Goexit()
		}
		fmt.Printf("msg: %v-%v\n", msg, i)

	}

}

func main() {
	go show_msg(100)
	runtime.Gosched()
	print("main end...")
}

```

**运行结果**

```
msg: 100-0
msg: 100-1
msg: 100-2
main end...
```



### Mutex实现同步

除了channel实现同步之外，还可以使用Mutex互斥锁实现同步。

```go
package main

import (
	"fmt"
	"sync"
)

var n int = 100
var lock sync.Mutex
var wg sync.WaitGroup

func add() {
	lock.Lock() //加锁
	n += 1
	lock.Unlock() //解锁
	wg.Add(-1)
}

func sub() {
	lock.Lock()
	n -= 1
	lock.Unlock()
	wg.Add(-1) //完成协程任务，计数减一
}

func main() {

	//当多个协程共同处理某一个数据时，如果不进行加锁，就会出现数据混乱的情况，例如：多个协程对一个数值进行加一和减一的操作
	//因此在多协程中，我们可以对数据进行的操作加锁，让同一时刻只能有一个协程对数据进行操作
	for i := 0; i < 100000; i++ {
		wg.Add(1) //新增协程任务，计数加一
		go add()
		wg.Add(1)
		go sub()
	}
	wg.Wait() //主协程等待子协程执行完毕
	fmt.Printf("n: %v\n", n)

}

```

**运行结果**

```
n: 100
```



### 并发编程之select

1. select是go中的一个控制结构，类似于switch语句，用于处理异步IO操作。select会监听case语句中的channel的读写操作，当case中channel读写操作为非阻塞状态时，会触发相应的动作。

   > select中case语句必须是一个channel操作（读或者写）
   >
   > select中default子句总是可以运行的

2. 如果有多个case都可运行，select会随机公平地选出一个执行，其他不会执行。

3. 如果没有可运行的case语句，且有default语句，那么就会执行default语句。

4. 如果没有可运行的case语句，且没有default语句，select将会阻塞。



```go
package main

import (
	"fmt"
	"time"
)

var chanName = make(chan string, 0)
var chanAge = make(chan int, 0)

func main() {

	go func() {
		for i := 0; i < 5; i++ {
			name := "Tom"
			chanName <- name
			chanAge <- i
		}
	}()

	for {
		select {
		case name := <-chanName: //读chanName
			fmt.Printf("name: %v\n", name)
		case age := <-chanAge:
			fmt.Printf("age: %v\n", age)

		default:
			print("default...\n")

		}
		time.Sleep(time.Second)
	}

}

```

**运行结果**

```
default...
name: Tom
age: 0
name: Tom
age: 1
name: Tom
age: 2
name: Tom
age: 3
name: Tom
age: 4
default...
default...
default...
default...
```



### 并发编程是Timer

Timer顾名思义，就是计时器的意思，可以实现一些定时操作，内部也是通过channel实现的。



```go
package main

import (
	"fmt"
	"time"
)

func main() {

	t1 := time.NewTimer(time.Second) //等待1秒钟
	fmt.Printf("time.Now(): %v\n", time.Now())
	res := <-t1.C //一直处于阻塞状态，直到到达等待时间
	fmt.Printf("res: %v\n", res)
	fmt.Printf("time.Now(): %v\n", time.Now())

	time.After(time.Second) //一秒后，After其实对NewTimer进行了一次封装
	fmt.Printf("time.Now(): %v\n", time.Now())

	//停止计时器
	t2 := time.NewTimer(time.Second * 2)

	go func() {
		print("执行到了这里...\n")
		<-t2.C
		print("匿名函数被执行...\n")
	}()

	t2.Stop() //如果调用了stop，上面的协程<-t2.C以后的语句都不会被执行
	time.Sleep(time.Second * 3)

	//修改计时时间
	t3 := time.NewTimer(time.Second)
	fmt.Printf("time.Now(): %v\n", time.Now())
	t3.Reset(time.Second * 2) //修改计时为2两秒
	<-t3.C
	fmt.Printf("time.Now(): %v\n", time.Now())

	fmt.Printf("main end...\n")
}

```

**运行结果**

```
time.Now(): 2022-10-05 15:04:46.380136682 +0800 CST m=+0.000102410
res: 2022-10-05 15:04:47.380331775 +0800 CST m=+1.000297497
time.Now(): 2022-10-05 15:04:47.380414771 +0800 CST m=+1.000380500
time.Now(): 2022-10-05 15:04:47.380453944 +0800 CST m=+1.000419671
执行到了这里...
time.Now(): 2022-10-05 15:04:50.381683725 +0800 CST m=+4.001649449
time.Now(): 2022-10-05 15:04:52.38280595 +0800 CST m=+6.002771675
main end...
```



### 并发编程之Ticker

Timer是到达时间后执行一次，而Ticker是周期性的执行，是一个定时器。

```go
package main

import (
	"fmt"
	"time"
)

func main() {

	// 每隔一秒打印一次当前时间
	t1 := time.NewTicker(time.Second)
	n := 0
	for _ = range t1.C {
		fmt.Printf("time.Now(): %v\n", time.Now())
		n += 1
		if n > 5 {
			t1.Stop()
			break
		}
	}

}

```

**运行结果**

```
time.Now(): 2022-10-05 15:20:32.53228156 +0800 CST m=+8.000943370
time.Now(): 2022-10-05 15:20:33.532483936 +0800 CST m=+9.001145744
time.Now(): 2022-10-05 15:20:34.531494225 +0800 CST m=+10.000156042
time.Now(): 2022-10-05 15:20:35.532310152 +0800 CST m=+11.000971959
time.Now(): 2022-10-05 15:20:36.532446657 +0800 CST m=+12.001108468
time.Now(): 2022-10-05 15:20:37.531670839 +0800 CST m=+13.000332645
```



```go
package main

import (
	"fmt"
	"time"
)

func main() {


	t2 := time.NewTicker(time.Second)
	chanInt := make(chan int)
	go func() {
		//每秒钟向chanInt随机写一个数据
		for _ = range t2.C {
			select {
			case chanInt <- 1:
			case chanInt <- 2:
			case chanInt <- 3:
			case chanInt <- 4:
			case chanInt <- 5:

			}
		}

	}()
	sum := 0
	for v := range chanInt {
		sum += v
		fmt.Printf("v: %v\n", v)
		if sum >= 10 {
			close(chanInt)
			break
		}
	}

	print("main end...\n")

}

```

**运行结果**

```
v: 1
v: 4
v: 1
v: 5
main end...
```



### 并发编程之原子操作

当多个协程操作同一个数据时，如果不加锁就可能会出现数据混乱的情况，除了使用mutex锁来解决这个问题外，golang还提供了原子操作这一概念来解决这个问题。



atomic提供的原子操作能够确保任一时刻只有一个goroutine对变量进行操作，善用atomic能够避程序中出现大量的锁操作。

atomic常见的操作有：

- 增减：Add
- 加载（读）：Load
- 存储（写）：Store
- 比较和交换：cas
- 交换：swap



**增减操作**

```go
// AddInt32 atomically adds delta to *addr and returns the new value.
func AddInt32(addr *int32, delta int32) (new int32)

// AddUint32 atomically adds delta to *addr and returns the new value.
// To subtract a signed positive constant value c from x, do AddUint32(&x, ^uint32(c-1)).
// In particular, to decrement x, do AddUint32(&x, ^uint32(0)).
func AddUint32(addr *uint32, delta uint32) (new uint32)

// AddInt64 atomically adds delta to *addr and returns the new value.
func AddInt64(addr *int64, delta int64) (new int64)

// AddUint64 atomically adds delta to *addr and returns the new value.
// To subtract a signed positive constant value c from x, do AddUint64(&x, ^uint64(c-1)).
// In particular, to decrement x, do AddUint64(&x, ^uint64(0)).
func AddUint64(addr *uint64, delta uint64) (new uint64)

// AddUintptr atomically adds delta to *addr and returns the new value.
func AddUintptr(addr *uintptr, delta uintptr) (new uintptr)
```



**加载操作**

```go
// LoadInt32 atomically loads *addr.
func LoadInt32(addr *int32) (val int32)

// LoadInt64 atomically loads *addr.
func LoadInt64(addr *int64) (val int64)

// LoadUint32 atomically loads *addr.
func LoadUint32(addr *uint32) (val uint32)

// LoadUint64 atomically loads *addr.
func LoadUint64(addr *uint64) (val uint64)

// LoadUintptr atomically loads *addr.
func LoadUintptr(addr *uintptr) (val uintptr)

// LoadPointer atomically loads *addr.
func LoadPointer(addr *unsafe.Pointer) (val unsafe.Pointer)
```



**存储操作**

```go
// StoreInt32 atomically stores val into *addr.
func StoreInt32(addr *int32, val int32)

// StoreInt64 atomically stores val into *addr.
func StoreInt64(addr *int64, val int64)

// StoreUint32 atomically stores val into *addr.
func StoreUint32(addr *uint32, val uint32)

// StoreUint64 atomically stores val into *addr.
func StoreUint64(addr *uint64, val uint64)

// StoreUintptr atomically stores val into *addr.
func StoreUintptr(addr *uintptr, val uintptr)

// StorePointer atomically stores val into *addr.
func StorePointer(addr *unsafe.Pointer, val unsafe.Pointer)
```



**比较和交换操作**

```go
// CompareAndSwapInt32 executes the compare-and-swap operation for an int32 value.
func CompareAndSwapInt32(addr *int32, old, new int32) (swapped bool)

// CompareAndSwapInt64 executes the compare-and-swap operation for an int64 value.
func CompareAndSwapInt64(addr *int64, old, new int64) (swapped bool)

// CompareAndSwapUint32 executes the compare-and-swap operation for a uint32 value.
func CompareAndSwapUint32(addr *uint32, old, new uint32) (swapped bool)

// CompareAndSwapUint64 executes the compare-and-swap operation for a uint64 value.
func CompareAndSwapUint64(addr *uint64, old, new uint64) (swapped bool)

// CompareAndSwapUintptr executes the compare-and-swap operation for a uintptr value.
func CompareAndSwapUintptr(addr *uintptr, old, new uintptr) (swapped bool)

// CompareAndSwapPointer executes the compare-and-swap operation for a unsafe.Pointer value.
func CompareAndSwapPointer(addr *unsafe.Pointer, old, new unsafe.Pointer) (swapped bool)
```



**交换操作**

```go
// SwapInt32 atomically stores new into *addr and returns the previous *addr value.
func SwapInt32(addr *int32, new int32) (old int32)

// SwapInt64 atomically stores new into *addr and returns the previous *addr value.
func SwapInt64(addr *int64, new int64) (old int64)

// SwapUint32 atomically stores new into *addr and returns the previous *addr value.
func SwapUint32(addr *uint32, new uint32) (old uint32)

// SwapUint64 atomically stores new into *addr and returns the previous *addr value.
func SwapUint64(addr *uint64, new uint64) (old uint64)

// SwapUintptr atomically stores new into *addr and returns the previous *addr value.
func SwapUintptr(addr *uintptr, new uintptr) (old uintptr)

// SwapPointer atomically stores new into *addr and returns the previous *addr value.
func SwapPointer(addr *unsafe.Pointer, new unsafe.Pointer) (old unsafe.Pointer)
```

 

```go
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var N int32 = 100
var wg sync.WaitGroup

func add() {
	atomic.AddInt32(&N, 1) //加一操作
	defer wg.Add(-1)
}

func sub() {
	atomic.AddInt32(&N, -1) //减一操作
	defer wg.Add(-1)
}
func main() {
	for i := 0; i < 100000; i++ {
		go add()
		wg.Add(1)
		go sub()
		wg.Add(1)
	}

	wg.Wait()
	fmt.Printf("N: %v\n", N)
}

```

**运行结果**

```
N: 100
```



```go
package main

import (
	"fmt"
	"sync/atomic"
)

func main() {

	var i int32 = 100
	fmt.Printf("i: %v\n", i)
	atomic.AddInt32(&i, 100) //加100
	fmt.Printf("i: %v\n", i)
	atomic.AddInt32(&i, -50) //减50
	fmt.Printf("i: %v\n", i)

	fmt.Printf("atomic.LoadInt32(&i): %v\n", atomic.LoadInt32(&i)) //读取i的值
	atomic.StoreInt32(&i, 1000)                                    //修改i的值为1000
	fmt.Printf("i: %v\n", i)

	//cas 其他操作的根基
	success := atomic.CompareAndSwapInt32(&i, i, 666) //修改i为666
	if success {

		fmt.Printf("修改成功！i: %v\n", i)
	}

	//swap 直接交换 比较暴力
	b := atomic.CompareAndSwapInt32(&i, i, 888)
	if b {
		fmt.Printf("修改成功！i: %v\n", i)
	}

}

```

**运行结果**

```
i: 100
i: 200
i: 150
atomic.LoadInt32(&i): 150
i: 1000
修改成功！i: 666
修改成功！i: 888
```



## 标准库

### log

golang内置了log包，实现了简单的日志服务。通过调用log包的函数，可以实现简单的日志打印功能。



log包中有3个系列的日志打印函数，分别是print系列，panic系列和fatal系列。



| 函数系列 | 作用                                                   |
| -------- | ------------------------------------------------------ |
| print    | 单纯打印日志                                           |
| panic    | 打印日志，抛出panic异常                                |
| fatal    | 打印日志，强制结束程序（os.Exit(1)）,defer函数不会执行 |



```go
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	//配置日志的输出前缀
	log.SetPrefix("Log:")

	//配置日志
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)

	//设置日志输出到文件
	f, _ := os.OpenFile("a.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModePerm)
	defer f.Close()
	log.SetOutput(f)

	//简单打印日志
	log.Print("hello golang...")
	defer fmt.Print("程序结束...")
	//panic日志
	// log.Panic("bye...") //defer会被执行

	//fatal日志
	// log.Fatal("致命错误...") //defer不会被执行

	//自定义logger
	my_logger := log.New(os.Stdout, "My Log:", log.Ldate|log.Ltime|log.Lmicroseconds|log.Llongfile)
	my_logger.Print("这个是自己定义的logger，一次性配置所有，会方便许多!")

}

```

**运行结果**

```
My Log:2022/11/05 17:27:35.150644 /home/lll/Desktop/go/lll08_标准库/os/lll07_日志相关的操作.go:33: 这个是自己定义的logger，一次性配置所有，会方便许多!
```



### builtin

这个包提供了一些类型声明，变量和常量声明，还有一些便利的函数，这个包不需要导入，这些变量和函数就可以直接使用。



#### panic

抛出一个panic异常。

```go
package main

import "fmt"

func main() {

	//panic 抛出异常
	panic("抛出异常!!!")
	
}

```



#### new和make

new和make的区别：

1. make只能用来分配及初始化类型为slice，map和chan的数据；new可以分配任意类型的数据
2. new分配返回的是指针，即类型为*T；make返回的是数据的值，即T
3. make分配后会对数据进行初始化，而new不会



```go
package main

import "fmt"

func main() {


	s := new(string)
	fmt.Printf("s: %T\n", s) //*string
	fmt.Printf("s: %v\n", *s)

	i2 := new([]int)
	fmt.Printf("i2: %T\n", i2) //*[]int
	fmt.Printf("i2: %v\n", *i2)

	i3 := make([]int, 10, 100) //初始化容量为100，长度为10
	fmt.Printf("i3: %T\n", i3)
	fmt.Printf("i3: %v\n", i3)

}

```

**运行结果**

```
s: *string
s: 
i2: *[]int
i2: []
i3: []int
i3: [0 0 0 0 0 0 0 0 0 0]
```

> 内建函数make(T,args)与new(T)的用途不一样。它只用来创建slice，map和chan，并且返回一个初始化后的类型为T的数据。之所以不同，是因为这三个类型的背后引用了使用前必须初始化的数据结构。例如：slice是一个三元描述符，包含一个指向数据的指针，长度，以及容量，在这些项被初始化之前，slice都是nil。对于slice，map和chan，make初始化这些内部数据结构，并准备好可用的值。



### bytes

bytes提供了对字节切片进行读写操作的一系列函数，字节切片处理函数比较多分为基本处理函数，比较函数，后缀检查函数，索引函数，分割函数，大小写处理函数和子切片处理函数。



```go
package main

import (
	"bytes"
	"fmt"
)

func main() {

	s1 := "hello world!"
	b1 := []byte("你好，世界！")
	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("b1: %v\n", b1)

	//bytes和string的相互转换
	//1.bytes转string
	fmt.Printf("string(b1): %v\n", string(b1))
	//2.string转bytes
	fmt.Printf("[]byte(s1): %v\n", []byte(s1))

	//contains：检查bytes中是否包含子bytes
	fmt.Printf("bytes.Contains(b1, []byte(\"世界\")): %v\n", bytes.Contains(b1, []byte("世界")))

	//count：统计某个bytes出现的次数
	fmt.Printf("bytes.Count([]byte(s1), []byte(\"l\")): %v\n", bytes.Count([]byte(s1), []byte("l")))

	//compare:比较两个bytes The result will be 0 if a == b, -1 if a < b, and +1 if a > b.
	fmt.Printf("bytes.Compare(b1, []byte(\"hello\")): %v\n", bytes.Compare(b1, []byte("hello")))

	//分割bytes
	before, after, _ := bytes.Cut(b1, []byte("，"))
	fmt.Printf("before: %v\n", string(before))
	fmt.Printf("after: %v\n", string(after))

	//连接bytes
	b := bytes.Join([][]byte{b1, []byte(s1)}, []byte("==="))
	fmt.Printf("b: %v\n", string(b))

	//runes 转成utf8编码 这样能够正确计算中文长度
	r := bytes.Runes(b1)
	fmt.Printf("bytes.Runes(b1): %v\n", r)
	fmt.Printf("len(r): %v\n", len(r))

}

```

**运行结果**

```
s1: hello world!

b1: [228 189 160 229 165 189 239 188 140 228 184 150 231 149 140 239 188 129]

string(b1): 你好，世界！

[]byte(s1): [104 101 108 108 111 32 119 111 114 108 100 33]

bytes.Contains(b1, []byte("世界")): true

bytes.Count([]byte(s1), []byte("l")): 3

bytes.Compare(b1, []byte("hello")): 1

before: 你好

after: 世界！

b: 你好，世界！===hello world!

bytes.Runes(b1): [20320 22909 65292 19990 30028 65281]

len(r): 6
```



### errors

errors包实现了操作错误的函数。go语言使用error类型来返回函数执行过程中遇到的错误，如果返回的error值为nil，则表示未遇到错误，否则error会返回一个字符串，用于说明遇到了什么错误。



error的结构

```go
type error interface {
    Error() string
}
```

**你可以使用任何类型去实现它（只要添加一个Error()方法即可）**，也就是说，error可以是任何类型，这意味着，函数返回的error值实际可以包含任意信息，不一定是字符串。

error不一定表示一个错误，它可以表示任何信息，比如io包中就用error类型的io.EOF表示数据读取结束，而不是遇到了什么错误。

errors包实现了一个最简单的error类型，只包含了一个字符串，它可以记录大多数情况下遇到的错误信息。errors包的用法很简单，只有一个New函数，用于生成一个简单的error对象：

```go
func New(text string) error
```



```go
package main

import (
	"errors"
	"fmt"
	"time"
)

//自定义errors
type MyError struct {
	When time.Time
	What string
}

/*
type error interface {
	Error() string
}


error是一个接口，只要实现了Error方法，就可以是一个error

*/

func (e MyError) Error() string {
	return fmt.Sprintf("%v : %v", e.When, e.What)
}

//检测字符串是否为空
func check_str(s string) (err error) {

	if s == "" {
		err = errors.New("字符串不能为空...")
	} else {
		err = MyError{
			When: time.Date(2022, 11, 11, 11, 11, 11, 11, time.UTC),
			What: fmt.Sprintf("%v 不是一个空字符串...", s),
		}
	}
	return

}
func main() {

	s := ""
	err := check_str(s)
	fmt.Printf("err: %v\n", err)
	s2 := "hello python"
	err2 := check_str(s2)
	fmt.Printf("err2: %v\n", err2)

}

```

**运行结果**

```
err: 字符串不能为空...
err2: 2022-11-11 11:11:11.000000011 +0000 UTC : hello python 不是一个空字符串...
```



### sort

sort包提供了排序切片和用户自定义数据集以及相关功能的函数。

sort包主要针对[]int，[]float64，[]string以及其他自定义的切片排序，对于任何类型的切片，只要实现了**Len**，**Less**和**Swap**接口就可以对其进行排序。

sort.Sort默认使用快速排序方法，因此平均时间复杂度为nlog(n)

```go
// Sort sorts data in ascending order as determined by the Less method.
// It makes one call to data.Len to determine n and O(n*log(n)) calls to
// data.Less and data.Swap. The sort is not guaranteed to be stable.
func Sort(data Interface) {
	n := data.Len()
	quickSort(data, 0, n, maxDepth(n))
}
```



 ```go
package main

import (
	"fmt"
	"sort"
)

type Person struct {
	name   string
	age    int
	weight float64
}

type PersonSlice []Person

//实现Len方法，返回切片的长度
func (ps PersonSlice) Len() int {
	return len(ps)
}

//实现Less方法，定义比较规则
func (ps PersonSlice) Less(i int, j int) bool {
	return ps[i].age < ps[j].age //按照年龄进行比较
}

//实现Swap方法，定义交换规则
func (ps PersonSlice) Swap(i int, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func main() {

	//对int切片进行排序 []float64和[]string的使用方法同此
	var a []int
	a = []int{5, 1, 3, 2, 4}
	fmt.Printf("排序前:a: %v\n", a)
	sort.Ints(a)
	fmt.Printf("排序后:a: %v\n", a)
	fmt.Printf("sort.IntsAreSorted(a): %v\n", sort.IntsAreSorted(a)) //判断[]int序列是不是增序序列
	fmt.Printf("sort.SearchInts(a, 3): %v\n", sort.SearchInts(a, 3)) //查找某一个元素的位置 默认是二分法进行查找

	//对自定义数据类型的切片进行排序
	tom := Person{name: "Tom", age: 18, weight: 66.7}
	jerry := Person{name: "Jerry", age: 16, weight: 56.7}
	jack := Person{name: "Jack", age: 19, weight: 78.1}
	hank := Person{name: "Hank", age: 18, weight: 61.7}
	marry := Person{name: "Marry", age: 20, weight: 55.7}
	var ps PersonSlice
	ps = append(ps, tom, jerry, jack, hank, marry)
	fmt.Printf("排序前ps: %v\n", ps)

	sort.Sort(ps)
	fmt.Printf("排序后ps: %v\n", ps) //按年龄是增序序列

}

 ```

**运行结果**

```
排序前:a: [5 1 3 2 4]
排序后:a: [1 2 3 4 5]
sort.IntsAreSorted(a): true
sort.SearchInts(a, 3): 2
排序前ps: [{Tom 18 66.7} {Jerry 16 56.7} {Jack 19 78.1} {Hank 18 61.7} {Marry 20 55.7}]
排序后ps: [{Jerry 16 56.7} {Tom 18 66.7} {Hank 18 61.7} {Jack 19 78.1} {Marry 20 55.7}]
```



### time

Package time provides functionality for measuring and displaying time.（用于时间的测量和显示）



#### 基本使用

```go
//获取当前时间
	now := time.Now()
	fmt.Printf("now: %v\n", now)
	year := now.Year()          //年
	month := now.Month()        //月
	day := now.Day()            //日
	hour := now.Hour()          //时
	minute := now.Minute()      //分
	second := now.Second()      //秒
	nsecond := now.Nanosecond() //纳秒
	fmt.Printf("%v-%v-%v:%v:%v:%v:%v\n", year, month, day, hour, minute, second, nsecond)
```

**运行结果**

```
now: 2022-11-14 21:53:00.435701792 +0800 CST m=+0.000115447
2022-November-14:21:53:0:435701792
```



#### 时间戳

时间戳是自1970年1月1日（08:00:00GMT）至当前时间的总秒数。它也被称为Unix时间戳。

```go
//时间戳
fmt.Printf("now.Unix(): %v\n", now.Unix())           //秒数
fmt.Printf("now.UnixMicro(): %v\n", now.UnixMicro()) //毫秒
fmt.Printf("now.UnixNano(): %v\n", now.UnixNano())   //纳秒

//时间戳转成时间
fmt.Printf("time.Unix(now.Unix(), 0): %v\n", time.Unix(now.Unix(), 0))
```

**运行结果**

```
now.Unix(): 1668434792
now.UnixMicro(): 1668434792741556
now.UnixNano(): 1668434792741556203
time.Unix(now.Unix(), 0): 2022-11-14 22:06:32 +0800 CST
```



#### 时间的"运算"

```go
//时间的运算
	today := now
	tomorrow := today.Add(time.Hour * 24) //在当前时间上增加24小时
	fmt.Printf("today: %v\n", today)
	fmt.Printf("tomorrow: %v\n", tomorrow)

	dif := today.Sub(tomorrow) //今天与昨天相差的时间
	fmt.Printf("dif: %v\n", dif)

	//比较两个时间是否相同
	fmt.Printf("today.Equal(tomorrow): %v\n", today.Equal(tomorrow))

	//判断当前时间是否在目标时间之前
	fmt.Printf("today.Before(tomorrow): %v\n", today.Before(tomorrow))

	//判断当前时间是否在目标时间之后
	fmt.Printf("today.After(tomorrow): %v\n", today.After(tomorrow))
```

**运行结果**

```
today: 2022-11-14 22:16:35.620284622 +0800 CST m=+0.000048094
tomorrow: 2022-11-15 22:16:35.620284622 +0800 CST m=+86400.000048094
dif: -24h0m0s
today.Equal(tomorrow): false
today.Before(tomorrow): true
today.After(tomorrow): false
```



#### Ticker和Timer

```go
//定时器Ticker
	c := time.Tick(time.Second) //设置一个间隔一秒的定时器
	for i := range c {
		fmt.Printf("now: %v\n", i)             //每隔一秒打印一下当前时间
		if i.After(now.Add(time.Second * 3)) { //3秒后停止
			break
		}
	}

	//Timer
	t := time.NewTimer(time.Second * 2)
	fmt.Printf("now2: %v\n", time.Now())
	<-t.C
	fmt.Printf("now2 after 2 seconds: %v\n", time.Now()) //两秒后执行到这里
```

**运行结果**

```
now: 2022-11-14 22:34:08.697290905 +0800 CST m=+1.000982288
now: 2022-11-14 22:34:09.697399128 +0800 CST m=+2.001090512
now: 2022-11-14 22:34:10.697623176 +0800 CST m=+3.001314560
now2: 2022-11-14 22:34:10.697739262 +0800 CST m=+3.001430647
now2 after 2 seconds: 2022-11-14 22:34:12.698702444 +0800 CST m=+5.002393830
```



#### 时间格式化

时间类型有一个自带的Format方法，需要注意的是Go语言中格式化的模板不是常见的Y-m-d H:M:S，而是使用Go的诞生时间2006年1月2号15点04分（记忆口诀为：2006 1 2 3 4）。



```go
//时间格式化
fmt.Printf("now.Format(\"2006-01-02 15:04:05.000 Mon Jan\"): %v\n", now.Format("2006/01/02 15:04:05.000 Mon Jan")) //24小时制
fmt.Printf("now.Format(\"Mon Jan 2006-01-02 3:4:4 PM\"): %v\n", now.Format("Mon Jan 2006-01-02 3:4:4.000 PM"))     //12小时制
```

**运行结果**

```
now.Format("2006-01-02 15:04:05.000 Mon Jan"): 2022/11/14 22:44:56.882 Mon Nov
now.Format("Mon Jan 2006-01-02 3:4:4 PM"): Mon Nov 2022-11-14 10:44:44.882 PM
```



#### 解析字符串格式的时间

```go
//解析字符串格式的时间
loc, _ := time.LoadLocation("Asia/Shanghai")

//第一个参数指定格式，第二参数为字符串格式的时间，第三个参数指定时区
t2, _ := time.ParseInLocation("2006/01/02 15:04:05.000 Mon Jan", "2022/11/11 22:44:56.882 Mon Nov", loc)
fmt.Printf("t2: %T\n", t2)
fmt.Printf("t2: %v\n", t2)
```

**运行结果**

```
t2: time.Time
t2: 2022-11-11 22:44:56.882 +0800 CST
```


### json

这个包可以实现json的编码和解码，即实现json对象和struct之间相互转换。



**核心的两个函数**

```go
func Marshal(v interface{}) ([]byte ,error)  //将go语言中的struct编码成json,返回json字符串的字节切片和错误信息
```

```go
func Unmarshal(data []byte, v interface{}) error //将json解码成go语言中的struct，返回错误信息
```



```go
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person1 struct {
	Name   string
	Age    int
	PetDog Dog
}

type Dog struct {
	Name string
	Age  int
}

func main() {

	erha := Dog{Name: "二哈", Age: 3}
	tom := Person1{Name: "Tom", Age: 18, PetDog: erha}
	fmt.Printf("tom: %v\n", tom)

	//1.将一个struct实例转换成一个json对象
	//Marshal接收一个任何类型的对象，然后会返回对应json字符串的字节切片
	b, _ := json.Marshal(tom)
	json_b := string(b)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("json_b: %v\n", json_b)

	//2.将一个json字符串转成一个struct Unmarshal接收一个json字符串的字节切片和一个任意struct对象的地址
	//Unmarshal会将json字符串对象的值相应的赋给这个struct对象
	var tom2 Person1
	json.Unmarshal([]byte(json_b), &tom2)
	fmt.Printf("tom2: %v\n", tom2)

	//3.从io流中获取json字符串，然后转成struct
	f, _ := os.Open("demo.json")
	defer f.Close()
	json_decoder := json.NewDecoder(f)
	var tom3 map[string]interface{}
	json_decoder.Decode(&tom3)
	for k, v := range tom3 {
		fmt.Printf("k:%v,v:%v\n", k, v)
	}

	//4.将struct实例转成json后写入文件
	jerry := Person1{Name: "Jerry", Age: 16, PetDog: erha}
	f2, _ := os.OpenFile("demo2.json", os.O_RDWR|os.O_CREATE, 0777)
	defer f2.Close()
	json_encoder := json.NewEncoder(f2)
	json_encoder.Encode(jerry)

	//从json文件中读取写入的内容
	res := make([]byte, 100)
	f3, _ := os.Open("demo2.json")
	f3.Read(res)
	defer f3.Close()
    fmt.Printf("res: %v\n", res)


}

```

**运行结果**

```
tom: {Tom 18 {二哈 3}}
b: [123 34 78 97 109 101 34 58 34 84 111 109 34 44 34 65 103 101 34 58 49 56 44 34 80 101 116 68 111 103 34 58 123 34 78 97 109 101 34 58 34 228 186 140 229 147 136 34 44 34 65 103 101 34 58 51 125 125]
json_b: {"Name":"Tom","Age":18,"PetDog":{"Name":"二哈","Age":3}}
tom2: {Tom 18 {二哈 3}}
k:PetDog,v:map[Age:3 Name:二哈]
k:Name,v:Tom
k:Age,v:18
res: {"Name":"Jerry","Age":16,"PetDog":{"Name":"二哈","Age":3}}
```



### xml

和对json的编码和解码类似，golang也提供了对xml的编码和解码，并且用法极其相似。

```go
package main

import (
	"encoding/xml"
	"fmt"
)

type Person2 struct {
	XMLName xml.Name `xml:"Person"` //xml中的根元素
	Name    string   `xml:"name"`
	Age     int      `xml:"age"`
	Email   string   `xml:"email"`
}

func main() {

	p := Person2{
		Name:  "Tom",
		Age:   18,
		Email: "Tom@110.com",
	}
	//1.将一个struct实例转成一个xml
	// xml.Marshal(&p) //这个是没有缩进的
	b, _ := xml.MarshalIndent(&p, "", " ")
	fmt.Printf("%v\n", string(b))

	//2.将一个xml转成一个struct
	var p2 Person2
	xml.Unmarshal(b, &p2)
	fmt.Printf("p2: %v\n", p2)

}

```

**运行结果**

```
<Person>
 <name>Tom</name>
 <age>18</age>
 <email>Tom@110.com</email>
</Person>
p2: {{ Person} Tom 18 Tom@110.com}
```















### math

## 操作数据库

### 操作mysql

### 操作MongoDB

## gorm

## git

## docker













































































## 参考

golang入门到项目实战：https://www.bilibili.com/video/BV1zR4y1t7Wj?p=1&vd_source=95ef87e61d1c37fc15117824ffba69f5

Golang 25 个保留关键字：https://blog.csdn.net/K346K346/article/details/90265765



























