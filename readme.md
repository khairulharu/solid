# Belajar SOLID principle 

Single Responsibilty Principle 

Jadi Single Responsibility Principle hanya memiliki sebuah method tunggal agar struct nya berubah,
Seperti namanya single, hanya memiliki satu buah cara agar si struct atau class nya berubah,
Struct nya tidak harus berubah, bisa saja method itu hanya mengembalikan sebuah nilai tapi dengan catatan,
Hanya boleh memiliki satu fungsional atau kegunaan saja.

Pada Struct Lemari Memiliki Method Masukkan Baju 
Spertin inlah cara penggunakan SOLid Principle yang hanya memiliki satu fungsi agar bisa berubah
Contoh penggunaan SRP pada struct Lemari
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

Pada Penggunaan Strutc Lemari Seperti ini
```go

func main() {
	lemariAnton := LemariBaru()
	//Lemari anton hanya boleh memasukkan baju ke dalam sebuah lemari
	lemariAnton.MasukkanBaju("louis vuitton", "lemari anton", 1)

	fmt.Println(lemariAnton.IsiLemari)
}
```