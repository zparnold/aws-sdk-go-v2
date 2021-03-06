// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelbuildingservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelbuildingservice/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns slot type information as follows:
//
//     * If you specify the nameContains
// field, returns the $LATEST version of all slot types that contain the specified
// string.
//
//     * If you don't specify the nameContains field, returns information
// about the $LATEST version of all slot types.
//
// The operation requires permission
// for the lex:GetSlotTypes action.
func (c *Client) GetSlotTypes(ctx context.Context, params *GetSlotTypesInput, optFns ...func(*Options)) (*GetSlotTypesOutput, error) {
	stack := middleware.NewStack("GetSlotTypes", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetSlotTypesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetSlotTypes(options.Region), middleware.Before)
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
			OperationName: "GetSlotTypes",
			Err:           err,
		}
	}
	out := result.(*GetSlotTypesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSlotTypesInput struct {

	// The maximum number of slot types to return in the response. The default is 10.
	MaxResults *int32

	// Substring to match in slot type names. A slot type will be returned if any part
	// of its name matches the substring. For example, "xyz" matches both "xyzabc" and
	// "abcxyz."
	NameContains *string

	// A pagination token that fetches the next page of slot types. If the response to
	// this API call is truncated, Amazon Lex returns a pagination token in the
	// response. To fetch next page of slot types, specify the pagination token in the
	// next request.
	NextToken *string
}

type GetSlotTypesOutput struct {

	// If the response is truncated, it includes a pagination token that you can
	// specify in your next request to fetch the next page of slot types.
	NextToken *string

	// An array of objects, one for each slot type, that provides information such as
	// the name of the slot type, the version, and a description.
	SlotTypes []*types.SlotTypeMetadata

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetSlotTypesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetSlotTypes{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetSlotTypes{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetSlotTypes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "GetSlotTypes",
	}
}
