// Code generated by smithy-go-codegen DO NOT EDIT.

package dynamodb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds or removes replicas in the specified global table. The global table must
// already exist to be able to use this operation. Any replica to be added must be
// empty, have the same name as the global table, have the same key schema, have
// DynamoDB Streams enabled, and have the same provisioned and maximum write
// capacity units. Although you can use UpdateGlobalTable to add replicas and
// remove replicas in a single request, for simplicity we recommend that you issue
// separate requests for adding or removing replicas. If global secondary indexes
// are specified, then the following conditions must also be met:
//
//     * The global
// secondary indexes must have the same name.
//
//     * The global secondary indexes
// must have the same hash key and sort key (if present).
//
//     * The global
// secondary indexes must have the same provisioned and maximum write capacity
// units.
func (c *Client) UpdateGlobalTable(ctx context.Context, params *UpdateGlobalTableInput, optFns ...func(*Options)) (*UpdateGlobalTableOutput, error) {
	if params == nil {
		params = &UpdateGlobalTableInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateGlobalTable", params, optFns, addOperationUpdateGlobalTableMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateGlobalTableOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateGlobalTableInput struct {

	// The global table name.
	//
	// This member is required.
	GlobalTableName *string

	// A list of Regions that should be added or removed from the global table.
	//
	// This member is required.
	ReplicaUpdates []*types.ReplicaUpdate
}

type UpdateGlobalTableOutput struct {

	// Contains the details of the global table.
	GlobalTableDescription *types.GlobalTableDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateGlobalTableMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpUpdateGlobalTable{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdateGlobalTable{}, middleware.After)
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
	addOpUpdateGlobalTableValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateGlobalTable(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	addValidateResponseChecksum(stack, options)
	addAcceptEncodingGzip(stack, options)
	return nil
}

func newServiceMetadataMiddleware_opUpdateGlobalTable(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dynamodb",
		OperationName: "UpdateGlobalTable",
	}
}
