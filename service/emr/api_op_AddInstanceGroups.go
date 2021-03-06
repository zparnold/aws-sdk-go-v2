// Code generated by smithy-go-codegen DO NOT EDIT.

package emr

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds one or more instance groups to a running cluster.
func (c *Client) AddInstanceGroups(ctx context.Context, params *AddInstanceGroupsInput, optFns ...func(*Options)) (*AddInstanceGroupsOutput, error) {
	stack := middleware.NewStack("AddInstanceGroups", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpAddInstanceGroupsMiddlewares(stack)
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
	addOpAddInstanceGroupsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAddInstanceGroups(options.Region), middleware.Before)
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
			OperationName: "AddInstanceGroups",
			Err:           err,
		}
	}
	out := result.(*AddInstanceGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input to an AddInstanceGroups call.
type AddInstanceGroupsInput struct {

	// Instance groups to add.
	//
	// This member is required.
	InstanceGroups []*types.InstanceGroupConfig

	// Job flow in which to add the instance groups.
	//
	// This member is required.
	JobFlowId *string
}

// Output from an AddInstanceGroups call.
type AddInstanceGroupsOutput struct {

	// The Amazon Resource Name of the cluster.
	ClusterArn *string

	// Instance group IDs of the newly created instance groups.
	InstanceGroupIds []*string

	// The job flow ID in which the instance groups are added.
	JobFlowId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpAddInstanceGroupsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpAddInstanceGroups{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpAddInstanceGroups{}, middleware.After)
}

func newServiceMetadataMiddleware_opAddInstanceGroups(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticmapreduce",
		OperationName: "AddInstanceGroups",
	}
}
