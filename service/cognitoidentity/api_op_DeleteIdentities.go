// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentity

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes identities from an identity pool. You can specify a list of 1-60
// identities that you want to delete. You must use AWS Developer credentials to
// call this API.
func (c *Client) DeleteIdentities(ctx context.Context, params *DeleteIdentitiesInput, optFns ...func(*Options)) (*DeleteIdentitiesOutput, error) {
	if params == nil {
		params = &DeleteIdentitiesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteIdentities", params, optFns, addOperationDeleteIdentitiesMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteIdentitiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input to the DeleteIdentities action.
type DeleteIdentitiesInput struct {

	// A list of 1-60 identities that you want to delete.
	//
	// This member is required.
	IdentityIdsToDelete []*string
}

// Returned in response to a successful DeleteIdentities operation.
type DeleteIdentitiesOutput struct {

	// An array of UnprocessedIdentityId objects, each of which contains an ErrorCode
	// and IdentityId.
	UnprocessedIdentityIds []*types.UnprocessedIdentityId

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteIdentitiesMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteIdentities{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteIdentities{}, middleware.After)
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
	addOpDeleteIdentitiesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteIdentities(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteIdentities(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-identity",
		OperationName: "DeleteIdentities",
	}
}
