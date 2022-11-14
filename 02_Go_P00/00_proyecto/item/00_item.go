package item

type GetterDescripcion interface {
	GetDescripcion() string
}

type GetterPrecio interface {
	GetPrecio() float64
}

type GetterId interface {
	GetId() *int
}

type Item interface {
	GetterDescripcion
	GetterPrecio
	GetterId
}

type Items []Item
