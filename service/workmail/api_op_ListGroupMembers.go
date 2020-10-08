// Code generated by smithy-go-codegen DO NOT EDIT.

package workmail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workmail/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns an overview of the members of a group. Users and groups can be members
// of a group.
func (c *Client) ListGroupMembers(ctx context.Context, params *ListGroupMembersInput, optFns ...func(*Options)) (*ListGroupMembersOutput, error) {
	if params == nil {
		params = &ListGroupMembersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListGroupMembers", params, optFns, addOperationListGroupMembersMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*ListGroupMembersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListGroupMembersInput struct {

	// The identifier for the group to which the members (users or groups) are
	// associated.
	//
	// This member is required.
	GroupId *string

	// The identifier for the organization under which the group exists.
	//
	// This member is required.
	OrganizationId *string

	// The maximum number of results to return in a single call.
	MaxResults *int32

	// The token to use to retrieve the next page of results. The first call does not
	// contain any tokens.
	NextToken *string
}

type ListGroupMembersOutput struct {

	// The members associated to the group.
	Members []*types.Member

	// The token to use to retrieve the next page of results. The first call does not
	// contain any tokens.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListGroupMembersMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListGroupMembers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListGroupMembers{}, middleware.After)
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
	addOpListGroupMembersValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListGroupMembers(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListGroupMembers(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workmail",
		OperationName: "ListGroupMembers",
	}
}
