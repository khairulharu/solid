package main

import "fmt"

//Liskov Subsitusion Principle
//menggunakan interface sebagai implementasi

// Kita menggunakan perumpaan paket sebagai interface untuk mewaliki semua tipe jasa kirim
type Paket interface {
	Kirim() string
	HitungHarga() int64
}

// Contoh Provider Jne
type JNE struct {
	Berat              int64
	StatusPaketDikirim bool
}

// Implementasi methos Kirim
func (j *JNE) Kirim() string {
	j.StatusPaketDikirim = true

	if j.StatusPaketDikirim {
		return fmt.Sprint("Paket Sudah Dikirim")
	}

	return fmt.Sprint("PaketBelumDikirim")
}

// Implementasi methos HitungHarga pada Interface Paket
func (j *JNE) HitungHarga() int64 {
	return j.Berat * 7000
}

// Contoh Di Provider Sicepat
type Sicepat struct {
	Berat              int64
	StatusPaketDikirim bool
}

// Implementasi method Kirim
func (s *Sicepat) Kirim() string {
	s.StatusPaketDikirim = true
	if s.StatusPaketDikirim {
		return fmt.Sprint("Paket Sudah Dikirim")
	}

	return fmt.Sprintf("PaketBelumDikirim")
}

// Implementasi Method HitungHarga
func (s *Sicepat) HitungHarga() int64 {
	return s.Berat * 7000
}

// Disini kegunaan conthoh kegunaan liskov subsitution Principle
// Sebuah struct hanya perlu mengimplementasi sebuah interface agar bisa di print
func PrintStatusDanHarga(p Paket) {
	fmt.Println("Harga: ", p.HitungHarga())
	fmt.Println("Status Pengiriman: ", p.Kirim())
}
func main() {
	//Inisiate Sebuah variable Denga Struct JNE
	jne := JNE{
		StatusPaketDikirim: false,
		Berat:              9,
	}

	//Disin hanya perlu memasukkan struct yang sudah mengimplementasi interface Paket
	PrintStatusDanHarga(&jne)
}
