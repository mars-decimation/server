package net

import (
	"errors"

	"./tcp"
)

// LoginStatus is one of the status codes that can be returned as a result of a login packet
type LoginStatus byte

const (
	// Success represents a successful action
	Success LoginStatus = 0x00
	// InvalidCredentials means either a username, password, or email address was incorrect
	InvalidCredentials LoginStatus = 0x01
	// UnverifiedAccount means the email address associated with the account has not yet been verified, so the login
	// cannot complete
	UnverifiedAccount LoginStatus = 0x02
	// EmailAlreadyTaken is returned when a user is trying to register an account with an email address that already has
	// an account
	EmailAlreadyTaken LoginStatus = 0x80
	// UsernameAlreadyTaken is returned when a user is trying to register an account with a username that is already
	// associated with another account
	UsernameAlreadyTaken LoginStatus = 0x81
	// InsecurePassword is returned when the password a user tries to set is too insecure and therefore should not be
	// used.  An insecure password will not be accepted by the server.
	InsecurePassword LoginStatus = 0x82
)

// DoLogin performs the login protocol and blocks until the login either fails or completes successfully.  In the event
// of a successful return from this method, the client has been authenticated.
func DoLogin(client *tcp.Client) error {
	for {
		select {
		case msg := <-client.Stream:
			var res LoginStatus
			switch msg[0] {
			case 0x00: // Login
				usernameLen := int(msg[1])
				passwordLen := 256*int(msg[2]) + int(msg[3])
				username := string(msg[4 : 4+usernameLen])
				password := string(msg[4+usernameLen : 4+usernameLen+passwordLen])
				res = TryLogin(username, password)
				break
			case 0x01: // Register
				usernameLen := int(msg[1])
				emailLen := 256*int(msg[2]) + int(msg[3])
				passwordLen := 256*int(msg[4]) + int(msg[5])
				username := string(msg[6 : 6+usernameLen])
				email := string(msg[6+usernameLen : 6+usernameLen+emailLen])
				password := string(msg[6+usernameLen+emailLen : 6+usernameLen+emailLen+passwordLen])
				res = RegisterAccount(username, email, password)
				break
			case 0x02: // Forgot Password (email)
				emailLen := 256*int(msg[1]) + int(msg[2])
				email := string(msg[3 : 3+emailLen])
				res = SendPasswordEmailByEmail(email)
				break
			case 0x03: // Forgot Password (username)
				usernameLen := int(msg[1])
				username := string(msg[2 : 2+usernameLen])
				res = SendPasswordEmailByUser(username)
				break
			case 0x04: // Password Recovery
				passwordLen := 256*int(msg[1]) + int(msg[2])
				recoveryCode := string(msg[3:19])
				password := string(msg[19 : 19+passwordLen])
				res = RecoverPassword(recoveryCode, password)
				break
			default:
				client.Close()
				return errors.New("Invalid protocol")
			}
			client.Stream <- []byte{byte(res)}
			if res == 0 {
				return nil
			}
			break
		case err := <-client.OnDisconnect:
			return err
		}
	}
}

// TryLogin handles a login packet
func TryLogin(username string, password string) LoginStatus {
	return Success
}

// RegisterAccount handles an account registration packet
func RegisterAccount(username string, email string, password string) LoginStatus {
	return Success
}

// SendPasswordEmailByUser handles a forgotten password packet when the remembered information is an email address
func SendPasswordEmailByUser(username string) LoginStatus {
	return Success
}

// SendPasswordEmailByEmail handles a forgotten password packet when the remembered information is a username
func SendPasswordEmailByEmail(email string) LoginStatus {
	return Success
}

// RecoverPassword handles a password recovery packet
func RecoverPassword(token string, password string) LoginStatus {
	return Success
}
