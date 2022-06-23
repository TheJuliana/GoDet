package main

import (
	"C"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
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
	var numbers = [4]int32{2, 2, 48, 5}
	ret, _, err := det.Call(uintptr(unsafe.Pointer(&numbers)), 4)
	if err != nil {
		fmt.Fprintf(w, "<a> a11 = %d</a>", numbers[0])
		fmt.Fprintf(w, "<a> a12 = %d</a>", numbers[1])
		fmt.Fprintf(w, "<a> a21 = %d</a>", numbers[2])
		fmt.Fprintf(w, "<a> a22 = %d</a>", numbers[3])
		fmt.Fprintf(w, "<a>Determinant:  %d</a>", ret)
	}
}

type Matrix struct {
	a11 int32
	a12 int32
	a21 int32
	a22 int32
}

func matrixHandler2(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("matrix2_page.html")
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}
	r.ParseForm()
	a11, _ := strconv.ParseInt(r.FormValue("a11"), 10, 32)
	a12, _ := strconv.ParseInt(r.FormValue("a12"), 10, 32)
	a21, _ := strconv.ParseInt(r.FormValue("a21"), 10, 32)
	a22, _ := strconv.ParseInt(r.FormValue("a22"), 10, 32)
	mat := Matrix{
		int32(a11),
		int32(a12),
		int32(a21),
		int32(a22),
	}
	fmt.Println("a11 = ", mat.a11)
	fmt.Println("a12 = ", mat.a12)
	fmt.Println("a21 = ", mat.a21)
	fmt.Println("a22 = ", mat.a22)
	fmt.Println(mat)
	var numbers = [4]int32{mat.a11, mat.a12, mat.a21, mat.a22}
	ret, _, err := det.Call(uintptr(unsafe.Pointer(&numbers)), 4)
	if err != nil {
		fmt.Println("DETERMINANT= ", ret)
	}
}

func getDet(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //Parse url parameters passed, then parse the response packet for the POST body (request body)
	// attention: If you do not call ParseForm method, the following data can not be obtained form
	fmt.Println(r.Form) // print information on server side.
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") // write data to response
}

func matrixHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("matrix2_page.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// logic part of log in
		fmt.Println("a12:", r.Form["a12"])
		fmt.Println("a22:", r.Form["a22"])
	}
}

func main() {
	http.HandleFunc("/", matrixHandler2)
	http.HandleFunc("/matrix", getDet)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/main", mainHandler)
	http.HandleFunc("/info", infoHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
