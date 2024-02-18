package repository

import (
	"log"
	"testing"

	"github.com/duchoang206h/send-server/database"
	"github.com/stretchr/testify/assert"
	"github.com/strikesecurity/strikememongo"
)

var (
	mongoServer *strikememongo.Server
	dbName      string
)

func TestFileRepo(t *testing.T) {
	assert := assert.New(t)
	var err error
	mongoServer, err = strikememongo.Start("4.0.5")
	if err != nil {
		log.Fatal(err)
	}
	defer mongoServer.Stop()
	dbName = strikememongo.RandomDatabase()
	err = database.ConnectMongo(mongoServer.URI(), dbName)
	assert.Nil(err)
	mongo := database.GetMongo()
	fileCollection := mongo.Collection("file")
	fileRepo := NewFileRepository(fileCollection)
	t.Run("Create new file", func(t *testing.T) {
		fileId := "123"
		file, err := fileRepo.CreateFile(fileId)
		assert.Nil(err)
		assert.Equal(file.FileID, fileId)
		assert.NotEmpty(file.ID)
	})
	t.Run("Find exist file", func(t *testing.T) {
		fileId := "123"
		file := fileRepo.FindById(fileId)
		assert.NotNil(file)
		assert.Equal(file.FileID, fileId)
	})
	t.Run("Find not exist file", func(t *testing.T) {
		fileId := "1234"
		file := fileRepo.FindById(fileId)
		assert.Nil(file)
	})
}
