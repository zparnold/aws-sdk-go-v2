// Code generated by smithy-go-codegen DO NOT EDIT.

package emr

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Provides more detail about the cluster step.
func (c *Client) DescribeStep(ctx context.Context, params *DescribeStepInput, optFns ...func(*Options)) (*DescribeStepOutput, error) {
	if params == nil {
		params = &DescribeStepInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeStep", params, optFns, addOperationDescribeStepMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeStepOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// This input determines which step to describe.
type DescribeStepInput struct {

	// The identifier of the cluster with steps to describe.
	//
	// This member is required.
	ClusterId *string

	// The identifier of the step to describe.
	//
	// This member is required.
	StepId *string
}

// This output contains the description of the cluster step.
type DescribeStepOutput struct {

	// The step details for the requested step identifier.
	Step *types.Step

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeStepMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeStep{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeStep{}, middleware.After)
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
	addOpDescribeStepValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeStep(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeStep(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticmapreduce",
		OperationName: "DescribeStep",
	}
}
