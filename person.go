package reactive

type Person struct {
	Observers map[string]Observer
	Name string

}

func (p *Person) AddObserver(observer Observer){
	if len(p.Observers) == 0 {
		p.Observers = make(map[string]Observer)
	}

	p.Observers[name] = observer
}

func (p *Person) RemoveObserver(name string){
	delete(p.Observers, name)
}

func (p *Person) NotifyObservers(data string){
	for _, v := range p.Observers {
		v.Notify(p.Name)
	}
}
