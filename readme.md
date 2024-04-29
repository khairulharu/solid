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

Terbuka untuk menambahkan service baru tetapi tertutup untuk merubah codingan yang sudah di buat.

Pada saat notification.SendNotification() maka akan me manggil sesuai dengan apa yang di masukkan pada struct tersebut jika Email maka akan return email, begitupun dengan Cart 


Sebagai interface untuk semua tipe implementasi yang ada
```go

type NotificationService interface {
	SendNotification(message *string) (*string, error)
}

```

Email Notification Service

```go
type EmailNotificationService struct{}

func (e *EmailNotificationService) SendNotification(message *string) (*string, error) {

	fmt.Printf("Email Notifcation: %s \n", *message)
	return nil, nil
}

```

Cart Notifcation Service 
```go

type CartNotificationService struct{}

func (c *CartNotificationService) SendNotification(message *string) (*string, error) {
	fmt.Printf("Cart Notifcation: %s \n", *message)
	return nil, nil
}


```

SendNotification struct bertugas sebagai perantara untuk pengiriman segala NotifcationService.SendNotification
```go
type NotificationSender struct {
	notificationService NotificationService
}

func (n *NotificationSender) SendNotification(message *string) (*string, error) {
	return n.notificationService.SendNotification(message)
}

```

```go

func main(){
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
}
```