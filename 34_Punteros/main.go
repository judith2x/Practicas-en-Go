//Video 34. Punteros
//Autor: Judith Valdez Martinez

package main

import (
  "fmt"
)

func incrementar (numero *int){
  *numero++
  fmt.Println("Valor de numero dentro de la función incrementar: ", *numero)

}

func main(){
  //a es de tipo int
  a := 25
  fmt.Println("Valor de a: ", a)
  fmt.Println("Dirección de a: ", &a)
  //b es un puntero de a, por lo que es de tipo int
  b := &a
  //se imprimir el contenido de b
  fmt.Println(b)
  //Se imprime el contenido de la varible a, que es donde apunta b
  fmt.Println(*b)
  //Le asignamos un nuevo valor a (a) a traves de b
  *b=32
  fmt.Println("Valor de a: ", a)
  a++
  //imprimimos *b
  fmt.Println("Valor al que apunta b: ", *b)

  //Valor 0 de los punteros es nil
  if b!=nil{
    fmt.Println("b es diferente de nil")
  }

  //Los punteros son comparables
  c := &a

  if b==c{
    fmt.Println("b y c son iguales")
  }

  //utilizar new()
  d := new(int)//*int
  fmt.Println("Dirección de d: ", d)
  fmt.Println("Valor de d: ", *d)
  d = b
  fmt.Println("Direccion de d:", d)
  fmt.Println("Valor de d: ", *d)
  fmt.Println("Valor de a: ", a)
  fmt.Println("Valor de b: ", *b)
  fmt.Println("Valor de c: ", *c)

  //Usar punteros en funciones
  numero := 2
  fmt.Println("Numero antes de incrementar: ", numero)
  incrementar(&numero)
  fmt.Println("Numero despues de incrementar: ", numero)

}
