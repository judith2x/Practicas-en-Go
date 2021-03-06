//Funciones Parte 5 Variadic Closure
//Autor: Judith Valdez Martinez
package main

import (
  "fmt"
)

//Closure
func print(cadena string){
  fmt.Println(cadena)
}
func print2(cadena string){
  fmt.Println(cadena)
}
func print3(cadena1, cadena2 string){
  fmt.Println(cadena1+cadena2)
}

func print4(fprint func(string)){
  fprint("Hola mundo desde print4")
}

func incrementar() func() int{
  i := 0
  return func()(r int){
    r = i
    i += 2
    return
  }
}

func incremento(){
  i := 0
  i++
  fmt.Println(i)
}



func main(){
  cadena := "Hola mundo"

  imprimir := print

  imprimir(cadena)

  imprimir2 := func ()  {
    fmt.Println(cadena)
  }
  imprimir2()

  imprimir=print2
  imprimir("Hola mundo 2")

  fmt.Printf("Función print1: %T\n", print)
  fmt.Printf("Función print2: %T\n", print2)
  fmt.Printf("Función print3: %T\n", print3)

  print4(print)

//Las funciones son comparadas con nil
  var fb func()
  if fb==nil {
      fmt.Println("fb es igual a nil")
  }

  inc := incrementar()

  fmt.Println("Valor de i; ", inc())
  fmt.Println("Valor de i; ", inc())
  fmt.Println("Valor de i; ", inc())
  fmt.Println("Valor de i; ", inc())

  incremento()
  incremento()
  incremento()

}
