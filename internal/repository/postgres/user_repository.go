package postgres

import (
	"bank/internal/domain"
	"database/sql"
	"log"

	"github.com/gofrs/uuid"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	dbCon *sql.DB
}

func NewUserRepository(db_co *sql.DB) *UserRepository {
	return &UserRepository{
		dbCon: db_co,
	}
}

func (repository *UserRepository) Save(user *domain.User) (uuid.UUID, error) {
	_, err := repository.dbCon.Exec("INSERT INTO bank_api.users( id, login, password, fullname) VALUES ( $1, $2, $3, $4,$5);", user.ID(), user.Login(), user.Password(), user.Fullname())
	if err != nil {
		log.Fatalf("Error: Unable to execute query: %v", err)
	}
	return user.ID(), err
}

func (repository *UserRepository) FindUser(login string, password string) (domain.User, error) {
	rows, err := repository.dbCon.Query("SELECT * FROM bank_api.users where login=$1;", login)
	if err != nil {
		log.Fatalf("Error: Unable to execute query: %v", err)
	}
	var (
		id          string
		newLogin    string
		newPassword string
		fullname    string
	)
	for rows.Next() {
		rows.Scan(&id, &newLogin, &newPassword, &password, &fullname)
	}

	idUUID, _ := uuid.FromString(id)
	newUser := domain.NewUser(idUUID, login, password, fullname)

	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(newUser.Password()))
	if err != nil {
		return domain.User{}, err
	}
	return *newUser, err
}
