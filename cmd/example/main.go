package main

import (
  "flag"
  "fmt"
  "io"
  "os"
  "strings"

  lab2 "github.com/horidor/architectrue-lab-2"
)

var (
  inputExpression = flag.String("e", "", "Expression to compute")
  flagOutput      = flag.String("o", "", "Output file destination (stdout if empty)")
  flagInput       = flag.String("f", "", "Input file path")
)

func main() {
  flag.Parse()

  var Input io.Reader
  var Output io.Writer = os.Stdout

  if len(*flagInput) > 0 && len(*inputExpression) > 0 {
    fmt.Fprintf(os.Stderr, "Specify only one way of input")
    return
  }

  Input = strings.NewReader(*inputExpression)

  if len(*flagInput) > 0 {
    file, err := os.Open(*flagInput)
    if err != nil {
      fmt.Fprintf(os.Stderr, "Specified file is not found")
      return
    }
    defer file.Close()
    Input = file
  }

  if len(*flagOutput) > 0 {
    _, err := os.Stat(*flagOutput)
    if !os.IsNotExist(err) {
      fmt.Fprintf(os.Stderr, "Improper or existing file path given")
    }
    file, err := os.OpenFile(*flagOutput, os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
      fmt.Fprintf(os.Stderr, "Error while creating new file")
    }
    defer file.Close()
    Output = file
  }

  handler := &lab2.ComputeHandler{
    Input:  Input,
    Output: Output,
  }
  
  err := handler.Compute()
  if err != nil {
    fmt.Fprintf(os.Stderr, err.Error())
    return
  }
}