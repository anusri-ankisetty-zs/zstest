// Hello Word in Go by Vivek Gite
package main

// Import OS and fmt packages
import (
	"fmt"
	"os"
)

// Let us start
func main() {
	fmt.Println("Hello, world!")                          // Print simple text on screen
	fmt.Println(os.Getenv("USER"), ", Let's be friends!") // Read Linux $USER environment variable
<<<<<<< HEAD
	fmt.Println("This is Anusri")
=======
	fmt.Println("Hello")
>>>>>>> 376a96e2a8998b8c23fcd270fecb714a3101b1cd
}
