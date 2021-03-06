// Code generated by smithy-go-codegen DO NOT EDIT.

package workspaces

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the specified image from your account. To delete an image, you must
// first delete any bundles that are associated with the image and unshare the
// image if it is shared with other accounts.
func (c *Client) DeleteWorkspaceImage(ctx context.Context, params *DeleteWorkspaceImageInput, optFns ...func(*Options)) (*DeleteWorkspaceImageOutput, error) {
	stack := middleware.NewStack("DeleteWorkspaceImage", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteWorkspaceImageMiddlewares(stack)
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
	addOpDeleteWorkspaceImageValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteWorkspaceImage(options.Region), middleware.Before)
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
			OperationName: "DeleteWorkspaceImage",
			Err:           err,
		}
	}
	out := result.(*DeleteWorkspaceImageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteWorkspaceImageInput struct {

	// The identifier of the image.
	//
	// This member is required.
	ImageId *string
}

type DeleteWorkspaceImageOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteWorkspaceImageMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteWorkspaceImage{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteWorkspaceImage{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteWorkspaceImage(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workspaces",
		OperationName: "DeleteWorkspaceImage",
	}
}
