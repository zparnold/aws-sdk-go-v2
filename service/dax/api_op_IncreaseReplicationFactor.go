// Code generated by smithy-go-codegen DO NOT EDIT.

package dax

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/dax/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds one or more nodes to a DAX cluster.
func (c *Client) IncreaseReplicationFactor(ctx context.Context, params *IncreaseReplicationFactorInput, optFns ...func(*Options)) (*IncreaseReplicationFactorOutput, error) {
	stack := middleware.NewStack("IncreaseReplicationFactor", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpIncreaseReplicationFactorMiddlewares(stack)
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
	addOpIncreaseReplicationFactorValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opIncreaseReplicationFactor(options.Region), middleware.Before)
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
			OperationName: "IncreaseReplicationFactor",
			Err:           err,
		}
	}
	out := result.(*IncreaseReplicationFactorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type IncreaseReplicationFactorInput struct {

	// The name of the DAX cluster that will receive additional nodes.
	//
	// This member is required.
	ClusterName *string

	// The new number of nodes for the DAX cluster.
	//
	// This member is required.
	NewReplicationFactor *int32

	// The Availability Zones (AZs) in which the cluster nodes will be created. All
	// nodes belonging to the cluster are placed in these Availability Zones. Use this
	// parameter if you want to distribute the nodes across multiple AZs.
	AvailabilityZones []*string
}

type IncreaseReplicationFactorOutput struct {

	// A description of the DAX cluster. with its new replication factor.
	Cluster *types.Cluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpIncreaseReplicationFactorMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpIncreaseReplicationFactor{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpIncreaseReplicationFactor{}, middleware.After)
}

func newServiceMetadataMiddleware_opIncreaseReplicationFactor(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dax",
		OperationName: "IncreaseReplicationFactor",
	}
}
