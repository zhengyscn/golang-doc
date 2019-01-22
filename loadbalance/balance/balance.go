package balance

type Balance interface {
	DoBalance([]*Instnace) (*Instnace, error)
}
