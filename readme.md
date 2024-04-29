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


OCP
Open Close Principle

Terbuka untuk menambahkan service baru tetapi tertutup untuk merubah codingan yang sudah di buat 

```go

//Buat variable untuk message yang akan di kirim
	message := "HALO"

	//inisiate struct EmailNotificationService yang sudah mengimplementasi interface Notification
	emailNotificationService := &EmailNotificationService{}

	//Cart juga sudah memiliki method yang sama pada interface  Notification
	cartNotficationService := &CartNotificationService{}

	//kita Panggil Struct Notification Service yang berisi interface tadi
	//Dan pada bagian notificaionService diisi dengan emailNotificationService
	notificationService := &NotificationSender{notificationService: emailNotificationService}

	//Sekarang notification service akan berisi emailNotificationservice
	notificationService.SendNotification(&message)
	//Output yang di keluarkan "Email Notifcation: HALO"

	//Pada saat Seperti ini kita tidak harus Menginisiate ulang Notification Service karena kenapa?
	//KArena cart Notification memiliki method ynag sama jadi kita hanya harus mengubah
	//notificationSerivce pada struct nya, tanpa harus mengubah keseluruhan kode
	//JIka pada saat nanti ingin membuat Service baru dengan contoh
	//MessageNotifcaionService dia hanya perlu implement interface NotificationService
	notificationService.notificationService = cartNotficationService

	//Memnaggil untuk print Message
	notificationService.SendNotification(&message)
	// Cart Notifcation: HALO
```