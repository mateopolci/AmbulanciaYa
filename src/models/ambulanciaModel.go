package models

// Modelo
type Ambulancia struct {
	Id              string `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Patente         string `json:"patente" gorm:"type:varchar(20);not null"`
	Inventario      bool   `json:"inventario" gorm:"type:boolean;not null"`
	Vtv             bool   `json:"vtv" gorm:"type:boolean;not null"`
	Seguro          bool   `json:"seguro" gorm:"type:boolean;not null"`
	ChoferId        string `json:"choferId" gorm:"column:choferid;type:uuid;not null"`
	ParamedicoId    string `json:"paramedicoId" gorm:"column:paramedicoid;type:uuid;not null"`
	Base            bool   `json:"base" gorm:"type:boolean;not null"`
	Cadenas         bool   `json:"cadenas" gorm:"type:boolean;not null"`
	Antinieblas     bool   `json:"antinieblas" gorm:"type:boolean;not null"`
	CubiertasLluvia bool   `json:"cubiertasLluvia" gorm:"column:cubiertaslluvia;type:boolean;not null"`
}

// DTO
type AmbulanciaDTO struct {
	Id              string `json:"id"`
	Patente         string `json:"patente"`
	Inventario      bool   `json:"inventario"`
	Vtv             bool   `json:"vtv"`
	Seguro          bool   `json:"seguro"`
	ChoferId        string `json:"choferId"`
	ParamedicoId    string `json:"paramedicoId"`
	Base            bool   `json:"base"`
	Cadenas         bool   `json:"cadenas"`
	Antinieblas     bool   `json:"antinieblas"`
	CubiertasLluvia bool   `json:"cubiertasLluvia"`
}

// Método DTO de ambulancia
func (a *Ambulancia) AmbulanciaToDTO() AmbulanciaDTO {
	return AmbulanciaDTO{
		Id:              a.Id,
		Patente:         a.Patente,
		Inventario:      a.Inventario,
		Vtv:             a.Vtv,
		Seguro:          a.Seguro,
		ChoferId:        a.ChoferId,
		ParamedicoId:    a.ParamedicoId,
		Base:            a.Base,
		Cadenas:         a.Cadenas,
		Antinieblas:     a.Antinieblas,
		CubiertasLluvia: a.CubiertasLluvia,
	}
}

//Especificacion del nombre de la tabla para GORM
func (Ambulancia) TableName() string {
	return "ambulancias"
}

// DTO para service de descripcion de Ambulancias
type AmbulanciaDescDTO struct {
	Id              string `json:"id"`
	Patente         string `json:"patente"`
	Inventario      bool   `json:"inventario"`
	Vtv             bool   `json:"vtv"`
	Seguro          bool   `json:"seguro"`
	Chofer          string `json:"chofer"`
	Paramedico      string `json:"paramedico"`
	Base            bool   `json:"base"`
	Cadenas         bool   `json:"cadenas"`
	Antinieblas     bool   `json:"antinieblas"`
	CubiertasLluvia bool   `json:"cubiertasLluvia"`
}

// DTO para pedido de ambulancia
type AmbulanciaPedidoDTO struct {
	Nombre      string `json:"nombre"`
	Telefono    string `json:"telefono"`
	Direccion   string `json:"direccion"`
	Descripcion string `json:"descripcion"`
}
