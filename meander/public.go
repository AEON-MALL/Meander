package meander

type Facade interface {
	Public() any
}

func Public(o any) any {
	if p, ok := o.(Facade); ok {
		return p.Public()
	}
	return o
}
