// Code generated by smithy-go-codegen DO NOT EDIT.

package ses

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes an email template. You can execute this operation no more than once per
// second.
func (c *Client) DeleteTemplate(ctx context.Context, params *DeleteTemplateInput, optFns ...func(*Options)) (*DeleteTemplateOutput, error) {
	if params == nil {
		params = &DeleteTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteTemplate", params, optFns, addOperationDeleteTemplateMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to delete an email template. For more information, see the
// Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/send-personalized-email-api.html).
type DeleteTemplateInput struct {

	// The name of the template to be deleted.
	//
	// This member is required.
	TemplateName *string
}

type DeleteTemplateOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteTemplateMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeleteTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteTemplate{}, middleware.After)
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
	addOpDeleteTemplateValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteTemplate(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteTemplate(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "DeleteTemplate",
	}
}
