package main
import (
  "net/http"
  "os"
  "fmt"
  "flag"
)

var port = flag.String("p", "8080", "specify the port number")

func main() {
  flag.Parse()
  cwd, _ := os.Getwd()
  fmt.Println("Serving files from", cwd)
  fmt.Println("Listening on port", *port)
  serving := ":" + *port
  panic(http.ListenAndServe(serving, http.FileServer(http.Dir(cwd))))
}
