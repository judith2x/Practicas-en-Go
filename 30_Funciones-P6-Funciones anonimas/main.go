//Funciones Parte 6 Funciones anonimas
//Autor: Judith Valdez Martinez
package main

import (
  "fmt"
  "strings"
)

func main(){
  cadena := "123456789"

  cadena = strings.Map(func(r rune) rune{
    return r+1
  }, cadena)
  fmt.Println(cadena)
}
