        package main

        import (
            "fmt"
            "myproject/mypackage" // Import your local package
        )

        func main() {
            mypackage.Hello() // Call an exported function from your package
            fmt.Println(mypackage.ExportedVar) // Access an exported variable
        }
