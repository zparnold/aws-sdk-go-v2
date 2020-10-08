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

// Sets the logging level.
func (c *Client) SetV2LoggingLevel(ctx context.Context, params *SetV2LoggingLevelInput, optFns ...func(*Options)) (*SetV2LoggingLevelOutput, error) {
	if params == nil {
		params = &SetV2LoggingLevelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SetV2LoggingLevel", params, optFns, addOperationSetV2LoggingLevelMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*SetV2LoggingLevelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SetV2LoggingLevelInput struct {

	// The log level.
	//
	// This member is required.
	LogLevel types.LogLevel

	// The log target.
	//
	// This member is required.
	LogTarget *types.LogTarget
}

type SetV2LoggingLevelOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationSetV2LoggingLevelMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSetV2LoggingLevel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSetV2LoggingLevel{}, middleware.After)
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
	addOpSetV2LoggingLevelValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSetV2LoggingLevel(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opSetV2LoggingLevel(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "SetV2LoggingLevel",
	}
}
