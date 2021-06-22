//Video 32. Defer
//Autor: Judith Valdez Martinez
package main

import (
  "fmt"
  "os" //libreria sirve para tener acceso al OS
)

func primera (){
  fmt.Println("Primera")
}
func segunda(){
  fmt.Println("Segunda")
}

func main(){
  defer primera()
  segunda()

  //Abrimos el archivo
  f, err := os.Open("texto.txt")
  //Verificamos que no haya ocurrido ningún error
  if err != nil{
    panic(err)
  }
  //Creamos un Slice para almacenar lo que leemos del archivo
  data := make([]byte,175)
  c, err := f.Read(data)
  //Verificamos que no haya ocurrido ningun error
  if err != nil{
    panic(err)
  }
  //Imprimimos lo leido
  fmt.Printf("Cantidad de byte leidos: %d\n Texto leido: \n%q \nerror: %v", c, data, err)
  f.Close()

  for i := 0; i<5; i++{
    defer fmt.Println(i)
  }
}
