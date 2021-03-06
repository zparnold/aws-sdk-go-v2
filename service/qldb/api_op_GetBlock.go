// Code generated by smithy-go-codegen DO NOT EDIT.

package qldb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/qldb/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a block object at a specified address in a journal. Also returns a proof
// of the specified block for verification if DigestTipAddress is provided. For
// information about the data contents in a block, see Journal contents
// (https://docs.aws.amazon.com/qldb/latest/developerguide/journal-contents.html)
// in the Amazon QLDB Developer Guide. If the specified ledger doesn't exist or is
// in DELETING status, then throws ResourceNotFoundException. If the specified
// ledger is in CREATING status, then throws ResourcePreconditionNotMetException.
// If no block exists with the specified address, then throws
// InvalidParameterException.
func (c *Client) GetBlock(ctx context.Context, params *GetBlockInput, optFns ...func(*Options)) (*GetBlockOutput, error) {
	stack := middleware.NewStack("GetBlock", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetBlockMiddlewares(stack)
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
	addOpGetBlockValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetBlock(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "GetBlock",
			Err:           err,
		}
	}
	out := result.(*GetBlockOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBlockInput struct {

	// The location of the block that you want to request. An address is an Amazon Ion
	// structure that has two fields: strandId and sequenceNo. For example:
	// {strandId:"BlFTjlSXze9BIh1KOszcE3",sequenceNo:14}
	//
	// This member is required.
	BlockAddress *types.ValueHolder

	// The name of the ledger.
	//
	// This member is required.
	Name *string

	// The latest block location covered by the digest for which to request a proof. An
	// address is an Amazon Ion structure that has two fields: strandId and sequenceNo.
	// For example: {strandId:"BlFTjlSXze9BIh1KOszcE3",sequenceNo:49}
	DigestTipAddress *types.ValueHolder
}

type GetBlockOutput struct {

	// The block data object in Amazon Ion format.
	//
	// This member is required.
	Block *types.ValueHolder

	// The proof object in Amazon Ion format returned by a GetBlock request. A proof
	// contains the list of hash values required to recalculate the specified digest
	// using a Merkle tree, starting with the specified block.
	Proof *types.ValueHolder

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetBlockMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetBlock{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetBlock{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetBlock(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "qldb",
		OperationName: "GetBlock",
	}
}
