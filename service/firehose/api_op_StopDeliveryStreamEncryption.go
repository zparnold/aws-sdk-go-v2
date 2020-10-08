// Code generated by smithy-go-codegen DO NOT EDIT.

package firehose

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Disables server-side encryption (SSE) for the delivery stream. This operation is
// asynchronous. It returns immediately. When you invoke it, Kinesis Data Firehose
// first sets the encryption status of the stream to DISABLING, and then to
// DISABLED. You can continue to read and write data to your stream while its
// status is DISABLING. It can take up to 5 seconds after the encryption status
// changes to DISABLED before all records written to the delivery stream are no
// longer subject to encryption. To find out whether a record or a batch of records
// was encrypted, check the response elements PutRecordOutput$Encrypted () and
// PutRecordBatchOutput$Encrypted (), respectively. To check the encryption state
// of a delivery stream, use DescribeDeliveryStream (). If SSE is enabled using a
// customer managed CMK and then you invoke StopDeliveryStreamEncryption, Kinesis
// Data Firehose schedules the related KMS grant for retirement and then retires it
// after it ensures that it is finished delivering records to the destination. The
// StartDeliveryStreamEncryption and StopDeliveryStreamEncryption operations have a
// combined limit of 25 calls per delivery stream per 24 hours. For example, you
// reach the limit if you call StartDeliveryStreamEncryption 13 times and
// StopDeliveryStreamEncryption 12 times for the same delivery stream in a 24-hour
// period.
func (c *Client) StopDeliveryStreamEncryption(ctx context.Context, params *StopDeliveryStreamEncryptionInput, optFns ...func(*Options)) (*StopDeliveryStreamEncryptionOutput, error) {
	if params == nil {
		params = &StopDeliveryStreamEncryptionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StopDeliveryStreamEncryption", params, optFns, addOperationStopDeliveryStreamEncryptionMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*StopDeliveryStreamEncryptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StopDeliveryStreamEncryptionInput struct {

	// The name of the delivery stream for which you want to disable server-side
	// encryption (SSE).
	//
	// This member is required.
	DeliveryStreamName *string
}

type StopDeliveryStreamEncryptionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStopDeliveryStreamEncryptionMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStopDeliveryStreamEncryption{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStopDeliveryStreamEncryption{}, middleware.After)
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
	addOpStopDeliveryStreamEncryptionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStopDeliveryStreamEncryption(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opStopDeliveryStreamEncryption(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "firehose",
		OperationName: "StopDeliveryStreamEncryption",
	}
}
