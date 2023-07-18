package util

type RateplanReq struct {
	Isdn     string `json:"isdn"`
	Rateplan string `json:"rateplan"`
}

func a() rateplanReq {
	var rer = RateplanReq{
		Isdn:     "",
		Rateplan: "",
	}
}
