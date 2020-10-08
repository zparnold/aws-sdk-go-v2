// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a model package that you can use to create Amazon SageMaker models or
// list on AWS Marketplace. Buyers can subscribe to model packages listed on AWS
// Marketplace to create models in Amazon SageMaker. To create a model package by
// specifying a Docker container that contains your inference code and the Amazon
// S3 location of your model artifacts, provide values for InferenceSpecification.
// To create a model from an algorithm resource that you created or subscribed to
// in AWS Marketplace, provide a value for SourceAlgorithmSpecification.
func (c *Client) CreateModelPackage(ctx context.Context, params *CreateModelPackageInput, optFns ...func(*Options)) (*CreateModelPackageOutput, error) {
	if params == nil {
		params = &CreateModelPackageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateModelPackage", params, optFns, addOperationCreateModelPackageMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateModelPackageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateModelPackageInput struct {

	// The name of the model package. The name must have 1 to 63 characters. Valid
	// characters are a-z, A-Z, 0-9, and - (hyphen).
	//
	// This member is required.
	ModelPackageName *string

	// Whether to certify the model package for listing on AWS Marketplace.
	CertifyForMarketplace *bool

	// Specifies details about inference jobs that can be run with models based on this
	// model package, including the following:
	//
	//     * The Amazon ECR paths of
	// containers that contain the inference code and model artifacts.
	//
	//     * The
	// instance types that the model package supports for transform jobs and real-time
	// endpoints used for inference.
	//
	//     * The input and output content formats that
	// the model package supports for inference.
	InferenceSpecification *types.InferenceSpecification

	// A description of the model package.
	ModelPackageDescription *string

	// Details about the algorithm that was used to create the model package.
	SourceAlgorithmSpecification *types.SourceAlgorithmSpecification

	// Specifies configurations for one or more transform jobs that Amazon SageMaker
	// runs to test the model package.
	ValidationSpecification *types.ModelPackageValidationSpecification
}

type CreateModelPackageOutput struct {

	// The Amazon Resource Name (ARN) of the new model package.
	//
	// This member is required.
	ModelPackageArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateModelPackageMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateModelPackage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateModelPackage{}, middleware.After)
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
	addOpCreateModelPackageValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateModelPackage(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateModelPackage(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "CreateModelPackage",
	}
}
