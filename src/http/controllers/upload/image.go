package v1uploadcontroller

import (
	userrepository "github.com/Dwibi/beli-mang/src/repositories/users"
	uploadsusecase "github.com/Dwibi/beli-mang/src/usecases/upload"
	"github.com/gofiber/fiber/v2"
)

type returnData struct {
	ImageUrl string `json:"imageUrl"`
}

type resultUpload struct {
	Message string     `json:"message"`
	Data    returnData `json:"data"`
}

func (i V1Upload) Image(c *fiber.Ctx) error {
	// Accessing userId from auth middleware
	userId := c.Locals("userId").(int)

	// fmt.Println(userId)

	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error retrieving the file",
		})
	}

	if file.Size < 10*1024 || file.Size > 2*1024*1024 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "File size must be between 10KB and 2MB",
		})
	}

	fileType := file.Header.Get("Content-Type")
	if fileType != "image/jpeg" && fileType != "image/jpg" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "File must be in *.jpg or *.jpeg format",
		})
	}

	uu := uploadsusecase.New(
		userrepository.New(i.DB), i.Uploader,
	)

	imageUrl, status, err := uu.Create(userId, file)

	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	data := returnData{
		ImageUrl: imageUrl,
	}

	return c.Status(fiber.StatusCreated).JSON(resultUpload{
		Message: "File uploaded sucessfully",
		Data:    data,
	})

}
