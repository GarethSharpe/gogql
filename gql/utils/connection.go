package utils

import (
	// "os"

	// f "github.com/nimajalali/go-force/force"
	// "github.com/subosito/gotenv"
)

// func GetConnection(accessToken string) (f.ForceApi, error) {
// 	gotenv.Load()
// 	connection, err := f.CreateWithAccessToken(
// 		os.Getenv(VERSION),
// 		os.Getenv(DEV_CONSUMER_KEY),
// 		accessToken,
// 		os.Getenv(DEV_INSTANCE_URL),
// 	)
// 	return *connection, err
// }

func GetConnection(accessToken string) (string, error) {
	return accessToken, nil
}