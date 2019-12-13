// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListComplianceItemsInput struct {
	_ struct{} `type:"structure"`

	// One or more compliance filters. Use a filter to return a more specific list
	// of results.
	Filters []ComplianceStringFilter `type:"list"`

	// The maximum number of items to return for this call. The call also returns
	// a token that you can specify in a subsequent call to get the next set of
	// results.
	MaxResults *int64 `min:"1" type:"integer"`

	// A token to start the list. Use this token to get the next set of results.
	NextToken *string `type:"string"`

	// The ID for the resources from which to get compliance information. Currently,
	// you can only specify one resource ID.
	ResourceIds []string `min:"1" type:"list"`

	// The type of resource from which to get compliance information. Currently,
	// the only supported resource type is ManagedInstance.
	ResourceTypes []string `min:"1" type:"list"`
}

// String returns the string representation
func (s ListComplianceItemsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListComplianceItemsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListComplianceItemsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.ResourceIds != nil && len(s.ResourceIds) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceIds", 1))
	}
	if s.ResourceTypes != nil && len(s.ResourceTypes) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceTypes", 1))
	}
	if s.Filters != nil {
		for i, v := range s.Filters {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Filters", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListComplianceItemsOutput struct {
	_ struct{} `type:"structure"`

	// A list of compliance information for the specified resource ID.
	ComplianceItems []ComplianceItem `type:"list"`

	// The token for the next set of items to return. Use this token to get the
	// next set of results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListComplianceItemsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListComplianceItems = "ListComplianceItems"

// ListComplianceItemsRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// For a specified resource ID, this API action returns a list of compliance
// statuses for different resource types. Currently, you can only specify one
// resource ID per call. List results depend on the criteria specified in the
// filter.
//
//    // Example sending a request using ListComplianceItemsRequest.
//    req := client.ListComplianceItemsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/ListComplianceItems
func (c *Client) ListComplianceItemsRequest(input *ListComplianceItemsInput) ListComplianceItemsRequest {
	op := &aws.Operation{
		Name:       opListComplianceItems,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListComplianceItemsInput{}
	}

	req := c.newRequest(op, input, &ListComplianceItemsOutput{})
	return ListComplianceItemsRequest{Request: req, Input: input, Copy: c.ListComplianceItemsRequest}
}

// ListComplianceItemsRequest is the request type for the
// ListComplianceItems API operation.
type ListComplianceItemsRequest struct {
	*aws.Request
	Input *ListComplianceItemsInput
	Copy  func(*ListComplianceItemsInput) ListComplianceItemsRequest
}

// Send marshals and sends the ListComplianceItems API request.
func (r ListComplianceItemsRequest) Send(ctx context.Context) (*ListComplianceItemsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListComplianceItemsResponse{
		ListComplianceItemsOutput: r.Request.Data.(*ListComplianceItemsOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListComplianceItemsResponse is the response type for the
// ListComplianceItems API operation.
type ListComplianceItemsResponse struct {
	*ListComplianceItemsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListComplianceItems request.
func (r *ListComplianceItemsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}