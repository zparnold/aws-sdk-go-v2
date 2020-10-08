// Code generated by smithy-go-codegen DO NOT EDIT.

package directconnect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Deletes the specified connection. Deleting a connection only stops the AWS
// Direct Connect port hour and data transfer charges. If you are partnering with
// any third parties to connect with the AWS Direct Connect location, you must
// cancel your service with them separately.
func (c *Client) DeleteConnection(ctx context.Context, params *DeleteConnectionInput, optFns ...func(*Options)) (*DeleteConnectionOutput, error) {
	if params == nil {
		params = &DeleteConnectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteConnection", params, optFns, addOperationDeleteConnectionMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteConnectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteConnectionInput struct {

	// The ID of the connection.
	//
	// This member is required.
	ConnectionId *string
}

// Information about an AWS Direct Connect connection.
type DeleteConnectionOutput struct {

	// The Direct Connect endpoint on which the physical connection terminates.
	AwsDevice *string

	// The Direct Connect endpoint on which the physical connection terminates.
	AwsDeviceV2 *string

	// The bandwidth of the connection.
	Bandwidth *string

	// The ID of the connection.
	ConnectionId *string

	// The name of the connection.
	ConnectionName *string

	// The state of the connection. The following are the possible values:
	//
	//     *
	// ordering: The initial state of a hosted connection provisioned on an
	// interconnect. The connection stays in the ordering state until the owner of the
	// hosted connection confirms or declines the connection order.
	//
	//     * requested:
	// The initial state of a standard connection. The connection stays in the
	// requested state until the Letter of Authorization (LOA) is sent to the
	// customer.
	//
	//     * pending: The connection has been approved and is being
	// initialized.
	//
	//     * available: The network link is up and the connection is
	// ready for use.
	//
	//     * down: The network link is down.
	//
	//     * deleting: The
	// connection is being deleted.
	//
	//     * deleted: The connection has been deleted.
	//
	//
	// * rejected: A hosted connection in the ordering state enters the rejected state
	// if it is deleted by the customer.
	//
	//     * unknown: The state of the connection is
	// not available.
	ConnectionState types.ConnectionState

	// Indicates whether the connection supports a secondary BGP peer in the same
	// address family (IPv4/IPv6).
	HasLogicalRedundancy types.HasLogicalRedundancy

	// Indicates whether jumbo frames (9001 MTU) are supported.
	JumboFrameCapable *bool

	// The ID of the LAG.
	LagId *string

	// The time of the most recent call to DescribeLoa () for this connection.
	LoaIssueTime *time.Time

	// The location of the connection.
	Location *string

	// The ID of the AWS account that owns the connection.
	OwnerAccount *string

	// The name of the AWS Direct Connect service provider associated with the
	// connection.
	PartnerName *string

	// The name of the service provider associated with the connection.
	ProviderName *string

	// The AWS Region where the connection is located.
	Region *string

	// The tags associated with the connection.
	Tags []*types.Tag

	// The ID of the VLAN.
	Vlan *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteConnectionMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteConnection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteConnection{}, middleware.After)
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
	addOpDeleteConnectionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteConnection(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteConnection(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "directconnect",
		OperationName: "DeleteConnection",
	}
}
