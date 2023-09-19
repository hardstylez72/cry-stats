package account

import (
	"strings"
)

var Promo = map[string]float64{
	"alf": 10.00, // альфред
	"cry": 10.00, // Босс! Ты нашел меня, теперь я твой, ломай меня полностью
}

func PromoBonus(promo string) float64 {

	s := strings.ToLower(promo)
	v, ok := Promo[s]
	if !ok {
		return 0
	}
	return v
}
