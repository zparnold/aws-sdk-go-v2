// Code generated by smithy-go-codegen DO NOT EDIT.

package workmail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a paginated call to list the aliases associated with a given entity.
func (c *Client) ListAliases(ctx context.Context, params *ListAliasesInput, optFns ...func(*Options)) (*ListAliasesOutput, error) {
	if params == nil {
		params = &ListAliasesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAliases", params, optFns, addOperationListAliasesMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAliasesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAliasesInput struct {

	// The identifier for the entity for which to list the aliases.
	//
	// This member is required.
	EntityId *string

	// The identifier for the organization under which the entity exists.
	//
	// This member is required.
	OrganizationId *string

	// The maximum number of results to return in a single call.
	MaxResults *int32

	// The token to use to retrieve the next page of results. The first call does not
	// contain any tokens.
	NextToken *string
}

type ListAliasesOutput struct {

	// The entity's paginated aliases.
	Aliases []*string

	// The token to use to retrieve the next page of results. The value is "null" when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListAliasesMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListAliases{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListAliases{}, middleware.After)
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
	addOpListAliasesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListAliases(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListAliases(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workmail",
		OperationName: "ListAliases",
	}
}
