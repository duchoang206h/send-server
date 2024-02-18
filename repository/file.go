package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/duchoang206h/send-server/config"
	"github.com/duchoang206h/send-server/model"
	"github.com/duchoang206h/send-server/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type FileRepository interface {
	CreateFile(fileId string) (*model.File, error)
	FindByHash(hash string) *model.File
	FindById(fileId string) *model.File
	Save(file *model.File) error
	FormatHashToUrl(hash string) string
}
type fileRepository struct {
	collection *mongo.Collection
}

func NewFileRepository(collection *mongo.Collection) FileRepository {
	return &fileRepository{
		collection: collection,
	}
}

func generateUniqueHash(ctx context.Context, findHash chan<- string, findHashFunc func(string) *model.File) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			hash := util.RandomString(8)
			exist := findHashFunc(hash)
			if exist == nil {
				fmt.Println("exist", exist)
				findHash <- hash
				return
			}
			fmt.Println("exist", exist)
		}
	}
}

func (fr *fileRepository) CreateFile(fileId string) (*model.File, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	findHashCh := make(chan string, 1)
	go generateUniqueHash(ctx, findHashCh, fr.FindByHash)
	select {
	case hash := <-findHashCh:
		file := model.File{
			FileID: fileId,
			Hash:   hash,
		}
		err := fr.Save(&file)
		if err != nil {
			return nil, err
		}
		return &file, nil
	case <-ctx.Done():
		// Timeout occurred, handle accordingly
		return nil, ctx.Err()
	}
}

func (fr *fileRepository) FindByHash(hash string) *model.File {
	query := bson.M{"hash": hash}
	var file model.File
	result := fr.collection.FindOne(context.TODO(), query)
	if result.Err() != nil {
		return nil
	}
	if err := result.Decode(&file); err != nil {
		return nil
	}
	return &file
}

func (fr *fileRepository) FindById(fileId string) *model.File {
	query := bson.M{"fileId": fileId}
	var file model.File
	result := fr.collection.FindOne(context.TODO(), query)
	if result.Err() != nil {
		return nil
	}
	if err := result.Decode(&file); err != nil {
		return nil
	}
	return &file
}

func (fr *fileRepository) Save(file *model.File) error {
	result, err := fr.collection.InsertOne(context.Background(), file)
	if err != nil {
		return err
	}
	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return errors.New("failed to convert inserted ID to ObjectID")
	}
	file.ID = insertedID
	return nil
}

func (fr *fileRepository) FormatHashToUrl(hash string) string {
	return fmt.Sprintf("%s/api/file/%s", config.Config("PROXY_URL"), hash)
}
