# Basic Programming (Variable, Operator, Logical, Function)

---

Tugas praktikum mandiri memiliki tujuan untuk menyelesaikan masalah:

- Menentukan bilangan Prima
- Menentukan bilangan kelipatan 7
- Menghitung luas bangun trapesium

---

## Penyelesaian penentuan bilangan prima

`func IsPrime(num int) bool {
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
} `
