package scripts

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"os"
	"testing"
	"time"

	"github.com/HsiaoCz/beast-clone/reader/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(os.Getenv("MD5SECRET")))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

func TestCreateUser(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	users := []models.User{
		{Username: "zhangsan", Password: encryptPassword("zhangsan123"), Email: "zhangsan@gmail.com", Content: "are you ok", Avatar: "./avatar/users/sssas.jpg", ReadTime: "0", IsAdmin: false},
		{Username: "lisi", Password: encryptPassword("lsisss123"), Email: "lisi@gmail.com", Content: "somethis we lost", Avatar: "./avatar/users/12321.jpg", ReadTime: "0", IsAdmin: false},
		{Username: "wangwu", Password: encryptPassword("wangsuws123"), Email: "wangwu@gmail.com", Content: "故事你还在听吗", Avatar: "./avatar/users/12323.jpg", ReadTime: "0", IsAdmin: false},
		{Username: "zhaoliu", Password: encryptPassword("zhsnags123"), Email: "zhaoliu@gmail.com", Content: "角落那窗口,闻得到玫瑰花香", Avatar: "./avatar/users/122322.jpg", ReadTime: "0", IsAdmin: false},
		{Username: "zhouqiba", Password: encryptPassword("zhasuos123"), Email: "zhouqiba@gmail.com", Content: "被你一说是有些印象", Avatar: "./avatar/users/212231.jpg", ReadTime: "0", IsAdmin: false},
		{Username: "qiansijiu", Password: encryptPassword("quasss123"), Email: "qiansijiu@gmail.com", Content: "我没有说谎,我何必说谎", Avatar: "./avatar/users/78722.jpg", ReadTime: "0", IsAdmin: false},
		{Username: "zhendongguo", Password: encryptPassword("zhaosusn123"), Email: "zhendongguo@gmail.com", Content: "你懂我的,我对你从来都不会假装", Avatar: "./avatar/users/787282.jpg", ReadTime: "0", IsAdmin: false},
		{Username: "duzhili", Password: encryptPassword("duzhishs123"), Email: "duzhili@gmail.com", Content: "全都是泡沫,只一刹那花火", Avatar: "./avatar/users/90922.jpg", ReadTime: "0", IsAdmin: false},
		{Username: "zhatians", Password: encryptPassword("zhashshs123"), Email: "zhatians@gmail.com", Content: "你所有承诺,全部都太脆弱", Avatar: "./avatar/users/2328382.jpg", ReadTime: "0", IsAdmin: false},
		{Username: "bengbuzhu", Password: encryptPassword("basnszhus123"), Email: "bengbuzhu@gmail.com", Content: "而你的轮廓,怪我没有看破", Avatar: "./avatar/users/898282.jpg", ReadTime: "0", IsAdmin: false},
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

func TestDeleteUser(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	feed, err := Newfeed()
	if err != nil {
		t.Fatal(err)
	}
	uid, err := primitive.ObjectIDFromHex("665d0982130a037e828f0b0f")
	if err != nil {
		t.Fatal(err)
	}
	if err := feed.DeleteUser(ctx, uid); err != nil {
		t.Fatal(err)
	}
	t.Log("delete user success")
}
