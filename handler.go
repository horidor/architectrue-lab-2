package lab2

import (
  "io"
  "strings"
)

type ComputeHandler struct {
  Input  io.Reader
  Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
  prefix := new(strings.Builder)
  io.Copy(prefix, ch.Input)

  res, err := PrefixToPostfix(prefix.String())
  if err != nil {
    return err
  }

  ch.Output.Write([]byte(res))
  return nil
}