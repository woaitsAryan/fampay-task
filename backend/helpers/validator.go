package helpers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	schemas "github.com/woaitsAryan/fampay-task/backend/schema"
)

var validate *validator.Validate

func ValidateVideoURLParams(c *fiber.Ctx) (*schemas.VideoFetchSchema, error) {
	validate = validator.New()

	schema := new(schemas.VideoFetchSchema)
	
	if schema.Limit == 0 {
		schema.Limit = 20
	}

	err := c.QueryParser(schema)
	if err != nil {
		return nil, err
	}

	err = validate.Struct(schema)
	if err != nil {
		return nil, err
	}

	return schema, nil
}
