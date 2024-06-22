package user

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/e-lua/go-nextjs-ts-chat/util"
)

type service struct {
	Repository
	timeout time.Duration
}

func NewService(repository Repository) Service {
	return &service{
		repository,
		time.Duration(2) * time.Second,
	}
}

func (s *service) CreateUser(c context.Context, req *CreateUserReq) (*CreateUserRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	//TODO: haspassword
	hasheddPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	u := &User{
		Username: req.Username,
		Email:    req.Email,
		Password: hasheddPassword,
	}

	r, err := s.Repository.CreateUser(ctx, u)
	if err != nil {
		return nil, err
	}

	res := &CreateUserRes{
		ID:       strconv.Itoa(int(r.ID)),
		Username: r.Username,
		Email:    r.Email,
	}

	return res, nil
}

type MyJWTClaims struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	ExpiresAt time.Time `json:"exp"`
}

func (s *service) Login(c context.Context, req *LoginUserReq) (*LoginUserRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	u, err := s.Repository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return &LoginUserRes{}, err
	}

	err = util.CheckPassword(req.Password, u.Password)
	if err != nil {
		return &LoginUserRes{}, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       strconv.Itoa(int(u.ID)),
		"username": u.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenKey := []byte("TokenGenerasdsa$$asdas..23c1qweadorRestoner")
	ss, err := token.SignedString([]byte(tokenKey))

	if err != nil {
		return &LoginUserRes{}, err
	}

	log.Println("5")

	return &LoginUserRes{accessToken: ss, Username: u.Username, ID: strconv.Itoa(int(u.ID))}, nil
}
