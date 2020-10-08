// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the resource record sets in a specified hosted zone that were created
// based on the settings in a specified traffic policy version. When you update a
// traffic policy instance, Amazon Route 53 continues to respond to DNS queries for
// the root resource record set name (such as example.com) while it replaces one
// group of resource record sets with another. Route 53 performs the following
// operations:
//
//     * Route 53 creates a new group of resource record sets based on
// the specified traffic policy. This is true regardless of how significant the
// differences are between the existing resource record sets and the new resource
// record sets.
//
//     * When all of the new resource record sets have been created,
// Route 53 starts to respond to DNS queries for the root resource record set name
// (such as example.com) by using the new resource record sets.
//
//     * Route 53
// deletes the old group of resource record sets that are associated with the root
// resource record set name.
func (c *Client) UpdateTrafficPolicyInstance(ctx context.Context, params *UpdateTrafficPolicyInstanceInput, optFns ...func(*Options)) (*UpdateTrafficPolicyInstanceOutput, error) {
	if params == nil {
		params = &UpdateTrafficPolicyInstanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateTrafficPolicyInstance", params, optFns, addOperationUpdateTrafficPolicyInstanceMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateTrafficPolicyInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A complex type that contains information about the resource record sets that you
// want to update based on a specified traffic policy instance.
type UpdateTrafficPolicyInstanceInput struct {

	// The ID of the traffic policy instance that you want to update.
	//
	// This member is required.
	Id *string

	// The TTL that you want Amazon Route 53 to assign to all of the updated resource
	// record sets.
	//
	// This member is required.
	TTL *int64

	// The ID of the traffic policy that you want Amazon Route 53 to use to update
	// resource record sets for the specified traffic policy instance.
	//
	// This member is required.
	TrafficPolicyId *string

	// The version of the traffic policy that you want Amazon Route 53 to use to update
	// resource record sets for the specified traffic policy instance.
	//
	// This member is required.
	TrafficPolicyVersion *int32
}

// A complex type that contains information about the resource record sets that
// Amazon Route 53 created based on a specified traffic policy.
type UpdateTrafficPolicyInstanceOutput struct {

	// A complex type that contains settings for the updated traffic policy instance.
	//
	// This member is required.
	TrafficPolicyInstance *types.TrafficPolicyInstance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateTrafficPolicyInstanceMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpUpdateTrafficPolicyInstance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpUpdateTrafficPolicyInstance{}, middleware.After)
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
	addOpUpdateTrafficPolicyInstanceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateTrafficPolicyInstance(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateTrafficPolicyInstance(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "UpdateTrafficPolicyInstance",
	}
}
