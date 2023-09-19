package account

var Promo = map[string]float64{
	"ALF": 10.00,
}

func PromoBonus(promo string) float64 {
	v, ok := Promo[promo]
	if !ok {
		return 0
	}
	return v
}
