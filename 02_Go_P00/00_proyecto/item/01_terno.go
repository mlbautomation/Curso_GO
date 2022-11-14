package item

// Abstracción a través de estructuras
type Terno struct {
	id     int
	talla  string
	color  string
	precio float64
}

// Encapsulación a través de indentificador no exportado

// Funcion constructora
func NewTerno(id *int, talla string, color string, precio float64) (terno Terno) {
	terno.id = *id + 1
	terno.talla = talla
	terno.color = color
	terno.precio = precio
	*id++
	return terno
}

// Metodos SETTER
func (t *Terno) SetId(id *int)            { t.id = *id }
func (t *Terno) SetTalla(talla string)    { t.talla = talla }
func (t *Terno) SetColor(color string)    { t.color = color }
func (t *Terno) SetPrecio(precio float64) { t.precio = precio }

// Metodos GETTER
func (t *Terno) Id() *int        { return &t.id }
func (t *Terno) Talla() string   { return t.talla }
func (t *Terno) Color() string   { return t.color }
func (t *Terno) Precio() float64 { return t.precio }

// Metodos de la interfaz
func (t *Terno) GetId() *int            { return &t.id }
func (t *Terno) GetDescripcion() string { return "TER" + t.Talla() + t.Color() }
func (t *Terno) GetPrecio() float64     { return t.Precio() }
