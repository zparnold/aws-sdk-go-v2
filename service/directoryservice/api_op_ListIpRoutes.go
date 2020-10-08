// Code generated by smithy-go-codegen DO NOT EDIT.

package directoryservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/directoryservice/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the address blocks that you have added to a directory.
func (c *Client) ListIpRoutes(ctx context.Context, params *ListIpRoutesInput, optFns ...func(*Options)) (*ListIpRoutesOutput, error) {
	if params == nil {
		params = &ListIpRoutesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListIpRoutes", params, optFns, addOperationListIpRoutesMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*ListIpRoutesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListIpRoutesInput struct {

	// Identifier (ID) of the directory for which you want to retrieve the IP
	// addresses.
	//
	// This member is required.
	DirectoryId *string

	// Maximum number of items to return. If this value is zero, the maximum number of
	// items is specified by the limitations of the operation.
	Limit *int32

	// The ListIpRoutes.NextToken value from a previous call to ListIpRoutes (). Pass
	// null if this is the first call.
	NextToken *string
}

type ListIpRoutesOutput struct {

	// A list of IpRoute ()s.
	IpRoutesInfo []*types.IpRouteInfo

	// If not null, more results are available. Pass this value for the NextToken
	// parameter in a subsequent call to ListIpRoutes () to retrieve the next set of
	// items.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListIpRoutesMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListIpRoutes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListIpRoutes{}, middleware.After)
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
	addOpListIpRoutesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListIpRoutes(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListIpRoutes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ds",
		OperationName: "ListIpRoutes",
	}
}
