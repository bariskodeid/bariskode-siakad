package repositories

import (
	"database/sql"
	"time"

	"github.com/bariskodeid/bariskode-siakad/siakad-service-auth/src/models"
	"gorm.io/gorm"
)

type UserRepository struct {
    DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
    return &UserRepository{DB: db}
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
    var user models.User
    if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *UserRepository) FindById(id string) (*models.User, error) {
	var user models.User
	if err := r.DB.Where("uuid = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	if err := r.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByPhone(phone string) (*models.User, error) {
	var user models.User
	if err := r.DB.Where("phone = ?", phone).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Logout(uuid string) (*models.User, error) {
	var user models.User
	if err := r.DB.Where("uuid = ?", uuid).First(&user).Error; err == nil {
		user.Status = "inactive"
		user.AccessToken = ""
		user.AccessTokenExpiredAt = sql.NullTime{}
		user.UpdatedAt = time.Now()
		if err := r.DB.Save(&user).Error; err != nil {
			return nil, err
		}
		return &user, nil
	}
	return &user, nil
}

func (r *UserRepository) Create(user *models.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepository) Update(user *models.User) error {
	return r.DB.Save(user).Error
}

func (r *UserRepository) Delete(user *models.User) error {
	return r.DB.Delete(user).Error
}