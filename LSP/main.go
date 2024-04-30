package main

import "fmt"

//Liskov Subsitusion Principle
//menggunakan interface sebagai implementasi

// Kita menggunakan perumpaan paket sebagai interface untuk mewaliki semua tipe jasa kirim
type Paket interface {
	Kirim()
	HitungHarga()
}

type JNE struct {
	Harga       int64
	StatusPaket bool
}

func (j *JNE) Kirim() {
	j.StatusPaket = true
	fmt.Printf("Paket sedang di : %v", j.StatusPaket)
}

func (j *JNE) HitungHarga() {
	fmt.Println(j.Harga * 5)
}

type Sicepat struct {
	Harga       int64
	StatusPaket bool
}

func (s *Sicepat) Kirim() {
	s.StatusPaket = true
	fmt.Printf("Paket sedang di : %v", s.StatusPaket)
}

func (s *Sicepat) HitungHarga() {
	fmt.Println(s.Harga * 7)
}

func PrintStatusDanHarga(p Paket) {
	p.HitungHarga()
	p.Kirim()
}
func main() {
}
