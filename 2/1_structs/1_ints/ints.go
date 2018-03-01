package main

import (
	"fmt"
	"sort"
)

type MyStruct struct {
	Num  int
	Name string
}

func NewMyStruct() func(string) MyStruct {
    id := 0
return func(name string)MyStruct{
    id = id + 1
return MyStruct{id, name}
}
}

type MyInt int

type withFiles bool

func (m MyInt) showYourSelf() {
	fmt.Printf("%T %v\n", m, m)
}

func (m *MyInt) add(i MyInt) {
	*m = *m + MyInt(i)
}

type mySliceStruct []MyStruct
func (m mySliceStruct) Less(i, j int) bool{
return m[i].Name < m[j].Name
}
func (m mySliceStruct) Len() int{
return len(m)
}
func (m mySliceStruct) Swap(i, j int){
m[i], m[j] = m[j], m[i]
}
func (m mySliceStruct) Sort() {
sort.Sort(m)
}

func main() {
	i := MyInt(0)

	i.add(3)
	i.showYourSelf()
    var slc mySliceStruct
    slc = make([]MyStruct, 0, 3)
    constructor := NewMyStruct()
slc = append(slc, constructor("Вася"))
slc = append(slc, constructor("Слава"))
slc = append(slc, constructor("Витя"))
fmt.Println(slc)
slc.Sort()
fmt.Println(slc)
}

