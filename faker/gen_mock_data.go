package BlogBackendFaker

import (
	sw "github.com/DzeCin/blog-backend/go"
	faker2 "github.com/ddosify/go-faker/faker"
	"github.com/grokify/html-strip-tags-go"
	"go.mongodb.org/mongo-driver/bson"
	"io/ioutil"
	"log"
	"net/http"
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

		var post = sw.Post{
			Id:          faker.RandomUUID().String(),
			Title:       faker.RandomJobTitle(),
			Tags:        []string{faker.RandomCatchPhraseAdjective(), faker.RandomCatchPhraseAdjective(), faker.RandomCatchPhraseAdjective()},
			Header:      faker.RandomLoremSentence(),
			Content:     strippedHtml,
			Author:      faker.RandomPersonFullName(),
			DateCreated: faker.RandomDatePast(),
			DateUpdated: faker.RandomDateFuture(),
		}

		toBson, err := bson.Marshal(post)
		if err != nil {
			panic(err)
		}

		posts = append(posts, toBson)

	}

	return posts

}
