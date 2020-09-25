// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the lifecycle configuration from the specified bucket. Amazon S3 removes
// all the lifecycle configuration rules in the lifecycle subresource associated
// with the bucket. Your objects never expire, and Amazon S3 no longer
// automatically deletes any objects on the basis of rules contained in the deleted
// lifecycle configuration. To use this operation, you must have permission to
// perform the s3:PutLifecycleConfiguration action. By default, the bucket owner
// has this permission and the bucket owner can grant this permission to others.
// <p>There is usually some time lag before lifecycle configuration deletion is
// fully propagated to all the Amazon S3 systems.</p> <p>For more information about
// the object expiration, see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/intro-lifecycle-rules.html#intro-lifecycle-rules-actions">Elements
// to Describe Lifecycle Actions</a>.</p> <p>Related actions include:</p> <ul> <li>
// <p> <a>PutBucketLifecycleConfiguration</a> </p> </li> <li> <p>
// <a>GetBucketLifecycleConfiguration</a> </p> </li> </ul>
func (c *Client) DeleteBucketLifecycle(ctx context.Context, params *DeleteBucketLifecycleInput, optFns ...func(*Options)) (*DeleteBucketLifecycleOutput, error) {
	stack := middleware.NewStack("DeleteBucketLifecycle", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpDeleteBucketLifecycleMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpDeleteBucketLifecycleValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteBucketLifecycle(options.Region), middleware.Before)
	addUpdateEndpointMiddleware(stack, options)
	addResponseErrorMiddleware(stack)
	addMetadataRetrieverMiddleware(stack)
	v4.AddContentSHA256HeaderMiddleware(stack)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "DeleteBucketLifecycle",
			Err:           err,
		}
	}
	out := result.(*DeleteBucketLifecycleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteBucketLifecycleInput struct {
	// The bucket name of the lifecycle to delete.
	Bucket *string
}

type DeleteBucketLifecycleOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpDeleteBucketLifecycleMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpDeleteBucketLifecycle{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpDeleteBucketLifecycle{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteBucketLifecycle(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "DeleteBucketLifecycle",
	}
}
