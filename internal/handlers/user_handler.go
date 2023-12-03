package handlers

import (
	"bank/internal/usecase"
	"encoding/json"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

type POSTUserHandler struct {
	useCase *usecase.CreateUserUseCase
}

type GETUserHandler struct {
	useCase *usecase.ReadUserUseCase
}

func NewPOSTUsersHandler(useCase *usecase.CreateUserUseCase) *POSTUserHandler {
	return &POSTUserHandler{
		useCase: useCase,
	}
}

func NewGETUsersHandler(useCase *usecase.ReadUserUseCase) *GETUserHandler {
	return &GETUserHandler{
		useCase: useCase,
	}
}

type GETUserResponse struct {
	id       uuid.UUID
	fullname string
}

func (response *GETUserResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ID       uuid.UUID
		Fullname string
	}{
		ID:       response.id,
		Fullname: response.fullname,
	})
}

type POSTUserRequest struct {
	Login    string
	Password string
	Fullname string
}

type POSTUserResponse struct {
	id       uuid.UUID
	login    string
	password string
	fullname string
}

func (response *POSTUserResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ID       uuid.UUID
		Login    string
		Password string
		Fullname string
	}{
		ID:       response.id,
		Login:    response.login,
		Password: response.password,
		Fullname: response.fullname,
	})
}

func (handler *POSTUserHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var body POSTUserRequest
	err := json.NewDecoder(request.Body).Decode(&body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}
	password := []byte(body.Password)
	passwordHash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadGateway)
	}
	command := &usecase.CreateUserCommand{
		Login:    body.Login,
		Password: string(passwordHash),
		Fullname: body.Fullname,
	}
	user, err := handler.useCase.Handle(command)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	response := &POSTUserResponse{
		id:       user.ID(),
		login:    user.Login(),
		password: user.Password(),
		fullname: user.Fullname(),
	}
	writer.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(writer).Encode(response)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

func (handler *GETUserHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	login, _ := vars["login"]
	password, _ := vars["password"]
	command := &usecase.ReadUserCommand{
		Login:    login,
		Password: password,
	}
	user, err := handler.useCase.Handle(command)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	response := &GETUserResponse{
		id:       user.ID(),
		fullname: user.Fullname(),
	}
	writer.WriteHeader(http.StatusOK)

	err = json.NewEncoder(writer).Encode(response)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}
