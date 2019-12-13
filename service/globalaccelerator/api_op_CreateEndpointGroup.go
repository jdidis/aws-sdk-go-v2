// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package globalaccelerator

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateEndpointGroupInput struct {
	_ struct{} `type:"structure"`

	// The list of endpoint objects.
	EndpointConfigurations []EndpointConfiguration `type:"list"`

	// The name of the AWS Region where the endpoint group is located. A listener
	// can have only one endpoint group in a specific Region.
	//
	// EndpointGroupRegion is a required field
	EndpointGroupRegion *string `type:"string" required:"true"`

	// The time—10 seconds or 30 seconds—between each health check for an endpoint.
	// The default value is 30.
	HealthCheckIntervalSeconds *int64 `min:"10" type:"integer"`

	// If the protocol is HTTP/S, then this specifies the path that is the destination
	// for health check targets. The default value is slash (/).
	HealthCheckPath *string `type:"string"`

	// The port that AWS Global Accelerator uses to check the health of endpoints
	// that are part of this endpoint group. The default port is the listener port
	// that this endpoint group is associated with. If listener port is a list of
	// ports, Global Accelerator uses the first port in the list.
	HealthCheckPort *int64 `min:"1" type:"integer"`

	// The protocol that AWS Global Accelerator uses to check the health of endpoints
	// that are part of this endpoint group. The default value is TCP.
	HealthCheckProtocol HealthCheckProtocol `type:"string" enum:"true"`

	// A unique, case-sensitive identifier that you provide to ensure the idempotency—that
	// is, the uniqueness—of the request.
	//
	// IdempotencyToken is a required field
	IdempotencyToken *string `type:"string" required:"true"`

	// The Amazon Resource Name (ARN) of the listener.
	//
	// ListenerArn is a required field
	ListenerArn *string `type:"string" required:"true"`

	// The number of consecutive health checks required to set the state of a healthy
	// endpoint to unhealthy, or to set an unhealthy endpoint to healthy. The default
	// value is 3.
	ThresholdCount *int64 `min:"1" type:"integer"`

	// The percentage of traffic to send to an AWS Region. Additional traffic is
	// distributed to other endpoint groups for this listener.
	//
	// Use this action to increase (dial up) or decrease (dial down) traffic to
	// a specific Region. The percentage is applied to the traffic that would otherwise
	// have been routed to the Region based on optimal routing.
	//
	// The default value is 100.
	TrafficDialPercentage *float64 `type:"float"`
}

// String returns the string representation
func (s CreateEndpointGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateEndpointGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateEndpointGroupInput"}

	if s.EndpointGroupRegion == nil {
		invalidParams.Add(aws.NewErrParamRequired("EndpointGroupRegion"))
	}
	if s.HealthCheckIntervalSeconds != nil && *s.HealthCheckIntervalSeconds < 10 {
		invalidParams.Add(aws.NewErrParamMinValue("HealthCheckIntervalSeconds", 10))
	}
	if s.HealthCheckPort != nil && *s.HealthCheckPort < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("HealthCheckPort", 1))
	}

	if s.IdempotencyToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("IdempotencyToken"))
	}

	if s.ListenerArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ListenerArn"))
	}
	if s.ThresholdCount != nil && *s.ThresholdCount < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("ThresholdCount", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateEndpointGroupOutput struct {
	_ struct{} `type:"structure"`

	// The information about the endpoint group that was created.
	EndpointGroup *EndpointGroup `type:"structure"`
}

// String returns the string representation
func (s CreateEndpointGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateEndpointGroup = "CreateEndpointGroup"

// CreateEndpointGroupRequest returns a request value for making API operation for
// AWS Global Accelerator.
//
// Create an endpoint group for the specified listener. An endpoint group is
// a collection of endpoints in one AWS Region. To see an AWS CLI example of
// creating an endpoint group, scroll down to Example.
//
//    // Example sending a request using CreateEndpointGroupRequest.
//    req := client.CreateEndpointGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/globalaccelerator-2018-08-08/CreateEndpointGroup
func (c *Client) CreateEndpointGroupRequest(input *CreateEndpointGroupInput) CreateEndpointGroupRequest {
	op := &aws.Operation{
		Name:       opCreateEndpointGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateEndpointGroupInput{}
	}

	req := c.newRequest(op, input, &CreateEndpointGroupOutput{})
	return CreateEndpointGroupRequest{Request: req, Input: input, Copy: c.CreateEndpointGroupRequest}
}

// CreateEndpointGroupRequest is the request type for the
// CreateEndpointGroup API operation.
type CreateEndpointGroupRequest struct {
	*aws.Request
	Input *CreateEndpointGroupInput
	Copy  func(*CreateEndpointGroupInput) CreateEndpointGroupRequest
}

// Send marshals and sends the CreateEndpointGroup API request.
func (r CreateEndpointGroupRequest) Send(ctx context.Context) (*CreateEndpointGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateEndpointGroupResponse{
		CreateEndpointGroupOutput: r.Request.Data.(*CreateEndpointGroupOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateEndpointGroupResponse is the response type for the
// CreateEndpointGroup API operation.
type CreateEndpointGroupResponse struct {
	*CreateEndpointGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateEndpointGroup request.
func (r *CreateEndpointGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}