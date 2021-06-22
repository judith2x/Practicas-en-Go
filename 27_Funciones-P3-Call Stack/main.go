//Funciones parte 3-Call Stack
//Autor: Judith Valdez Martinez
package main

import (
  "fmt"
)

//Call stack

func f1(){
  fmt.Println("Entrando en F1")
  f2()
  fmt.Println("Saliendo de F1")
}
func f2(){
  fmt.Println("Entrando en F2")
  f3()
  fmt.Println("Saliendo de F2")
}
func f3(){
  fmt.Println("Entrando en F3")

  fmt.Println("Saliendo de F3")
}

//Funcion main
func main()  {
  fmt.Println("Entrando en main")
  f1()
  fmt.Println("Saliendo de main")
}
