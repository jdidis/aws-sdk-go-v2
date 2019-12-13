// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package athena

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListNamedQueriesInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of queries to return in this request.
	MaxResults *int64 `type:"integer"`

	// The token that specifies where to start pagination if a previous request
	// was truncated.
	NextToken *string `min:"1" type:"string"`

	// The name of the workgroup from which the named queries are being returned.
	WorkGroup *string `type:"string"`
}

// String returns the string representation
func (s ListNamedQueriesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListNamedQueriesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListNamedQueriesInput"}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListNamedQueriesOutput struct {
	_ struct{} `type:"structure"`

	// The list of unique query IDs.
	NamedQueryIds []string `min:"1" type:"list"`

	// A token to be used by the next request if this request is truncated.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListNamedQueriesOutput) String() string {
	return awsutil.Prettify(s)
}

const opListNamedQueries = "ListNamedQueries"

// ListNamedQueriesRequest returns a request value for making API operation for
// Amazon Athena.
//
// Provides a list of available query IDs only for queries saved in the specified
// workgroup. Requires that you have access to the workgroup.
//
// For code samples using the AWS SDK for Java, see Examples and Code Samples
// (http://docs.aws.amazon.com/athena/latest/ug/code-samples.html) in the Amazon
// Athena User Guide.
//
//    // Example sending a request using ListNamedQueriesRequest.
//    req := client.ListNamedQueriesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/athena-2017-05-18/ListNamedQueries
func (c *Client) ListNamedQueriesRequest(input *ListNamedQueriesInput) ListNamedQueriesRequest {
	op := &aws.Operation{
		Name:       opListNamedQueries,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListNamedQueriesInput{}
	}

	req := c.newRequest(op, input, &ListNamedQueriesOutput{})
	return ListNamedQueriesRequest{Request: req, Input: input, Copy: c.ListNamedQueriesRequest}
}

// ListNamedQueriesRequest is the request type for the
// ListNamedQueries API operation.
type ListNamedQueriesRequest struct {
	*aws.Request
	Input *ListNamedQueriesInput
	Copy  func(*ListNamedQueriesInput) ListNamedQueriesRequest
}

// Send marshals and sends the ListNamedQueries API request.
func (r ListNamedQueriesRequest) Send(ctx context.Context) (*ListNamedQueriesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListNamedQueriesResponse{
		ListNamedQueriesOutput: r.Request.Data.(*ListNamedQueriesOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListNamedQueriesRequestPaginator returns a paginator for ListNamedQueries.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListNamedQueriesRequest(input)
//   p := athena.NewListNamedQueriesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListNamedQueriesPaginator(req ListNamedQueriesRequest) ListNamedQueriesPaginator {
	return ListNamedQueriesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListNamedQueriesInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// ListNamedQueriesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListNamedQueriesPaginator struct {
	aws.Pager
}

func (p *ListNamedQueriesPaginator) CurrentPage() *ListNamedQueriesOutput {
	return p.Pager.CurrentPage().(*ListNamedQueriesOutput)
}

// ListNamedQueriesResponse is the response type for the
// ListNamedQueries API operation.
type ListNamedQueriesResponse struct {
	*ListNamedQueriesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListNamedQueries request.
func (r *ListNamedQueriesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}