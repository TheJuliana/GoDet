package main

import (
	"C"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"syscall"
	"unsafe"
)

var mod = syscall.NewLazyDLL("libmylib.dll")
var add = mod.NewProc("Addiction")
var inc = mod.NewProc("AddToMas")
var det = mod.NewProc("Determinant")

func sumHandler(w http.ResponseWriter, r *http.Request) {
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

type Matrix2 struct {
	a11 int32
	a12 int32
	a21 int32
	a22 int32
}

type Page1 struct {
	Determinant uintptr
	Done        bool
	El11        int32
	El12        int32
	El21        int32
	El22        int32
}
type Matrix3 struct {
	a11 int32
	a12 int32
	a13 int32
	a21 int32
	a22 int32
	a23 int32
	a31 int32
	a32 int32
	a33 int32
}
type Page2 struct {
	Determinant uintptr
	Done        bool
	El11        int32
	El12        int32
	El13        int32
	El21        int32
	El22        int32
	El23        int32
	El31        int32
	El32        int32
	El33        int32
}

func matrixHandler2(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("pages/matrix2_page.html")
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}
	r.ParseForm()
	a11, _ := strconv.ParseInt(r.FormValue("a11"), 10, 32)
	a12, _ := strconv.ParseInt(r.FormValue("a12"), 10, 32)
	a21, _ := strconv.ParseInt(r.FormValue("a21"), 10, 32)
	a22, _ := strconv.ParseInt(r.FormValue("a22"), 10, 32)
	mat := Matrix2{
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
	d := Page1{ret, true, numbers[0], numbers[1], numbers[2], numbers[3]}

	tmpl.Execute(w, d)
	if err != nil {
		fmt.Println("DETERMINANT= ", ret)
	}

}

func matrixHandler3(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("pages/matrix3_page.html")
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}
	r.ParseForm()
	a11, _ := strconv.ParseInt(r.FormValue("a11"), 10, 32)
	a12, _ := strconv.ParseInt(r.FormValue("a12"), 10, 32)
	a13, _ := strconv.ParseInt(r.FormValue("a13"), 10, 32)
	a21, _ := strconv.ParseInt(r.FormValue("a21"), 10, 32)
	a22, _ := strconv.ParseInt(r.FormValue("a22"), 10, 32)
	a23, _ := strconv.ParseInt(r.FormValue("a23"), 10, 32)
	a31, _ := strconv.ParseInt(r.FormValue("a31"), 10, 32)
	a32, _ := strconv.ParseInt(r.FormValue("a32"), 10, 32)
	a33, _ := strconv.ParseInt(r.FormValue("a33"), 10, 32)
	mat := Matrix3{
		int32(a11),
		int32(a12),
		int32(a13),
		int32(a21),
		int32(a22),
		int32(a23),
		int32(a31),
		int32(a32),
		int32(a33),
	}
	fmt.Println("333333")
	fmt.Println("a11 = ", mat.a11)
	fmt.Println("a12 = ", mat.a12)
	fmt.Println("a13 = ", mat.a13)
	fmt.Println("a21 = ", mat.a21)
	fmt.Println("a22 = ", mat.a22)
	fmt.Println("a23 = ", mat.a23)
	fmt.Println(mat)
	var numbers = [9]int32{
		mat.a11, mat.a12, mat.a13,
		mat.a21, mat.a22, mat.a23,
		mat.a31, mat.a32, mat.a33,
	}
	ret, _, err := det.Call(uintptr(unsafe.Pointer(&numbers)), 9)
	d := Page2{ret,
		true,
		numbers[0],
		numbers[1],
		numbers[2],
		numbers[3],
		numbers[4],
		numbers[5],
		numbers[6],
		numbers[7],
		numbers[8],
	}
	tmpl.Execute(w, d)
	if err != nil {
		fmt.Println("DETERMINANT= ", ret)
	}
	http.Redirect(w, r, "/matrix3", http.StatusFound)
}
func mainHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("pages/main_page.html")
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}
}
func main() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.HandleFunc("/matrix2", matrixHandler2)
	http.HandleFunc("/matrix3", matrixHandler3)
	http.HandleFunc("/", mainHandler)
	/*http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/main", mainHandler)
	http.HandleFunc("/info", infoHandler)*/
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
