package handler

import (
	"Nikhils-179/go-swagger/models"
	db "Nikhils-179/go-swagger/mongo-db-atlas"
	"Nikhils-179/go-swagger/response"
	"context"
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)


// GetUser returns a user by ID from the database
// @Summary Get a user by ID
// @Description Returns a user by ID from the database
// @Tags Users
// @Param id path int true "User ID"
// @Success 200 {object} response.JSONSuccessResult
// @Failure 400 {object} response.JSONBadReqResult
// @Failure 500 {object} response.JSONIntServerErrReqResult
// @Router /detail/{id} [get]
func GetUser(c *gin.Context) {
	id := c.Param("id")

	//convert and validate param
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.FailResponse(c, http.StatusBadRequest, "Wrong Parameter")
		return
	}

	user, err := fetchUser(idInt)
	if err != nil {
		response.FailResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.SuccessResponse(c, user)

}

// FetchUser is a helper function for GetUser
func fetchUser(id int64) (models.User, error) {
	var user models.User

	collection := db.Database.Collection("users")
	filter := bson.D{{"id", id}}

	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return user, errors.New("user not found")
		}
		return user, err
	}
	return user, nil
}

// @description CreateUser API endpoint creates User
// @tags Users
// @ummary Creates User
// @accept json
// @produce json
// @Param users body models.UserCreateParam true "User Data"
// @success 200 {object} response.JSONSuccessResult
// @failure 400 {object} response.JSONBadReqResult 
// @failure 500 {object} response.JSONIntServerErrReqResult
// @Router /create [post]
func CreateUser(c *gin.Context) {
	var param models.UserCreateParam

	err := c.BindJSON(&param)
	if err != nil {
		response.FailResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := saveUser(param)
	if err != nil {
		response.FailResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.SuccessResponse(c, models.User{
		ID:      id,
		Address: param.Address,
		Age:     param.Age,
		Name:    param.Name,
	})
}

//saveUser is helper function for CreateUser
func saveUser(newUser models.UserCreateParam) (int64, error) {
	collection := db.Database.Collection("users")
	user := models.User{
		ID:      time.Now().Unix(),
		Name:    newUser.Name,
		Age:     newUser.Age,
		Address: newUser.Address,
	}

	_, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}

// Update updates an existing user
// @Summary Updates an existing user
// @Description UpdateUser API endpoint updates an existing user. only the provided fields will be updated.
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body models.UserUpdateParam true "User Data"
// @Success 200 {object} response.JSONSuccessResult
// @Failure 400 {object} response.JSONBadReqResult
// @Failure 500 {object} response.JSONIntServerErrReqResult
// @Router /update/{id} [put]
func UpdateUser(c *gin.Context){
	id := c.Param("id")
	
	var param models.UserUpdateParam
	err := c.BindJSON(&param)
	if err != nil {
		response.FailResponse(c , http.StatusInternalServerError , " Failed to convert to json object")
		return
	}
	log.Printf("Updating user with ID: %s, Data: %+v\n", id, param) // Debugging log

	err = updateUser(param,id) 
	if err != nil {
		response.FailResponse(c,http.StatusMethodNotAllowed , "failed to call update user function")
		return
	}
	response.SuccessResponse(c , "User updated succesfully")
}

//updateUser is helper function for UpdateUser
func updateUser(updatedUsermodel models.UserUpdateParam , id string) error{
	collection := db.Database.Collection("users")
	//query the mongoDB user table to find value of id string

	userID , err := strconv.ParseInt(id , 10 , 64)
	if err != nil {
		return errors.New("invalid user id")
	}
	// filter to find the user by ID
	filter := bson.D{{"id", userID}}

	// Create update fields
	update := bson.D{}
	
	if updatedUsermodel.Name != nil{
		update = append(update , bson.E{"name" , *updatedUsermodel.Name})
	}
	if updatedUsermodel.Address != nil {
		update = append(update, bson.E{"address" , *updatedUsermodel.Address})
	}
	if updatedUsermodel.Age != nil {
		update = append(update , bson.E{"age" , *updatedUsermodel.Age})
	}

	if len(update) == 0 {
		return errors.New("no fields provided for update")
	}

	
	_, err = collection.UpdateOne(context.TODO(), filter, bson.D{{"$set", update}})

	if err != nil {
		return err
	}
	return nil
}