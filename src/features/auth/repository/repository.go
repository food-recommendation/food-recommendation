package repository

import (
	"gorm.io/gorm"
)

type SignupAuthRepository struct {
	GormDB *gorm.DB
}
type SigninAuthRepository struct {
	GormDB *gorm.DB
}

type LogoutAuthRepository struct {
	GormDB *gorm.DB
}

type ReissueAuthRepository struct {
	GormDB *gorm.DB
}

type GoogleOauthAuthRepository struct {
	GormDB *gorm.DB
}
type GoogleOauthCallbackAuthRepository struct {
	GormDB *gorm.DB
}

type RequestPasswordAuthRepository struct {
	GormDB *gorm.DB
}

type ValidatePasswordAuthRepository struct {
	GormDB *gorm.DB
}
type CheckEmailAuthRepository struct {
	GormDB *gorm.DB
}

type GuestAuthRepository struct {
	GormDB *gorm.DB
}

type V02GoogleOauthCallbackAuthRepository struct {
	GormDB *gorm.DB
}
type V02GoogleOauthAuthRepository struct {
	GormDB *gorm.DB
}

type KakaoOauthAuthRepository struct {
	GormDB *gorm.DB
}

type NaverOauthAuthRepository struct {
	GormDB *gorm.DB
}

type RequestSignupAuthRepository struct {
	GormDB *gorm.DB
}
