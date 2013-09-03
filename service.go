package main

import (
    "github.com/loz/getapubcrawl/pubcrawl"
    "fmt"
    "net/http"
    "os"
)

func main() {
    http.HandleFunc("/", handler)
      fmt.Println("listening...")
      err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
      if err != nil {
        panic(err)
      }
}

func handler(w http.ResponseWriter, r *http.Request) {
  location := r.URL.Query().Get("Body")
  pubs := pubcrawl.FindPubs(location)
  fmt.Println("Found: ", pubs)
  route, distance := pubcrawl.FindRoute(pubs)
  writePretendyXMLStart(w)
  writeResults(w, location, route, distance)
  writePretendyXMLEnd(w)
}

func writePretendyXMLStart(w http.ResponseWriter) {
  fmt.Fprint(w, "<Response><Sms>")
}

func writePretendyXMLEnd(w http.ResponseWriter) {
  fmt.Fprint(w, "</Sms></Response>")
}

func writeResults(w http.ResponseWriter, location string, route []*pubcrawl.Pub, distance float64) {
  for _, pub := range route {
    fmt.Fprintf(w, "%v, ", pub.Name())
  }
  fmt.Fprintf(w, "%.1fkm", distance)
}

