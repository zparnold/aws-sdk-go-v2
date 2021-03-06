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

// Deletes the specified version from the specified managed policy. You cannot
// delete the default version from a policy using this API. To delete the default
// version from a policy, use DeletePolicy (). To find out which version of a
// policy is marked as the default version, use ListPolicyVersions (). For
// information about versions for managed policies, see Versioning for Managed
// Policies
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-versions.html)
// in the IAM User Guide.
func (c *Client) DeletePolicyVersion(ctx context.Context, params *DeletePolicyVersionInput, optFns ...func(*Options)) (*DeletePolicyVersionOutput, error) {
	stack := middleware.NewStack("DeletePolicyVersion", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDeletePolicyVersionMiddlewares(stack)
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
	addOpDeletePolicyVersionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeletePolicyVersion(options.Region), middleware.Before)
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
			OperationName: "DeletePolicyVersion",
			Err:           err,
		}
	}
	out := result.(*DeletePolicyVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeletePolicyVersionInput struct {

	// The Amazon Resource Name (ARN) of the IAM policy from which you want to delete a
	// version. For more information about ARNs, see Amazon Resource Names (ARNs) and
	// AWS Service Namespaces
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in
	// the AWS General Reference.
	//
	// This member is required.
	PolicyArn *string

	// The policy version to delete. This parameter allows (through its regex pattern
	// (http://wikipedia.org/wiki/regex)) a string of characters that consists of the
	// lowercase letter 'v' followed by one or two digits, and optionally followed by a
	// period '.' and a string of letters and digits. For more information about
	// managed policy versions, see Versioning for Managed Policies
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-versions.html)
	// in the IAM User Guide.
	//
	// This member is required.
	VersionId *string
}

type DeletePolicyVersionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDeletePolicyVersionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDeletePolicyVersion{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDeletePolicyVersion{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeletePolicyVersion(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "DeletePolicyVersion",
	}
}
