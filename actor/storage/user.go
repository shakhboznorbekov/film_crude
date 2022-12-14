package storage

import (
	"database/sql"
	// "fmt"

	"github.com/shakhboznorbekov/mytasks/30-dars/golang_crud-master/models"
)

func Create(db *sql.DB, user models.User) (string, error) {

	var (
		id    string
		query string
	)

	query = `
		INSERT INTO 
			actor (first_name, last_name,last_update)
		VALUES ( $1, $2,$3 )
		RETURNING actor_id
	`
	err := db.QueryRow(
		query,
		user.FirstName,
		user.LastName,
		user.LastUpdate,
	).Scan(&id)

	if err != nil {
		return "", err
	}

	return id, nil
}

func GetById(db *sql.DB, id string) (models.User, error) {

	var (
		user  models.User
		query string
	)

	query = `
		SELECT
			actor_id,
			first_name,
			last_name,
			last_update
		FROM
			actor
		WHERE actor_id = $1
	`
	err := db.QueryRow(
		query,
		id,
	).Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.LastUpdate,
	)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func GetList(db *sql.DB) ([]models.User, error) {

	var (
		users []models.User
		query string
	)

	query = `
		SELECT
			actor_id,
			first_name,
			last_name,
			last_update
		FROM
			actor
	`
	rows, err := db.Query(query)

	if err != nil {
		return []models.User{}, err
	}

	for rows.Next() {
		var user models.User

		err = rows.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.LastUpdate,
		)

		if err != nil {
			return []models.User{}, err
		}

		users = append(users, user)
	}

	return users, nil
}

func Update(db *sql.DB, user models.User) (int64, error) {

	var (
		query string
	)

	query = `
		UPDATE
			actor
		SET
			first_name = $2,
			last_name = $3,
			last_update=$4
		WHERE
			actor_id = $1
	`

	result, err := db.Exec(
		query,
		user.Id,
		user.FirstName,
		user.LastName,
		user.LastUpdate,
	)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func Delete(db *sql.DB, id string) error {

	_, err := db.Exec(`DELETE FROM actor WHERE actor_id = $1`, id)
	if err != nil {
		return err
	}

	return nil
}
