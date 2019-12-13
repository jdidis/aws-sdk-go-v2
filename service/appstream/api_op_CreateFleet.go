// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appstream

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateFleetInput struct {
	_ struct{} `type:"structure"`

	// The desired capacity for the fleet.
	//
	// ComputeCapacity is a required field
	ComputeCapacity *ComputeCapacity `type:"structure" required:"true"`

	// The description to display.
	Description *string `type:"string"`

	// The amount of time that a streaming session remains active after users disconnect.
	// If users try to reconnect to the streaming session after a disconnection
	// or network interruption within this time interval, they are connected to
	// their previous session. Otherwise, they are connected to a new session with
	// a new streaming instance.
	//
	// Specify a value between 60 and 360000.
	DisconnectTimeoutInSeconds *int64 `type:"integer"`

	// The fleet name to display.
	DisplayName *string `type:"string"`

	// The name of the directory and organizational unit (OU) to use to join the
	// fleet to a Microsoft Active Directory domain.
	DomainJoinInfo *DomainJoinInfo `type:"structure"`

	// Enables or disables default internet access for the fleet.
	EnableDefaultInternetAccess *bool `type:"boolean"`

	// The fleet type.
	//
	// ALWAYS_ON
	//
	// Provides users with instant-on access to their apps. You are charged for
	// all running instances in your fleet, even if no users are streaming apps.
	//
	// ON_DEMAND
	//
	// Provide users with access to applications after they connect, which takes
	// one to two minutes. You are charged for instance streaming when users are
	// connected and a small hourly fee for instances that are not streaming apps.
	FleetType FleetType `type:"string" enum:"true"`

	// The Amazon Resource Name (ARN) of the IAM role to apply to the fleet. To
	// assume a role, a fleet instance calls the AWS Security Token Service (STS)
	// AssumeRole API operation and passes the ARN of the role to use. The operation
	// creates a new session with temporary credentials. AppStream 2.0 retrieves
	// the temporary credentials and creates the AppStream_Machine_Role credential
	// profile on the instance.
	//
	// For more information, see Using an IAM Role to Grant Permissions to Applications
	// and Scripts Running on AppStream 2.0 Streaming Instances (https://docs.aws.amazon.com/appstream2/latest/developerguide/using-iam-roles-to-grant-permissions-to-applications-scripts-streaming-instances.html)
	// in the Amazon AppStream 2.0 Administration Guide.
	IamRoleArn *string `type:"string"`

	// The amount of time that users can be idle (inactive) before they are disconnected
	// from their streaming session and the DisconnectTimeoutInSeconds time interval
	// begins. Users are notified before they are disconnected due to inactivity.
	// If they try to reconnect to the streaming session before the time interval
	// specified in DisconnectTimeoutInSeconds elapses, they are connected to their
	// previous session. Users are considered idle when they stop providing keyboard
	// or mouse input during their streaming session. File uploads and downloads,
	// audio in, audio out, and pixels changing do not qualify as user activity.
	// If users continue to be idle after the time interval in IdleDisconnectTimeoutInSeconds
	// elapses, they are disconnected.
	//
	// To prevent users from being disconnected due to inactivity, specify a value
	// of 0. Otherwise, specify a value between 60 and 3600. The default value is
	// 0.
	//
	// If you enable this feature, we recommend that you specify a value that corresponds
	// exactly to a whole number of minutes (for example, 60, 120, and 180). If
	// you don't do this, the value is rounded to the nearest minute. For example,
	// if you specify a value of 70, users are disconnected after 1 minute of inactivity.
	// If you specify a value that is at the midpoint between two different minutes,
	// the value is rounded up. For example, if you specify a value of 90, users
	// are disconnected after 2 minutes of inactivity.
	IdleDisconnectTimeoutInSeconds *int64 `type:"integer"`

	// The ARN of the public, private, or shared image to use.
	ImageArn *string `type:"string"`

	// The name of the image used to create the fleet.
	ImageName *string `min:"1" type:"string"`

	// The instance type to use when launching fleet instances. The following instance
	// types are available:
	//
	//    * stream.standard.medium
	//
	//    * stream.standard.large
	//
	//    * stream.compute.large
	//
	//    * stream.compute.xlarge
	//
	//    * stream.compute.2xlarge
	//
	//    * stream.compute.4xlarge
	//
	//    * stream.compute.8xlarge
	//
	//    * stream.memory.large
	//
	//    * stream.memory.xlarge
	//
	//    * stream.memory.2xlarge
	//
	//    * stream.memory.4xlarge
	//
	//    * stream.memory.8xlarge
	//
	//    * stream.graphics-design.large
	//
	//    * stream.graphics-design.xlarge
	//
	//    * stream.graphics-design.2xlarge
	//
	//    * stream.graphics-design.4xlarge
	//
	//    * stream.graphics-desktop.2xlarge
	//
	//    * stream.graphics-pro.4xlarge
	//
	//    * stream.graphics-pro.8xlarge
	//
	//    * stream.graphics-pro.16xlarge
	//
	// InstanceType is a required field
	InstanceType *string `min:"1" type:"string" required:"true"`

	// The maximum amount of time that a streaming session can remain active, in
	// seconds. If users are still connected to a streaming instance five minutes
	// before this limit is reached, they are prompted to save any open documents
	// before being disconnected. After this time elapses, the instance is terminated
	// and replaced by a new instance.
	//
	// Specify a value between 600 and 360000.
	MaxUserDurationInSeconds *int64 `type:"integer"`

	// A unique name for the fleet.
	//
	// Name is a required field
	Name *string `type:"string" required:"true"`

	// The tags to associate with the fleet. A tag is a key-value pair, and the
	// value is optional. For example, Environment=Test. If you do not specify a
	// value, Environment=.
	//
	// If you do not specify a value, the value is set to an empty string.
	//
	// Generally allowed characters are: letters, numbers, and spaces representable
	// in UTF-8, and the following special characters:
	//
	// _ . : / = + \ - @
	//
	// For more information, see Tagging Your Resources (https://docs.aws.amazon.com/appstream2/latest/developerguide/tagging-basic.html)
	// in the Amazon AppStream 2.0 Administration Guide.
	Tags map[string]string `min:"1" type:"map"`

	// The VPC configuration for the fleet.
	VpcConfig *VpcConfig `type:"structure"`
}

