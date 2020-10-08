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

// Creates a classifier in the user's account. This can be a GrokClassifier, an
// XMLClassifier, a JsonClassifier, or a CsvClassifier, depending on which field of
// the request is present.
func (c *Client) CreateClassifier(ctx context.Context, params *CreateClassifierInput, optFns ...func(*Options)) (*CreateClassifierOutput, error) {
	if params == nil {
		params = &CreateClassifierInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateClassifier", params, optFns, addOperationCreateClassifierMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateClassifierOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateClassifierInput struct {

	// A CsvClassifier object specifying the classifier to create.
	CsvClassifier *types.CreateCsvClassifierRequest

	// A GrokClassifier object specifying the classifier to create.
	GrokClassifier *types.CreateGrokClassifierRequest

	// A JsonClassifier object specifying the classifier to create.
	JsonClassifier *types.CreateJsonClassifierRequest

	// An XMLClassifier object specifying the classifier to create.
	XMLClassifier *types.CreateXMLClassifierRequest
}

type CreateClassifierOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateClassifierMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateClassifier{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateClassifier{}, middleware.After)
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
	addOpCreateClassifierValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateClassifier(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateClassifier(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "CreateClassifier",
	}
}
