// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehend

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/comprehend/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Determines the dominant language of the input text for a batch of documents. For
// a list of languages that Amazon Comprehend can detect, see Amazon Comprehend
// Supported Languages
// (https://docs.aws.amazon.com/comprehend/latest/dg/how-languages.html).
func (c *Client) BatchDetectDominantLanguage(ctx context.Context, params *BatchDetectDominantLanguageInput, optFns ...func(*Options)) (*BatchDetectDominantLanguageOutput, error) {
	stack := middleware.NewStack("BatchDetectDominantLanguage", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpBatchDetectDominantLanguageMiddlewares(stack)
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
	addOpBatchDetectDominantLanguageValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opBatchDetectDominantLanguage(options.Region), middleware.Before)
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
			OperationName: "BatchDetectDominantLanguage",
			Err:           err,
		}
	}
	out := result.(*BatchDetectDominantLanguageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchDetectDominantLanguageInput struct {

	// A list containing the text of the input documents. The list can contain a
	// maximum of 25 documents. Each document should contain at least 20 characters and
	// must contain fewer than 5,000 bytes of UTF-8 encoded characters.
	//
	// This member is required.
	TextList []*string
}

type BatchDetectDominantLanguageOutput struct {

	// A list containing one object for each document that contained an error. The
	// results are sorted in ascending order by the Index field and match the order of
	// the documents in the input list. If there are no errors in the batch, the
	// ErrorList is empty.
	//
	// This member is required.
	ErrorList []*types.BatchItemError

	// A list of objects containing the results of the operation. The results are
	// sorted in ascending order by the Index field and match the order of the
	// documents in the input list. If all of the documents contain an error, the
	// ResultList is empty.
	//
	// This member is required.
	ResultList []*types.BatchDetectDominantLanguageItemResult

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpBatchDetectDominantLanguageMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpBatchDetectDominantLanguage{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchDetectDominantLanguage{}, middleware.After)
}

func newServiceMetadataMiddleware_opBatchDetectDominantLanguage(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "comprehend",
		OperationName: "BatchDetectDominantLanguage",
	}
}
