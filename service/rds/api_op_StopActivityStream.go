// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Stops a database activity stream that was started using the AWS console, the
// start-activity-stream AWS CLI command, or the StartActivityStream action. For
// more information, see Database Activity Streams
// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/DBActivityStreams.html)
// in the Amazon Aurora User Guide.
func (c *Client) StopActivityStream(ctx context.Context, params *StopActivityStreamInput, optFns ...func(*Options)) (*StopActivityStreamOutput, error) {
	if params == nil {
		params = &StopActivityStreamInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StopActivityStream", params, optFns, addOperationStopActivityStreamMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*StopActivityStreamOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StopActivityStreamInput struct {

	// The Amazon Resource Name (ARN) of the DB cluster for the database activity
	// stream. For example, arn:aws:rds:us-east-1:12345667890:cluster:das-cluster.
	//
	// This member is required.
	ResourceArn *string

	// Specifies whether or not the database activity stream is to stop as soon as
	// possible, regardless of the maintenance window for the database.
	ApplyImmediately *bool
}

type StopActivityStreamOutput struct {

	// The name of the Amazon Kinesis data stream used for the database activity
	// stream.
	KinesisStreamName *string

	// The AWS KMS key identifier used for encrypting messages in the database activity
	// stream.
	KmsKeyId *string

	// The status of the database activity stream.
	Status types.ActivityStreamStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStopActivityStreamMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpStopActivityStream{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpStopActivityStream{}, middleware.After)
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
	addOpStopActivityStreamValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStopActivityStream(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opStopActivityStream(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "StopActivityStream",
	}
}
