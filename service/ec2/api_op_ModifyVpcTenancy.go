// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Modifies the instance tenancy attribute of the specified VPC. You can change the
// instance tenancy attribute of a VPC to default only. You cannot change the
// instance tenancy attribute to dedicated. After you modify the tenancy of the
// VPC, any new instances that you launch into the VPC have a tenancy of default,
// unless you specify otherwise during launch. The tenancy of any existing
// instances in the VPC is not affected. For more information, see Dedicated
// Instances
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/dedicated-instance.html) in
// the Amazon Elastic Compute Cloud User Guide.
func (c *Client) ModifyVpcTenancy(ctx context.Context, params *ModifyVpcTenancyInput, optFns ...func(*Options)) (*ModifyVpcTenancyOutput, error) {
	if params == nil {
		params = &ModifyVpcTenancyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyVpcTenancy", params, optFns, addOperationModifyVpcTenancyMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyVpcTenancyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyVpcTenancyInput struct {

	// The instance tenancy attribute for the VPC.
	//
	// This member is required.
	InstanceTenancy types.VpcTenancy

	// The ID of the VPC.
	//
	// This member is required.
	VpcId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type ModifyVpcTenancyOutput struct {

	// Returns true if the request succeeds; otherwise, returns an error.
	ReturnValue *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationModifyVpcTenancyMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpModifyVpcTenancy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpModifyVpcTenancy{}, middleware.After)
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
	addOpModifyVpcTenancyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opModifyVpcTenancy(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opModifyVpcTenancy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ModifyVpcTenancy",
	}
}
