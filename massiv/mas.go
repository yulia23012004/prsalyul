package main

import (
"fmt"
"math/rand"
"time"
)

var size int

func main() {
var a [5][5]int
gen(a)
vivod(a)
stroka := calc(a)
fmt.Println("Строка с самым большим числом", stroka)
fmt.Print(rand.Intn(5))
}
func gen(a [5][5]int) {

for i := 0; i < 5; i++ {
rand.Seed(time.Now().UnixNano())
for j := 0; j < 5; j++ {

a[i][j] = rand.Intn(10)
}
}
}
func vivod(a [5][5]int) {
for i := 0; i < 5; i++ {
for j := 0; j < 5; j++ {
fmt.Print(a[i][j])
fmt.Print(" ")
}
fmt.Print("\n\n")
}
}
func calc(a [5][5]int) int {
var k int
max := 0
for i := 0; i < 5; i++ {
for j := 0; j < 5; j++ {
if a[i][j] > max {
max = a[i][j]
k = i
}
}
}
return k
}