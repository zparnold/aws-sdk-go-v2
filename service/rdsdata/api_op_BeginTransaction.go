// Code generated by smithy-go-codegen DO NOT EDIT.

package rdsdata

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Starts a SQL transaction.  <important> <p>A transaction can run for a maximum of
// 24 hours. A transaction is terminated and rolled back automatically after 24
// hours.</p> <p>A transaction times out if no calls use its transaction ID in
// three minutes. If a transaction times out before it's committed, it's rolled
// back automatically.</p> <p>DDL statements inside a transaction cause an implicit
// commit. We recommend that you run each DDL statement in a separate
// <code>ExecuteStatement</code> call with <code>continueAfterTimeout</code>
// enabled.</p> </important>
func (c *Client) BeginTransaction(ctx context.Context, params *BeginTransactionInput, optFns ...func(*Options)) (*BeginTransactionOutput, error) {
	if params == nil {
		params = &BeginTransactionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BeginTransaction", params, optFns, addOperationBeginTransactionMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*BeginTransactionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request parameters represent the input of a request to start a SQL
// transaction.
type BeginTransactionInput struct {

	// The Amazon Resource Name (ARN) of the Aurora Serverless DB cluster.
	//
	// This member is required.
	ResourceArn *string

	// The name or ARN of the secret that enables access to the DB cluster.
	//
	// This member is required.
	SecretArn *string

	// The name of the database.
	Database *string

	// The name of the database schema.
	Schema *string
}

// The response elements represent the output of a request to start a SQL
// transaction.
type BeginTransactionOutput struct {

	// The transaction ID of the transaction started by the call.
	TransactionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationBeginTransactionMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBeginTransaction{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBeginTransaction{}, middleware.After)
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
	addOpBeginTransactionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opBeginTransaction(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opBeginTransaction(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "",
		OperationName: "BeginTransaction",
	}
}
