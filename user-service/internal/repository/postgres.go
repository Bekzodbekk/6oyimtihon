package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"user-service/genproto/userpb"
	"user-service/internal/pkg/load"
	"user-service/storage"
	"user-service/token"

	"gopkg.in/gomail.v2"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/exp/rand"
)

type UserRepo struct {
	db  *storage.Queries
	rdb *redis.Client
	cfg load.Config
}

func NewUserRepo(cfg load.Config, db *storage.Queries, rdb *redis.Client) IUserRepository {
	return &UserRepo{
		db:  db,
		rdb: rdb,
		cfg: cfg,
	}
}

func (u *UserRepo) Register(ctx context.Context, req *userpb.CreateUserReq) (*userpb.CreateUserResp, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	_, err = u.db.CheckUser(ctx, storage.CheckUserParams{
		Username: req.Username,
		Email:    req.Email,
	})

	if err != nil {
		if err == sql.ErrNoRows {
			return &userpb.CreateUserResp{
				Status:  false,
				Message: "Bunday user mavjud",
			}, nil
		}
	}

	expiredKey := u.GenerateRandomNumber()
	err = u.rdb.HSet(ctx, req.Email, map[string]interface{}{
		"username":   req.Username,
		"password":   passwordHash,
		"email":      req.Email,
		"expiredKey": expiredKey,
	}).Err()
	if err != nil {
		return nil, err
	}

	err = u.rdb.Expire(ctx, req.Email, 60*time.Second).Err()
	if err != nil {
		return nil, err
	}

	err = u.SendEmail(req.Email, expiredKey)
	if err != nil {
		return nil, err
	}

	return &userpb.CreateUserResp{
		Status:  true,
		Message: "Malumotlar Redisga saqlandi",
	}, nil
}
func (u *UserRepo) UpdateUser(ctx context.Context, req *userpb.UpdateUserReq) (*userpb.UpdateUserResp, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	result, err := u.db.UpdateUser(ctx, storage.UpdateUserParams{
		ID:        id,
		Username:  req.Username,
		Password:  string(passwordHash),
		Email:     req.Email,
		UpdatedAt: sql.NullString{String: time.Now().String(), Valid: true},
	})
	if err != nil {
		return nil, err
	}

	return &userpb.UpdateUserResp{
		Status:  true,
		Message: "User Successfuly updated",
		User: &userpb.User{
			Id:       result.ID.String(),
			Username: result.Username,
			Email:    result.Email,
			Cud: &userpb.CUD{
				CreatedAt: result.CreatedAt.String,
				UpdatedAt: result.UpdatedAt.String,
				DeletedAt: int64(result.DeletedAt.Int32),
			},
		},
	}, nil
}
func (u *UserRepo) DeleteUser(ctx context.Context, req *userpb.DeleteUserReq) (*userpb.DeleteUserResp, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	err = u.db.DeleteUser(ctx, storage.DeleteUserParams{
		ID: id,
	})
	if err != nil {
		return nil, err
	}
	return &userpb.DeleteUserResp{
		Status:  true,
		Message: "User success deleted!",
	}, nil
}

func (u *UserRepo) GetUserById(ctx context.Context, req *userpb.GetUserByIdReq) (*userpb.GetUserByIdResp, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}

	result, err := u.db.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &userpb.GetUserByIdResp{
		Status:  true,
		Message: "Get User By ID",
		User: &userpb.User{
			Id:       result.ID.String(),
			Username: result.Username,
			Email:    result.Email,
			Cud: &userpb.CUD{
				CreatedAt: result.CreatedAt.String,
				UpdatedAt: result.UpdatedAt.String,
				DeletedAt: int64(result.DeletedAt.Int32),
			},
		},
	}, nil
}
func (u *UserRepo) Login(ctx context.Context, req *userpb.LoginReq) (*userpb.LoginResp, error) {
	result, err := u.db.Login(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(req.Password))
	if err != nil {
		return nil, err
	}
	accessToken, err := token.CreateTokens(result.ID.String(), u.cfg.SecretKey)
	if err != nil {
		return nil, err
	}
	return &userpb.LoginResp{
		Status:  true,
		Message: "Login successfuly",
		Token:   accessToken,
	}, nil
}
func (u *UserRepo) VerifyUser(ctx context.Context, req *userpb.VerifyUserReq) (*userpb.VerifyUserResp, error) {
	userRedis, err := u.rdb.HGetAll(ctx, req.Email).Result()
	if err != nil {
		return nil, err
	}

	if userRedis["expiredKey"] != req.Password {
		return &userpb.VerifyUserResp{
			Status:  false,
			Message: "Verify password wrong!",
		}, nil
	}

	user, err := u.db.CreateUser(ctx, storage.CreateUserParams{
		Username: userRedis["username"],
		Password: userRedis["password"],
		Email:    req.Email,
	})
	if err != nil {
		return nil, err
	}

	return &userpb.VerifyUserResp{
		Status:  true,
		Message: "User successfuly created",
		User: &userpb.User{
			Id:       user.ID.String(),
			Username: user.Username,
			Email:    user.Email,
			Cud: &userpb.CUD{
				CreatedAt: user.CreatedAt.String,
				UpdatedAt: user.UpdatedAt.String,
				DeletedAt: int64(user.DeletedAt.Int32),
			},
		},
	}, nil
}

func (u *UserRepo) GenerateRandomNumber() int {
	rand.Seed(uint64(time.Now().UnixNano()))
	return rand.Intn(9000) + 1000
}

func (u *UserRepo) SendEmail(to string, code int) error {
	subject := "----Welcome buddy----"

	body := fmt.Sprintf(`
	  <!DOCTYPE html>
	  <html>
	  <head>
		<style>
		  .container {
			font-family: Arial, sans-serif;
			background-color: #f4f4f4;
			padding: 20px;
			border-radius: 10px;
			width: 80%%;
			margin: 0 auto;
			color: #333;
		  }
		  .header {
			background-color: #4CAF50;
			color: white;
			padding: 10px;
			border-radius: 10px 10px 0 0;
			text-align: center;
		  }
		  .content {
			padding: 20px;
			background-color: white;
			border-radius: 0 0 10px 10px;
		  }
		  .code {
			font-size: 24px;
			font-weight: bold;
			color: #4CAF50;
			text-align: center;
			margin: 20px 0;
		  }
		  .footer {
			text-align: center;
			padding: 10px;
			font-size: 12px;
			color: #777;
		  }
		</style>
	  </head>
	  <body>
		<div class="container">
		  <div class="header">
			<h1>Welcome to Our Service!</h1>
		  </div>
		  <div class="content">
			<p>Dear user,</p>
			<p>Thank you for signing up. To complete your registration, please use the following confirmation code:</p>
			<div class="code">%d</div>
			<p>If you didn't sign up, please ignore this email.</p>
		  </div>
		  <div class="footer">
			<p>This is an automated message, please do not reply.</p>
		  </div>
		</div>
	  </body>
	  </html>
	`, code)

	m := gomail.NewMessage()
	m.SetHeader("From", "bekzodnematov709@gmail.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	d := gomail.NewDialer("smtp.gmail.com", 587, u.cfg.Email, u.cfg.AccessKey)

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
