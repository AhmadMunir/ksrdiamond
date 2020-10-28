package domain

type Diamond struct {
	Id    int8   `json:"id_diamond"`
	Name  string `json:"name"`
	Dm    string `json:"diamond"`
	Price int64  `json:"price"`
}

type Diamonds struct {
	Diamonds []Diamond `json:"diamond"`
}
