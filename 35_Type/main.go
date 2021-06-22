//Video 35. Type
//Autor: Judith Valdez Martinez

package main

import (
  "fmt"
)

//Type

//Declarar nuestro propio tipo

//Dinero tipo declarado por nosotros
type Dinero int

//Declaramos el método String para el tipo Dinero
func (d Dinero) String() string{
  return fmt.Sprintf("$%d", d)
}

func main(){
  var sueldo Dinero
  sueldo=25000
  fmt.Println("Sueldo: ", sueldo)

  aumento := 10000
  sueldo += Dinero(aumento)

  fmt.Println("Sueldo + Aumento", sueldo)

}
