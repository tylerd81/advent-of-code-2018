package main

import (
  "fmt"
  "strings"
  "io/ioutil"
)

type GuardData struct {
  Action string
  Guard string
  Minute string
}

func main() {
  lines, err := getInput("test_input.txt")
  if err != 0 {
    fmt.Println("there was an error reading the file")
    return
  }

  for _, line := range lines {
    fmt.Println(line)
  }
  for _, line := range lines {
    guard := parseLine(line)
    fmt.Println(guard)
  }
}

func getInput(fileName string) ([]string, int) {
  data, err := ioutil.ReadFile(fileName)

  if err != nil {
    fmt.Println("Error reading the file")
    return nil, 1
  }
  stringData := string(data)
  stringData = strings.Trim(stringData, "\n ")
  lines := strings.Split(stringData, "\n")
  return lines,  0

}
func parseLine(line string) GuardData {
  var guard GuardData

  if strings.Contains(line, "Guard") == true {
    guard.Action = "Guard Change"
  }else {
    guard.Action = "Other"
  }

  return guard
}
