package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID                   int    `gorm:"primaryKey"`
	Uuid                 string `gorm:"uniqueIndex"`
	FirstName            string
	LastName             string
	Email                string `gorm:"uniqueIndex;not null"`
	Phone                string `gorm:"uniqueIndex"`
	Username             string `gorm:"uniqueIndex;not null"`
	Password             string
	Role                 string
	Status               string
	AccessToken          string
	AccessTokenExpiredAt sql.NullTime
	EmailVerifiedAt      sql.NullTime
	PhoneVerifiedAt      sql.NullTime
	LoginAttempt         int    
	LastSuccessfulLogin  sql.NullTime
	LastFailedLogin      sql.NullTime
	CreatedAt            time.Time `gorm:"autoCreateTime"`
	UpdatedAt            time.Time `gorm:"autoUpdateTime"`
}
