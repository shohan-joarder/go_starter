package repositories

import (
	"github.com/go-playground/validator/v10"
	"github.com/shohan-joarder/go_pos/internal/models"
	"github.com/shohan-joarder/go_pos/internal/utils"
	"gorm.io/gorm"
)

var validate = validator.New()

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := r.DB.Find(&users).Error
	return users, err
}

func (r *UserRepository) CreateUser(user *models.User) error {

	validate = validator.New()

	// Register custom validation
	validate.RegisterValidation("phone", utils.PhoneValidator)
	validate.RegisterValidation("unique", utils.UniqueValidator(r.DB, false))

	err := validate.Struct(user)
	if err != nil {
		return err
	}

	return r.DB.Create(user).Error
}

func (r *UserRepository) GetUserByID(id string) (*models.User, error) {
	var user models.User
	err := r.DB.First(&user, "id = ?", id).Error
	return &user, err
}

func (r *UserRepository) UpdateUser(user *models.User) error {
	return r.DB.Save(user).Error
}

func (r *UserRepository) DeleteUser(user *models.User) error {
	return r.DB.Delete(user).Error
}
