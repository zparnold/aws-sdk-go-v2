// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes tag keys from the specified resource.
func (c *Client) RemoveTagsFromResource(ctx context.Context, params *RemoveTagsFromResourceInput, optFns ...func(*Options)) (*RemoveTagsFromResourceOutput, error) {
	stack := middleware.NewStack("RemoveTagsFromResource", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpRemoveTagsFromResourceMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpRemoveTagsFromResourceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRemoveTagsFromResource(options.Region), middleware.Before)

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
			OperationName: "RemoveTagsFromResource",
			Err:           err,
		}
	}
	out := result.(*RemoveTagsFromResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RemoveTagsFromResourceInput struct {
	// The ID of the resource from which you want to remove tags. For example:
	// ManagedInstance: mi-012345abcde MaintenanceWindow: mw-012345abcde PatchBaseline:
	// pb-012345abcde For the Document and Parameter values, use the name of the
	// resource. The ManagedInstance type for this API action is only for on-premises
	// managed instances. Specify the name of the managed instance in the following
	// format: mi-ID_number. For example, mi-1a2b3c4d5e6f.
	ResourceId *string
	// Tag keys that you want to remove from the specified resource.
	TagKeys []*string
	// The type of resource from which you want to remove a tag. The ManagedInstance
	// type for this API action is only for on-premises managed instances. Specify the
	// name of the managed instance in the following format: mi-ID_number. For example,
	// mi-1a2b3c4d5e6f.
	ResourceType types.ResourceTypeForTagging
}

type RemoveTagsFromResourceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpRemoveTagsFromResourceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpRemoveTagsFromResource{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpRemoveTagsFromResource{}, middleware.After)
}

func newServiceMetadataMiddleware_opRemoveTagsFromResource(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "RemoveTagsFromResource",
	}
}
