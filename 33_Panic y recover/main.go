//Video 33. Panic y recover
//Autor: Judith Valdez Martinez

package main

import (
  "fmt"
)

func imprimir(){
  fmt.Println("Hola Alejando")

  defer func(){
    cadena := recover()
    fmt.Println(cadena)
  }()
  panic("Error")

}

func main(){
  imprimir()
  fmt.Println("Hola Main")
}
