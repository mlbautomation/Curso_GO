package item

type GetterId interface {
	GetId() int
}

type GetterDescripcion interface {
	GetDescripcion() string
}

type GetterPrecio interface {
	GetPrecio() float64
}

type Item interface {
	GetterDescripcion
	GetterPrecio
	GetterId
}

type Items []Item
