// Code generated by smithy-go-codegen DO NOT EDIT.

package fms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/fms/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves violations for a resource based on the specified AWS Firewall Manager
// policy and AWS account.
func (c *Client) GetViolationDetails(ctx context.Context, params *GetViolationDetailsInput, optFns ...func(*Options)) (*GetViolationDetailsOutput, error) {
	if params == nil {
		params = &GetViolationDetailsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetViolationDetails", params, optFns, addOperationGetViolationDetailsMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*GetViolationDetailsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetViolationDetailsInput struct {

	// The AWS account ID that you want the details for.
	//
	// This member is required.
	MemberAccount *string

	// The ID of the AWS Firewall Manager policy that you want the details for. This
	// currently only supports security group content audit policies.
	//
	// This member is required.
	PolicyId *string

	// The ID of the resource that has violations.
	//
	// This member is required.
	ResourceId *string

	// The resource type. This is in the format shown in the AWS Resource Types
	// Reference
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-template-resource-type-ref.html).
	// Supported resource types are: AWS::EC2::Instance, AWS::EC2::NetworkInterface, or
	// AWS::EC2::SecurityGroup.
	//
	// This member is required.
	ResourceType *string
}

type GetViolationDetailsOutput struct {

	// Violation detail for a resource.
	ViolationDetail *types.ViolationDetail

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetViolationDetailsMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetViolationDetails{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetViolationDetails{}, middleware.After)
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
	addOpGetViolationDetailsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetViolationDetails(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetViolationDetails(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "fms",
		OperationName: "GetViolationDetails",
	}
}
