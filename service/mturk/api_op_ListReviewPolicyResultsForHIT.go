// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mturk

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListReviewPolicyResultsForHITInput struct {
	_ struct{} `type:"structure"`

	// The unique identifier of the HIT to retrieve review results for.
	//
	// HITId is a required field
	HITId *string `min:"1" type:"string" required:"true"`

	// Limit the number of results returned.
	MaxResults *int64 `min:"1" type:"integer"`

	// Pagination token
	NextToken *string `min:"1" type:"string"`

	// The Policy Level(s) to retrieve review results for - HIT or Assignment. If
	// omitted, the default behavior is to retrieve all data for both policy levels.
	// For a list of all the described policies, see Review Policies.
	PolicyLevels []ReviewPolicyLevel `type:"list"`

	// Specify if the operation should retrieve a list of the actions taken executing
	// the Review Policies and their outcomes.
	RetrieveActions *bool `type:"boolean"`

	// Specify if the operation should retrieve a list of the results computed by
	// the Review Policies.
	RetrieveResults *bool `type:"boolean"`
}

// String returns the string representation
func (s ListReviewPolicyResultsForHITInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListReviewPolicyResultsForHITInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListReviewPolicyResultsForHITInput"}

	if s.HITId == nil {
		invalidParams.Add(aws.NewErrParamRequired("HITId"))
	}
	if s.HITId != nil && len(*s.HITId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("HITId", 1))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListReviewPolicyResultsForHITOutput struct {
	_ struct{} `type:"structure"`

	// The name of the Assignment-level Review Policy. This contains only the PolicyName
	// element.
	AssignmentReviewPolicy *ReviewPolicy `type:"structure"`

	// Contains both ReviewResult and ReviewAction elements for an Assignment.
	AssignmentReviewReport *ReviewReport `type:"structure"`

	// The HITId of the HIT for which results have been returned.
	HITId *string `min:"1" type:"string"`

	// The name of the HIT-level Review Policy. This contains only the PolicyName
	// element.
	HITReviewPolicy *ReviewPolicy `type:"structure"`

	// Contains both ReviewResult and ReviewAction elements for a particular HIT.
	HITReviewReport *ReviewReport `type:"structure"`

	// If the previous response was incomplete (because there is more data to retrieve),
	// Amazon Mechanical Turk returns a pagination token in the response. You can
	// use this pagination token to retrieve the next set of results.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListReviewPolicyResultsForHITOutput) String() string {
	return awsutil.Prettify(s)
}

const opListReviewPolicyResultsForHIT = "ListReviewPolicyResultsForHIT"

// ListReviewPolicyResultsForHITRequest returns a request value for making API operation for
// Amazon Mechanical Turk.
//
// The ListReviewPolicyResultsForHIT operation retrieves the computed results
// and the actions taken in the course of executing your Review Policies for
// a given HIT. For information about how to specify Review Policies when you
// call CreateHIT, see Review Policies. The ListReviewPolicyResultsForHIT operation
// can return results for both Assignment-level and HIT-level review results.
//
//    // Example sending a request using ListReviewPolicyResultsForHITRequest.
//    req := client.ListReviewPolicyResultsForHITRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mturk-requester-2017-01-17/ListReviewPolicyResultsForHIT
func (c *Client) ListReviewPolicyResultsForHITRequest(input *ListReviewPolicyResultsForHITInput) ListReviewPolicyResultsForHITRequest {
	op := &aws.Operation{
		Name:       opListReviewPolicyResultsForHIT,
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
		input = &ListReviewPolicyResultsForHITInput{}
	}

	req := c.newRequest(op, input, &ListReviewPolicyResultsForHITOutput{})
	return ListReviewPolicyResultsForHITRequest{Request: req, Input: input, Copy: c.ListReviewPolicyResultsForHITRequest}
}

// ListReviewPolicyResultsForHITRequest is the request type for the
// ListReviewPolicyResultsForHIT API operation.
type ListReviewPolicyResultsForHITRequest struct {
	*aws.Request
	Input *ListReviewPolicyResultsForHITInput
	Copy  func(*ListReviewPolicyResultsForHITInput) ListReviewPolicyResultsForHITRequest
}

// Send marshals and sends the ListReviewPolicyResultsForHIT API request.
func (r ListReviewPolicyResultsForHITRequest) Send(ctx context.Context) (*ListReviewPolicyResultsForHITResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListReviewPolicyResultsForHITResponse{
		ListReviewPolicyResultsForHITOutput: r.Request.Data.(*ListReviewPolicyResultsForHITOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListReviewPolicyResultsForHITRequestPaginator returns a paginator for ListReviewPolicyResultsForHIT.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListReviewPolicyResultsForHITRequest(input)
//   p := mturk.NewListReviewPolicyResultsForHITRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListReviewPolicyResultsForHITPaginator(req ListReviewPolicyResultsForHITRequest) ListReviewPolicyResultsForHITPaginator {
	return ListReviewPolicyResultsForHITPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListReviewPolicyResultsForHITInput
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

// ListReviewPolicyResultsForHITPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListReviewPolicyResultsForHITPaginator struct {
	aws.Pager
}

func (p *ListReviewPolicyResultsForHITPaginator) CurrentPage() *ListReviewPolicyResultsForHITOutput {
	return p.Pager.CurrentPage().(*ListReviewPolicyResultsForHITOutput)
}

// ListReviewPolicyResultsForHITResponse is the response type for the
// ListReviewPolicyResultsForHIT API operation.
type ListReviewPolicyResultsForHITResponse struct {
	*ListReviewPolicyResultsForHITOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListReviewPolicyResultsForHIT request.
func (r *ListReviewPolicyResultsForHITResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}