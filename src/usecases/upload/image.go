package uploadsusecase

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (i sUploadUseCase) Create(userId int, file *multipart.FileHeader) (string, int, error) {
	// Validate user who make a request is admin or not
	userRole, err := i.userRepository.FindUserRole(userId)

	if err != nil {
		return "", fiber.StatusInternalServerError, err
	}

	if userRole != "admin" {
		return "", fiber.StatusUnauthorized, errors.New("only admin can this route")
	}

	// Generate a UUID for the file name
	uuid := uuid.New().String()
	ext := filepath.Ext(file.Filename)
	newFileName := fmt.Sprintf("%s%s", uuid, ext)

	// Define S3 bucket and key
	bucket := os.Getenv("AWS_S3_BUCKET_NAME")
	key := newFileName

	// Open the uploaded file
	fileContent, err := file.Open()

	if err != nil {
		return "", fiber.StatusInternalServerError, errors.New("error opening the file")
	}

	defer fileContent.Close()

	// Upload the file to S3
	result, err := i.uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		ACL:    "public-read",
		Body:   fileContent,
	})

	if err != nil {
		return "", fiber.StatusInternalServerError, err
	}

	return result.Location, fiber.StatusCreated, nil
}
