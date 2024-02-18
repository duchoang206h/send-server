package service

import (
	"fmt"
	"log"
	"os"
	"path"
	"time"

	"github.com/duchoang206h/send-server/config"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

type Service struct {
	Resource *dockertest.Resource
	Mongo    *dockertest.Resource
	Storage  *dockertest.Resource
	Network  *dockertest.Network
	Pool     *dockertest.Pool
}

func New() (*Service, error) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Printf("Could not construct pool: %s", err)
		return nil, err
	}
	if err = pool.Client.Ping(); err != nil {
		log.Printf(`could not connect to docker: %s`, err)
		return nil, err
	}
	network, err := pool.CreateNetwork(fmt.Sprintf("test_network_%d", time.Now().UnixNano()))
	if err != nil {
		log.Printf(`could not connect to docker: %s`, err)
		return nil, err
	}
	storage, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "duchoang206h/telebot-storage",
		Tag:        "latest",
		PortBindings: map[docker.Port][]docker.PortBinding{
			"3002/tcp": {{HostPort: "3002"}},
		},
		Env: []string{
			"BOT_TOKEN=" + config.Config("BOT_TOKEN_TEST"),
			"CHAT_ID=" + config.Config("CHAT_ID_TEST"),
			"APP_PORT=:3002",
		},
	})
	if err != nil {
		log.Printf(`could not start storage: %s`, err)
		return nil, err
	}
	if err := storage.ConnectToNetwork(network); err != nil {
		log.Printf(`storage could not connect to network: %s`, err)
		return nil, err
	}
	mongo, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "mongo",
		Tag:        "latest",
		PortBindings: map[docker.Port][]docker.PortBinding{
			"27017/tcp": {{HostPort: "27017"}},
		},
		Env: []string{
			"MONGO_INITDB_ROOT_USERNAME=admin",
			"MONGO_INITDB_ROOT_PASSWORD=admin",
		},
	})
	if err != nil {
		log.Printf("Could not start resource: %s", err)
		return nil, err
	}
	if err := mongo.ConnectToNetwork(network); err != nil {
		log.Printf(`mongo could not connect to network: %s`, err)
		return nil, err
	}
	dirname, _ := os.Getwd()
	dockerPath := path.Join(dirname, "../../Dockerfile")
	fmt.Println("dockerPath::", dockerPath)
	resource, err := pool.BuildAndRunWithOptions(
		dockerPath,
		&dockertest.RunOptions{
			Hostname:  "server",
			NetworkID: network.Network.ID,
			Name:      "test-application",
			PortBindings: map[docker.Port][]docker.PortBinding{
				"3000/tcp": {{HostPort: "3000"}},
			},
			Env: []string{
				"DB_USER=admin",
				"DB_PASSWORD=admin",
				"DB_NAME=send",
				"DB_HOST=" + mongo.GetIPInNetwork(network),
				"DB_PORT=27017",
				"SERVER_URL=http://localhost:3002",
				"STORAGE_URL=http://localhost:3000",
				"PROXY_URL=http://localhost:3000",
			},
		},
	)
	if err != nil {
		log.Printf("Could not start resource: %s", err)
		return nil, err
	}
	return &Service{resource, mongo, storage, network, pool}, nil
}
