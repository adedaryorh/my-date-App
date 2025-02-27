package models

// QrCodeLoginDto object used to log user in via qr code
type QrCodeLoginDto struct {
	Id    string `json:"id"`
	Token string `json:"token"`
}
