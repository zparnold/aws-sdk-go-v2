// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the principals associated with the specified thing. A principal can be
// X.509 certificates, IAM users, groups, and roles, Amazon Cognito identities or
// federated identities.
func (c *Client) ListThingPrincipals(ctx context.Context, params *ListThingPrincipalsInput, optFns ...func(*Options)) (*ListThingPrincipalsOutput, error) {
	if params == nil {
		params = &ListThingPrincipalsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListThingPrincipals", params, optFns, addOperationListThingPrincipalsMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*ListThingPrincipalsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the ListThingPrincipal operation.
type ListThingPrincipalsInput struct {

	// The name of the thing.
	//
	// This member is required.
	ThingName *string
}

// The output from the ListThingPrincipals operation.
type ListThingPrincipalsOutput struct {

	// The principals associated with the thing.
	Principals []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListThingPrincipalsMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListThingPrincipals{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListThingPrincipals{}, middleware.After)
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
	addOpListThingPrincipalsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListThingPrincipals(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListThingPrincipals(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "ListThingPrincipals",
	}
}
