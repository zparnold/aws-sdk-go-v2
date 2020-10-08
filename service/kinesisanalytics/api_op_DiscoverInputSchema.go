// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisanalytics

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kinesisanalytics/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This documentation is for version 1 of the Amazon Kinesis Data Analytics API,
// which only supports SQL applications. Version 2 of the API supports SQL and Java
// applications. For more information about version 2, see Amazon Kinesis Data
// Analytics API V2 Documentation (). Infers a schema by evaluating sample records
// on the specified streaming source (Amazon Kinesis stream or Amazon Kinesis
// Firehose delivery stream) or S3 object. In the response, the operation returns
// the inferred schema and also the sample records that the operation used to infer
// the schema. You can use the inferred schema when configuring a streaming source
// for your application. For conceptual information, see Configuring Application
// Input
// (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/how-it-works-input.html).
// Note that when you create an application using the Amazon Kinesis Analytics
// console, the console uses this operation to infer a schema and show it in the
// console user interface. This operation requires permissions to perform the
// kinesisanalytics:DiscoverInputSchema action.
func (c *Client) DiscoverInputSchema(ctx context.Context, params *DiscoverInputSchemaInput, optFns ...func(*Options)) (*DiscoverInputSchemaOutput, error) {
	if params == nil {
		params = &DiscoverInputSchemaInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DiscoverInputSchema", params, optFns, addOperationDiscoverInputSchemaMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DiscoverInputSchemaOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DiscoverInputSchemaInput struct {

	// The InputProcessingConfiguration
	// (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_InputProcessingConfiguration.html)
	// to use to preprocess the records before discovering the schema of the records.
	InputProcessingConfiguration *types.InputProcessingConfiguration

	// Point at which you want Amazon Kinesis Analytics to start reading records from
	// the specified streaming source discovery purposes.
	InputStartingPositionConfiguration *types.InputStartingPositionConfiguration

	// Amazon Resource Name (ARN) of the streaming source.
	ResourceARN *string

	// ARN of the IAM role that Amazon Kinesis Analytics can assume to access the
	// stream on your behalf.
	RoleARN *string

	// Specify this parameter to discover a schema from data in an Amazon S3 object.
	S3Configuration *types.S3Configuration
}

//
type DiscoverInputSchemaOutput struct {

	// Schema inferred from the streaming source. It identifies the format of the data
	// in the streaming source and how each data element maps to corresponding columns
	// in the in-application stream that you can create.
	InputSchema *types.SourceSchema

	// An array of elements, where each element corresponds to a row in a stream record
	// (a stream record can have more than one row).
	ParsedInputRecords [][]*string

	// Stream data that was modified by the processor specified in the
	// InputProcessingConfiguration parameter.
	ProcessedInputRecords []*string

	// Raw stream data that was sampled to infer the schema.
	RawInputRecords []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDiscoverInputSchemaMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDiscoverInputSchema{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDiscoverInputSchema{}, middleware.After)
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
	addOpDiscoverInputSchemaValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDiscoverInputSchema(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDiscoverInputSchema(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesisanalytics",
		OperationName: "DiscoverInputSchema",
	}
}
