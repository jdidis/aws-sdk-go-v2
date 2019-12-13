// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rekognition

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteFacesInput struct {
	_ struct{} `type:"structure"`

	// Collection from which to remove the specific faces.
	//
	// CollectionId is a required field
	CollectionId *string `min:"1" type:"string" required:"true"`

	// An array of face IDs to delete.
	//
	// FaceIds is a required field
	FaceIds []string `min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s DeleteFacesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteFacesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteFacesInput"}

	if s.CollectionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("CollectionId"))
	}
	if s.CollectionId != nil && len(*s.CollectionId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CollectionId", 1))
	}

	if s.FaceIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("FaceIds"))
	}
	if s.FaceIds != nil && len(s.FaceIds) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("FaceIds", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteFacesOutput struct {
	_ struct{} `type:"structure"`

	// An array of strings (face IDs) of the faces that were deleted.
	DeletedFaces []string `min:"1" type:"list"`
}

// String returns the string representation
func (s DeleteFacesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteFaces = "DeleteFaces"

// DeleteFacesRequest returns a request value for making API operation for
// Amazon Rekognition.
//
// Deletes faces from a collection. You specify a collection ID and an array
// of face IDs to remove from the collection.
//
// This operation requires permissions to perform the rekognition:DeleteFaces
// action.
//
//    // Example sending a request using DeleteFacesRequest.
//    req := client.DeleteFacesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DeleteFacesRequest(input *DeleteFacesInput) DeleteFacesRequest {
	op := &aws.Operation{
		Name:       opDeleteFaces,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteFacesInput{}
	}

	req := c.newRequest(op, input, &DeleteFacesOutput{})
	return DeleteFacesRequest{Request: req, Input: input, Copy: c.DeleteFacesRequest}
}

// DeleteFacesRequest is the request type for the
// DeleteFaces API operation.
type DeleteFacesRequest struct {
	*aws.Request
	Input *DeleteFacesInput
	Copy  func(*DeleteFacesInput) DeleteFacesRequest
}

// Send marshals and sends the DeleteFaces API request.
func (r DeleteFacesRequest) Send(ctx context.Context) (*DeleteFacesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteFacesResponse{
		DeleteFacesOutput: r.Request.Data.(*DeleteFacesOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteFacesResponse is the response type for the
// DeleteFaces API operation.
type DeleteFacesResponse struct {
	*DeleteFacesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteFaces request.
func (r *DeleteFacesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}