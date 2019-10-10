
package main

import (
  "fmt"
  "time"
)

type MyError struct{
  When time.Time
  What string
}

type Calc struct{
  result float64
}

func (e *MyError) Error() string{
  return fmt.Sprintf("at %v, %s", e.When, e.What)
}


func (c *Calc) div(x, y float64) (float64, error){
  if y != 0{
    c.result = x/y
    return c.result, nil
  } else{
    return 0, &MyError{time.Now(), "you can't divided zero"}
  }
}

func (c *Calc) mul(x, y float64) float64{
  c.result = x*y
  return c.result
}

func (c *Calc) add(x, y float64) float64{
  c.result = x+y
  return c.result
}

func (c *Calc) sub(x, y float64) float64{
  c.result = x - y
  return c.result
}

func main(){
  //calc := Calc{0}
  calc := new(Calc)
  if value, err := calc.div(1,0); err != nil {
    fmt.Println(err)
  } else{
    fmt.Println(value)
  }
  //fmt.Println(calc.add(1,2))
}
