//Funciones parte 2.
//Autor: Judith Valdez Martinez
package main

import (
  "fmt"
)

//Declarar funciones
func imprimirNombre(nombre string){
  fmt.Println("Fuera de Main")
  fmt.Println("El nombre es: ", nombre)
}

func main()  {
  //Llamar una función
  imprimirNombre("Jose")
  fmt.Println("Dentro de main")

  fmt.Println(suma(25,66))
  fmt.Println(resta(88,66))

  //Firma de una función
  fmt.Printf("Función suma: %T\n", suma)
  fmt.Printf("Función resta: %T\n", resta)

  //package math
  func Sin(x float64) float64 //Implementado en lenguaje ensamblador
}
  //Declarar funciones 2
func suma(n1 int, n2 int) int{
  return n1+n2
}

//Declarar funciones 3
func resta(n1,n2 int) (r int){
  r=n1-n2
  return
}
