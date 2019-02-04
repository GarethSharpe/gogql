package utils

import (
	"os"

	f "github.com/nimajalali/go-force/force"
	"github.com/subosito/gotenv"
)

func GetConnection(accessToken string) (*f.ForceApi, error) {
	gotenv.Load()
	consumerKey := os.Getenv(DEV_CONSUMER_KEY)
	version := os.Getenv(VERSION)
	instanceURL := os.Getenv(DEV_INSTANCE_URL)
	connection, err := f.CreateWithAccessToken(
		version,
		consumerKey,
		accessToken,
		instanceURL,
	)
	return connection, err
}