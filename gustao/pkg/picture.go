package pkg

import (
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

func GetPicture(name string) string {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal(err)
	}
	pPath := os.Getenv(name)
	pictures := strings.Split(pPath, " ")
	i := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(pPath))
	return pictures[i]
}
