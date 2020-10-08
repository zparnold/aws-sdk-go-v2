// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists logging levels.
func (c *Client) ListV2LoggingLevels(ctx context.Context, params *ListV2LoggingLevelsInput, optFns ...func(*Options)) (*ListV2LoggingLevelsOutput, error) {
	if params == nil {
		params = &ListV2LoggingLevelsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListV2LoggingLevels", params, optFns, addOperationListV2LoggingLevelsMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*ListV2LoggingLevelsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListV2LoggingLevelsInput struct {

	// The maximum number of results to return at one time.
	MaxResults *int32

	// The token used to get the next set of results, or null if there are no
	// additional results.
	NextToken *string

	// The type of resource for which you are configuring logging. Must be THING_Group.
	TargetType types.LogTargetType
}

type ListV2LoggingLevelsOutput struct {

	// The logging configuration for a target.
	LogTargetConfigurations []*types.LogTargetConfiguration

	// The token used to get the next set of results, or null if there are no
	// additional results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListV2LoggingLevelsMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListV2LoggingLevels{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListV2LoggingLevels{}, middleware.After)
	if err != nil {
		return err
	}
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListV2LoggingLevels(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListV2LoggingLevels(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "ListV2LoggingLevels",
	}
}
