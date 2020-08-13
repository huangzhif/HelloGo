package utils

type Tcconfigs struct {
	Tcconfigs []Tcconfig `json:TCconfigs`
}

type Tcconfig struct {
	AK       string `json:"AK"`
	SK       string `json:"SK"`
	Mainpart string `json:"mainpart"`
}

