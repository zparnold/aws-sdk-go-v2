// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehendmedical

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/comprehendmedical/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets the properties associated with a medical entities detection job. Use this
// operation to get the status of a detection job.
func (c *Client) DescribeEntitiesDetectionV2Job(ctx context.Context, params *DescribeEntitiesDetectionV2JobInput, optFns ...func(*Options)) (*DescribeEntitiesDetectionV2JobOutput, error) {
	stack := middleware.NewStack("DescribeEntitiesDetectionV2Job", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeEntitiesDetectionV2JobMiddlewares(stack)
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
	addOpDescribeEntitiesDetectionV2JobValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEntitiesDetectionV2Job(options.Region), middleware.Before)
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
			OperationName: "DescribeEntitiesDetectionV2Job",
			Err:           err,
		}
	}
	out := result.(*DescribeEntitiesDetectionV2JobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeEntitiesDetectionV2JobInput struct {

	// The identifier that Amazon Comprehend Medical generated for the job. The
	// StartEntitiesDetectionV2Job operation returns this identifier in its response.
	//
	// This member is required.
	JobId *string
}

type DescribeEntitiesDetectionV2JobOutput struct {

	// An object that contains the properties associated with a detection job.
	ComprehendMedicalAsyncJobProperties *types.ComprehendMedicalAsyncJobProperties

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeEntitiesDetectionV2JobMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeEntitiesDetectionV2Job{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeEntitiesDetectionV2Job{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeEntitiesDetectionV2Job(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "comprehendmedical",
		OperationName: "DescribeEntitiesDetectionV2Job",
	}
}
