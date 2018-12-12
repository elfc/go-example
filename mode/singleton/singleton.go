package singleton

import (
	"fmt"
	"sync"
)

type singleton struct {
	Age int
}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}

func main() {
	insA := GetInstance()
	insA.Age = 1
	insB := GetInstance()
	insB.Age = 2
	fmt.Println(&insA.Age, insA)
	fmt.Println(&insB.Age, insB)
	fmt.Printf("%p %T\n", insA, insA)
	fmt.Printf("%p %T\n", insB, insB)
}
