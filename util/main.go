package main

func main() {
	var p = f()
	println(*p)
	println(p == f())
	len := new([]int)
	print(len)
}

func f() *int {
	v := 1
	return &v
}
