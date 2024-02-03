package main

import "fmt"

func main() {
  var a [4]int
  a[0] = 1
  a[1] = 2
  a[2] = 3
  a[3] = 4

  for j := 0; j < len(a); j++ {
    fmt.Println(a[j])
  }

  // var names [5]string = [5]string{"Nima", "Pema", "Dema", "John", "Doe"}
  names := [5]string{"Nima", "Pema", "Dema", "John", "Doe"}

  for i := 0; i < len(names); i++ {
    fmt.Println(names[i])
  }

  var twoD [5][5]int

  for i := 0; i < len(twoD); i++ {
    for j := 0; j< len(twoD); j++ {
      twoD[i][j] = i + j;
    } 
  }
  fmt.Print("twoD:", twoD)

  
}

