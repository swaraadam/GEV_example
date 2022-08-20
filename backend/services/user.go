package services

import (
	"context"
	db "gev_example/database/connection"
	"gev_example/helpers"
	token "gev_example/helpers/auth"
	"gev_example/helpers/util"
	"gev_example/models/protobuff"
	sqlc "gev_example/models/sqlc"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	protobuff.UnimplementedUserServiceServer
}

func (service *UserService) Register(ctx context.Context, request *protobuff.RegisterRequest) (*protobuff.RegisterResponse, error) {
	var user sqlc.MUser

	if request.Email == "" || request.Name == "" || request.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "missing datas")
	}

	user, _ = db.PG.Query.GetUserByEmail(ctx, request.Email)
	if user.Uuid != uuid.Nil {
		return nil, status.Error(codes.PermissionDenied, "User already exist!")
	}

	hashedPassword, err := util.HashPassword(request.Password)
	if err != nil {
		return nil, status.Error(codes.PermissionDenied, "Fail Hashing Password!")
	}

	user, _ = db.PG.Query.Register(ctx, sqlc.RegisterParams{
		Name:      request.Name,
		Password:  hashedPassword,
		Email:     request.Email,
		UpdatedAt: time.Time{},
	})

	//generate JWT Token
	paseto, err := token.NewJWTMaker()
	if err != nil {
		return nil, status.Error(codes.Aborted, "Undocumented!")
	}
	token, err := paseto.CreateToken(user.Uuid.String(), int64(helpers.Env.TokenDuration))
	if err != nil {
		return nil, status.Error(codes.Aborted, "Undocumented")
	}
	// claim, err := paseto.VerifyToken(token)
	// if err != nil {
	// 	return nil, err
	// }

	return &protobuff.RegisterResponse{
		AccessToken:   token,
		TokenDuration: int64(helpers.Env.TokenDuration),
	}, nil
}

func (service *UserService) Login(ctx context.Context, request *protobuff.LoginRequest) (*protobuff.LoginResponse, error) {
	return &protobuff.LoginResponse{
		Status:      0,
		Message:     "",
		AccessToken: "",
		User:        &protobuff.User{},
	}, nil
}
