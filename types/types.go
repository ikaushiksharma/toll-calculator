package types
type OBUData struct{
	OBUID int 				`json:"obuID"`
	Latitiude float64		`json:"latitiude"`
	Longitude float64		`json:"longitude"`
}

type Distance struct {
	Value float64 `json:"value"`
	OBUID int     `json:"obuID"`
	Unix  int64   `json:"unix"`
}
