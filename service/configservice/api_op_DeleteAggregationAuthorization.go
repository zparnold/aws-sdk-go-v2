// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the authorization granted to the specified configuration aggregator
// account in a specified region.
func (c *Client) DeleteAggregationAuthorization(ctx context.Context, params *DeleteAggregationAuthorizationInput, optFns ...func(*Options)) (*DeleteAggregationAuthorizationOutput, error) {
	stack := middleware.NewStack("DeleteAggregationAuthorization", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteAggregationAuthorizationMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	addRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpDeleteAggregationAuthorizationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteAggregationAuthorization(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)

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
			OperationName: "DeleteAggregationAuthorization",
			Err:           err,
		}
	}
	out := result.(*DeleteAggregationAuthorizationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteAggregationAuthorizationInput struct {

	// The 12-digit account ID of the account authorized to aggregate data.
	//
	// This member is required.
	AuthorizedAccountId *string

	// The region authorized to collect aggregated data.
	//
	// This member is required.
	AuthorizedAwsRegion *string
}

type DeleteAggregationAuthorizationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteAggregationAuthorizationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteAggregationAuthorization{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteAggregationAuthorization{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteAggregationAuthorization(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "DeleteAggregationAuthorization",
	}
}
