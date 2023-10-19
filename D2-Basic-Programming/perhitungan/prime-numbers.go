package perhitungan

func IsPrime(num int) bool {
	if num > 1 && num <= 3 {
		return true
	}
	if num <= 1 || num%2 == 0 || num%3 == 0 {
		return false
	}

	for i := 5; i*i <= num; i += 6 {
		if num%i == 0 || num%(i+2) == 0 {
			return false
		}
	}
	return true
}
