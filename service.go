package main

import (
    "github.com/loz/getapubcrawl/pubcrawl"
    "fmt"
    "net/http"
    "encoding/xml"
    "os"
)

type Response struct {
  Sms string
}

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
  sms := stringResults(location, route, distance)
  response := Response{sms}
  content, err := xml.MarshalIndent(response, "", "  ")
  if err == nil {
    fmt.Fprint(w, string(content))
  }
}

func writePretendyXMLStart(w http.ResponseWriter) {
  fmt.Fprint(w, "<Response><Sms>")
}

func writePretendyXMLEnd(w http.ResponseWriter) {
  fmt.Fprint(w, "</Sms></Response>")
}

func stringResults(location string, route []*pubcrawl.Pub, distance float64) string {
  results := ""
  for _, pub := range route {
    results += pub.Name() + ", "
  }
  results += fmt.Sprintf("%.1fkm", distance)
  return results
}

