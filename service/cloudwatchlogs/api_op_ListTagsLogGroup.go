// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchlogs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the tags for the specified log group.
func (c *Client) ListTagsLogGroup(ctx context.Context, params *ListTagsLogGroupInput, optFns ...func(*Options)) (*ListTagsLogGroupOutput, error) {
	if params == nil {
		params = &ListTagsLogGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTagsLogGroup", params, optFns, addOperationListTagsLogGroupMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTagsLogGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTagsLogGroupInput struct {

	// The name of the log group.
	//
	// This member is required.
	LogGroupName *string
}

type ListTagsLogGroupOutput struct {

	// The tags for the log group.
	Tags map[string]*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListTagsLogGroupMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListTagsLogGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListTagsLogGroup{}, middleware.After)
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
	addOpListTagsLogGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListTagsLogGroup(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListTagsLogGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "logs",
		OperationName: "ListTagsLogGroup",
	}
}
