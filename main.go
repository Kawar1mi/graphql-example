package main

import (
	"context"
	"log"

	"github.com/machinebox/graphql"
)

type ResponseStruct struct {
	User struct {
		Publication struct {
			Posts []struct {
				Id        string `json:"_id"`
				Title     string
				dateAdded string
			}
		}
	}
}

func main() {
	client := graphql.NewClient("https://api.hashnode.com/graphql")

	// make a request
	req := graphql.NewRequest(`
	query getUserPosts($username: String!) {
		user(username: $username) {
		  publication {
			posts {
			  _id
			  title
			  dateAdded
			}
		  }
		}
	  }
`)

	// set any variables
	req.Var("username", "codingpastor")

	// set header fields
	req.Header.Set("Cache-Control", "no-cache")

	// define a Context for the request
	ctx := context.Background()

	// run it and capture the response
	var respData ResponseStruct
	if err := client.Run(ctx, req, &respData); err != nil {
		log.Fatal(err)
	}

	for _, el := range respData.User.Publication.Posts {
		log.Println(el.Title)
	}

}
