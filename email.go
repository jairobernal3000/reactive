package reactive

import "fmt"

type Email struct {

}

func (e *Email) Notify(data string){
	//http post servicio
	fmt.Println("soy email, ha cambiado el nombre", data)
}