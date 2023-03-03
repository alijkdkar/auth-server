package Initilzer

import models "github.com/alijkdkar/auth-server-go/Models"

var SecretKey []byte

func InitilzerMethod() {
	ConnectToDB()
	DB.AutoMigrate(&models.MyUser{})
	SecretKey = []byte("̣㕋⬏ꖼ荾٥睻ꚏ媳⣦줨ꤏ蒲퉐⺿폻媒툆䟯隊듃睢㴢㴮ⅲԮ伈샍髵某肣")
}
