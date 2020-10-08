// Code generated by smithy-go-codegen DO NOT EDIT.

package iotsitewise

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Retrieves information about a project.
func (c *Client) DescribeProject(ctx context.Context, params *DescribeProjectInput, optFns ...func(*Options)) (*DescribeProjectOutput, error) {
	if params == nil {
		params = &DescribeProjectInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeProject", params, optFns, addOperationDescribeProjectMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeProjectOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeProjectInput struct {

	// The ID of the project.
	//
	// This member is required.
	ProjectId *string
}

type DescribeProjectOutput struct {

	// The ID of the portal that the project is in.
	//
	// This member is required.
	PortalId *string

	// The ARN
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of
	// the project, which has the following format.
	// arn:${Partition}:iotsitewise:${Region}:${Account}:project/${ProjectId}
	//
	// This member is required.
	ProjectArn *string

	// The date the project was created, in Unix epoch time.
	//
	// This member is required.
	ProjectCreationDate *time.Time

	// The ID of the project.
	//
	// This member is required.
	ProjectId *string

	// The date the project was last updated, in Unix epoch time.
	//
	// This member is required.
	ProjectLastUpdateDate *time.Time

	// The name of the project.
	//
	// This member is required.
	ProjectName *string

	// The project's description.
	ProjectDescription *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeProjectMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeProject{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeProject{}, middleware.After)
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
	addOpDescribeProjectValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeProject(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeProject(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotsitewise",
		OperationName: "DescribeProject",
	}
}
