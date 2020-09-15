// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticache

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Modifies the settings for a Global Datastore.
func (c *Client) ModifyGlobalReplicationGroup(ctx context.Context, params *ModifyGlobalReplicationGroupInput, optFns ...func(*Options)) (*ModifyGlobalReplicationGroupOutput, error) {
	stack := middleware.NewStack("ModifyGlobalReplicationGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpModifyGlobalReplicationGroupMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpModifyGlobalReplicationGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opModifyGlobalReplicationGroup(options.Region), middleware.Before)

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
			OperationName: "ModifyGlobalReplicationGroup",
			Err:           err,
		}
	}
	out := result.(*ModifyGlobalReplicationGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyGlobalReplicationGroupInput struct {
	// A valid cache node type that you want to scale this Global Datastore to.
	CacheNodeType *string
	// The upgraded version of the cache engine to be run on the clusters in the Global
	// Datastore.
	EngineVersion *string
	// The name of the Global Datastore
	GlobalReplicationGroupId *string
	// A description of the Global Datastore
	GlobalReplicationGroupDescription *string
	// This parameter causes the modifications in this request and any pending
	// modifications to be applied, asynchronously and as soon as possible.
	// Modifications to Global Replication Groups cannot be requested to be applied in
	// PreferredMaintenceWindow.
	ApplyImmediately *bool
	// Determines whether a read replica is automatically promoted to read/write
	// primary if the existing primary encounters a failure.
	AutomaticFailoverEnabled *bool
}

type ModifyGlobalReplicationGroupOutput struct {
	// Consists of a primary cluster that accepts writes and an associated secondary
	// cluster that resides in a different AWS region. The secondary cluster accepts
	// only reads. The primary cluster automatically replicates updates to the
	// secondary cluster.  <ul> <li> <p>The <b>GlobalReplicationGroupIdSuffix</b>
	// represents the name of the Global Datastore, which is what you use to associate
	// a secondary cluster.</p> </li> </ul>
	GlobalReplicationGroup *types.GlobalReplicationGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpModifyGlobalReplicationGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpModifyGlobalReplicationGroup{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyGlobalReplicationGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opModifyGlobalReplicationGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticache",
		OperationName: "ModifyGlobalReplicationGroup",
	}
}
