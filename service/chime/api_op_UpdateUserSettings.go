// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

type UpdateUserSettingsInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Chime account ID.
	//
	// AccountId is a required field
	AccountId *string `location:"uri" locationName:"accountId" type:"string" required:"true"`

	// The user ID.
	//
	// UserId is a required field
	UserId *string `location:"uri" locationName:"userId" type:"string" required:"true"`

	// The user settings to update.
	//
	// UserSettings is a required field
	UserSettings *UserSettings `type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdateUserSettingsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateUserSettingsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateUserSettingsInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}

	if s.UserId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserId"))
	}

	if s.UserSettings == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserSettings"))
	}
	if s.UserSettings != nil {
		if err := s.UserSettings.Validate(); err != nil {
			invalidParams.AddNested("UserSettings", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateUserSettingsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.UserSettings != nil {
		v := s.UserSettings

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "UserSettings", v, metadata)
	}
	if s.AccountId != nil {
		v := *s.AccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "accountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.UserId != nil {
		v := *s.UserId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "userId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type UpdateUserSettingsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateUserSettingsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateUserSettingsOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opUpdateUserSettings = "UpdateUserSettings"

// UpdateUserSettingsRequest returns a request value for making API operation for
// Amazon Chime.
//
// Updates the settings for the specified user, such as phone number settings.
//
//    // Example sending a request using UpdateUserSettingsRequest.
//    req := client.UpdateUserSettingsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/UpdateUserSettings
func (c *Client) UpdateUserSettingsRequest(input *UpdateUserSettingsInput) UpdateUserSettingsRequest {
	op := &aws.Operation{
		Name:       opUpdateUserSettings,
		HTTPMethod: "PUT",
		HTTPPath:   "/accounts/{accountId}/users/{userId}/settings",
	}

	if input == nil {
		input = &UpdateUserSettingsInput{}
	}

	req := c.newRequest(op, input, &UpdateUserSettingsOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return UpdateUserSettingsRequest{Request: req, Input: input, Copy: c.UpdateUserSettingsRequest}
}

// UpdateUserSettingsRequest is the request type for the
// UpdateUserSettings API operation.
type UpdateUserSettingsRequest struct {
	*aws.Request
	Input *UpdateUserSettingsInput
	Copy  func(*UpdateUserSettingsInput) UpdateUserSettingsRequest
}

// Send marshals and sends the UpdateUserSettings API request.
func (r UpdateUserSettingsRequest) Send(ctx context.Context) (*UpdateUserSettingsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateUserSettingsResponse{
		UpdateUserSettingsOutput: r.Request.Data.(*UpdateUserSettingsOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateUserSettingsResponse is the response type for the
// UpdateUserSettings API operation.
type UpdateUserSettingsResponse struct {
	*UpdateUserSettingsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateUserSettings request.
func (r *UpdateUserSettingsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}