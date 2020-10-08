// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchevents

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This operation is used by SaaS partners to delete a partner event source. This
// operation is not used by AWS customers. When you delete an event source, the
// status of the corresponding partner event bus in the AWS customer account
// becomes DELETED.
func (c *Client) DeletePartnerEventSource(ctx context.Context, params *DeletePartnerEventSourceInput, optFns ...func(*Options)) (*DeletePartnerEventSourceOutput, error) {
	if params == nil {
		params = &DeletePartnerEventSourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeletePartnerEventSource", params, optFns, addOperationDeletePartnerEventSourceMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DeletePartnerEventSourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeletePartnerEventSourceInput struct {

	// The AWS account ID of the AWS customer that the event source was created for.
	//
	// This member is required.
	Account *string

	// The name of the event source to delete.
	//
	// This member is required.
	Name *string
}

type DeletePartnerEventSourceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeletePartnerEventSourceMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeletePartnerEventSource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeletePartnerEventSource{}, middleware.After)
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
	addOpDeletePartnerEventSourceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeletePartnerEventSource(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeletePartnerEventSource(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "events",
		OperationName: "DeletePartnerEventSource",
	}
}
