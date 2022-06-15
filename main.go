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
var det = mod.NewProc("Determinant")

func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("dll: ", mod.Name)
	ret, _, err := add.Call(3, 4)
	if err != nil {
		fmt.Fprintf(w, "<a href='/'>Sum: %d</a>", ret)
	}
}
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("dll: ", mod.Name)
	var numbers [3]int32 = [3]int32{2, 34, 4}
	ret, _, err := inc.Call(uintptr(unsafe.Pointer(&numbers)), 3)
	if err != nil {
		fmt.Fprintf(w, "<a href='/'>Sums: %d</a>", ret)
	}
}
func infoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("dll: ", mod.Name)
	var numbers [4]int32 = [4]int32{2, 2, 4, 5}
	ret, _, err := det.Call(uintptr(unsafe.Pointer(&numbers)), 4)
	if err != nil {
		fmt.Fprintf(w, "<a href='/'>Determinant:  %d</a>", ret)
	}
}

func main() {
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/info", infoHandler)
	http.ListenAndServe(":8080", nil)
}
