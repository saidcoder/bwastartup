package user

import "gorm.io/gorm"

type Repository interface {
	//sebuah contrant antara repository dan entity
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRespository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error

	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) FindByEmail(email string) (User, error) {
	var user User
	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
