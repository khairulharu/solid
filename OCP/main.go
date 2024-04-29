package main

func main() {
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
