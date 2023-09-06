package manager

import (
	"fmt"
	"github.com/jutionck/golang-todo-apps/config"
	"github.com/jutionck/golang-todo-apps/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type InfraManager interface {
	Conn() *gorm.DB
	Migrate() error
}

type infraManager struct {
	db  *gorm.DB
	cfg *config.Config
}

func (i *infraManager) Migrate() error {
	if err := i.db.Debug().AutoMigrate(
		&domain.User{},
		&domain.Todo{},
	); err != nil {
		return err
	}
	return nil
}

func (i *infraManager) initDb() error {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		i.cfg.Host,
		i.cfg.Port,
		i.cfg.User,
		i.cfg.Password,
		i.cfg.Name,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	if i.cfg.FileConfig.Env == "migration" {
		i.db = db
		err := i.Migrate()
		if err != nil {
			return err
		}
	} else if i.cfg.FileConfig.Env == "dev" {
		i.db = db.Debug()
	} else {
		i.db = db
	}

	return nil
}

func (i *infraManager) Conn() *gorm.DB {
	return i.db
}

func NewInfraManager(cfg *config.Config) (InfraManager, error) {
	conn := &infraManager{cfg: cfg}

	if err := conn.initDb(); err != nil {
		return nil, err
	}
	return conn, nil
}
