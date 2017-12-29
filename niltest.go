package main

type T struct{}

func do() *T {
	var err *T
	return err
}

func wrapDo() *T {
	return do()
}

func main() {
	err := do()
	println(err, err == nil)
	err = wrapDo()
	println(err, err == nil)
}
