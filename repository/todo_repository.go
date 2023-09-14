package repository

import (
	"github.com/jutionck/golang-todo-apps/domain"
	"github.com/jutionck/golang-todo-apps/utils/commons"
	"github.com/jutionck/golang-todo-apps/utils/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TodoRepository interface {
	Save(payload *domain.Todo) error
	List(requestQueryParams model.RequestQueryParams, excludeUserID string) ([]domain.Todo, model.Paging, error)
	Get(id string) (*domain.Todo, error)
	Delete(id string) error
}

type todoRepository struct {
	db *gorm.DB
}

func (t *todoRepository) Save(payload *domain.Todo) error {
	return t.db.Save(payload).Error
}

func (t *todoRepository) List(requestQueryParams model.RequestQueryParams, excludeUserID string) (
	[]domain.Todo,
	model.Paging,
	error,
) {
	paginationQuery, orderQuery := commons.PagingValidate(requestQueryParams)
	var todos []domain.Todo
	var result error
	if excludeUserID != "" {
		result = t.db.
			Preload(clause.Associations).
			Order(orderQuery).
			Limit(paginationQuery.Take).
			Offset(paginationQuery.Skip).
			Where(
				"user_id=?",
				excludeUserID,
			).Find(&todos).Error
	} else {
		result = t.db.
			Preload(clause.Associations).
			Order(orderQuery).
			Limit(paginationQuery.Take).
			Offset(paginationQuery.Skip).
			Find(&todos).Error
	}
	if result != nil {
		return nil, model.Paging{}, result
	}
	var totalRows int64
	result = t.db.Model(&domain.User{}).Count(&totalRows).Error
	if result != nil {
		return nil, model.Paging{}, result
	}
	return todos, commons.Paginate(
		paginationQuery.Page,
		paginationQuery.Take,
		int(totalRows),
	), nil
}

func (t *todoRepository) Get(id string) (*domain.Todo, error) {
	var todo domain.Todo
	result := t.db.
		Where("id=?", id).
		Preload("User").
		First(&todo).Error
	if result != nil {
		return nil, result
	}
	return &todo, nil
}

func (t *todoRepository) Delete(id string) error {
	return t.db.Where("id=?", id).Delete(&domain.Todo{}).Error
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoRepository{db: db}
}
