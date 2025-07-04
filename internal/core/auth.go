// File: internal/core/auth.go
// Purpose: Core logic for login, registration, token refresh, logout, and password management.

package core

import (
	"errors"

	"shorty/internal/db"
	"shorty/internal/lib/jwt"
	"shorty/internal/models"
	"shorty/internal/utils"
	"shorty/internal/utils/mailer"
)

func AuthLogin(ctx *models.Context, req models.LoginRequest) (*models.AuthResponse, error) {
	user, err := db.GetUserByEmail(ctx, req.Email)
	if err != nil || !utils.CheckPassword(req.Password, user.Password) {
		return nil, errors.New("invalid email or password")
	}

	token, err := jwt.GenerateToken(user.ID)
	if err != nil {
		return nil, err
	}

	return &models.AuthResponse{
		Token: token,
		User:  user.Sanitized(),
	}, nil
}

func AuthLogout(ctx *models.Context) error {
	return jwt.InvalidateToken(ctx.Token)
}

func RegisterUser(ctx *models.Context, req models.RegisterRequest) (*models.UserPublic, error) {
	if !utils.IsEmail(req.Email) {
		return nil, errors.New("invalid email")
	}

	hash, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Email:    req.Email,
		Password: hash,
		Name:     req.Name,
	}

	if err := db.CreateUser(ctx, user); err != nil {
		return nil, err
	}

	return user.Sanitized(), nil
}

func RefreshToken(ctx *models.Context) (*models.AuthResponse, error) {
	newToken, err := jwt.GenerateToken(ctx.User.ID)
	if err != nil {
		return nil, err
	}

	return &models.AuthResponse{
		Token: newToken,
		User:  ctx.User.Sanitized(),
	}, nil
}

func ForgotPassword(ctx *models.Context, email string) error {
	user, err := db.GetUserByEmail(ctx, email)
	if err != nil {
		return err
	}

	token, err := jwt.GeneratePasswordResetToken(user.ID)
	if err != nil {
		return err
	}

	return mailer.SendResetEmail(user.Email, token)
}

func ResetPassword(ctx *models.Context, req models.ResetPasswordRequest) error {
	userID, err := jwt.ValidateResetToken(req.Token)
	if err != nil {
		return err
	}

	user, err := db.GetUserByID(ctx, userID)
	if err != nil {
		return err
	}

	hash, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		return err
	}

	return db.UpdateUserPassword(ctx, user.ID, hash)
}
