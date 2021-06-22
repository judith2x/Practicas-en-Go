//Funciones parte 1.
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

}
  //Declarar funciones 2
func suma(n1 int, n2 int) int{
  return n1+n2
}
