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
	fmt.Printf("sort.SearchInts(a, 3): %v\n", sort.SearchInts(a, 3)) //查找某一个元素的位置

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