// String returns the string representation
func (s CreateFleetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateFleetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateFleetInput"}

	if s.ComputeCapacity == nil {
		invalidParams.Add(aws.NewErrParamRequired("ComputeCapacity"))
	}
	if s.ImageName != nil && len(*s.ImageName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ImageName", 1))
	}

	if s.InstanceType == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceType"))
	}
	if s.InstanceType != nil && len(*s.InstanceType) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("InstanceType", 1))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Tags != nil && len(s.Tags) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Tags", 1))
	}
	if s.ComputeCapacity != nil {
		if err := s.ComputeCapacity.Validate(); err != nil {
			invalidParams.AddNested("ComputeCapacity", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateFleetOutput struct {
	_ struct{} `type:"structure"`

	// Information about the fleet.
	Fleet *Fleet `type:"structure"`
}

// String returns the string representation
func (s CreateFleetOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateFleet = "CreateFleet"

// CreateFleetRequest returns a request value for making API operation for
// Amazon AppStream.
//
// Creates a fleet. A fleet consists of streaming instances that run a specified
// image.
//
//    // Example sending a request using CreateFleetRequest.
//    req := client.CreateFleetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appstream-2016-12-01/CreateFleet
func (c *Client) CreateFleetRequest(input *CreateFleetInput) CreateFleetRequest {
	op := &aws.Operation{
		Name:       opCreateFleet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateFleetInput{}
	}

	req := c.newRequest(op, input, &CreateFleetOutput{})
	return CreateFleetRequest{Request: req, Input: input, Copy: c.CreateFleetRequest}
}

// CreateFleetRequest is the request type for the
// CreateFleet API operation.
type CreateFleetRequest struct {
	*aws.Request
	Input *CreateFleetInput
	Copy  func(*CreateFleetInput) CreateFleetRequest
}

// Send marshals and sends the CreateFleet API request.
func (r CreateFleetRequest) Send(ctx context.Context) (*CreateFleetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateFleetResponse{
		CreateFleetOutput: r.Request.Data.(*CreateFleetOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateFleetResponse is the response type for the
// CreateFleet API operation.
type CreateFleetResponse struct {
	*CreateFleetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateFleet request.
func (r *CreateFleetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}