package controllers

import (
	"context"
	"fiber-mongo-api/configs"
	"fiber-mongo-api/models"
	"fiber-mongo-api/responses"
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection("users")
var validate = validator.New()

// CreateUser godoc
//
//	@Summary		Create a user
//	@Description	Create a informed user in database
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			request	body		models.User	true	" "
//	@Success		200		{object}	responses.UserResponse{data=models.User}
//
//	@Failure		400		{object}	responses.UserResponse
//	@Failure		404		{object}	responses.UserResponse
//	@Failure		500		{object}	responses.UserResponse
//
//	@Router			/users [post]
func CreateUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user models.User
	defer cancel()

	//validate the request body
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: fiber.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validatorErr := validate.Struct(&user); validatorErr != nil {
		return c.Status(http.StatusBadRequest).JSON(
			responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: fiber.Map{"data": validatorErr.Error()}})
	}

	newUser := models.User{
		Id:       primitive.NewObjectID(),
		Name:     user.Name,
		Location: user.Location,
		Title:    user.Title,
	}

	_, err := userCollection.InsertOne(ctx, newUser)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusCreated).JSON(responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: fiber.Map{"data": newUser}})

}

// GetAUser godoc
//
//	@Summary		Get a user
//	@Description	Find a user by userId
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			userId	path		string	true	"User ID"
//	@Success		200		{object}	responses.UserResponse{data=models.User}
//
//	@Failure		400		{object}	responses.UserResponse
//	@Failure		404		{object}	responses.UserResponse
//	@Failure		500		{object}	responses.UserResponse
//
//	@Router			/users/{userId} [get]
func GetAUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	userId := c.Params("userId")
	var user models.User
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(userId)

	err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&user)
	if err != nil {
		fmt.Println(err)
		return c.Status(http.StatusNotFound).JSON(responses.UserResponse{Status: http.StatusNotFound, Message: "error", Data: fiber.Map{"data": "user not found"}})
	}

	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: fiber.Map{"data": user}})
}

// EdiAUser godoc
//
//	@Summary		Edit a user
//	@Description	Edit user data
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			userId	path		string		true	"User ID"
//	@Param			request	body		models.User	true	" "
//	@Success		200		{object}	responses.UserResponse{data=models.User}
//	@Failure		400		{object}	responses.UserResponse
//	@Failure		404		{object}	responses.UserResponse
//	@Failure		500		{object}	responses.UserResponse
//
//	@Router			/users/{userId} [put]
func EditAUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	userId := c.Params("userId")
	var user models.User
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(userId)

	//validate the request body
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: fiber.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&user); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: fiber.Map{"data": validationErr.Error()}})
	}

	update := bson.M{"name": user.Name, "location": user.Location, "title": user.Title}

	result, err := userCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: fiber.Map{"data": err.Error()}})
	}

	//get updated user details
	var updatedUser models.User
	if result.MatchedCount == 1 {
		err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedUser)

		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: fiber.Map{"data": err.Error()}})
		}
	}

	return c.Status(http.StatusOK).JSON(responses.UserResponse{Status: http.StatusOK, Message: "success", Data: fiber.Map{"data": updatedUser}})

}

// DeleteAUser godoc
//
//	@Summary		Delete a user
//	@Description	Delete a user data
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			userId	path		string	true	"User ID"
//	@Success		200		{object}	responses.UserResponse
//	@Failure		400		{object}	responses.UserResponse
//	@Failure		404		{object}	responses.UserResponse
//	@Failure		500		{object}	responses.UserResponse
//
//	@Router			/users/{userId} [delete]
func DeleteAUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	userId := c.Params("userId")
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(userId)

	result, err := userCollection.DeleteOne(ctx, bson.M{"id": objId})
	if err != nil {
		fmt.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: fiber.Map{"data": err.Error()}})
	}

	if result.DeletedCount < 1 {
		return c.Status(http.StatusNotFound).JSON(
			responses.UserResponse{Status: http.StatusNotFound, Message: "error", Data: fiber.Map{"data": "User with specified ID not found!"}},
		)
	}

	return c.Status(http.StatusOK).JSON(
		responses.UserResponse{Status: http.StatusOK, Message: "success", Data: fiber.Map{"data": "User successfully deleted!"}},
	)
}

// GetAllUsers godoc
//
//	@Summary		List users
//	@Description	List all of users
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		responses.UserResponse{data=[]models.User}
//	@Failure		400	{object}	responses.UserResponse
//	@Failure		404	{object}	responses.UserResponse
//	@Failure		500	{object}	responses.UserResponse
//
//	@Router			/users/ [get]
func GetAllUsers(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var users []models.User
	defer cancel()

	results, err := userCollection.Find(ctx, bson.M{})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: fiber.Map{"data": err.Error()}})
	}

	//reading from the db in an optimal way
	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleUser models.User
		if err = results.Decode(&singleUser); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: fiber.Map{"data": err.Error()}})
		}

		users = append(users, singleUser)
	}

	return c.Status(http.StatusOK).JSON(
		responses.UserResponse{Status: http.StatusOK, Message: "success", Data: fiber.Map{"data": users}},
	)
}
