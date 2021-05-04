package controllers

import (
	"context"
	"fmt"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"main.go/config"
	"main.go/pkg/db/model"
	"net/http"
)

func GetTeacher(c echo.Context) error {
	email := c.QueryParam("email")
	var teacher bson.M
	collection := config.GetTeachersCollection()

	collection.FindOne(
		context.TODO(),
		bson.M{"email": email},
	).Decode(&teacher)
	if teacher["Status"] == "deleted" {
		return c.String(http.StatusOK, "teacher Not Found! :(")
	}
	if len(teacher) > 0 {
		return c.JSON(http.StatusOK, teacher)
	} else {
		return c.String(http.StatusOK, "teacher Not Found! :(")
	}
}

func AddTeacher(c echo.Context) error {
	teacher := &models.Teacher{
		Status: "created",
	}
	if err := c.Bind(teacher); err != nil {
		return err
	}

	collection := config.GetTeachersCollection()

	var teacher1 bson.M
	collection.FindOne(
		context.TODO(),
		bson.M{"email": teacher.Email},
	).Decode(&teacher1)
	fmt.Println(teacher1["email"])
	if len(teacher1) > 0 {
		return c.JSON(http.StatusOK, "Email address taken\n Use another :(")
	} else {
		result, err := collection.InsertOne(context.TODO(), teacher)
		log.Println(result.InsertedID)
		log.Println(err)
		var returnMessage string
		if err != nil {
			returnMessage = "Something went wrong! \nteacher addition Unsuccessful :("
			log.Fatal(err)
		} else if result.InsertedID == 0 {
			returnMessage = "Something went wrong! \nteacher addition Unsuccessful :("
		} else {
			returnMessage = "teacher: " + teacher.FullName + " added successfully to database :)"
		}
		return c.String(http.StatusCreated, returnMessage)
	}
}

func UpdateTeacher(c echo.Context) error {
	teacher:= &models.Teacher{}
	if err := c.Bind(teacher); err != nil {
		return err
	}

	email := teacher.Email
	collection := config.GetTeachersCollection()

	var teacher1 bson.M
	collection.FindOne(
		context.TODO(),
		bson.M{"email": email},
	).Decode(&teacher1)
	if teacher1["Status"] == "deleted" {
		return c.String(http.StatusOK, "Student Not Found! :(")
	}
	if len(teacher1) > 0 {

		result, err := collection.UpdateOne(
			context.TODO(),
			bson.M{"email": email},
			bson.D{
				{"$set", bson.D{{"fullName", teacher.FullName}}},
			})
		var returnMessage string
		if err != nil {
			returnMessage = "Something went wrong! \nUpdate Unsuccessful :("
			log.Fatal(err)
		} else if result.ModifiedCount == 0 {
			returnMessage = "teacher not found :("
		} else {
			returnMessage = teacher.Email + " updated successfully :)"
		}
		return c.String(http.StatusCreated, returnMessage)
	} else {
		return c.String(http.StatusOK, "teacher Not Found! :(")
	}
}

func DeleteTeacher(c echo.Context) error {
	email := c.QueryParam("email")

	collection := config.GetTeachersCollection()
	var teacher bson.M
	collection.FindOne(
		context.TODO(),
		bson.M{"email": email},
	).Decode(&teacher)
	fmt.Println(teacher["fullName"])
	if len(teacher) > 0 {
		_, err := collection.UpdateOne(
			context.TODO(),
			bson.M{"email": email},
			bson.D{
				{"$set", bson.D{{"Status", "deleted"}}},
			})

		var returnMessage string
		if err != nil {
			returnMessage = "Something went wrong! \nDelete Unsuccessful :("
			log.Fatal(err)
		} else {
			returnMessage = email + " deleted from database successfully :)"
		}
		return c.String(http.StatusCreated, returnMessage)
	} else {
		return c.String(http.StatusOK, "teacher Not Found! :(")
	}
}
