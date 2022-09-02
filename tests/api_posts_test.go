package blogBackendTests

import (
	sw "blog-backend/go"
	"bytes"
	"context"
	"encoding/json"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var (
	client *mongo.Database
	ctx    context.Context
)

var newPost = sw.Post{
	Id:      "25",
	UserId:  "dzenan",
	Header:  "This is a blog about",
	Content: "Does it work ?",
	Author:  "Dzenan",
	Date:    "25 Janv",
}

func TestMain(m *testing.M) {

	log.Printf("Setting testing context")

	var (
		username     string
		password     string
		databaseHost string
		databaseName string
	)

	err := godotenv.Load(".test.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	username = os.Getenv("DB_USERNAME")
	password = os.Getenv("DB_PASSWORD")
	databaseHost = os.Getenv("DB_HOST")
	databaseName = os.Getenv("DB_NAME")

	key := "db"

	// create db client

	client := sw.NewDatabaseCLI(username, password, databaseHost, databaseName)

	ctx = context.Background()

	ctx = context.WithValue(ctx, key, client)

	code := m.Run()

	os.Exit(code)

}

func TestHealthCheckHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(sw.ContextHandler(sw.HealthCheck, &ctx))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestAddPostHandler(t *testing.T) {

	rawData, err := json.Marshal(newPost)

	if err != nil {
		panic(err)
	}

	body := bytes.NewReader(rawData)

	req, err := http.NewRequest("POST", "posts", body)

	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(sw.ContextHandler(sw.AddPost, &ctx))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestDeletePostHandler(t *testing.T) {

	req, err := http.NewRequest("DELETE", "/posts/"+newPost.Id, nil)
	if err != nil {
		t.Fatal(err)
	}

	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(sw.ContextHandler(sw.DeletePost, &ctx))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}
