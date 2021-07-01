// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new trigger.
func (c *Client) CreateTrigger(ctx context.Context, params *CreateTriggerInput, optFns ...func(*Options)) (*CreateTriggerOutput, error) {
	if params == nil {
		params = &CreateTriggerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateTrigger", params, optFns, c.addOperationCreateTriggerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateTriggerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateTriggerInput struct {

	// The actions initiated by this trigger when it fires.
	//
	// This member is required.
	Actions []types.Action

	// The name of the trigger.
	//
	// This member is required.
	Name *string

	// The type of the new trigger.
	//
	// This member is required.
	Type types.TriggerType

	// A description of the new trigger.
	Description *string

	// A predicate to specify when the new trigger should fire. This field is required
	// when the trigger type is CONDITIONAL.
	Predicate *types.Predicate

	// A cron expression used to specify the schedule (see Time-Based Schedules for
	// Jobs and Crawlers
	// (https://docs.aws.amazon.com/glue/latest/dg/monitor-data-warehouse-schedule.html).
	// For example, to run something every day at 12:15 UTC, you would specify: cron(15
	// 12 * * ? *). This field is required when the trigger type is SCHEDULED.
	Schedule *string

	// Set to true to start SCHEDULED and CONDITIONAL triggers when created. True is
	// not supported for ON_DEMAND triggers.
	StartOnCreation bool

	// The tags to use with this trigger. You may use tags to limit access to the
	// trigger. For more information about tags in Glue, see Amazon Web Services Tags
	// in Glue (https://docs.aws.amazon.com/glue/latest/dg/monitor-tags.html) in the
	// developer guide.
	Tags map[string]string

	// The name of the workflow associated with the trigger.
	WorkflowName *string
}

type CreateTriggerOutput struct {

	// The name of the trigger.
	Name *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationCreateTriggerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateTrigger{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateTrigger{}, middleware.After)
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
	if err = addOpCreateTriggerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTrigger(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateTrigger(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "CreateTrigger",
	}
}
