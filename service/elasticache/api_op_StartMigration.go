// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticache

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Start the migration of data.
func (c *Client) StartMigration(ctx context.Context, params *StartMigrationInput, optFns ...func(*Options)) (*StartMigrationOutput, error) {
	stack := middleware.NewStack("StartMigration", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpStartMigrationMiddlewares(stack)
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
	addOpStartMigrationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartMigration(options.Region), middleware.Before)
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
			OperationName: "StartMigration",
			Err:           err,
		}
	}
	out := result.(*StartMigrationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartMigrationInput struct {

	// List of endpoints from which data should be migrated. For Redis (cluster mode
	// disabled), list should have only one element.
	//
	// This member is required.
	CustomerNodeEndpointList []*types.CustomerNodeEndpoint

	// The ID of the replication group to which data should be migrated.
	//
	// This member is required.
	ReplicationGroupId *string
}

type StartMigrationOutput struct {

	// Contains all of the attributes of a specific Redis replication group.
	ReplicationGroup *types.ReplicationGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpStartMigrationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpStartMigration{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpStartMigration{}, middleware.After)
}

func newServiceMetadataMiddleware_opStartMigration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticache",
		OperationName: "StartMigration",
	}
}
