// Code generated by smithy-go-codegen DO NOT EDIT.

package workmail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes a member from the resource's set of delegates.
func (c *Client) DisassociateDelegateFromResource(ctx context.Context, params *DisassociateDelegateFromResourceInput, optFns ...func(*Options)) (*DisassociateDelegateFromResourceOutput, error) {
	if params == nil {
		params = &DisassociateDelegateFromResourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisassociateDelegateFromResource", params, optFns, addOperationDisassociateDelegateFromResourceMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DisassociateDelegateFromResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateDelegateFromResourceInput struct {

	// The identifier for the member (user, group) to be removed from the resource's
	// delegates.
	//
	// This member is required.
	EntityId *string

	// The identifier for the organization under which the resource exists.
	//
	// This member is required.
	OrganizationId *string

	// The identifier of the resource from which delegates' set members are removed.
	//
	// This member is required.
	ResourceId *string
}

type DisassociateDelegateFromResourceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDisassociateDelegateFromResourceMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDisassociateDelegateFromResource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDisassociateDelegateFromResource{}, middleware.After)
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
	addOpDisassociateDelegateFromResourceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateDelegateFromResource(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDisassociateDelegateFromResource(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workmail",
		OperationName: "DisassociateDelegateFromResource",
	}
}
