package matematica

func Media(lista ...float64)float64{
	var total float64

	for _, valor := range lista{
		total += valor
	}

	return total / float64(len(lista))
}