// Code generated by smithy-go-codegen DO NOT EDIT.

package sesv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Displays the template object (which includes the subject line, HTML part and
// text part) for the template you specify.  <p>You can execute this operation no
// more than once per second.</p>
func (c *Client) GetEmailTemplate(ctx context.Context, params *GetEmailTemplateInput, optFns ...func(*Options)) (*GetEmailTemplateOutput, error) {
	if params == nil {
		params = &GetEmailTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetEmailTemplate", params, optFns, addOperationGetEmailTemplateMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*GetEmailTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to display the template object (which includes the subject
// line, HTML part and text part) for the template you specify.
type GetEmailTemplateInput struct {

	// The name of the template you want to retrieve.
	//
	// This member is required.
	TemplateName *string
}

// The following element is returned by the service.
type GetEmailTemplateOutput struct {

	// The content of the email template, composed of a subject line, an HTML part, and
	// a text-only part.
	//
	// This member is required.
	TemplateContent *types.EmailTemplateContent

	// The name of the template you want to retrieve.
	//
	// This member is required.
	TemplateName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetEmailTemplateMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetEmailTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetEmailTemplate{}, middleware.After)
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
	addOpGetEmailTemplateValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetEmailTemplate(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetEmailTemplate(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "GetEmailTemplate",
	}
}
