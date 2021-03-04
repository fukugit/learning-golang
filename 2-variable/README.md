# 変数について

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
目次

- [学んだこと](#%E5%AD%A6%E3%82%93%E3%81%A0%E3%81%93%E3%81%A8)
  - [明示的に型指定](#%E6%98%8E%E7%A4%BA%E7%9A%84%E3%81%AB%E5%9E%8B%E6%8C%87%E5%AE%9A)
  - [変数をまとめて指定 その1](#%E5%A4%89%E6%95%B0%E3%82%92%E3%81%BE%E3%81%A8%E3%82%81%E3%81%A6%E6%8C%87%E5%AE%9A-%E3%81%9D%E3%81%AE1)
  - [変数をまとめて指定 その2](#%E5%A4%89%E6%95%B0%E3%82%92%E3%81%BE%E3%81%A8%E3%82%81%E3%81%A6%E6%8C%87%E5%AE%9A-%E3%81%9D%E3%81%AE2)
  - [暗黙的な定義](#%E6%9A%97%E9%BB%99%E7%9A%84%E3%81%AA%E5%AE%9A%E7%BE%A9)
  - [](#)
  - [](#-1)
  - [](#-2)
  - [](#-3)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->
<br>

## 学んだこと
### 明示的に型指定
```go
func main() {
  var i int = 100
  fmt.Println(i)
  
  var str string = "test"
  fmt.Println(str)
}
```
<br>

### 変数をまとめて指定 その1
```go
func main() {
  var t, f bool = true, false
  fmt.Println(t)
  fmt.Println(f)
}
```
<br>

### 変数をまとめて指定 その2
```go
func main() {
  var (
    i2 int = 200
    s2 string = "test2"
  )
  fmt.Println(i2, s2)
}
```
<br>

### 暗黙的な定義
```go
func main() {
  // 暗黙的な定義は再定義はNG
  // 暗黙的な定義は関数外で定義はNG
  i4 := 400
  s4 := "test4"
  fmt.Println(i4, s4)
}
```
<br>

### 
```go
```
<br>

### 
```go
```
<br>

### 
```go
```
<br>

### 
```go
```
<br>
