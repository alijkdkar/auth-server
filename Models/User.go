package Models

type MyUser struct {
	ID       uint
	Email    *string `gorm:"unique"`
	Password string
}
