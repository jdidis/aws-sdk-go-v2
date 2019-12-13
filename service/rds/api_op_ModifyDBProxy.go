// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ModifyDBProxyInput struct {
	_ struct{} `type:"structure"`

	// The new authentication settings for the DBProxy.
	Auth []UserAuthConfig `type:"list"`

	// The identifier for the DBProxy to modify.
	//
	// DBProxyName is a required field
	DBProxyName *string `type:"string" required:"true"`

	// Whether the proxy includes detailed information about SQL statements in its
	// logs. This information helps you to debug issues involving SQL behavior or
	// the performance and scalability of the proxy connections. The debug information
	// includes the text of SQL statements that you submit through the proxy. Thus,
	// only enable this setting when needed for debugging, and only when you have
	// security measures in place to safeguard any sensitive information that appears
	// in the logs.
	DebugLogging *bool `type:"boolean"`

	// The number of seconds that a connection to the proxy can be inactive before
	// the proxy disconnects it. You can set this value higher or lower than the
	// connection timeout limit for the associated database.
	IdleClientTimeout *int64 `type:"integer"`

	// The new identifier for the DBProxy. An identifier must begin with a letter
	// and must contain only ASCII letters, digits, and hyphens; it can't end with
	// a hyphen or contain two consecutive hyphens.
	NewDBProxyName *string `type:"string"`

	// Whether Transport Layer Security (TLS) encryption is required for connections
	// to the proxy. By enabling this setting, you can enforce encrypted TLS connections
	// to the proxy, even if the associated database doesn't use TLS.
	RequireTLS *bool `type:"boolean"`

	// The Amazon Resource Name (ARN) of the IAM role that the proxy uses to access
	// secrets in AWS Secrets Manager.
	RoleArn *string `type:"string"`

	// The new list of security groups for the DBProxy.
	SecurityGroups []string `type:"list"`
}

// String returns the string representation
func (s ModifyDBProxyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyDBProxyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifyDBProxyInput"}

	if s.DBProxyName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBProxyName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ModifyDBProxyOutput struct {
	_ struct{} `type:"structure"`

	// The DBProxy object representing the new settings for the proxy.
	DBProxy *DBProxy `type:"structure"`
}

// String returns the string representation
func (s ModifyDBProxyOutput) String() string {
	return awsutil.Prettify(s)
}

const opModifyDBProxy = "ModifyDBProxy"

// ModifyDBProxyRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
//
// This is prerelease documentation for the RDS Database Proxy feature in preview
// release. It is subject to change.
//
// Changes the settings for an existing DB proxy.
//
//    // Example sending a request using ModifyDBProxyRequest.
//    req := client.ModifyDBProxyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/ModifyDBProxy
func (c *Client) ModifyDBProxyRequest(input *ModifyDBProxyInput) ModifyDBProxyRequest {
	op := &aws.Operation{
		Name:       opModifyDBProxy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyDBProxyInput{}
	}

	req := c.newRequest(op, input, &ModifyDBProxyOutput{})
	return ModifyDBProxyRequest{Request: req, Input: input, Copy: c.ModifyDBProxyRequest}
}

// ModifyDBProxyRequest is the request type for the
// ModifyDBProxy API operation.
type ModifyDBProxyRequest struct {
	*aws.Request
	Input *ModifyDBProxyInput
	Copy  func(*ModifyDBProxyInput) ModifyDBProxyRequest
}

// Send marshals and sends the ModifyDBProxy API request.
func (r ModifyDBProxyRequest) Send(ctx context.Context) (*ModifyDBProxyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifyDBProxyResponse{
		ModifyDBProxyOutput: r.Request.Data.(*ModifyDBProxyOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifyDBProxyResponse is the response type for the
// ModifyDBProxy API operation.
type ModifyDBProxyResponse struct {
	*ModifyDBProxyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifyDBProxy request.
func (r *ModifyDBProxyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}