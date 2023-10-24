package handler

import (
	"encoding/json"
	"fmt"

	"github.com/duchoang206h/send-server/config"
	"github.com/duchoang206h/send-server/repository"
	"github.com/duchoang206h/send-server/util"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)
const (
	PROXY_RETRY = 3
)
type BodyResponse struct {
	Result string `json:"result"`
}
type FileHandler struct {
	fileRepo repository.FileRepository
}
func NewFileHandler (fileRepo repository.FileRepository) FileHandler {
return FileHandler{
	fileRepo: fileRepo,
}
}
func (fileHandler *FileHandler) HandleFileUpload (c *fiber.Ctx) error {
	if err := proxy.DoRedirects(c, config.GetStorageProxyUrl(), PROXY_RETRY); err != nil {
		return err
	}
	response := c.Response();
	var bodyRsp BodyResponse
	err := json.Unmarshal(response.Body(), &bodyRsp)
	fmt.Println("err::", err)
	if err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"result": nil,
		})
	}
	file, err := fileHandler.fileRepo.CreateFile(bodyRsp.Result)
	fmt.Println("err::", err)

	if err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"result": nil,
		})
	}
	fileUrl := fileHandler.fileRepo.FormatHashToUrl(file.Hash)
	shortenFileUrl, _ := util.ShortenUrl(fileUrl)
	return c.JSON(fiber.Map{
		"result": shortenFileUrl,
	})
} 
func (fileHandler *FileHandler) HandleDownloadFile (c *fiber.Ctx) error {
	hash := c.Params("hash")
	file := fileHandler.fileRepo.FindByHash(hash)
	//todo!: handle error message
	if file == nil {
		return c.Status(fiber.ErrNotFound.Code).JSON(fiber.Map{
			"result": nil,
		})
	}
	if err := proxy.DoRedirects(c, fmt.Sprintf("%s/%s", config.GetStorageProxyUrl(), file.FileID), PROXY_RETRY); err != nil {
		return err
	}
	response := c.Response();
	var bodyRsp BodyResponse
	err := json.Unmarshal(response.Body(), &bodyRsp)
	if err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"result": nil,
		})
	}
	c.Set("Content-Type", "application/octet-stream")
	c.Redirect(bodyRsp.Result)
	return nil
}