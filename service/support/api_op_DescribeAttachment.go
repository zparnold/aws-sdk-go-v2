// Code generated by smithy-go-codegen DO NOT EDIT.

package support

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/support/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the attachment that has the specified ID. Attachments can include
// screenshots, error logs, or other files that describe your issue. Attachment IDs
// are generated by the case management system when you add an attachment to a case
// or case communication. Attachment IDs are returned in the AttachmentDetails ()
// objects that are returned by the DescribeCommunications () operation.
//
//     * You
// must have a Business or Enterprise support plan to use the AWS Support API.
//
//
// * If you call the AWS Support API from an account that does not have a Business
// or Enterprise support plan, the SubscriptionRequiredException error message
// appears. For information about changing your support plan, see AWS Support
// (http://aws.amazon.com/premiumsupport/).
func (c *Client) DescribeAttachment(ctx context.Context, params *DescribeAttachmentInput, optFns ...func(*Options)) (*DescribeAttachmentOutput, error) {
	stack := middleware.NewStack("DescribeAttachment", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeAttachmentMiddlewares(stack)
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
	addOpDescribeAttachmentValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAttachment(options.Region), middleware.Before)
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
			OperationName: "DescribeAttachment",
			Err:           err,
		}
	}
	out := result.(*DescribeAttachmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAttachmentInput struct {

	// The ID of the attachment to return. Attachment IDs are returned by the
	// DescribeCommunications () operation.
	//
	// This member is required.
	AttachmentId *string
}

// The content and file name of the attachment returned by the DescribeAttachment
// () operation.
type DescribeAttachmentOutput struct {

	// This object includes the attachment content and file name. In the previous
	// response syntax, the value for the data parameter appears as blob, which is
	// represented as a base64-encoded string. The value for fileName is the name of
	// the attachment, such as troubleshoot-screenshot.png.
	Attachment *types.Attachment

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeAttachmentMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeAttachment{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeAttachment{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeAttachment(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "support",
		OperationName: "DescribeAttachment",
	}
}
