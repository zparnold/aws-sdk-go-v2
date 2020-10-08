// Code generated by smithy-go-codegen DO NOT EDIT.

package directoryservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// AWS Directory Service for Microsoft Active Directory allows you to configure and
// verify trust relationships. This action verifies a trust relationship between
// your AWS Managed Microsoft AD directory and an external domain.
func (c *Client) VerifyTrust(ctx context.Context, params *VerifyTrustInput, optFns ...func(*Options)) (*VerifyTrustOutput, error) {
	if params == nil {
		params = &VerifyTrustInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "VerifyTrust", params, optFns, addOperationVerifyTrustMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*VerifyTrustOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Initiates the verification of an existing trust relationship between an AWS
// Managed Microsoft AD directory and an external domain.
type VerifyTrustInput struct {

	// The unique Trust ID of the trust relationship to verify.
	//
	// This member is required.
	TrustId *string
}

// Result of a VerifyTrust request.
type VerifyTrustOutput struct {

	// The unique Trust ID of the trust relationship that was verified.
	TrustId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationVerifyTrustMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpVerifyTrust{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpVerifyTrust{}, middleware.After)
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
	addOpVerifyTrustValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opVerifyTrust(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opVerifyTrust(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ds",
		OperationName: "VerifyTrust",
	}
}
