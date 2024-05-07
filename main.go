package main

func main() {
}

func complexCalculation(total, tax float64) float64 {
	if total > 100 {
		return total * tax
	}
	if total > 50 {
		return total * tax * 0.5
	}
	return total * tax * 0.1
}
