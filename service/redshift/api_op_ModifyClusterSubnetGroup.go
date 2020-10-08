// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Modifies a cluster subnet group to include the specified list of VPC subnets.
// The operation replaces the existing list of subnets with the new list of
// subnets.
func (c *Client) ModifyClusterSubnetGroup(ctx context.Context, params *ModifyClusterSubnetGroupInput, optFns ...func(*Options)) (*ModifyClusterSubnetGroupOutput, error) {
	if params == nil {
		params = &ModifyClusterSubnetGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyClusterSubnetGroup", params, optFns, addOperationModifyClusterSubnetGroupMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyClusterSubnetGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type ModifyClusterSubnetGroupInput struct {

	// The name of the subnet group to be modified.
	//
	// This member is required.
	ClusterSubnetGroupName *string

	// An array of VPC subnet IDs. A maximum of 20 subnets can be modified in a single
	// request.
	//
	// This member is required.
	SubnetIds []*string

	// A text description of the subnet group to be modified.
	Description *string
}

type ModifyClusterSubnetGroupOutput struct {

	// Describes a subnet group.
	ClusterSubnetGroup *types.ClusterSubnetGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationModifyClusterSubnetGroupMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpModifyClusterSubnetGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyClusterSubnetGroup{}, middleware.After)
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
	addOpModifyClusterSubnetGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opModifyClusterSubnetGroup(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opModifyClusterSubnetGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "ModifyClusterSubnetGroup",
	}
}
