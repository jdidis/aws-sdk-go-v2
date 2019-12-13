// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DeleteEventStreamInput struct {
	_ struct{} `type:"structure"`

	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"application-id" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteEventStreamInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteEventStreamInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteEventStreamInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteEventStreamInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "application-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteEventStreamOutput struct {
	_ struct{} `type:"structure" payload:"EventStream"`

	// Specifies settings for publishing event data to an Amazon Kinesis data stream
	// or an Amazon Kinesis Data Firehose delivery stream.
	//
	// EventStream is a required field
	EventStream *EventStream `type:"structure" required:"true"`
}

// String returns the string representation
func (s DeleteEventStreamOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteEventStreamOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.EventStream != nil {
		v := s.EventStream

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "EventStream", v, metadata)
	}
	return nil
}

const opDeleteEventStream = "DeleteEventStream"

// DeleteEventStreamRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Deletes the event stream for an application.
//
//    // Example sending a request using DeleteEventStreamRequest.
//    req := client.DeleteEventStreamRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/DeleteEventStream
func (c *Client) DeleteEventStreamRequest(input *DeleteEventStreamInput) DeleteEventStreamRequest {
	op := &aws.Operation{
		Name:       opDeleteEventStream,
		HTTPMethod: "DELETE",
		HTTPPath:   "/v1/apps/{application-id}/eventstream",
	}

	if input == nil {
		input = &DeleteEventStreamInput{}
	}

	req := c.newRequest(op, input, &DeleteEventStreamOutput{})
	return DeleteEventStreamRequest{Request: req, Input: input, Copy: c.DeleteEventStreamRequest}
}

// DeleteEventStreamRequest is the request type for the
// DeleteEventStream API operation.
type DeleteEventStreamRequest struct {
	*aws.Request
	Input *DeleteEventStreamInput
	Copy  func(*DeleteEventStreamInput) DeleteEventStreamRequest
}

// Send marshals and sends the DeleteEventStream API request.
func (r DeleteEventStreamRequest) Send(ctx context.Context) (*DeleteEventStreamResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteEventStreamResponse{
		DeleteEventStreamOutput: r.Request.Data.(*DeleteEventStreamOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteEventStreamResponse is the response type for the
// DeleteEventStream API operation.
type DeleteEventStreamResponse struct {
	*DeleteEventStreamOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteEventStream request.
func (r *DeleteEventStreamResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}