// Code generated by smithy-go-codegen DO NOT EDIT.

package frauddetector

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets all labels or a specific label if name is provided. This is a paginated
// API. If you provide a null maxResults, this action retrieves a maximum of 50
// records per page. If you provide a maxResults, the value must be between 10 and
// 50. To get the next page results, provide the pagination token from the
// GetGetLabelsResponse as part of your request. A null pagination token fetches
// the records from the beginning.
func (c *Client) GetLabels(ctx context.Context, params *GetLabelsInput, optFns ...func(*Options)) (*GetLabelsOutput, error) {
	if params == nil {
		params = &GetLabelsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetLabels", params, optFns, addOperationGetLabelsMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*GetLabelsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetLabelsInput struct {

	// The maximum number of objects to return for the request.
	MaxResults *int32

	// The name of the label or labels to get.
	Name *string

	// The next token for the subsequent request.
	NextToken *string
}

type GetLabelsOutput struct {

	// An array of labels.
	Labels []*types.Label

	// The next page token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetLabelsMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetLabels{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetLabels{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetLabels(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetLabels(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "frauddetector",
		OperationName: "GetLabels",
	}
}
