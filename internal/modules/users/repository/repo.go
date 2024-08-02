package repository

import (
	"log"
	"pos-acen/internal/helper/pass_encryption"
	"pos-acen/internal/modules/users/entity"
	"pos-acen/internal/modules/users/ports"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

var validate = validator.New()

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) ports.UserRepository {
	return &userRepository{
		db,
	}
}

func (u *userRepository) RegisterUser(bReq entity.User) (*uuid.UUID, error) {
	err := validate.Struct(bReq)

	if err != nil {
		log.Println("Error Validation body Request Register User : ", err.Error())
		return nil, err
	}

	pass, err := pass_encryption.Encrypt(bReq.Password)

	if err != nil {
		log.Println("Error Pada Saat Encrypt Password : ", err.Error())
		return nil, err
	}

	tx, err := u.db.Begin()

	if err != nil {
		log.Println("Error Pada Saat Db Begin : ", err.Error())
		return nil, err
	}

	query := `
	INSERT INTO users (id, email, username, password, created_at)
	VALUES ($1, $2, $3, $4, now())
	RETURNING id`

	var id uuid.UUID

	if err := tx.QueryRow(query,
		bReq.Id,
		bReq.Email,
		bReq.Username,
		pass,
	).Scan(&id); err != nil {
		log.Println("Error Pada Saat Melakukan Query Register User : " + err.Error())
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &id, err
}
