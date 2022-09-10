package BlogBackendTests

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/DzeCin/blog-backend/faker"
	sw "github.com/DzeCin/blog-backend/go"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

const nbPosts = 10

var (
	ctx context.Context
)

var newPost = sw.Post{
	Id:          "5066a748-9a72-404d-94f7-512c0779ff8e",
	Title:       "How to setup",
	Tags:        []string{"devops", "python", "golang"},
	Header:      "This is a blog about",
	Content:     "Does it work ?",
	Author:      "Dzenan",
	DateCreated: "2022-05-31T22:58:40.653Z",
	DateUpdated: "2022-05-31T22:58:40.653Z",
}

var newBadPost = sw.Post{
	Id:          "badUID",
	Title:       "How to setup",
	Tags:        []string{"devops", "python", "golang"},
	Header:      "This is a blog about",
	Content:     "Does it work ?",
	Author:      "Dzenan",
	DateCreated: "2022-05-31T22:58:40.653Z",
	DateUpdated: "2022-05-31T22:58:40.653Z",
}

var updatedPost = sw.Post{
	Id:          newPost.Id,
	Title:       "How to setup",
	Tags:        []string{"tag3", "tag4"},
	Header:      "Edited header",
	Content:     "Edited content",
	Author:      "Edited author",
	DateCreated: newPost.DateCreated,
	DateUpdated: "2023-05-31T22:58:40.653Z",
}

var updatedBadPost = sw.Post{
	Id:          newPost.Id,
	Title:       "How to setup",
	Tags:        []string{"tag3", "tag4"},
	Header:      "",
	Content:     "Edited content",
	Author:      "Edited author",
	DateCreated: newPost.DateCreated,
	DateUpdated: "2023-05-31T22:58:40.653Z",
}

func TestMain(m *testing.M) {

	log.Printf("Setting testing context")

	var (
		username      string
		password      string
		databaseHost  string
		databaseName  string
		oauthProvider string
		oidcClientID  string
	)

	err := godotenv.Load(".test.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	username = os.Getenv("DB_USERNAME")
	password = os.Getenv("DB_PASSWORD")
	databaseHost = os.Getenv("DB_HOST")
	databaseName = os.Getenv("DB_NAME")
	oauthProvider = os.Getenv("OAUTH_PROVIDER")
	oidcClientID = os.Getenv("OIDC_CLIENT_ID")

	DBkey := "db"
	OAuthProviderKey := "oauthProvider"
	ClientIDOIDC := "clientIDOIDC"

	// create db client

	client := sw.NewDatabaseCLI(username, password, databaseHost, databaseName)

	client.Collection("posts").DeleteMany(context.Background(), bson.D{{}})

	_, err = client.Collection("posts").InsertMany(context.Background(), BlogBackendFaker.PostGenerator(nbPosts))

	ctx = context.Background()

	ctx = context.WithValue(ctx, DBkey, client)
	ctx = context.WithValue(ctx, OAuthProviderKey, oauthProvider)
	ctx = context.WithValue(ctx, ClientIDOIDC, oidcClientID)

	code := m.Run()

	client.Collection("posts").DeleteMany(context.Background(), bson.D{{}})

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

func TestGetPostsHandlers(t *testing.T) {

	req, err := http.NewRequest("GET", "/posts", nil)
	if err != nil {
		t.Fatal(err)
	}

	if err != nil {
		t.Fatal(err)
	}

	var posts []sw.Post
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(sw.ContextHandler(sw.GetPosts, &ctx))

	handler.ServeHTTP(rr, req)

	json.Unmarshal(rr.Body.Bytes(), &posts)

	if len(posts) != nbPosts {
		t.Errorf("number of documents received is %v want %v", len(posts), nbPosts)
	}

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestAddPostHandlerPostNotExists(t *testing.T) {

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

// Test with no auth token

func TestAddPostHandlerPostUnauthorizedNoAuthToken(t *testing.T) {

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

	handler := http.HandlerFunc(sw.HandleAuthorization(sw.AddPost, &ctx))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestAddPostHandlerPostBadID(t *testing.T) {

	rawData, err := json.Marshal(newBadPost)

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

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestAddPostHandlerPostExists(t *testing.T) {

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

	if status := rr.Code; status != http.StatusConflict {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestUpdatePostHandlerPostExists(t *testing.T) {

	rawData, err := json.Marshal(updatedPost)

	if err != nil {
		panic(err)
	}

	body := bytes.NewReader(rawData)

	req, err := http.NewRequest("PATCH", "posts/"+updatedPost.Id, body)

	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(sw.ContextHandler(sw.UpdatePost, &ctx))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestUpdatePostHandlerBadPostExists(t *testing.T) {

	rawData, err := json.Marshal(updatedBadPost)

	if err != nil {
		panic(err)
	}

	body := bytes.NewReader(rawData)

	req, err := http.NewRequest("PATCH", "posts/"+updatedBadPost.Id, body)

	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(sw.ContextHandler(sw.UpdatePost, &ctx))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
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
