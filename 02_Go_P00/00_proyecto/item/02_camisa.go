package item

// Abstracción a través de estructuras
type Camisa struct {
	id     int
	talla  string
	color  string
	precio float64
}

// Encapsulación a través de indentificador no exportado

// Funcion constructora
func NewCamisa(id *int, talla string, color string, precio float64) (camisa Camisa) {
	camisa.id = *id + 1
	camisa.talla = talla
	camisa.color = color
	camisa.precio = precio
	*id++
	return camisa
}

// Metodos SETTER
func (c *Camisa) SetId(id *int)            { c.id = *id }
func (c *Camisa) SetTalla(talla string)    { c.talla = talla }
func (c *Camisa) SetColor(color string)    { c.color = color }
func (c *Camisa) SetPrecio(precio float64) { c.precio = precio }

// Metodos GETTER
func (c *Camisa) Id() *int        { return &c.id }
func (c *Camisa) Talla() string   { return c.talla }
func (c *Camisa) Color() string   { return c.color }
func (c *Camisa) Precio() float64 { return c.precio }

// Metodos de la interfaz
func (c *Camisa) GetId() *int            { return &c.id }
func (c *Camisa) GetDescripcion() string { return "CAM" + c.Talla() + c.Color() }
func (c *Camisa) GetPrecio() float64     { return c.Precio() }
