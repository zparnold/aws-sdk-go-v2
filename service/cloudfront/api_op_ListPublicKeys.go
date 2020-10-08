// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// List all public keys that have been added to CloudFront for this account.
func (c *Client) ListPublicKeys(ctx context.Context, params *ListPublicKeysInput, optFns ...func(*Options)) (*ListPublicKeysOutput, error) {
	if params == nil {
		params = &ListPublicKeysInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPublicKeys", params, optFns, addOperationListPublicKeysMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPublicKeysOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPublicKeysInput struct {

	// Use this when paginating results to indicate where to begin in your list of
	// public keys. The results include public keys in the list that occur after the
	// marker. To get the next page of results, set the Marker to the value of the
	// NextMarker from the current page's response (which is also the ID of the last
	// public key on that page).
	Marker *string

	// The maximum number of public keys you want in the response body.
	MaxItems *string
}

type ListPublicKeysOutput struct {

	// Returns a list of all public keys that have been added to CloudFront for this
	// account.
	PublicKeyList *types.PublicKeyList

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListPublicKeysMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpListPublicKeys{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpListPublicKeys{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListPublicKeys(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListPublicKeys(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "ListPublicKeys",
	}
}
