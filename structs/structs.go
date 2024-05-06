package structs


type Mensaje struct {
	Code     int         `json:"code"`
	Message  string      `json:"message"`
	Method   string      `json:"method"`
	Resource string      `json:"resource"`
	Errors   []error     `json:"errors"`
	Data     interface{} `json:"data"`
	Pagina   Pagina      `json:"pagina"`
}

// Estructura para la paginación
type Pagina struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Sort   string `json:"sort"`
	Order  string `json:"order"`
}
