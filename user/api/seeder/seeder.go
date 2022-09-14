package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/richard-here/imx-ilv-land-hooks/user/models"
)

func Load(db *gorm.DB) {
	err := db.Debug().DropTableIfExists(&models.Subscription{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}

	err = db.Debug().AutoMigrate(&models.Subscription{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("cannot migrate tabel : %v", err)
	}

	err = db.Debug().Model(&models.Subscription{}).AddForeignKey("user_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}
}
