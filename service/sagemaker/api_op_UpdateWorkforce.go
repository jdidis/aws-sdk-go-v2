// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Use this operation to update your workforce. You can use this operation to
// require that workers use specific IP addresses to work on tasks and to update
// your OpenID Connect (OIDC) Identity Provider (IdP) workforce configuration. Use
// SourceIpConfig to restrict worker access to tasks to a specific range of IP
// addresses. You specify allowed IP addresses by creating a list of up to ten
// CIDRs (https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Subnets.html). By
// default, a workforce isn't restricted to specific IP addresses. If you specify a
// range of IP addresses, workers who attempt to access tasks using any IP address
// outside the specified range are denied and get a Not Found error message on the
// worker portal. Use OidcConfig to update the configuration of a workforce created
// using your own OIDC IdP. You can only update your OIDC IdP configuration when
// there are no work teams associated with your workforce. You can delete work
// teams using the operation. After restricting access to a range of IP addresses
// or updating your OIDC IdP configuration with this operation, you can view
// details about your update workforce using the operation. This operation only
// applies to private workforces.
func (c *Client) UpdateWorkforce(ctx context.Context, params *UpdateWorkforceInput, optFns ...func(*Options)) (*UpdateWorkforceOutput, error) {
	if params == nil {
		params = &UpdateWorkforceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateWorkforce", params, optFns, c.addOperationUpdateWorkforceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateWorkforceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateWorkforceInput struct {

	// The name of the private workforce that you want to update. You can find your
	// workforce name by using the operation.
	//
	// This member is required.
	WorkforceName *string

	// Use this parameter to update your OIDC Identity Provider (IdP) configuration for
	// a workforce made using your own IdP.
	OidcConfig *types.OidcConfig

	// A list of one to ten worker IP address ranges (CIDRs
	// (https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Subnets.html)) that can be
	// used to access tasks assigned to this workforce. Maximum: Ten CIDR values
	SourceIpConfig *types.SourceIpConfig
}

type UpdateWorkforceOutput struct {

	// A single private workforce. You can create one private work force in each Amazon
	// Web Services Region. By default, any workforce-related API operation used in a
	// specific region will apply to the workforce created in that region. To learn how
	// to create a private workforce, see Create a Private Workforce
	// (https://docs.aws.amazon.com/sagemaker/latest/dg/sms-workforce-create-private.html).
	//
	// This member is required.
	Workforce *types.Workforce

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationUpdateWorkforceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateWorkforce{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateWorkforce{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpUpdateWorkforceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateWorkforce(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opUpdateWorkforce(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "UpdateWorkforce",
	}
}
