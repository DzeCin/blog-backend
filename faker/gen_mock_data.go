package BlogBackendFaker

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"

	sw "github.com/DzeCin/blog-backend/go"
	faker2 "github.com/ddosify/go-faker/faker"
	"github.com/grokify/html-strip-tags-go"
	"go.mongodb.org/mongo-driver/bson"
)

func PostGenerator(number int) []interface{} {
	faker := faker2.NewFaker()

	var posts []interface{}

	for i := 0; i < number; i++ {

		resp, err := http.Get("https://jaspervdj.be/lorem-markdownum/markdown-html.html")
		if err != nil {
			log.Fatalln(err)
		}

		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Fatalln(err)
		}

		strippedHtml := strip.StripTags(string(body))

		dateCreated := time.Unix(faker.CurrentTimestamp(), 0).Format("2006-01-02")
		dateUpdated := time.Unix(faker.CurrentTimestamp() + 3600 * 24 * 50, 0).Format("2006-01-02")

		var post = sw.Post{
			Id:          faker.RandomUUID().String(),
			Title:       faker.RandomJobTitle(),
			Tags:        []string{faker.RandomCatchPhraseAdjective(), faker.RandomCatchPhraseAdjective(), faker.RandomCatchPhraseAdjective()},
			Header:      faker.RandomLoremSentence(),
			Content:     strippedHtml,
			Author:      faker.RandomPersonFullName(),
			DateCreated: dateCreated,
			DateUpdated: dateUpdated,
		}

		toBson, err := bson.Marshal(post)
		if err != nil {
			panic(err)
		}

		posts = append(posts, toBson)

	}

	return posts

}
