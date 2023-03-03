package Initilzer

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	//postgres://yjaaeohv:I1nOGXXIhvJ94tz1LZiK0OK4JAxsOyK6@trumpet.db.elephantsql.com/yjaaeohv
	dsn := "host=trumpet.db.elephantsql.com user=yjaaeohv password=I1nOGXXIhvJ94tz1LZiK0OK4JAxsOyK6 dbname=yjaaeohv sslmode=disable "
	dbname, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Faild To Connect to the DB")
	}
	DB = dbname

}
