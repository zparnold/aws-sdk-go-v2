// Code generated by smithy-go-codegen DO NOT EDIT.

package clouddirectory

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a directory. Only disabled directories can be deleted. A deleted
// directory cannot be undone. Exercise extreme caution when deleting directories.
func (c *Client) DeleteDirectory(ctx context.Context, params *DeleteDirectoryInput, optFns ...func(*Options)) (*DeleteDirectoryOutput, error) {
	if params == nil {
		params = &DeleteDirectoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteDirectory", params, optFns, addOperationDeleteDirectoryMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteDirectoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteDirectoryInput struct {

	// The ARN of the directory to delete.
	//
	// This member is required.
	DirectoryArn *string
}

type DeleteDirectoryOutput struct {

	// The ARN of the deleted directory.
	//
	// This member is required.
	DirectoryArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteDirectoryMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteDirectory{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteDirectory{}, middleware.After)
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
	addOpDeleteDirectoryValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteDirectory(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteDirectory(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "clouddirectory",
		OperationName: "DeleteDirectory",
	}
}
