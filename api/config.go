package api

import (
	"os"

	"github.com/joho/godotenv"
)

func envACCOUNTSID() string {
	println(godotenv.Load(".env"))
	err := godotenv.Load(".env")
	if err != nil {
		println("Error loading .env file")
		return ""
	}
	return os.Getenv("TWILIO_ACCOUNT_SID")
}

func envAUTHTOKEN() string {
	println(godotenv.Load(".env"))
	err := godotenv.Load(".env")
	if err != nil {
		println("Error loading .env file")
		return ""
	}
	return os.Getenv("TWILIO_AUTH_TOKEN")
}

func envSERVICESID() string {
	println(godotenv.Load(".env"))
	err := godotenv.Load(".env")
	if err != nil {
		println("Error loading .env file")
		return ""
	}
	return os.Getenv("TWILIO_SERVICE_SID")
}