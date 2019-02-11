package net

import (
	"errors"

	"./tcp"
)

type LoginStatus byte

const (
	Success              LoginStatus = 0x00
	InvalidCredentials   LoginStatus = 0x01
	UnverifiedAccount    LoginStatus = 0x02
	EmailAlreadyTaken    LoginStatus = 0x80
	UsernameAlreadyTaken LoginStatus = 0x81
	InsecurePassword     LoginStatus = 0x82
)

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

func TryLogin(username string, password string) LoginStatus {
	return Success
}

func RegisterAccount(username string, email string, password string) LoginStatus {
	return Success
}

func SendPasswordEmailByUser(username string) LoginStatus {
	return Success
}

func SendPasswordEmailByEmail(email string) LoginStatus {
	return Success
}

func RecoverPassword(token string, password string) LoginStatus {
	return Success
}
