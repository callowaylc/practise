package main

import(
  "log"
  "encoding/json"
  "io/ioutil"
)

type Record struct {
  Id string `json:"id"`
  Value string `json:"value"`
}

func main() {
  r := Record{}

  contents, err := ioutil.ReadFile("./json-example.json")
  if err != nil {
    log.Println(err)
  }

  json.Unmarshal(contents, &r)
  log.Printf("%+v", r)
}
