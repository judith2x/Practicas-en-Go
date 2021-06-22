//Video 36. Estructuras parte 1
//Autor: Judith Valdez Martinez

package main

import (
  "fmt"
)

type Persona struct{
  Nombre string
  Edad int
}

//Persona: estructura que representa una persona
func Older(p1, p2 Persona) (Persona, int){
  if p1.Edad > p2.Edad {
    return p1, p1.Edad - p2.Edad
  }
  return p2, p2.Edad - p1.Edad
}

func main(){
  //Declarar un variable de tipo Persona #1
  var p Persona
  p.Nombre="Alejandro"
  p.Edad=29
  fmt.Println("Estructura p de tipo persona: ", p)
  fmt.Println("Nombre p: ",p.Nombre)
  fmt.Println("Edad p: ", p.Edad)

  //Declarar una variables de tipo Persona #2
  p2 := Persona{Nombre: "Rafael", Edad: 25}
  fmt.Println("Nombre p2: ", p2.Nombre)
  fmt.Println("Edad p2: ", p2.Edad)

  //Declarar una variable de tipo Persona #3
  p3 := Persona{"Miguel", 18}
  fmt.Println("Nombre p3: ", p3.Nombre)
  fmt.Println("Edad p3: ", p3.Edad)

  //Inicializamos 3 variables de tipo persona
  tom := Persona{"Tom", 60}
  bob := Persona{"Bob", 25}
  paul := Persona{"Paul", 43}
  //Llamamos la función Order
  tbOlder, tbDiff := Older(tom, bob)
  tbOlder, tbDiff := Older(tom, paul)
  tbOlder, tbDiff := Older(bob, paul)

  //Imprimimos los resultados
  fmt.Printf("Entre %s y %s, %s es mayor que %d años\n",
  tom.Nombre, bob.Nombre, tbOlder.Nombre, tbDiff)

  fmt.Printf("Entre %s y %s, %s es mayor que %d años\n",
  tom.Nombre, paul.Nombre, tbOlder.Nombre, tbDiff)

  fmt.Printf("Entre %s y %s, %s es mayor que %d años\n",
  bob.Nombre, paul.Nombre, tbOlder.Nombre, tbDiff)
}
