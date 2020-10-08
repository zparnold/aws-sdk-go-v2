// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Associates the specified sign-in delegate groups with the specified Amazon Chime
// account.
func (c *Client) AssociateSigninDelegateGroupsWithAccount(ctx context.Context, params *AssociateSigninDelegateGroupsWithAccountInput, optFns ...func(*Options)) (*AssociateSigninDelegateGroupsWithAccountOutput, error) {
	if params == nil {
		params = &AssociateSigninDelegateGroupsWithAccountInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateSigninDelegateGroupsWithAccount", params, optFns, addOperationAssociateSigninDelegateGroupsWithAccountMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateSigninDelegateGroupsWithAccountOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateSigninDelegateGroupsWithAccountInput struct {

	// The Amazon Chime account ID.
	//
	// This member is required.
	AccountId *string

	// The sign-in delegate groups.
	//
	// This member is required.
	SigninDelegateGroups []*types.SigninDelegateGroup
}

type AssociateSigninDelegateGroupsWithAccountOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationAssociateSigninDelegateGroupsWithAccountMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpAssociateSigninDelegateGroupsWithAccount{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpAssociateSigninDelegateGroupsWithAccount{}, middleware.After)
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
	addOpAssociateSigninDelegateGroupsWithAccountValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateSigninDelegateGroupsWithAccount(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opAssociateSigninDelegateGroupsWithAccount(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "AssociateSigninDelegateGroupsWithAccount",
	}
}
