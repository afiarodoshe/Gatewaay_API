package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"io/ioutil"
	"log"
	"main.go/config"
	"main.go/model"
	"net/http"
	"os"
	"regexp"
)

func Get(c echo.Context) error {
	config.LoadEnvironments()
	email := c.QueryParam("Email")
	res, err := http.Get(os.Getenv("GET_REQUEST") + email)
	if err != nil {
		log.Fatal(err)
	}
	data, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	key := string(data)
	fmt.Printf(key)
	return c.JSON(http.StatusOK, key)
}

func Post(c echo.Context) error {
	config.LoadEnvironments()
	teacher := &models.Teacher{
	}
	if err := c.Bind(teacher); err != nil {
		return err
	}
	postBody, _ := json.Marshal(teacher)
	Email := teacher.Email
	if !isEmailValid(Email) {
		return c.JSON(http.StatusOK, "Invalid Email :(")
	}
	responseBody := bytes.NewBuffer(postBody)
	resp, err := http.Post(os.Getenv("POST_REQUEST"), "application/json", responseBody)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	key := string(body)
	log.Printf(key)
	return c.JSON(http.StatusOK, key)
}

func Put(c echo.Context) error{
	config.LoadEnvironments()
	teacher := &models.Teacher{
	}
	if err := c.Bind(teacher); err != nil {
		return err
	}
	postBody, err := json.Marshal(teacher)
	Email := teacher.Email
	responseBody := bytes.NewBuffer(postBody)
	req, err := http.NewRequest(http.MethodPut, os.Getenv("PUT_REQUEST") + Email, responseBody)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	body := string(bodyBytes)
	fmt.Println(body)
	return c.JSON(http.StatusOK, body)

}

func Delete(c echo.Context) error{
	config.LoadEnvironments()
	teacher := &models.Teacher{
	}
	if err := c.Bind(teacher); err != nil {
		return err
	}
	postBody, err := json.Marshal(teacher)
	Email := c.QueryParam("email")
	responseBody := bytes.NewBuffer(postBody)
	req, err := http.NewRequest(http.MethodDelete, os.Getenv("DELETE_REQUEST") + Email, responseBody)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	body := string(bodyBytes)
	fmt.Println(body)
	return c.JSON(http.StatusOK, body)

}

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func isEmailValid(e string) bool {
	if len(e) < 3 && len(e) > 254 {
		return false
	}
	return emailRegex.MatchString(e)
}
