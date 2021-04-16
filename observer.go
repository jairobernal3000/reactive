package reactive

type Observer interface {
	Notify(data string)
}