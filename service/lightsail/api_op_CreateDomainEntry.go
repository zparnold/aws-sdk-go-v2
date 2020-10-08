// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates one of the following entry records associated with the domain: Address
// (A), canonical name (CNAME), mail exchanger (MX), name server (NS), start of
// authority (SOA), service locator (SRV), or text (TXT). The create domain entry
// operation supports tag-based access control via resource tags applied to the
// resource identified by domain name. For more information, see the Lightsail Dev
// Guide
// (https://lightsail.aws.amazon.com/ls/docs/en/articles/amazon-lightsail-controlling-access-using-tags).
func (c *Client) CreateDomainEntry(ctx context.Context, params *CreateDomainEntryInput, optFns ...func(*Options)) (*CreateDomainEntryOutput, error) {
	if params == nil {
		params = &CreateDomainEntryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDomainEntry", params, optFns, addOperationCreateDomainEntryMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDomainEntryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDomainEntryInput struct {

	// An array of key-value pairs containing information about the domain entry
	// request.
	//
	// This member is required.
	DomainEntry *types.DomainEntry

	// The domain name (e.g., example.com) for which you want to create the domain
	// entry.
	//
	// This member is required.
	DomainName *string
}

type CreateDomainEntryOutput struct {

	// An array of objects that describe the result of the action, such as the status
	// of the request, the timestamp of the request, and the resources affected by the
	// request.
	Operation *types.Operation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateDomainEntryMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateDomainEntry{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateDomainEntry{}, middleware.After)
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
	addOpCreateDomainEntryValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDomainEntry(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateDomainEntry(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "CreateDomainEntry",
	}
}
