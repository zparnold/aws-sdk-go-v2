// Code generated by smithy-go-codegen DO NOT EDIT.

package imagebuilder

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a distribution configuration.
func (c *Client) DeleteDistributionConfiguration(ctx context.Context, params *DeleteDistributionConfigurationInput, optFns ...func(*Options)) (*DeleteDistributionConfigurationOutput, error) {
	if params == nil {
		params = &DeleteDistributionConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteDistributionConfiguration", params, optFns, addOperationDeleteDistributionConfigurationMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteDistributionConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteDistributionConfigurationInput struct {

	// The Amazon Resource Name (ARN) of the distribution configuration to delete.
	//
	// This member is required.
	DistributionConfigurationArn *string
}

type DeleteDistributionConfigurationOutput struct {

	// The Amazon Resource Name (ARN) of the distribution configuration that was
	// deleted.
	DistributionConfigurationArn *string

	// The request ID that uniquely identifies this request.
	RequestId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteDistributionConfigurationMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteDistributionConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteDistributionConfiguration{}, middleware.After)
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
	addOpDeleteDistributionConfigurationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteDistributionConfiguration(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteDistributionConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "imagebuilder",
		OperationName: "DeleteDistributionConfiguration",
	}
}
