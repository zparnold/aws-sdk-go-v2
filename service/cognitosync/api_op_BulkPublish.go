// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitosync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Initiates a bulk publish of all existing datasets for an Identity Pool to the
// configured stream. Customers are limited to one successful bulk publish per 24
// hours. Bulk publish is an asynchronous request, customers can see the status of
// the request via the GetBulkPublishDetails operation. This API can only be called
// with developer credentials. You cannot call this API with the temporary user
// credentials provided by Cognito Identity.
func (c *Client) BulkPublish(ctx context.Context, params *BulkPublishInput, optFns ...func(*Options)) (*BulkPublishOutput, error) {
	if params == nil {
		params = &BulkPublishInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BulkPublish", params, optFns, addOperationBulkPublishMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*BulkPublishOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the BulkPublish operation.
type BulkPublishInput struct {

	// A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE)
	// created by Amazon Cognito. GUID generation is unique within a region.
	//
	// This member is required.
	IdentityPoolId *string
}

// The output for the BulkPublish operation.
type BulkPublishOutput struct {

	// A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE)
	// created by Amazon Cognito. GUID generation is unique within a region.
	IdentityPoolId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationBulkPublishMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBulkPublish{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBulkPublish{}, middleware.After)
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
	addOpBulkPublishValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opBulkPublish(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opBulkPublish(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-sync",
		OperationName: "BulkPublish",
	}
}
