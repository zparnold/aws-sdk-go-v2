// Code generated by smithy-go-codegen DO NOT EDIT.

package emr

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// ModifyInstanceGroups modifies the number of nodes and configuration settings of
// an instance group. The input parameters include the new target instance count
// for the group and the instance group ID. The call will either succeed or fail
// atomically.
func (c *Client) ModifyInstanceGroups(ctx context.Context, params *ModifyInstanceGroupsInput, optFns ...func(*Options)) (*ModifyInstanceGroupsOutput, error) {
	if params == nil {
		params = &ModifyInstanceGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyInstanceGroups", params, optFns, addOperationModifyInstanceGroupsMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyInstanceGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Change the size of some instance groups.
type ModifyInstanceGroupsInput struct {

	// The ID of the cluster to which the instance group belongs.
	ClusterId *string

	// Instance groups to change.
	InstanceGroups []*types.InstanceGroupModifyConfig
}

type ModifyInstanceGroupsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationModifyInstanceGroupsMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpModifyInstanceGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpModifyInstanceGroups{}, middleware.After)
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
	addOpModifyInstanceGroupsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opModifyInstanceGroups(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opModifyInstanceGroups(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticmapreduce",
		OperationName: "ModifyInstanceGroups",
	}
}
