// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a list of the work teams that you are subscribed to in the Amazon Web
// Services Marketplace. The list may be empty if no work team satisfies the filter
// specified in the NameContains parameter.
func (c *Client) ListSubscribedWorkteams(ctx context.Context, params *ListSubscribedWorkteamsInput, optFns ...func(*Options)) (*ListSubscribedWorkteamsOutput, error) {
	if params == nil {
		params = &ListSubscribedWorkteamsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSubscribedWorkteams", params, optFns, c.addOperationListSubscribedWorkteamsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSubscribedWorkteamsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSubscribedWorkteamsInput struct {

	// The maximum number of work teams to return in each page of the response.
	MaxResults *int32

	// A string in the work team name. This filter returns only work teams whose name
	// contains the specified string.
	NameContains *string

	// If the result of the previous ListSubscribedWorkteams request was truncated, the
	// response includes a NextToken. To retrieve the next set of labeling jobs, use
	// the token in the next request.
	NextToken *string
}

type ListSubscribedWorkteamsOutput struct {

	// An array of Workteam objects, each describing a work team.
	//
	// This member is required.
	SubscribedWorkteams []types.SubscribedWorkteam

	// If the response is truncated, Amazon SageMaker returns this token. To retrieve
	// the next set of work teams, use it in the subsequent request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationListSubscribedWorkteamsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListSubscribedWorkteams{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListSubscribedWorkteams{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSubscribedWorkteams(options.Region), middleware.Before); err != nil {
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

// ListSubscribedWorkteamsAPIClient is a client that implements the
// ListSubscribedWorkteams operation.
type ListSubscribedWorkteamsAPIClient interface {
	ListSubscribedWorkteams(context.Context, *ListSubscribedWorkteamsInput, ...func(*Options)) (*ListSubscribedWorkteamsOutput, error)
}

var _ ListSubscribedWorkteamsAPIClient = (*Client)(nil)

// ListSubscribedWorkteamsPaginatorOptions is the paginator options for
// ListSubscribedWorkteams
type ListSubscribedWorkteamsPaginatorOptions struct {
	// The maximum number of work teams to return in each page of the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSubscribedWorkteamsPaginator is a paginator for ListSubscribedWorkteams
type ListSubscribedWorkteamsPaginator struct {
	options   ListSubscribedWorkteamsPaginatorOptions
	client    ListSubscribedWorkteamsAPIClient
	params    *ListSubscribedWorkteamsInput
	nextToken *string
	firstPage bool
}

// NewListSubscribedWorkteamsPaginator returns a new
// ListSubscribedWorkteamsPaginator
func NewListSubscribedWorkteamsPaginator(client ListSubscribedWorkteamsAPIClient, params *ListSubscribedWorkteamsInput, optFns ...func(*ListSubscribedWorkteamsPaginatorOptions)) *ListSubscribedWorkteamsPaginator {
	if params == nil {
		params = &ListSubscribedWorkteamsInput{}
	}

	options := ListSubscribedWorkteamsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSubscribedWorkteamsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSubscribedWorkteamsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListSubscribedWorkteams page.
func (p *ListSubscribedWorkteamsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSubscribedWorkteamsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListSubscribedWorkteams(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListSubscribedWorkteams(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "ListSubscribedWorkteams",
	}
}
