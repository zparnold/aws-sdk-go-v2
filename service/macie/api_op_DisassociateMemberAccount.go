// Code generated by smithy-go-codegen DO NOT EDIT.

package macie

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes the specified member account from Amazon Macie Classic.
func (c *Client) DisassociateMemberAccount(ctx context.Context, params *DisassociateMemberAccountInput, optFns ...func(*Options)) (*DisassociateMemberAccountOutput, error) {
	if params == nil {
		params = &DisassociateMemberAccountInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisassociateMemberAccount", params, optFns, addOperationDisassociateMemberAccountMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DisassociateMemberAccountOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateMemberAccountInput struct {

	// The ID of the member account that you want to remove from Amazon Macie Classic.
	//
	// This member is required.
	MemberAccountId *string
}

type DisassociateMemberAccountOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDisassociateMemberAccountMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDisassociateMemberAccount{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDisassociateMemberAccount{}, middleware.After)
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
	addOpDisassociateMemberAccountValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateMemberAccount(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDisassociateMemberAccount(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "macie",
		OperationName: "DisassociateMemberAccount",
	}
}
