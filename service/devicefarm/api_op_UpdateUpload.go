// Code generated by smithy-go-codegen DO NOT EDIT.

package devicefarm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates an uploaded test spec.
func (c *Client) UpdateUpload(ctx context.Context, params *UpdateUploadInput, optFns ...func(*Options)) (*UpdateUploadOutput, error) {
	if params == nil {
		params = &UpdateUploadInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateUpload", params, optFns, addOperationUpdateUploadMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateUploadOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateUploadInput struct {

	// The Amazon Resource Name (ARN) of the uploaded test spec.
	//
	// This member is required.
	Arn *string

	// The upload's content type (for example, application/x-yaml).
	ContentType *string

	// Set to true if the YAML file has changed and must be updated. Otherwise, set to
	// false.
	EditContent *bool

	// The upload's test spec file name. The name must not contain any forward slashes
	// (/). The test spec file name must end with the .yaml or .yml file extension.
	Name *string
}

type UpdateUploadOutput struct {

	// A test spec uploaded to Device Farm.
	Upload *types.Upload

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateUploadMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateUpload{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateUpload{}, middleware.After)
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
	addOpUpdateUploadValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateUpload(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateUpload(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devicefarm",
		OperationName: "UpdateUpload",
	}
}
