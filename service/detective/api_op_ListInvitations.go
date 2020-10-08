// Code generated by smithy-go-codegen DO NOT EDIT.

package detective

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/detective/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves the list of open and accepted behavior graph invitations for the
// member account. This operation can only be called by a member account. Open
// invitations are invitations that the member account has not responded to. The
// results do not include behavior graphs for which the member account declined the
// invitation. The results also do not include behavior graphs that the member
// account resigned from or was removed from.
func (c *Client) ListInvitations(ctx context.Context, params *ListInvitationsInput, optFns ...func(*Options)) (*ListInvitationsOutput, error) {
	if params == nil {
		params = &ListInvitationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListInvitations", params, optFns, addOperationListInvitationsMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*ListInvitationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListInvitationsInput struct {

	// The maximum number of behavior graph invitations to return in the response. The
	// total must be less than the overall limit on the number of results to return,
	// which is currently 200.
	MaxResults *int32

	// For requests to retrieve the next page of results, the pagination token that was
	// returned with the previous page of results. The initial request does not include
	// a pagination token.
	NextToken *string
}

type ListInvitationsOutput struct {

	// The list of behavior graphs for which the member account has open or accepted
	// invitations.
	Invitations []*types.MemberDetail

	// If there are more behavior graphs remaining in the results, then this is the
	// pagination token to use to request the next page of behavior graphs.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListInvitationsMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListInvitations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListInvitations{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListInvitations(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListInvitations(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "detective",
		OperationName: "ListInvitations",
	}
}
