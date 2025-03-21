package model

type Bond struct {
	Simbolo      string
	UltimoPrecio float64
	Moneda       string
	Type         string
}

type Quotes struct {
	Quotes      []Bond    `bson:"quotes"`
	Id          string    `bson:"_id"`
	OtherQuotes PreDolars `bson:"otherQuotes" json:"otherQuotes"`
}

type PreDolars struct {
	Quotes Dolars `bson:"quotes" json:"quotes"`
}

type Dolars struct {
	Mep float64 `bson:"dolarMep" json:"dolarMep"`
}
