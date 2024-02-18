package integration

import (
	"testing"

	"github.com/duchoang206h/send-server/test/integration/service"
	"github.com/stretchr/testify/assert"
)

func TestIntegration(t *testing.T) {
	assert := assert.New(t)
	srv, err := service.New()
	t.Run("Integration test", func(t *testing.T) {
		assert.Nil(err)
		assert.NotNil(srv)
	})
	/* if err := srv.Pool.Purge(srv.Mongo); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
	if err := srv.Pool.Purge(srv.Resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
	if err := srv.Pool.Purge(srv.Storage); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	} */
}
