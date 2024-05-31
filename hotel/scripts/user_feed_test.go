package scripts

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"os"
	"testing"
	"time"

	"github.com/HsiaoCz/beast-clone/hotel/types"
	"github.com/joho/godotenv"
)

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(os.Getenv("SECRET")))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

func TestCreateUser(t *testing.T) {
	if err := godotenv.Load(); err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	users := []types.User{
		{FirstName: "wudi", LastName: "wang", Email: "wangwudi@gmail.com", EncryptedPassword: encryptPassword("wangwudi123"), IsAdmin: false},
		{FirstName: "zhasa", LastName: "zhao", Email: "zhaozhasa@gmail.com", EncryptedPassword: encryptPassword("zhaozhasa123"), IsAdmin: false},
		{FirstName: "shengzong", LastName: "song", Email: "songshenzong@gmail.com", EncryptedPassword: encryptPassword("songshenzong123"), IsAdmin: false},
		{FirstName: "yuanzhang", LastName: "zhu", Email: "zhuyuanzhang@gmail.com", EncryptedPassword: encryptPassword("zhuyuanzhang123"), IsAdmin: false},
		{FirstName: "sixiang", LastName: "lis", Email: "lisixiang@gmail.com", EncryptedPassword: encryptPassword("lisixiang123"), IsAdmin: false},
		{FirstName: "dabing", LastName: "liu", Email: "liudabing@gmail.com", EncryptedPassword: encryptPassword("liudabing123"), IsAdmin: false},
		{FirstName: "kuisi", LastName: "qian", Email: "qiankuisi@gmail.com", EncryptedPassword: encryptPassword("qians234"), IsAdmin: false},
		{FirstName: "bsussns", LastName: "assfx", Email: "zksjsfs@gmail.com", EncryptedPassword: encryptPassword("moaoas12345"), IsAdmin: false},
		{FirstName: "bionss", LastName: "zssda", Email: "cssdscs@gmail.com", EncryptedPassword: encryptPassword("sjfjsjs2334"), IsAdmin: false},
		{FirstName: "sncssd", LastName: "slksks", Email: "scsdsc@gmail.com", EncryptedPassword: encryptPassword("scvds2w323"), IsAdmin: false},
	}
	feed, err := Newfeed()
	if err != nil {
		t.Fatal(err)
	}
	for _, user := range users {
		_, err := feed.CreateUser(ctx, &user)
		if err != nil {
			t.Fatal(err)
		}
	}
}
