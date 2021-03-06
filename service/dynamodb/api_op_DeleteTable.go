// Code generated by smithy-go-codegen DO NOT EDIT.

package dynamodb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// The DeleteTable operation deletes a table and all of its items. After a
// DeleteTable request, the specified table is in the DELETING state until DynamoDB
// completes the deletion. If the table is in the ACTIVE state, you can delete it.
// If a table is in CREATING or UPDATING states, then DynamoDB returns a
// ResourceInUseException. If the specified table does not exist, DynamoDB returns
// a ResourceNotFoundException. If table is already in the DELETING state, no error
// is returned. DynamoDB might continue to accept data read and write operations,
// such as GetItem and PutItem, on a table in the DELETING state until the table
// deletion is complete. When you delete a table, any indexes on that table are
// also deleted. If you have DynamoDB Streams enabled on the table, then the
// corresponding stream on that table goes into the DISABLED state, and the stream
// is automatically deleted after 24 hours.  <p>Use the <code>DescribeTable</code>
// action to check the status of the table. </p>
func (c *Client) DeleteTable(ctx context.Context, params *DeleteTableInput, optFns ...func(*Options)) (*DeleteTableOutput, error) {
	stack := middleware.NewStack("DeleteTable", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson10_serdeOpDeleteTableMiddlewares(stack)
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
	addOpDeleteTableValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteTable(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	addValidateResponseChecksum(stack, options)
	addAcceptEncodingGzip(stack, options)

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
			OperationName: "DeleteTable",
			Err:           err,
		}
	}
	out := result.(*DeleteTableOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a DeleteTable operation.
type DeleteTableInput struct {

	// The name of the table to delete.
	//
	// This member is required.
	TableName *string
}

// Represents the output of a DeleteTable operation.
type DeleteTableOutput struct {

	// Represents the properties of a table.
	TableDescription *types.TableDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson10_serdeOpDeleteTableMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson10_serializeOpDeleteTable{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson10_deserializeOpDeleteTable{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteTable(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dynamodb",
		OperationName: "DeleteTable",
	}
}
