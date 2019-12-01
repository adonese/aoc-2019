package main

func reduceFuel(mass float64, reduced []float64) []float64 {
	if mass <= 0 {
		return reduced
	}
	reduced = append(reduced, calculateFuel(mass))
	return reduceFuel(reduced[len(reduced)-1], reduced)
}

func everyMass(masses [][]float64) float64 {
	var res float64
	for _, v := range masses {
		res += sum(v)
	}
	return res
}

func sum(mass []float64) float64 {
	var sum float64
	for _, v := range mass {
		sum += v
	}
	return sum
}
