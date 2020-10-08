// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Detaches the specified principal from the specified thing. A principal can be
// X.509 certificates, IAM users, groups, and roles, Amazon Cognito identities or
// federated identities. This call is asynchronous. It might take several seconds
// for the detachment to propagate.
func (c *Client) DetachThingPrincipal(ctx context.Context, params *DetachThingPrincipalInput, optFns ...func(*Options)) (*DetachThingPrincipalOutput, error) {
	if params == nil {
		params = &DetachThingPrincipalInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DetachThingPrincipal", params, optFns, addOperationDetachThingPrincipalMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DetachThingPrincipalOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the DetachThingPrincipal operation.
type DetachThingPrincipalInput struct {

	// If the principal is a certificate, this value must be ARN of the certificate. If
	// the principal is an Amazon Cognito identity, this value must be the ID of the
	// Amazon Cognito identity.
	//
	// This member is required.
	Principal *string

	// The name of the thing.
	//
	// This member is required.
	ThingName *string
}

// The output from the DetachThingPrincipal operation.
type DetachThingPrincipalOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDetachThingPrincipalMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDetachThingPrincipal{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDetachThingPrincipal{}, middleware.After)
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
	addOpDetachThingPrincipalValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDetachThingPrincipal(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDetachThingPrincipal(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "DetachThingPrincipal",
	}
}
