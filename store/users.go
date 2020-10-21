// Author: dat.nguyen

package store

import "time"

// User -
type User struct {
	ID             int64     `json:"_id"`
	URL            string    `json:"url"`
	ExternalID     int       `json:"external_id"`
	Name           string    `json:"name"`
	Alias          []string  `json:"alias"`
	CreateAt       time.Time `json:"created_at"`
	Active         bool      `json:"active"`
	Verified       bool      `json:"verified"`
	Shared         bool      `json:"shared"`
	Locale         string    `json:"locale"`
	TimeZone       string    `json:"timezone"`
	LastLoginAt    time.Time `json:"last_login_at"`
	Email          string    `json:"email"`
	Phone          string    `json:"phone"`
	Signature      string    `json:"signature"`
	OrganizationID int       `json:"organization_id"`
	Tags           []string  `json:"tags"`
	Suspended      bool      `json:"suspended"`
	Role           bool      `json:"role"`
}
