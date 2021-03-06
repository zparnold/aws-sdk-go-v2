// Code generated by smithy-go-codegen DO NOT EDIT.

package codebuild

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codebuild/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets information about one or more build projects.
func (c *Client) BatchGetProjects(ctx context.Context, params *BatchGetProjectsInput, optFns ...func(*Options)) (*BatchGetProjectsOutput, error) {
	stack := middleware.NewStack("BatchGetProjects", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpBatchGetProjectsMiddlewares(stack)
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
	addOpBatchGetProjectsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opBatchGetProjects(options.Region), middleware.Before)
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
			OperationName: "BatchGetProjects",
			Err:           err,
		}
	}
	out := result.(*BatchGetProjectsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchGetProjectsInput struct {

	// The names or ARNs of the build projects. To get information about a project
	// shared with your AWS account, its ARN must be specified. You cannot specify a
	// shared project using its name.
	//
	// This member is required.
	Names []*string
}

type BatchGetProjectsOutput struct {

	// Information about the requested build projects.
	Projects []*types.Project

	// The names of build projects for which information could not be found.
	ProjectsNotFound []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpBatchGetProjectsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpBatchGetProjects{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchGetProjects{}, middleware.After)
}

func newServiceMetadataMiddleware_opBatchGetProjects(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codebuild",
		OperationName: "BatchGetProjects",
	}
}
