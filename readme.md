Belajar SOLID principle 

Single Responsibilty Principle 

Jadi solid principle hanya memiliki sebuah method tunggal agar struct nya berubah.

Pada Struct Lemari Memiliki Method Masukkan Baju 
Spertin inlah cara penggunakan SOLid Principle yang hanya memiliki satu fungsi agar bisa berubah
```go
// Lemari struct merepensetasikan sebuah Lemari yang sergin kita temukan di rumah
type Lemari struct {
	Nama      string
	IsiLemari interface{}
	Rak       int
}

func LemariBaru() *Lemari {
	return &Lemari{}
}

func (l *Lemari) MasukkanBaju(baju interface{}, nama string, rak int) {
	l.Nama = nama
	l.IsiLemari = baju
	l.Rak = rak
}
```