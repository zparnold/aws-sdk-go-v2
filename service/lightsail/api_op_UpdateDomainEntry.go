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

// Updates a domain recordset after it is created. The update domain entry
// operation supports tag-based access control via resource tags applied to the
// resource identified by domain name. For more information, see the Lightsail Dev
// Guide
// (https://lightsail.aws.amazon.com/ls/docs/en/articles/amazon-lightsail-controlling-access-using-tags).
func (c *Client) UpdateDomainEntry(ctx context.Context, params *UpdateDomainEntryInput, optFns ...func(*Options)) (*UpdateDomainEntryOutput, error) {
	if params == nil {
		params = &UpdateDomainEntryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateDomainEntry", params, optFns, addOperationUpdateDomainEntryMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateDomainEntryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateDomainEntryInput struct {

	// An array of key-value pairs containing information about the domain entry.
	//
	// This member is required.
	DomainEntry *types.DomainEntry

	// The name of the domain recordset to update.
	//
	// This member is required.
	DomainName *string
}

type UpdateDomainEntryOutput struct {

	// An array of objects that describe the result of the action, such as the status
	// of the request, the timestamp of the request, and the resources affected by the
	// request.
	Operations []*types.Operation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateDomainEntryMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateDomainEntry{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateDomainEntry{}, middleware.After)
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
	addOpUpdateDomainEntryValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDomainEntry(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateDomainEntry(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "UpdateDomainEntry",
	}
}
