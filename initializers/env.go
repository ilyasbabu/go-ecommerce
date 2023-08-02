package initializers

import (
	"fmt"

	"github.com/joho/godotenv"
)

func LoadEnv() {

	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading env vars.")
	}
}
