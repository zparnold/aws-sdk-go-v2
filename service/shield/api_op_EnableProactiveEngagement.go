// Code generated by smithy-go-codegen DO NOT EDIT.

package shield

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Authorizes the DDoS Response Team (DRT) to use email and phone to notify
// contacts about escalations to the DRT and to initiate proactive customer
// support.
func (c *Client) EnableProactiveEngagement(ctx context.Context, params *EnableProactiveEngagementInput, optFns ...func(*Options)) (*EnableProactiveEngagementOutput, error) {
	if params == nil {
		params = &EnableProactiveEngagementInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "EnableProactiveEngagement", params, optFns, addOperationEnableProactiveEngagementMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*EnableProactiveEngagementOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type EnableProactiveEngagementInput struct {
}

type EnableProactiveEngagementOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationEnableProactiveEngagementMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpEnableProactiveEngagement{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpEnableProactiveEngagement{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opEnableProactiveEngagement(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opEnableProactiveEngagement(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "shield",
		OperationName: "EnableProactiveEngagement",
	}
}
