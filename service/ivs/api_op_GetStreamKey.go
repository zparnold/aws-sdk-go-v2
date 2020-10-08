// Code generated by smithy-go-codegen DO NOT EDIT.

package ivs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ivs/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets stream-key information for a specified ARN.
func (c *Client) GetStreamKey(ctx context.Context, params *GetStreamKeyInput, optFns ...func(*Options)) (*GetStreamKeyOutput, error) {
	if params == nil {
		params = &GetStreamKeyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetStreamKey", params, optFns, addOperationGetStreamKeyMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*GetStreamKeyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetStreamKeyInput struct {

	// ARN for the stream key to be retrieved.
	//
	// This member is required.
	Arn *string
}

type GetStreamKeyOutput struct {

	// Object specifying a stream key.
	StreamKey *types.StreamKey

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetStreamKeyMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetStreamKey{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetStreamKey{}, middleware.After)
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
	addOpGetStreamKeyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetStreamKey(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetStreamKey(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ivs",
		OperationName: "GetStreamKey",
	}
}
