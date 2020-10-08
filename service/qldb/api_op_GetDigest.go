// Code generated by smithy-go-codegen DO NOT EDIT.

package qldb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/qldb/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the digest of a ledger at the latest committed block in the journal. The
// response includes a 256-bit hash value and a block address.
func (c *Client) GetDigest(ctx context.Context, params *GetDigestInput, optFns ...func(*Options)) (*GetDigestOutput, error) {
	if params == nil {
		params = &GetDigestInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDigest", params, optFns, addOperationGetDigestMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDigestOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDigestInput struct {

	// The name of the ledger.
	//
	// This member is required.
	Name *string
}

type GetDigestOutput struct {

	// The 256-bit hash value representing the digest returned by a GetDigest request.
	//
	// This member is required.
	Digest []byte

	// The latest block location covered by the digest that you requested. An address
	// is an Amazon Ion structure that has two fields: strandId and sequenceNo.
	//
	// This member is required.
	DigestTipAddress *types.ValueHolder

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetDigestMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetDigest{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetDigest{}, middleware.After)
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
	addOpGetDigestValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetDigest(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetDigest(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "qldb",
		OperationName: "GetDigest",
	}
}
