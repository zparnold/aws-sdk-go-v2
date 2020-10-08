// Code generated by smithy-go-codegen DO NOT EDIT.

package fms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Permanently deletes an AWS Firewall Manager protocols list.
func (c *Client) DeleteProtocolsList(ctx context.Context, params *DeleteProtocolsListInput, optFns ...func(*Options)) (*DeleteProtocolsListOutput, error) {
	if params == nil {
		params = &DeleteProtocolsListInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteProtocolsList", params, optFns, addOperationDeleteProtocolsListMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteProtocolsListOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteProtocolsListInput struct {

	// The ID of the protocols list that you want to delete. You can retrieve this ID
	// from PutProtocolsList, ListProtocolsLists, and GetProtocolsLost.
	//
	// This member is required.
	ListId *string
}

type DeleteProtocolsListOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteProtocolsListMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteProtocolsList{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteProtocolsList{}, middleware.After)
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
	addOpDeleteProtocolsListValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteProtocolsList(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteProtocolsList(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "fms",
		OperationName: "DeleteProtocolsList",
	}
}
