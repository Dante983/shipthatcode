package main

import "fmt"

func main() {
    var first, second string
    fmt.Scan(&first)
    fmt.Scan(&second)
	fmt.Println("Hello,", first)
	fmt.Println("Hello,", second)
    // Print one greeting per line, first name first.
}
