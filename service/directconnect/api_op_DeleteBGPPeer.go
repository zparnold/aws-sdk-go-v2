// Code generated by smithy-go-codegen DO NOT EDIT.

package directconnect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the specified BGP peer on the specified virtual interface with the
// specified customer address and ASN. You cannot delete the last BGP peer from a
// virtual interface.
func (c *Client) DeleteBGPPeer(ctx context.Context, params *DeleteBGPPeerInput, optFns ...func(*Options)) (*DeleteBGPPeerOutput, error) {
	stack := middleware.NewStack("DeleteBGPPeer", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteBGPPeerMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteBGPPeer(options.Region), middleware.Before)
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
			OperationName: "DeleteBGPPeer",
			Err:           err,
		}
	}
	out := result.(*DeleteBGPPeerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteBGPPeerInput struct {

	// The autonomous system (AS) number for Border Gateway Protocol (BGP)
	// configuration.
	Asn *int32

	// The ID of the BGP peer.
	BgpPeerId *string

	// The IP address assigned to the customer interface.
	CustomerAddress *string

	// The ID of the virtual interface.
	VirtualInterfaceId *string
}

type DeleteBGPPeerOutput struct {

	// The virtual interface.
	VirtualInterface *types.VirtualInterface

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteBGPPeerMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteBGPPeer{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteBGPPeer{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteBGPPeer(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "directconnect",
		OperationName: "DeleteBGPPeer",
	}
}
