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

// Creates a Git repository as a resource in your Amazon SageMaker account. You can
// associate the repository with notebook instances so that you can use Git source
// control for the notebooks you create. The Git repository is a resource in your
// Amazon SageMaker account, so it can be associated with more than one notebook
// instance, and it persists independently from the lifecycle of any notebook
// instances it is associated with. The repository can be hosted either in AWS
// CodeCommit
// (https://docs.aws.amazon.com/codecommit/latest/userguide/welcome.html) or in any
// other Git repository.
func (c *Client) CreateCodeRepository(ctx context.Context, params *CreateCodeRepositoryInput, optFns ...func(*Options)) (*CreateCodeRepositoryOutput, error) {
	stack := middleware.NewStack("CreateCodeRepository", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateCodeRepositoryMiddlewares(stack)
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
	addOpCreateCodeRepositoryValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCodeRepository(options.Region), middleware.Before)
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
			OperationName: "CreateCodeRepository",
			Err:           err,
		}
	}
	out := result.(*CreateCodeRepositoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateCodeRepositoryInput struct {

	// The name of the Git repository. The name must have 1 to 63 characters. Valid
	// characters are a-z, A-Z, 0-9, and - (hyphen).
	//
	// This member is required.
	CodeRepositoryName *string

	// Specifies details about the repository, including the URL where the repository
	// is located, the default branch, and credentials to use to access the repository.
	//
	// This member is required.
	GitConfig *types.GitConfig
}

type CreateCodeRepositoryOutput struct {

	// The Amazon Resource Name (ARN) of the new repository.
	//
	// This member is required.
	CodeRepositoryArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateCodeRepositoryMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateCodeRepository{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateCodeRepository{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateCodeRepository(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "CreateCodeRepository",
	}
}
