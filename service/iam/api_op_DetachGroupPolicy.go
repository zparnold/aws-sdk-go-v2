// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes the specified managed policy from the specified IAM group. A group can
// also have inline policies embedded with it. To delete an inline policy, use the
// DeleteGroupPolicy () API. For information about policies, see Managed Policies
// and Inline Policies
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html)
// in the IAM User Guide.
func (c *Client) DetachGroupPolicy(ctx context.Context, params *DetachGroupPolicyInput, optFns ...func(*Options)) (*DetachGroupPolicyOutput, error) {
	stack := middleware.NewStack("DetachGroupPolicy", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDetachGroupPolicyMiddlewares(stack)
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
	addOpDetachGroupPolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDetachGroupPolicy(options.Region), middleware.Before)
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
			OperationName: "DetachGroupPolicy",
			Err:           err,
		}
	}
	out := result.(*DetachGroupPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DetachGroupPolicyInput struct {

	// The name (friendly name, not ARN) of the IAM group to detach the policy from.
	// This parameter allows (through its regex pattern
	// (http://wikipedia.org/wiki/regex)) a string of characters consisting of upper
	// and lowercase alphanumeric characters with no spaces. You can also include any
	// of the following characters: _+=,.@-
	//
	// This member is required.
	GroupName *string

	// The Amazon Resource Name (ARN) of the IAM policy you want to detach. For more
	// information about ARNs, see Amazon Resource Names (ARNs) and AWS Service
	// Namespaces
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in
	// the AWS General Reference.
	//
	// This member is required.
	PolicyArn *string
}

type DetachGroupPolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDetachGroupPolicyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDetachGroupPolicy{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDetachGroupPolicy{}, middleware.After)
}

func newServiceMetadataMiddleware_opDetachGroupPolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "DetachGroupPolicy",
	}
}
