package main

import "fmt"

//Liskov Subsitusion Principle
//menggunakan interface sebagai implementasi

// Kita menggunakan perumpaan paket sebagai interface untuk mewaliki semua tipe jasa kirim
type Paket interface {
	Kirim()
	HitungHarga()
}

// Contoh Provider Jne
type JNE struct {
	Berat              int64
	StatusPaketDikirim bool
}

// Implementasi methos Kirim
func (j *JNE) Kirim() {
	j.StatusPaketDikirim = true
	if j.StatusPaketDikirim {
		fmt.Println("Paket Sudah Dikirim")
	}

	fmt.Println("PaketBelumDikirim")
}

// Implementasi methos HitungHarga pada Interface Paket
func (j *JNE) HitungHarga() {
	fmt.Println(j.Berat * 5000)
}

// Contoh Di Provider Sicepat
type Sicepat struct {
	Berat              int64
	StatusPaketDikirim bool
}

// Implementasi method Kirim
func (s *Sicepat) Kirim() {
	s.StatusPaketDikirim = true
	if s.StatusPaketDikirim {
		fmt.Println("Paket Sudah Dikirim")
	}

	fmt.Println("PaketBelumDikirim")
}

// Implementasi Method HitungHarga
func (s *Sicepat) HitungHarga() {
	fmt.Println(s.Berat * 7000)
}

// Disini kegunaan conthoh kegunaan liskov subsitution Principle
// Sebuah struct hanya perlu mengimplementasi sebuah interface agar bisa di print
func PrintStatusDanHarga(p Paket) {
	p.HitungHarga()
	p.Kirim()
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
