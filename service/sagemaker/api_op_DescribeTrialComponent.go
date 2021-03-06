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
	"time"
)

// Provides a list of a trials component's properties.
func (c *Client) DescribeTrialComponent(ctx context.Context, params *DescribeTrialComponentInput, optFns ...func(*Options)) (*DescribeTrialComponentOutput, error) {
	stack := middleware.NewStack("DescribeTrialComponent", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeTrialComponentMiddlewares(stack)
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
	addOpDescribeTrialComponentValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTrialComponent(options.Region), middleware.Before)
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
			OperationName: "DescribeTrialComponent",
			Err:           err,
		}
	}
	out := result.(*DescribeTrialComponentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeTrialComponentInput struct {

	// The name of the trial component to describe.
	//
	// This member is required.
	TrialComponentName *string
}

type DescribeTrialComponentOutput struct {

	// Who created the component.
	CreatedBy *types.UserContext

	// When the component was created.
	CreationTime *time.Time

	// The name of the component as displayed. If DisplayName isn't specified,
	// TrialComponentName is displayed.
	DisplayName *string

	// When the component ended.
	EndTime *time.Time

	// The input artifacts of the component.
	InputArtifacts map[string]*types.TrialComponentArtifact

	// Who last modified the component.
	LastModifiedBy *types.UserContext

	// When the component was last modified.
	LastModifiedTime *time.Time

	// The metrics for the component.
	Metrics []*types.TrialComponentMetricSummary

	// The output artifacts of the component.
	OutputArtifacts map[string]*types.TrialComponentArtifact

	// The hyperparameters of the component.
	Parameters map[string]*types.TrialComponentParameterValue

	// The Amazon Resource Name (ARN) of the source and, optionally, the job type.
	Source *types.TrialComponentSource

	// When the component started.
	StartTime *time.Time

	// The status of the component. States include:
	//
	//     * InProgress
	//
	//     *
	// Completed
	//
	//     * Failed
	Status *types.TrialComponentStatus

	// The Amazon Resource Name (ARN) of the trial component.
	TrialComponentArn *string

	// The name of the trial component.
	TrialComponentName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeTrialComponentMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeTrialComponent{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeTrialComponent{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeTrialComponent(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "DescribeTrialComponent",
	}
}
