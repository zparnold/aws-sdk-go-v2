// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// A list of inventory items returned by the request.
func (c *Client) ListInventoryEntries(ctx context.Context, params *ListInventoryEntriesInput, optFns ...func(*Options)) (*ListInventoryEntriesOutput, error) {
	if params == nil {
		params = &ListInventoryEntriesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListInventoryEntries", params, optFns, addOperationListInventoryEntriesMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*ListInventoryEntriesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListInventoryEntriesInput struct {

	// The instance ID for which you want inventory information.
	//
	// This member is required.
	InstanceId *string

	// The type of inventory item for which you want information.
	//
	// This member is required.
	TypeName *string

	// One or more filters. Use a filter to return a more specific list of results.
	Filters []*types.InventoryFilter

	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults *int32

	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string
}

type ListInventoryEntriesOutput struct {

	// The time that inventory information was collected for the instance(s).
	CaptureTime *string

	// A list of inventory items on the instance(s).
	Entries []map[string]*string

	// The instance ID targeted by the request to query inventory information.
	InstanceId *string

	// The token to use when requesting the next set of items. If there are no
	// additional items to return, the string is empty.
	NextToken *string

	// The inventory schema version used by the instance(s).
	SchemaVersion *string

	// The type of inventory item returned by the request.
	TypeName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListInventoryEntriesMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListInventoryEntries{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListInventoryEntries{}, middleware.After)
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
	addOpListInventoryEntriesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListInventoryEntries(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListInventoryEntries(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "ListInventoryEntries",
	}
}
