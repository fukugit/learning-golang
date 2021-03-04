# 変数について

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
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
