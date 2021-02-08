# Graphql-go-lambda


Lambda handler to work with [graphql-go](https://github.com/graphql-go/graphql) and AWS Serverless Lambda .


## Usage
- First : Handling your Schema 

```go
package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/graphql-go/graphql"
	lambdahandler "github.com/yasseralhendawy/graphql-go-lambda"
)

func lambdaPublicHandle(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Setting up you schema 
	publicQuery := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "PublicQuery",
			Fields: graphql.Fields{
				"hello": &graphql.Field{
					Type: graphql.String,
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return "world", nil
					},
				},
			},
		},
	)
	publicSchema, _ := graphql.NewSchema(
		graphql.SchemaConfig{
			Query: publicQuery,
		})

	// Usage of the handler 
	h := lambdahandler.Handler{
		Schema: publicSchema,
	}
	return h.Handle(ctx, request)
}
```
- Second the main function 
```go
func main() {

	lambda.Start(lambdaPublicHandle)
}

```

## References
- [graphql-go](https://godoc.org/github.com/graphql-go/graphql)
- [graphql-go/handler](https://godoc.org/github.com/graphql-go/handler) 
- [aws-lambda-go packages](https://godoc.org/github.com/aws/aws-lambda-go) 
