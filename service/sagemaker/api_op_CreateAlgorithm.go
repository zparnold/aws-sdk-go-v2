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

// Create a machine learning algorithm that you can use in Amazon SageMaker and
// list in the AWS Marketplace.
func (c *Client) CreateAlgorithm(ctx context.Context, params *CreateAlgorithmInput, optFns ...func(*Options)) (*CreateAlgorithmOutput, error) {
	stack := middleware.NewStack("CreateAlgorithm", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateAlgorithmMiddlewares(stack)
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
	addOpCreateAlgorithmValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAlgorithm(options.Region), middleware.Before)
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
			OperationName: "CreateAlgorithm",
			Err:           err,
		}
	}
	out := result.(*CreateAlgorithmOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAlgorithmInput struct {

	// The name of the algorithm.
	//
	// This member is required.
	AlgorithmName *string

	// Specifies details about training jobs run by this algorithm, including the
	// following:
	//
	//     * The Amazon ECR path of the container and the version digest of
	// the algorithm.
	//
	//     * The hyperparameters that the algorithm supports.
	//
	//     *
	// The instance types that the algorithm supports for training.
	//
	//     * Whether the
	// algorithm supports distributed training.
	//
	//     * The metrics that the algorithm
	// emits to Amazon CloudWatch.
	//
	//     * Which metrics that the algorithm emits can be
	// used as the objective metric for hyperparameter tuning jobs.
	//
	//     * The input
	// channels that the algorithm supports for training data. For example, an
	// algorithm might support train, validation, and test channels.
	//
	// This member is required.
	TrainingSpecification *types.TrainingSpecification

	// A description of the algorithm.
	AlgorithmDescription *string

	// Whether to certify the algorithm so that it can be listed in AWS Marketplace.
	CertifyForMarketplace *bool

	// Specifies details about inference jobs that the algorithm runs, including the
	// following:
	//
	//     * The Amazon ECR paths of containers that contain the inference
	// code and model artifacts.
	//
	//     * The instance types that the algorithm supports
	// for transform jobs and real-time endpoints used for inference.
	//
	//     * The input
	// and output content formats that the algorithm supports for inference.
	InferenceSpecification *types.InferenceSpecification

	// Specifies configurations for one or more training jobs and that Amazon SageMaker
	// runs to test the algorithm's training code and, optionally, one or more batch
	// transform jobs that Amazon SageMaker runs to test the algorithm's inference
	// code.
	ValidationSpecification *types.AlgorithmValidationSpecification
}

type CreateAlgorithmOutput struct {

	// The Amazon Resource Name (ARN) of the new algorithm.
	//
	// This member is required.
	AlgorithmArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateAlgorithmMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateAlgorithm{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateAlgorithm{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateAlgorithm(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "CreateAlgorithm",
	}
}
