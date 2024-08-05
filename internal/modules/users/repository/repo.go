package repository

import (
	"fmt"
	"log"
	"pos-acen/internal/helper/pass_encryption"
	"pos-acen/internal/modules/users/entity"
	"pos-acen/internal/modules/users/ports"
	"strings"

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
	INSERT INTO users (email, username, password, created_at)
	VALUES ($1, $2, $3, now())
	RETURNING id`

	var id uuid.UUID

	if err := tx.QueryRow(query,
		bReq.Email,
		bReq.Username,
		pass,
	).Scan(&id); err != nil {
		log.Println("Error Pada Saat Melakukan Query Register User : " + err.Error())
		return nil, err
	}

	// Commit akan menutup rows walaupun kita tidak melakukan defer rows.close
	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &id, err
}

func (u *userRepository) GetUserDetails(bReq entity.User) (*entity.User, error) {
	_, err := u.db.Begin()

	if err != nil {
		log.Println("Error pada saat Memulai Db Begin : " + err.Error())
		return nil, err
	}

	var queryConditionals []string
	var user *entity.User

	query := `SELECT *
	FROM users`

	queryConditionals = append(queryConditionals, " WHERE deleted_at IS null")

	if bReq.Email != "" {
		queryConditionals = append(queryConditionals, fmt.Sprintf("email = '%s'", bReq.Email))
	}

	if bReq.Username != "" {
		queryConditionals = append(queryConditionals, fmt.Sprintf("username = '%s'", bReq.Username))
	}

	if bReq.Id != uuid.Nil {
		queryConditionals = append(queryConditionals, fmt.Sprintf("id = '%v'", bReq.Id))
	}

	if len(queryConditionals) > 0 {
		query += strings.Join(queryConditionals, " AND ")
	}

	query += " Limit 1 "

	rows, err := u.db.Query(query)

	if err != nil {
		log.Println(query)
		log.Println("Error On GETUSERDETAILS : " + err.Error())
		return nil, err
	}

	// Kita harus menutup sqlx.Rows setiap kali ingin dipakai defer dalam hal ini berguna untuk
	// Menyatakan jalankan kode ini jika sebuah function sudah selesai

	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(
			&user.Id,
			&user.Email,
			&user.Username,
			&user.Password,
			&user.CreatedAt,
			&user.UpdatedAt,
		)

		if err != nil {
			log.Println("Error On GETUSERDETAILS Scan Rows : " + err.Error())
			return nil, err
		}

	}


	return user, nil
}
