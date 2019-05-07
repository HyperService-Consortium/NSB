package main

import "fmt"

type DeferFunc func(ck interface{})
type Environment struct {
	cb *int
}

func (env *Environment) fc(ck interface{}) {
	fmt.Println(env)
	err := recover()
	fmt.Println(ck, "catch", err)
	if err != nil {
		*env.cb = 2
	}
}
func test4() {
	panic("err")
}

func test3(fc DeferFunc) {
	defer fc("test3")
	test4()
}
func test2(env Environment) (cb int) {
	defer env.fc("test2")
	test3(env.fc)
	fmt.Println("gone 2")
	return 3
}
func test() (cb int) {
	var env = Environment{&cb}
	defer env.fc("test")
	ret := test2(env)
	fmt.Println(ret, cb)
	return ret
}

func main() {
	cb := test()
	fmt.Println(cb)
}