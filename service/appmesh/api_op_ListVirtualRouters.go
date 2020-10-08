// Code generated by smithy-go-codegen DO NOT EDIT.

package appmesh

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appmesh/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of existing virtual routers in a service mesh.
func (c *Client) ListVirtualRouters(ctx context.Context, params *ListVirtualRoutersInput, optFns ...func(*Options)) (*ListVirtualRoutersOutput, error) {
	if params == nil {
		params = &ListVirtualRoutersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListVirtualRouters", params, optFns, addOperationListVirtualRoutersMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*ListVirtualRoutersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type ListVirtualRoutersInput struct {

	// The name of the service mesh to list virtual routers in.
	//
	// This member is required.
	MeshName *string

	// The maximum number of results returned by ListVirtualRouters in paginated
	// output. When you use this parameter, ListVirtualRouters returns only limit
	// results in a single page along with a nextToken response element. You can see
	// the remaining results of the initial request by sending another
	// ListVirtualRouters request with the returned nextToken value. This value can be
	// between 1 and 100. If you don't use this parameter, ListVirtualRouters returns
	// up to 100 results and a nextToken value if applicable.
	Limit *int32

	// The AWS IAM account ID of the service mesh owner. If the account ID is not your
	// own, then it's the ID of the account that shared the mesh with your account. For
	// more information about mesh sharing, see Working with shared meshes
	// (https://docs.aws.amazon.com/app-mesh/latest/userguide/sharing.html).
	MeshOwner *string

	// The nextToken value returned from a previous paginated ListVirtualRouters
	// request where limit was used and the results exceeded the value of that
	// parameter. Pagination continues from the end of the previous results that
	// returned the nextToken value.
	NextToken *string
}

//
type ListVirtualRoutersOutput struct {

	// The list of existing virtual routers for the specified service mesh.
	//
	// This member is required.
	VirtualRouters []*types.VirtualRouterRef

	// The nextToken value to include in a future ListVirtualRouters request. When the
	// results of a ListVirtualRouters request exceed limit, you can use this value to
	// retrieve the next page of results. This value is null when there are no more
	// results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListVirtualRoutersMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListVirtualRouters{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListVirtualRouters{}, middleware.After)
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
	addOpListVirtualRoutersValidationMiddleware(stack)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}
