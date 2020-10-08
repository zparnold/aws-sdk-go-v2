// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchevents

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Tests whether the specified event pattern matches the provided event. Most
// services in AWS treat : or / as the same character in Amazon Resource Names
// (ARNs). However, EventBridge uses an exact match in event patterns and rules. Be
// sure to use the correct ARN characters when creating event patterns so that they
// match the ARN syntax in the event you want to match.
func (c *Client) TestEventPattern(ctx context.Context, params *TestEventPatternInput, optFns ...func(*Options)) (*TestEventPatternOutput, error) {
	if params == nil {
		params = &TestEventPatternInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "TestEventPattern", params, optFns, addOperationTestEventPatternMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*TestEventPatternOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type TestEventPatternInput struct {

	// The event, in JSON format, to test against the event pattern.
	//
	// This member is required.
	Event *string

	// The event pattern. For more information, see Events and Event Patterns
	// (https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-and-event-patterns.html)
	// in the Amazon EventBridge User Guide.
	//
	// This member is required.
	EventPattern *string
}

type TestEventPatternOutput struct {

	// Indicates whether the event matches the event pattern.
	Result *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationTestEventPatternMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpTestEventPattern{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpTestEventPattern{}, middleware.After)
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
	addOpTestEventPatternValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opTestEventPattern(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opTestEventPattern(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "events",
		OperationName: "TestEventPattern",
	}
}
