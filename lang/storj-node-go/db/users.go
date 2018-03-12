package db

import (
	"encoding/json"
	"github.com/boltdb/bolt"
	"github.com/kataras/iris"
	"github.com/satori/go.uuid"
)

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Uuid     string `json:"uuid"`
}

type DB struct {
	Bucket bolt.Bucket
}

func (db *DB) CreateUser(ctx iris.Context) {
	user := User{}
	err := ctx.ReadJSON(&user)

	uu, err := uuid.NewV4()
	uuidString := uu.String()
	var uuidBytes = []byte(uuidString)

	userBytes, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	db.Bucket.Put(uuidBytes, userBytes)

	if err != nil {
		ctx.JSON(iris.StatusNotAcceptable)
		return
	}
	ctx.JSON(user)
}
