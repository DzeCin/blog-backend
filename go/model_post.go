/*
 * Blog API
 *
 * This is a blog API
 *
 * API version: 1.0.0
 * Contact: dzenancindrak@gmail.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package blog

type Post struct {
	Id string `json:"id" bson:"_id"`

	Tags []string `json:"tags" bson:"tags"`

	Header string `json:"header" bson:"header"`

	Content string `json:"content" bson:"content"`

	Author string `json:"author" bson:"author"`

	DateCreated string `json:"dateCreated" bson:"dateCreated"`

	DateUpdated string `json:"dateUpdated" bson:"dateUpdated"`
}
