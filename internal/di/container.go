package di

import (
	"bank/internal/domain"
	"bank/internal/handlers"
	"bank/internal/repository/postgres"
	"bank/internal/usecase"
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

type Container struct {
	router http.Handler

	db *sql.DB

	createUsers      *usecase.CreateUserUseCase
	readUsers        *usecase.ReadUserUseCase
	usersRepository  *postgres.UserRepository
	postUsersHandler *handlers.POSTUserHandler
	getUsersHandler  *handlers.GETUserHandler
}

func NewContainer() *Container {
	return &Container{
		db: postgres.CreateConnection(),
	}
}

func (c *Container) POSTUserHandler() *handlers.POSTUserHandler {
	if c.postUsersHandler == nil {
		c.postUsersHandler = handlers.NewPOSTUsersHandler(c.CreateUsers())
	}
	return c.postUsersHandler
}

func (c *Container) GETUserHandler() *handlers.GETUserHandler {
	if c.getUsersHandler == nil {
		c.getUsersHandler = handlers.NewGETUsersHandler(c.ReadUser())
	}
	return c.getUsersHandler
}

func (c *Container) ReadUser() *usecase.ReadUserUseCase {
	if c.readUsers == nil {
		c.readUsers = usecase.NewReadUserUseCase(c.UsersRepository())
	}
	return c.readUsers
}
func (c *Container) CreateUsers() *usecase.CreateUserUseCase {
	if c.createUsers == nil {
		c.createUsers = usecase.NewCreateUserUseCase(c.UsersRepository())
	}
	return c.createUsers
}

func (c *Container) UsersRepository() domain.UserRepository {
	if c.usersRepository == nil {
		c.usersRepository = postgres.NewUserRepository(c.db)
	}
	return c.usersRepository
}

func (c *Container) HTTPRouter() http.Handler {
	if c.router != nil {
		return c.router
	}
	router := mux.NewRouter()
	router.Handle("/users", c.POSTUserHandler()).Methods(http.MethodPost)
	router.Handle("/users/{login}/{password}", c.GETUserHandler()).Methods(http.MethodGet)
	c.router = router
	return c.router

}
