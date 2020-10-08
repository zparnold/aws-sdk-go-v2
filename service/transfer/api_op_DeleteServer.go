// Code generated by smithy-go-codegen DO NOT EDIT.

package transfer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the file transfer protocol-enabled server that you specify.  <p>No
// response returns from this operation.</p>
func (c *Client) DeleteServer(ctx context.Context, params *DeleteServerInput, optFns ...func(*Options)) (*DeleteServerOutput, error) {
	if params == nil {
		params = &DeleteServerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteServer", params, optFns, addOperationDeleteServerMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteServerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteServerInput struct {

	// A unique system-assigned identifier for a file transfer protocol-enabled server
	// instance.
	//
	// This member is required.
	ServerId *string
}

type DeleteServerOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteServerMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteServer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteServer{}, middleware.After)
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
	addOpDeleteServerValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteServer(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteServer(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transfer",
		OperationName: "DeleteServer",
	}
}
