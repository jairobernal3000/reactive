package reactive

type Observable interface {
	AddObserver(name string, observer Observer)
	RemoveObserver(name string)
	NotifyObservers()
}

