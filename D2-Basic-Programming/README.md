# Basic Programming (Variable, Operator, Logical, Function)

---

Tugas praktikum mandiri memiliki tujuan untuk menyelesaikan masalah:

- Menentukan bilangan Prima
- Menentukan bilangan kelipatan 7
- Menghitung luas bangun trapesium

---

## Penyelesaian penentuan bilangan prima

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

> 1. Penyelesaian menggunakan pendekatan dengan melakukan perkondisian untuk menentukan angka 2 dan 3 sebagai bilangan prima
> 2. Selanjutnya kondisi digunakan untuk menyeleksi angka 1 dan angka di bawah 25 apakah habis dibagi 2 atau 3 yang berarti bukan bilangan prima
> 3. Perulangan digunakan untuk mengecek bilangan dengan mengambil angka 5 kemudian ditambakan dengan 6 untuk setiap perulangan. Pada setiap perulangan kode akan mengecek apakah angka habis dibagi 5, 7, 11, 13, dst.
> 4. Jika angka berhasil melewati semua pengecekan maka angka tersebut merupakan bilangan prima dan kode akan mengirimkan nilai kembalian _True_.

---

## Penyelesaian penentuan bilangan kelipatan 7

    func IsMutipleOfSeven(num int) bool {
        return num%7 == 0
    }

> Kode akan mengecek apakah angka yang diterima function merupakan nilai kelipatan 7 kemudian mengembalikan nilai _True_ atau _False_ . Perhitungan menggunakan operator _modulo_ untuk angka yang diberikan ke function dilakukan operasi modulo 7.

---

## Penyelesaian perhitungan luas bidang trapesium

    func AreaOfTrapezoid(a, b, t float64) float64 {
        area := (a + b) * t / 2
        return area
    }

> Kode akan menghitung luas trapesium dengan menggunakan function yang berisi rumus trapesium kemudian mengembalikan nilai luas. Tipe data yang digunakan adalah _float_ untuk mengakomodasi nilai desimal.
