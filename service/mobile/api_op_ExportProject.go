// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mobile

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Request structure used in requests to export project configuration details.
type ExportProjectInput struct {
	_ struct{} `type:"structure"`

	// Unique project identifier.
	//
	// ProjectId is a required field
	ProjectId *string `location:"uri" locationName:"projectId" type:"string" required:"true"`
}

// String returns the string representation
func (s ExportProjectInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ExportProjectInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ExportProjectInput"}

	if s.ProjectId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProjectId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ExportProjectInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ProjectId != nil {
		v := *s.ProjectId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "projectId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Result structure used for requests to export project configuration details.
type ExportProjectOutput struct {
	_ struct{} `type:"structure"`

	// URL which can be used to download the exported project configuation file(s).
	DownloadUrl *string `locationName:"downloadUrl" type:"string"`

	// URL which can be shared to allow other AWS users to create their own project
	// in AWS Mobile Hub with the same configuration as the specified project. This
	// URL pertains to a snapshot in time of the project configuration that is created
	// when this API is called. If you want to share additional changes to your
	// project configuration, then you will need to create and share a new snapshot
	// by calling this method again.
	ShareUrl *string `locationName:"shareUrl" type:"string"`

	// Unique identifier for the exported snapshot of the project configuration.
	// This snapshot identifier is included in the share URL.
	SnapshotId *string `locationName:"snapshotId" type:"string"`
}

// String returns the string representation
func (s ExportProjectOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ExportProjectOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DownloadUrl != nil {
		v := *s.DownloadUrl

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "downloadUrl", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ShareUrl != nil {
		v := *s.ShareUrl

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "shareUrl", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SnapshotId != nil {
		v := *s.SnapshotId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "snapshotId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opExportProject = "ExportProject"

// ExportProjectRequest returns a request value for making API operation for
// AWS Mobile.
//
// Exports project configuration to a snapshot which can be downloaded and shared.
// Note that mobile app push credentials are encrypted in exported projects,
// so they can only be shared successfully within the same AWS account.
//
//    // Example sending a request using ExportProjectRequest.
//    req := client.ExportProjectRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mobile-2017-07-01/ExportProject
func (c *Client) ExportProjectRequest(input *ExportProjectInput) ExportProjectRequest {
	op := &aws.Operation{
		Name:       opExportProject,
		HTTPMethod: "POST",
		HTTPPath:   "/exports/{projectId}",
	}

	if input == nil {
		input = &ExportProjectInput{}
	}

	req := c.newRequest(op, input, &ExportProjectOutput{})
	return ExportProjectRequest{Request: req, Input: input, Copy: c.ExportProjectRequest}
}

// ExportProjectRequest is the request type for the
// ExportProject API operation.
type ExportProjectRequest struct {
	*aws.Request
	Input *ExportProjectInput
	Copy  func(*ExportProjectInput) ExportProjectRequest
}

// Send marshals and sends the ExportProject API request.
func (r ExportProjectRequest) Send(ctx context.Context) (*ExportProjectResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ExportProjectResponse{
		ExportProjectOutput: r.Request.Data.(*ExportProjectOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ExportProjectResponse is the response type for the
// ExportProject API operation.
type ExportProjectResponse struct {
	*ExportProjectOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ExportProject request.
func (r *ExportProjectResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}