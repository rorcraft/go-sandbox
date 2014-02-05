package main
import (
  "net/http"
  "os"
  "fmt"
)
func main() {
  cwd, _ := os.Getwd()
  fmt.Println("Serving files from ", cwd)
  panic(http.ListenAndServe(":8080", http.FileServer(http.Dir(cwd))))
}
