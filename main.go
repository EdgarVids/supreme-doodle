package main

import "fmt"
//import "net/http"
import "os"

func main() {
    token := os.Getenv("SPACETRADERS")
    fmt.Println("Space Traders Client Practice")
    fmt.Println("Token:", token)


}
