// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists all classifier objects in the Data Catalog.
func (c *Client) GetClassifiers(ctx context.Context, params *GetClassifiersInput, optFns ...func(*Options)) (*GetClassifiersOutput, error) {
	if params == nil {
		params = &GetClassifiersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetClassifiers", params, optFns, addOperationGetClassifiersMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*GetClassifiersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetClassifiersInput struct {

	// The size of the list to return (optional).
	MaxResults *int32

	// An optional continuation token.
	NextToken *string
}

type GetClassifiersOutput struct {

	// The requested list of classifier objects.
	Classifiers []*types.Classifier

	// A continuation token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetClassifiersMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetClassifiers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetClassifiers{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetClassifiers(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetClassifiers(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "GetClassifiers",
	}
}
