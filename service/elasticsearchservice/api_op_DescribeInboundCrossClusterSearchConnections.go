// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticsearchservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists all the inbound cross-cluster search connections for a destination domain.
func (c *Client) DescribeInboundCrossClusterSearchConnections(ctx context.Context, params *DescribeInboundCrossClusterSearchConnectionsInput, optFns ...func(*Options)) (*DescribeInboundCrossClusterSearchConnectionsOutput, error) {
	stack := middleware.NewStack("DescribeInboundCrossClusterSearchConnections", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeInboundCrossClusterSearchConnectionsMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeInboundCrossClusterSearchConnections(options.Region), middleware.Before)

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
			OperationName: "DescribeInboundCrossClusterSearchConnections",
			Err:           err,
		}
	}
	out := result.(*DescribeInboundCrossClusterSearchConnectionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to the DescribeInboundCrossClusterSearchConnections
// () operation.
type DescribeInboundCrossClusterSearchConnectionsInput struct {
	// NextToken is sent in case the earlier API call results contain the NextToken. It
	// is used for pagination.
	NextToken *string
	// Set this value to limit the number of results returned. If not specified,
	// defaults to 100.
	MaxResults *int32
	// A list of filters used to match properties for inbound cross-cluster search
	// connection. Available Filter () names for this operation are:
	//
	//     *
	// cross-cluster-search-connection-id
	//
	//     * source-domain-info.domain-name
	//
	//     *
	// source-domain-info.owner-id
	//
	//     * source-domain-info.region
	//
	//     *
	// destination-domain-info.domain-name
	Filters []*types.Filter
}

// The result of a DescribeInboundCrossClusterSearchConnections () request.
// Contains the list of connections matching the filter criteria.
type DescribeInboundCrossClusterSearchConnectionsOutput struct {
	// If more results are available and NextToken is present, make the next request to
	// the same API with the received NextToken to paginate the remaining results.
	NextToken *string
	// Consists of list of InboundCrossClusterSearchConnection () matching the
	// specified filter criteria.
	CrossClusterSearchConnections []*types.InboundCrossClusterSearchConnection

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeInboundCrossClusterSearchConnectionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeInboundCrossClusterSearchConnections{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeInboundCrossClusterSearchConnections{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeInboundCrossClusterSearchConnections(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "es",
		OperationName: "DescribeInboundCrossClusterSearchConnections",
	}
}
