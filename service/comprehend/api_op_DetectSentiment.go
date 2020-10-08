// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehend

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/comprehend/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Inspects text and returns an inference of the prevailing sentiment (POSITIVE,
// NEUTRAL, MIXED, or NEGATIVE).
func (c *Client) DetectSentiment(ctx context.Context, params *DetectSentimentInput, optFns ...func(*Options)) (*DetectSentimentOutput, error) {
	if params == nil {
		params = &DetectSentimentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DetectSentiment", params, optFns, addOperationDetectSentimentMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DetectSentimentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DetectSentimentInput struct {

	// The language of the input documents. You can specify any of the primary
	// languages supported by Amazon Comprehend. All documents must be in the same
	// language.
	//
	// This member is required.
	LanguageCode types.LanguageCode

	// A UTF-8 text string. Each string must contain fewer that 5,000 bytes of UTF-8
	// encoded characters.
	//
	// This member is required.
	Text *string
}

type DetectSentimentOutput struct {

	// The inferred sentiment that Amazon Comprehend has the highest level of
	// confidence in.
	Sentiment types.SentimentType

	// An object that lists the sentiments, and their corresponding confidence levels.
	SentimentScore *types.SentimentScore

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDetectSentimentMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDetectSentiment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDetectSentiment{}, middleware.After)
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
	addOpDetectSentimentValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDetectSentiment(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDetectSentiment(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "comprehend",
		OperationName: "DetectSentiment",
	}
}
