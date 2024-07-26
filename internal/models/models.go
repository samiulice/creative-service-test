package models

import (
	"time"
)

const (
	// AustraliaRegex matches Australian mobile numbers with or without country code (+61)
	AustraliaRegex = `^(\+?61|0)4\d{8}$`

	// BangladeshRegex matches Bangladeshi mobile numbers with or without country code (+880)
	BangladeshRegex = `^\+?(880)?1[3-9]\d{8}$`

	// CanadaRegex matches Canadian phone numbers in various formats
	CanadaRegex = `^(\+?1)?[-.\s]?\(?\d{3}\)?[-.\s]?\d{3}[-.\s]?\d{4}$`

	// FranceRegex matches French phone numbers with or without country code (+33)
	FranceRegex = `^(?:(?:\+|00)33|0)\s*[1-9](?:[\s.-]*\d{2}){4}$`

	// GermanyRegex matches German phone numbers with or without country code (+49)
	GermanyRegex = `^(\+?49|0)(\d{3,4})?[ -]?(\d{3,4})?[ -]?(\d{4,6})$`

	// IndiaRegex matches Indian mobile numbers with or without country code (+91)
	IndiaRegex = `^\+?(91)?\d{10}$`

	// JapanRegex matches Japanese phone numbers with or without country code (+81)
	JapanRegex = `^\+?81[-.\s]?\d{1,4}[-.\s]?\d{1,4}[-.\s]?\d{4}$`

	// PakistanRegex matches Pakistani mobile numbers with or without country code (+92)
	PakistanRegex = `^\+?(92)?\d{10}$`

	// SriLankaRegex matches Sri Lankan mobile numbers with or without country code (+94)
	SriLankaRegex = `^\+?(94)?\d{9}$`

	// UKRegex matches UK phone numbers including landline, mobile, and toll-free numbers
	UKRegex = `^(?:(?:\+44\s?|0)(?:\d{5}\s?\d{5}|\d{4}\s?\d{4}\s?\d{4}|\d{3}\s?\d{3}\s?\d{4}|\d{2}\s?\d{4}\s?\d{4}|\d{4}\s?\d{4}|\d{4}\s?\d{3})|\d{5}\s?\d{4}\s?\d{4}|0800\s?\d{3}\s?\d{4})$`

	// USRegex matches US phone numbers in various formats
	USRegex = `^\+?1?[-.\s]?\(?\d{3}\)?[-.\s]?\d{3}[-.\s]?\d{4}$`
)

// user==3
// Employee==2
// Admin==1

// Order is the type for all orders
type Order struct {
	ClientName      string `json:"client-name"`
	ClientMobile    string `json:"client-mobile"`
	Division        string `json:"division"`
	District        string `json:"district"`
	Upazila         string `json:"upazila"`
	Area            string `json:"area"`
	Date            string `json:"date"`
	PreferredTime   string `json:"preferredtime"`
	Category        string `json:"category"`
	Service         string `json:"service"`
	ProblemSummary  string `json:"problem-summary"`
	ImgFile         string `json:"img_file"`
	LocationID int `json:"locaiton_id"`
	ServiceID int `json:"service_id"`
}

// Customer is the type for users
type Customer struct {
	ID              int       `json:"id"`
	UserName        string    `json:"user_name"`
	FirstName       string    `json:"first_name"`
	LastName        string    `json:"last_name"`
	Email           string    `json:"email"`
	Password        string    `json:"password"`
	ImageLink       string    `json:"image_link"` //username_profile_id_yy-mm-dd_hh-mm-ss.jpf
	AccountType     string    `json:"acc_type"`
	AccountStatusID int       `json:"account_status_id"` //0 = deleted, 1 = active, 2 = deactivated
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
