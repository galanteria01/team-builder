package controllers

import (
	"context"
	"example/go-auth/configs"
	"example/go-auth/models"
	"example/go-auth/responses"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "user")

func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var user models.User
		defer cancel()

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponses{
				Status:  http.StatusBadRequest,
				Message: "Error during request",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		newUser := models.User{
			Id:    primitive.NewObjectID(),
			Name:  user.Name,
			Email: user.Email,
		}

		result, err := userCollection.InsertOne(ctx, newUser)

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponses{
				Status:  http.StatusInternalServerError,
				Message: "Internal Server Error",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		c.JSON(http.StatusCreated, responses.UserResponses{
			Status:  http.StatusCreated,
			Message: "Request Success",
			Data:    map[string]interface{}{"data": result},
		})
	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get a user code goes here
	}
}

func EditUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		//edit a user code goes here
	}
}

func DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		//delete a user code goes here
	}
}

func GetAllUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var users []models.User
		defer cancel()

		results, err := userCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponses{
				Status:  http.StatusInternalServerError,
				Message: "Error fetching data",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		defer results.Close(ctx)

		for results.Next(ctx) {
			var singleUser models.User
			if err = results.Decode(&singleUser); err != nil {
				c.JSON(http.StatusInternalServerError, responses.UserResponses{
					Status:  http.StatusInternalServerError,
					Message: "Error fetching data",
					Data:    map[string]interface{}{"data": err.Error()},
				})
			}
			users = append(users, singleUser)
		}
		c.JSON(http.StatusOK, responses.UserResponses{
			Status:  http.StatusOK,
			Message: "Successfully fetched users",
			Data:    map[string]interface{}{"data": users},
		})
	}
}