package main

import "fmt"

func main() {
	fmt.Println("Hello world")
}

/*
how do we run the code in the project ?
 - go run file_name.go => Compiles and immediately executes one or two files
 - go build => Compiles Go program
 - go fmt => Formats all the code in each file in the current directory
 - go install => Compiles and installs a pkg
 - go get => Downloads the raw source code of someone else's kg
 - go test => Runs any test associated with current project
====================================
what does "package main" mean ?
 - The work "main" is used to make an executable type package - executable pkg => Must have a func called 'main'
 - Other custome pkg name is used to defines a pkg that can be used as a dependency (helper code) - reusuable pkg
====================================
what does "import "fmt" " mean ?
 - Give our pkg access to code written in another pkg
 - fmt: standard library of go
====================================
What's that 'func' thing ?
 - func: tell go to create a function
 - main: function name
====================================
How is the main.go file organized
 package main: Package declaration
 import "fmt": Import other pkgs that we need
 func main() {}: Declare functions
*/
