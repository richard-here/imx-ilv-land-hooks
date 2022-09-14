package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type Subscription struct {
	ID      uint64 `gorm:"primary_key;auto_increment" json:"id"`
	AssetID uint64 `gorm:"not null" json:"asset_id"`
	User    User   `json:"user"`
	UserID  uint32 `gorm:"not null" json:"user_id"`
}

func (s *Subscription) Prepare() {
	s.ID = 0
	s.User = User{}
}

func (s *Subscription) Validate() error {
	if s.AssetID < 1 {
		return errors.New("required asset ID")
	}
	if s.UserID < 1 {
		return errors.New("required user")
	}
	return nil
}

func (s *Subscription) Subscribe(db *gorm.DB) (*Subscription, error) {
	var err error
	err = db.Debug().Model(&Subscription{}).Create(&s).Error
	if err != nil {
		return &Subscription{}, err
	}
	if s.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", s.UserID).Take(&s.User).Error
		if err != nil {
			return &Subscription{}, err
		}
	}
	return s, nil
}

func (s *Subscription) FindAllSubscriptions(db *gorm.DB) (*[]Subscription, error) {
	var err error
	subscriptions := []Subscription{}
	err = db.Debug().Model(&Subscription{}).Find(&subscriptions).Error
	if err != nil {
		return &[]Subscription{}, err
	}

	if len(subscriptions) > 0 {
		for i := range subscriptions {
			err := db.Debug().Model(&User{}).Where("id = ?", subscriptions[i].UserID).
				Take(&subscriptions[i].User).Error
			if err != nil {
				return &[]Subscription{}, err
			}
		}
	}

	return &subscriptions, nil
}

func (s *Subscription) Unsubscribe(db *gorm.DB, sid uint64, uid uint32) (int64, error) {
	db = db.Debug().Model(&Subscription{}).Where("id = ? and user_id = ?", sid, uid).
		Take(&Subscription{}).Delete(&Subscription{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("subscription to that asset not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
