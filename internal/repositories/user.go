package repositories

import (
	"fmt"

	"github.com/go-api-rest/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepo(DB_CONNECTION *gorm.DB) *UserRepository {
	return &UserRepository{
		db: DB_CONNECTION,
	}
}

func (r *UserRepository) FindByID(id int) (*models.User, error) {
	var user models.User

	// Usar First correctamente y manejar el error
	result := r.db.Where("id = ?", id).First(&user)

	if result.Error != nil {

		// Imprimir el error para depuración
		fmt.Printf("Error fetching user: %v\n", gorm.ErrRecordNotFound)
		return nil, gorm.ErrRecordNotFound
	}

	return &user, nil
}

func (r *UserRepository) AddNewUser(user *models.User) (*models.User, error) {
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}

	// Retornar el usuario creado
	return user, nil
}
