// Code generated by smithy-go-codegen DO NOT EDIT.

package efs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the FileSystemPolicy for the specified file system. The default
// FileSystemPolicy goes into effect once the existing policy is deleted. For more
// information about the default file system policy, see Using Resource-based
// Policies with EFS
// (https://docs.aws.amazon.com/efs/latest/ug/res-based-policies-efs.html). This
// operation requires permissions for the elasticfilesystem:DeleteFileSystemPolicy
// action.
func (c *Client) DeleteFileSystemPolicy(ctx context.Context, params *DeleteFileSystemPolicyInput, optFns ...func(*Options)) (*DeleteFileSystemPolicyOutput, error) {
	stack := middleware.NewStack("DeleteFileSystemPolicy", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDeleteFileSystemPolicyMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpDeleteFileSystemPolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteFileSystemPolicy(options.Region), middleware.Before)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "DeleteFileSystemPolicy",
			Err:           err,
		}
	}
	out := result.(*DeleteFileSystemPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteFileSystemPolicyInput struct {
	// Specifies the EFS file system for which to delete the FileSystemPolicy.
	FileSystemId *string
}

type DeleteFileSystemPolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDeleteFileSystemPolicyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDeleteFileSystemPolicy{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteFileSystemPolicy{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteFileSystemPolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticfilesystem",
		OperationName: "DeleteFileSystemPolicy",
	}
}