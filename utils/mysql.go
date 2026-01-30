package utils

import (
	"context"
	"shared-modules/common"
	"sync"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var DB *gorm.DB
var DBs map[common.DatabaseRole]*gorm.DB = make(map[common.DatabaseRole]*gorm.DB, 4)
var m sync.Mutex

type GormMySQLConfig struct {
	Name                  string // db key，例如: "main", "read", "snapshot"
	DSN                   string
	MaxOpenConns          int
	MaxIdleConns          int
	ConnMaxLifetimeSecond time.Duration
}

func RegisterDB(cfg GormMySQLConfig) error {
	db, err := gorm.Open(mysql.Open(cfg.DSN), &gorm.Config{})

	if err != nil {
		return err
	}

	db = db.Debug()

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(cfg.ConnMaxLifetimeSecond * time.Second)
	m.Lock()
	defer m.Unlock()
	DBs[(common.DatabaseRole)(cfg.Name)] = db
	if cfg.Name == string(common.DATABASE_ROLE_MAIN) {
		DB = db
		DBs[""] = db
	}

	return nil
}

func InitDb(dsn string, maxConn int) error {
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	DB = DB.Debug()
	db, err := DB.DB()

	if err != nil {
		return err
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(maxConn)
	db.SetConnMaxLifetime(time.Hour)

	return nil
}

func GetDB(ctx context.Context, name ...common.DatabaseRole) *gorm.DB {

	txDB := ctx.Value("db")

	if txDB == nil && (len(name) == 0 || name[0] == "") {
		return DB
	}

	if len(name) == 0 {
		return txDB.(*gorm.DB)
	}

	if db, ok := DBs[name[0]]; ok {
		if txDB != nil && txDB.(*gorm.DB).Dialector.(mysql.Dialector).DSN == db.Dialector.(mysql.Dialector).Config.DSN {
			return txDB.(*gorm.DB)
		}
		return db
	}

	return txDB.(*gorm.DB)
}

type Clause interface {
	Scope() func(db *gorm.DB) *gorm.DB
}

type IgnoreClause struct {
}

var _ Clause = (*IgnoreClause)(nil)

func (IgnoreClause) Scope() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Clauses(clause.Insert{Modifier: "IGNORE"})
	}
}
