package integration

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"github.com/duchoang206h/send-server/test/integration/service"
	"github.com/stretchr/testify/assert"
)

type APIResponse struct {
	Result string `json:"result"`
}

func TestIntegration(t *testing.T) {
	assert := assert.New(t)
	srv, err := service.New()
	t.Run("Setup services", func(t *testing.T) {
		assert.Nil(err)
		assert.NotNil(srv)
	})
	t.Run("Upload file", func(t *testing.T) {
		tempUploadsDir := "./temp"
		os.Mkdir(tempUploadsDir, os.ModePerm)
		defer os.RemoveAll(tempUploadsDir)
		content := []byte("Integration test content")
		filePath := filepath.Join(tempUploadsDir, "test.txt")
		err := os.WriteFile(filePath, content, 0o644)
		if err != nil {
			t.Fatal(err)
		}
		defer os.Remove(filePath)
		file, err := os.Open(filePath)
		if err != nil {
			t.Fatal(err)
		}
		defer file.Close()

		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		part, err := writer.CreateFormFile("file", filepath.Base(filePath))
		if err != nil {
			t.Fatal(err)
		}
		_, err = io.Copy(part, file)
		if err != nil {
			t.Fatal(err)
		}
		writer.Close()
		// Send the file upload request to the test server
		resp, err := http.Post(
			"http://localhost:3000/api/file",
			writer.FormDataContentType(),
			body,
		)
		if err != nil {
			t.Fatal(err)
		}
		defer resp.Body.Close()
		assert.Equal(http.StatusOK, resp.StatusCode)
		responseBody, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}
		result := APIResponse{}
		err = json.Unmarshal(responseBody, &result)
		if err != nil {
			t.Fatal(err)
		}
		assert.NotEmpty(result.Result)
		fileResp, err := http.Get(
			result.Result,
		)
		if err != nil {
			t.Fatal(err)
		}
		defer fileResp.Body.Close()
		if fileResp.StatusCode != http.StatusOK {
			t.Fatalf("unexpected status code: %d", resp.StatusCode)
		}
		outputPath := filepath.Join(tempUploadsDir, "test_download.txt")
		// Create the output file
		out, err := os.Create(outputPath)
		if err != nil {
			t.Fatal(err)
		}
		defer out.Close()
		_, err = io.Copy(out, fileResp.Body)
		if err != nil {
			t.Fatal(err)
		}
		outFile, err := os.Open(outputPath)
		if err != nil {
			t.Fatal(err)
		}
		defer outFile.Close()

		// Get the file size
		stat, err := outFile.Stat()
		if err != nil {
			t.Fatal(err)
		}
		bs := make([]byte, stat.Size())
		_, err = bufio.NewReader(outFile).Read(bs)
		if err != nil && err != io.EOF {
			t.Fatal(err)
		}
		assert.Equal(bs, content)
	})
	if err := srv.Pool.Purge(srv.Mongo); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
	if err := srv.Pool.Purge(srv.Resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
	if err := srv.Pool.Purge(srv.Storage); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
}
