// Code generated by smithy-go-codegen DO NOT EDIT.

package applicationinsights

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes the specified application from monitoring. Does not delete the
// application.
func (c *Client) DeleteApplication(ctx context.Context, params *DeleteApplicationInput, optFns ...func(*Options)) (*DeleteApplicationOutput, error) {
	if params == nil {
		params = &DeleteApplicationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteApplication", params, optFns, addOperationDeleteApplicationMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteApplicationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteApplicationInput struct {

	// The name of the resource group.
	//
	// This member is required.
	ResourceGroupName *string
}

type DeleteApplicationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteApplicationMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteApplication{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteApplication{}, middleware.After)
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
	addOpDeleteApplicationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteApplication(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteApplication(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "applicationinsights",
		OperationName: "DeleteApplication",
	}
}
