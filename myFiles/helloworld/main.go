// go run executes files- compiles and executes
// go build- a bunch of source files- compiles but doesn't execute
// go fmt- formats all the code in each file in the current directory
// go install- compiles and installs a package
// go get- downloads the raw source code of someone else's package
// go test- runs any tests associated with the current project

// package is a where all the files belong
// main is used to make an executeable package- makes runable file
// MUST have func called main

// package calculator- defines a package that can be used as a dependency (helper code)
// package uploader- defines a package that can be used as a dependency (helper code)
package main

// access to some code in another library
// fmt- format package
// main has no access to other packages outside main unless imported
import "fmt"

func main() {
	fmt.Println("Hi there!")
}
