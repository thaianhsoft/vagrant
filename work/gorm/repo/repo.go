package repo

import (
	"database/sql"
	"github.com/thaianhsoft/gorm/schema"
)

type RepoService interface {
	Save(schema schema.Schema, tx *sql.Tx) error
	Delete(schema schema.Schema, tx *sql.Tx) error
	Update(schema schema.Schema, tx *sql.Tx) error
	GetById(schema schema.Schema, tx *sql.Tx) error
	RawQuery(query string, schemaCastBack interface{}) error
}

type CrudRepoService struct {
	db *sql.DB
}



