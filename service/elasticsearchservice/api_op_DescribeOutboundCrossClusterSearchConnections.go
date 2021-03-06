// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticsearchservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists all the outbound cross-cluster search connections for a source domain.
func (c *Client) DescribeOutboundCrossClusterSearchConnections(ctx context.Context, params *DescribeOutboundCrossClusterSearchConnectionsInput, optFns ...func(*Options)) (*DescribeOutboundCrossClusterSearchConnectionsOutput, error) {
	stack := middleware.NewStack("DescribeOutboundCrossClusterSearchConnections", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeOutboundCrossClusterSearchConnectionsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeOutboundCrossClusterSearchConnections(options.Region), middleware.Before)
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
			OperationName: "DescribeOutboundCrossClusterSearchConnections",
			Err:           err,
		}
	}
	out := result.(*DescribeOutboundCrossClusterSearchConnectionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to the
// DescribeOutboundCrossClusterSearchConnections () operation.
type DescribeOutboundCrossClusterSearchConnectionsInput struct {

	// A list of filters used to match properties for outbound cross-cluster search
	// connection. Available Filter () names for this operation are:
	//
	//     *
	// cross-cluster-search-connection-id
	//
	//     * destination-domain-info.domain-name
	//
	//
	// * destination-domain-info.owner-id
	//
	//     * destination-domain-info.region
	//
	//     *
	// source-domain-info.domain-name
	Filters []*types.Filter

	// Set this value to limit the number of results returned. If not specified,
	// defaults to 100.
	MaxResults *int32

	// NextToken is sent in case the earlier API call results contain the NextToken. It
	// is used for pagination.
	NextToken *string
}

// The result of a DescribeOutboundCrossClusterSearchConnections () request.
// Contains the list of connections matching the filter criteria.
type DescribeOutboundCrossClusterSearchConnectionsOutput struct {

	// Consists of list of OutboundCrossClusterSearchConnection () matching the
	// specified filter criteria.
	CrossClusterSearchConnections []*types.OutboundCrossClusterSearchConnection

	// If more results are available and NextToken is present, make the next request to
	// the same API with the received NextToken to paginate the remaining results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeOutboundCrossClusterSearchConnectionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeOutboundCrossClusterSearchConnections{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeOutboundCrossClusterSearchConnections{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeOutboundCrossClusterSearchConnections(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "es",
		OperationName: "DescribeOutboundCrossClusterSearchConnections",
	}
}
