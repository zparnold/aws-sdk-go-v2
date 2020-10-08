// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Enables EBS encryption by default for your account in the current Region. After
// you enable encryption by default, the EBS volumes that you create are are always
// encrypted, either using the default CMK or the CMK that you specified when you
// created each volume. For more information, see Amazon EBS Encryption
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html) in the
// Amazon Elastic Compute Cloud User Guide. You can specify the default CMK for
// encryption by default using ModifyEbsDefaultKmsKeyId () or
// ResetEbsDefaultKmsKeyId (). Enabling encryption by default has no effect on the
// encryption status of your existing volumes. After you enable encryption by
// default, you can no longer launch instances using instance types that do not
// support encryption. For more information, see Supported Instance Types
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html#EBSEncryption_supported_instances).
func (c *Client) EnableEbsEncryptionByDefault(ctx context.Context, params *EnableEbsEncryptionByDefaultInput, optFns ...func(*Options)) (*EnableEbsEncryptionByDefaultOutput, error) {
	if params == nil {
		params = &EnableEbsEncryptionByDefaultInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "EnableEbsEncryptionByDefault", params, optFns, addOperationEnableEbsEncryptionByDefaultMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*EnableEbsEncryptionByDefaultOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type EnableEbsEncryptionByDefaultInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type EnableEbsEncryptionByDefaultOutput struct {

	// The updated status of encryption by default.
	EbsEncryptionByDefault *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationEnableEbsEncryptionByDefaultMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpEnableEbsEncryptionByDefault{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpEnableEbsEncryptionByDefault{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opEnableEbsEncryptionByDefault(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opEnableEbsEncryptionByDefault(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "EnableEbsEncryptionByDefault",
	}
}
