package main

import (
  "fmt"
  "math/cmplx"
)

var (
  ToBe   bool       = false
  MaxInt uint64     = 1<<64 - 1
  z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
  const f = "%T(%v)\n"
  fmt.Printf(f, ToBe, ToBe)
  fmt.Printf(f, MaxInt, MaxInt)
  fmt.Printf(f, z, z)
}

// Tipos básicos
// Os tipos básicos de Go são
// bool
// string
// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr
// byte // pseudônimo para uint8
// rune // pseudônimo para int32
//      // representa um ponto de código Unicode
// float32 float64
// complex64 complex128
// O exemplo mostra vários tipos de variáveis e também que as declarações de variáveis podem ser "construidas" em blocos, como com as declarações de importação.
