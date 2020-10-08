// Code generated by smithy-go-codegen DO NOT EDIT.

package directoryservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/directoryservice/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// For the specified directory, lists all the certificates registered for a secured
// LDAP connection.
func (c *Client) ListCertificates(ctx context.Context, params *ListCertificatesInput, optFns ...func(*Options)) (*ListCertificatesOutput, error) {
	if params == nil {
		params = &ListCertificatesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCertificates", params, optFns, addOperationListCertificatesMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCertificatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCertificatesInput struct {

	// The identifier of the directory.
	//
	// This member is required.
	DirectoryId *string

	// The number of items that should show up on one page
	Limit *int32

	// A token for requesting another page of certificates if the NextToken response
	// element indicates that more certificates are available. Use the value of the
	// returned NextToken element in your request until the token comes back as null.
	// Pass null if this is the first call.
	NextToken *string
}

type ListCertificatesOutput struct {

	// A list of certificates with basic details including certificate ID, certificate
	// common name, certificate state.
	CertificatesInfo []*types.CertificateInfo

	// Indicates whether another page of certificates is available when the number of
	// available certificates exceeds the page limit.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListCertificatesMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListCertificates{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListCertificates{}, middleware.After)
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
	addOpListCertificatesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListCertificates(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListCertificates(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ds",
		OperationName: "ListCertificates",
	}
}
