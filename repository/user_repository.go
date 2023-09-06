package repository

import (
	"github.com/jutionck/golang-todo-apps/config"
	"github.com/jutionck/golang-todo-apps/domain"
	"github.com/jutionck/golang-todo-apps/utils/commons"
	"github.com/jutionck/golang-todo-apps/utils/exception"
	"github.com/jutionck/golang-todo-apps/utils/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository interface {
	Save(payload *domain.User) error
	List(requestQueryParams model.RequestQueryParams) ([]domain.User, model.Paging, error)
	Get(id string) (*domain.User, error)
	GetEmail(email string) (*domain.User, error)
	Delete(id string) error
	Init() (string, error)
}

type userRepository struct {
	db *gorm.DB
}

func (u *userRepository) Save(payload *domain.User) error {
	return u.db.Save(payload).Error
}

func (u *userRepository) List(requestQueryParams model.RequestQueryParams) ([]domain.User, model.Paging, error) {
	paginationQuery, orderQuery := commons.PagingValidate(requestQueryParams)
	var users []domain.User
	result := u.db.Order(orderQuery).Limit(paginationQuery.Take).Offset(paginationQuery.Skip).Preload(clause.Associations).Find(&users).Error
	if result != nil {
		return nil, model.Paging{}, result
	}
	var totalRows int64
	result = u.db.Model(&domain.User{}).Count(&totalRows).Error
	if result != nil {
		return nil, model.Paging{}, result
	}
	return users, commons.Paginate(paginationQuery.Page, paginationQuery.Take, int(totalRows)), nil
}

func (u *userRepository) Get(id string) (*domain.User, error) {
	var user domain.User
	result := u.db.Where("id=?", id).Preload(clause.Associations).First(&user).Error
	if result != nil {
		return nil, result
	}
	return &user, nil
}

func (u *userRepository) GetEmail(email string) (*domain.User, error) {
	var user domain.User
	result := u.db.Where("email=?", email).First(&user).Error
	if result != nil {
		return nil, result
	}
	return &user, nil
}

func (u *userRepository) Delete(id string) error {
	return u.db.Where("id=?", id).Delete(&domain.User{}).Error
}

func (u *userRepository) Init() (string, error) {
	password, err := commons.HashPassword("password")
	exception.CheckError(err)
	users := []domain.User{
		{
			Email:    "admin@mail.com",
			Password: password,
			Role:     config.ADMIN,
		},
		{
			Email:    "user1@mail.com",
			Password: password,
			Role:     config.USER,
		},
		{
			Email:    "user2@mail.com",
			Password: password,
			Role:     config.USER,
		},
	}

	// Check user
	var data []domain.User
	isExists := false
	u.db.Find(&data)
	for _, v := range data {
		if v.Role == config.ADMIN {
			isExists = true
		}
	}

	if !isExists {
		result := u.db.Save(&users).Error
		if result != nil {
			return "", result
		}
	}
	return "Ok", nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}
