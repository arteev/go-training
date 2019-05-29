package main

// extern void call_go();
import "C"
import "fmt"

func main() {
	C.call_go()
}

//export CallFromC
func CallFromC() {
	fmt.Println("called from C")
}
