package lambdahandler

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/graphql-go/graphql"
)

//Handler definition for handle function
type Handler struct {
	Schema graphql.Schema
}

//GraphQLRequest for lambda
type GraphQLRequest struct {
	Query         string                 `json:"query"`
	OperationName string                 `json:"operationName"`
	Variables     map[string]interface{} `json:"variables"`
}

//Handle to handle Lambda function request and response
func (h Handler) Handle(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	graphQLRequest := func(r events.APIGatewayProxyRequest) GraphQLRequest {
		query := r.QueryStringParameters["query"]
		operationName := r.QueryStringParameters["operationName"]
		variablesString := r.QueryStringParameters["variables"]
		variables := make(map[string]interface{})
		_ = json.Unmarshal([]byte(variablesString), &variables)
		graphqlRequest := GraphQLRequest{
			Query:         query,
			OperationName: operationName,
			Variables:     variables,
		}
		if query == "" {
			var graphqlRequest GraphQLRequest
			_ = json.Unmarshal([]byte(r.Body), &graphqlRequest)
			return graphqlRequest
		}
		return graphqlRequest
	}(request)

	var p = graphql.Params{
		Schema:         h.Schema,
		RequestString:  graphQLRequest.Query,
		OperationName:  graphQLRequest.OperationName,
		VariableValues: graphQLRequest.Variables,
		Context:        ctx,
	}
	var body string
	jsonBody, _ := json.Marshal(graphql.Do(p))
	body = string(jsonBody)
	return events.APIGatewayProxyResponse{
		Body:       body,
		StatusCode: 200,
		Headers: map[string]string{
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Methods": "POST, GET, OPTIONS, PUT, DELETE",
			"Access-Control-Allow-Headers": "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization",
		},
	}, nil
}
