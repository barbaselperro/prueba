package main

/*Realizar la carga del lado de un cuadrado, mostrar por pantalla el perímetro del 
mismo (El perímetro de un cuadrado se calcula multiplicando el valor del lado por cuatro)"*/

import "fmt"

func main() {
	var lado, perimetro int;
	fmt.Println("escriba el valor del lado del cuadrado: ")
	fmt.Scan(&lado)
	perimetro = lado * 4
	fmt.Println("el perimetro del cuadrado es: ",perimetro)
}

