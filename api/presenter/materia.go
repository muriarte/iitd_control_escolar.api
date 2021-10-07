package presenter

//Student data
type Materia struct {
	ID            int    `json:"id"`
	Nombre        string `json:"nombre"`
	Observaciones string `json:"observaciones"`
	Activo        string `json:"activo"`
}
