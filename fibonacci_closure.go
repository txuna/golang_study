package main

import "fmt"
//return f(i-1) * f(i-2)
func fibonacci() func() uint64{
  var fn1 uint64 = 0
  var fn0 uint64 = 0
  return func() uint64{
    fn2 := fn1 + fn0
    fn1, fn0 = fn2, fn1
    if fn2 == 0{
      fn1 = 1
    }
    return fn2
  }
}

/*
fn2 = 1
fn1 = 1
fn0 = 1
*/

func main(){
  f := fibonacci()
  for i:=0;i<50;i++{
    fmt.Println(f())
  }
}
