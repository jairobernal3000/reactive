package reactive

import "fmt"

type Slack struct {

}

func (s *Slack) Notify(data string){
	fmt.Println("Soy slack","Se ha cambiado el nombre", data)
}