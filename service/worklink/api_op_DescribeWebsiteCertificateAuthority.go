// Code generated by smithy-go-codegen DO NOT EDIT.

package worklink

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Provides information about the certificate authority.
func (c *Client) DescribeWebsiteCertificateAuthority(ctx context.Context, params *DescribeWebsiteCertificateAuthorityInput, optFns ...func(*Options)) (*DescribeWebsiteCertificateAuthorityOutput, error) {
	if params == nil {
		params = &DescribeWebsiteCertificateAuthorityInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeWebsiteCertificateAuthority", params, optFns, addOperationDescribeWebsiteCertificateAuthorityMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeWebsiteCertificateAuthorityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeWebsiteCertificateAuthorityInput struct {

	// The ARN of the fleet.
	//
	// This member is required.
	FleetArn *string

	// A unique identifier for the certificate authority.
	//
	// This member is required.
	WebsiteCaId *string
}

type DescribeWebsiteCertificateAuthorityOutput struct {

	// The root certificate of the certificate authority.
	Certificate *string

	// The time that the certificate authority was added.
	CreatedTime *time.Time

	// The certificate name to display.
	DisplayName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeWebsiteCertificateAuthorityMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeWebsiteCertificateAuthority{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeWebsiteCertificateAuthority{}, middleware.After)
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
	addOpDescribeWebsiteCertificateAuthorityValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeWebsiteCertificateAuthority(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeWebsiteCertificateAuthority(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "worklink",
		OperationName: "DescribeWebsiteCertificateAuthority",
	}
}
