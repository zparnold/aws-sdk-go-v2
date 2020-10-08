// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpoint

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a journey from an application.
func (c *Client) DeleteJourney(ctx context.Context, params *DeleteJourneyInput, optFns ...func(*Options)) (*DeleteJourneyOutput, error) {
	if params == nil {
		params = &DeleteJourneyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteJourney", params, optFns, addOperationDeleteJourneyMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteJourneyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteJourneyInput struct {

	// The unique identifier for the application. This identifier is displayed as the
	// Project ID on the Amazon Pinpoint console.
	//
	// This member is required.
	ApplicationId *string

	// The unique identifier for the journey.
	//
	// This member is required.
	JourneyId *string
}

type DeleteJourneyOutput struct {

	// Provides information about the status, configuration, and other settings for a
	// journey.
	//
	// This member is required.
	JourneyResponse *types.JourneyResponse

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteJourneyMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteJourney{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteJourney{}, middleware.After)
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
	addOpDeleteJourneyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteJourney(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteJourney(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mobiletargeting",
		OperationName: "DeleteJourney",
	}
}
