# Hello world

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
目次

  - [Goの実行](#go%E3%81%AE%E5%AE%9F%E8%A1%8C)
- [学んだこと](#%E5%AD%A6%E3%82%93%E3%81%A0%E3%81%93%E3%81%A8)
  - [とりあえずHello world](#%E3%81%A8%E3%82%8A%E3%81%82%E3%81%88%E3%81%9Ahello-world)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->
<br>

### Goの実行
```
go run main.go
```

<br>

## 学んだこと
### とりあえずHello world
```go
package main

/* 
  importは以下のようにまとめて行います。
 */
import (
  "fmt"
  "time"
)

/* 
  メイン関数です
 */
func main() {
  fmt.Println("hello")
  fmt.Println(time.Now())
}
```

<br>
