// Code generated by smithy-go-codegen DO NOT EDIT.

package directoryservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Rejects a directory sharing request that was sent from the directory owner
// account.
func (c *Client) RejectSharedDirectory(ctx context.Context, params *RejectSharedDirectoryInput, optFns ...func(*Options)) (*RejectSharedDirectoryOutput, error) {
	if params == nil {
		params = &RejectSharedDirectoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RejectSharedDirectory", params, optFns, addOperationRejectSharedDirectoryMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*RejectSharedDirectoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RejectSharedDirectoryInput struct {

	// Identifier of the shared directory in the directory consumer account. This
	// identifier is different for each directory owner account.
	//
	// This member is required.
	SharedDirectoryId *string
}

type RejectSharedDirectoryOutput struct {

	// Identifier of the shared directory in the directory consumer account.
	SharedDirectoryId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRejectSharedDirectoryMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRejectSharedDirectory{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRejectSharedDirectory{}, middleware.After)
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
	addOpRejectSharedDirectoryValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRejectSharedDirectory(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opRejectSharedDirectory(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ds",
		OperationName: "RejectSharedDirectory",
	}
}
