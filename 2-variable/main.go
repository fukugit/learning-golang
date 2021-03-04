package main

import (
  "fmt"
)

func main() {
  // 明示的に型指定
  var i int = 100
  fmt.Println(i)
  
  var str string = "test"
  fmt.Println(str)

  // 変数をまとめて指定 その1
  var t, f bool = true, false
  fmt.Println(t)
  fmt.Println(f)

  // 変数をまとめて指定 その2
  var (
    i2 int = 200
    s2 string = "test2"
  )
  fmt.Println(i2, s2)

  // 暗黙的な定義
  i4 := 400
  s4 := "test4"
  fmt.Println(i4, s4)
}