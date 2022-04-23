package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/boro8eyy/go-restapi-struct/configs"
	"github.com/boro8eyy/go-restapi-struct/models"
	"github.com/boro8eyy/go-restapi-struct/responses"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var nameCollection *mongo.Collection = configs.GetCollection(configs.DB,"yourNameCollection")
// var validate = validator.New()

func GetNames(c *fiber.Ctx) error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    var names []models.NameModel
    defer cancel()
  
    results, err := nameCollection.Find(ctx, bson.M{})
  
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(responses.NameResponse{Status: http.StatusInternalServerError, Message: "error", Payload: &fiber.Map{"data": err.Error()}})
    }
  
    //reading from the db in an optimal way
    defer results.Close(ctx)
    for results.Next(ctx) {
        var singleName models.NameModel
        if err = results.Decode(&singleName); err != nil {
            return c.Status(http.StatusInternalServerError).JSON(responses.NameResponse{Status: http.StatusInternalServerError, Message: "error", Payload: &fiber.Map{"data": err.Error()}})
        }
      
        names = append(names, singleName)
    }
  
    return c.Status(http.StatusOK).JSON(
        responses.NameResponse{Status: http.StatusOK, Message: "success", Payload: &fiber.Map{"data": names}},
    )
}