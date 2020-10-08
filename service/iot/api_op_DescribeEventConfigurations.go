// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Describes event configurations.
func (c *Client) DescribeEventConfigurations(ctx context.Context, params *DescribeEventConfigurationsInput, optFns ...func(*Options)) (*DescribeEventConfigurationsOutput, error) {
	if params == nil {
		params = &DescribeEventConfigurationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeEventConfigurations", params, optFns, addOperationDescribeEventConfigurationsMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeEventConfigurationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeEventConfigurationsInput struct {
}

type DescribeEventConfigurationsOutput struct {

	// The creation date of the event configuration.
	CreationDate *time.Time

	// The event configurations.
	EventConfigurations map[string]*types.Configuration

	// The date the event configurations were last modified.
	LastModifiedDate *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeEventConfigurationsMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeEventConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeEventConfigurations{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEventConfigurations(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeEventConfigurations(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "DescribeEventConfigurations",
	}
}
