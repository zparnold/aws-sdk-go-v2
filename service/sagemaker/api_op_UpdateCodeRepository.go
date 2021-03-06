// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the specified Git repository with the specified values.
func (c *Client) UpdateCodeRepository(ctx context.Context, params *UpdateCodeRepositoryInput, optFns ...func(*Options)) (*UpdateCodeRepositoryOutput, error) {
	stack := middleware.NewStack("UpdateCodeRepository", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateCodeRepositoryMiddlewares(stack)
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
	addOpUpdateCodeRepositoryValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateCodeRepository(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)

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
			OperationName: "UpdateCodeRepository",
			Err:           err,
		}
	}
	out := result.(*UpdateCodeRepositoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateCodeRepositoryInput struct {

	// The name of the Git repository to update.
	//
	// This member is required.
	CodeRepositoryName *string

	// The configuration of the git repository, including the URL and the Amazon
	// Resource Name (ARN) of the AWS Secrets Manager secret that contains the
	// credentials used to access the repository. The secret must have a staging label
	// of AWSCURRENT and must be in the following format: {"username": UserName,
	// "password": Password}
	GitConfig *types.GitConfigForUpdate
}

type UpdateCodeRepositoryOutput struct {

	// The ARN of the Git repository.
	//
	// This member is required.
	CodeRepositoryArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateCodeRepositoryMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateCodeRepository{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateCodeRepository{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateCodeRepository(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "UpdateCodeRepository",
	}
}
