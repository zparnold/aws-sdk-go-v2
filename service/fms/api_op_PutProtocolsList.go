// Code generated by smithy-go-codegen DO NOT EDIT.

package fms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/fms/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an AWS Firewall Manager protocols list.
func (c *Client) PutProtocolsList(ctx context.Context, params *PutProtocolsListInput, optFns ...func(*Options)) (*PutProtocolsListOutput, error) {
	stack := middleware.NewStack("PutProtocolsList", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpPutProtocolsListMiddlewares(stack)
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
	addOpPutProtocolsListValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutProtocolsList(options.Region), middleware.Before)
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
			OperationName: "PutProtocolsList",
			Err:           err,
		}
	}
	out := result.(*PutProtocolsListOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutProtocolsListInput struct {

	// The details of the AWS Firewall Manager protocols list to be created.
	//
	// This member is required.
	ProtocolsList *types.ProtocolsListData

	// The tags associated with the resource.
	TagList []*types.Tag
}

type PutProtocolsListOutput struct {

	// The details of the AWS Firewall Manager protocols list.
	ProtocolsList *types.ProtocolsListData

	// The Amazon Resource Name (ARN) of the protocols list.
	ProtocolsListArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpPutProtocolsListMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpPutProtocolsList{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutProtocolsList{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutProtocolsList(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "fms",
		OperationName: "PutProtocolsList",
	}
}
