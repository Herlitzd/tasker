// Writing a basic HTTP server is easy using the
// `net/http` package.
package main

import "fmt"

func main() {
	ls := Task{Name: "Do List", Program: "ls"}
	pwd := Task{Name: "Do Working Directory", Program: "pwd"}
	p := Pipeline{Task: &ls, OnSuccess: &Pipeline{Task: &pwd}}
	p.Begin()
	fmt.Print(p.String())
}
