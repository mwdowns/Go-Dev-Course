package db

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/supabase-community/supabase-go"
)

func Client() (*supabase.Client, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}
	client, err := supabase.NewClient(os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_KEY"), nil)
	if err != nil {
		return nil, err
	}
	return client, nil
}
