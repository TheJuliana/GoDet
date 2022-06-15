package main

import (
	"C"
	"fmt"
	"net/http"
	"syscall"
	"unsafe"
)

var mod = syscall.NewLazyDLL("libmylib.dll")
var add = mod.NewProc("Addiction")
var inc = mod.NewProc("AddToMas")

func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("dll: ", mod.Name)
	ret, _, err := add.Call(3, 4)
	if err != nil {
		fmt.Fprintf(w, "<a href='/'>Result %d</a>", ret)
	}
}
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("dll: ", mod.Name)
	var numbers [3]int32 = [3]int32{2, 34, 4}
	ret, _, err := inc.Call(uintptr(unsafe.Pointer(&numbers)), 3)
	if err != nil {
		fmt.Fprintf(w, "<a href='/'>Result %d</a>", ret)
	}
}
func main() {
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/about", aboutHandler)
	http.ListenAndServe(":8080", nil)
}
