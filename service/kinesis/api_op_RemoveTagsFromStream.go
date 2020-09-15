// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesis

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes tags from the specified Kinesis data stream. Removed tags are deleted
// and cannot be recovered after this operation successfully completes. If you
// specify a tag that does not exist, it is ignored. RemoveTagsFromStream () has a
// limit of five transactions per second per account.
func (c *Client) RemoveTagsFromStream(ctx context.Context, params *RemoveTagsFromStreamInput, optFns ...func(*Options)) (*RemoveTagsFromStreamOutput, error) {
	stack := middleware.NewStack("RemoveTagsFromStream", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpRemoveTagsFromStreamMiddlewares(stack)
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
	addOpRemoveTagsFromStreamValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRemoveTagsFromStream(options.Region), middleware.Before)

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
			OperationName: "RemoveTagsFromStream",
			Err:           err,
		}
	}
	out := result.(*RemoveTagsFromStreamOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for RemoveTagsFromStream.
type RemoveTagsFromStreamInput struct {
	// A list of tag keys. Each corresponding tag is removed from the stream.
	TagKeys []*string
	// The name of the stream.
	StreamName *string
}

type RemoveTagsFromStreamOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpRemoveTagsFromStreamMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpRemoveTagsFromStream{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpRemoveTagsFromStream{}, middleware.After)
}

func newServiceMetadataMiddleware_opRemoveTagsFromStream(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesis",
		OperationName: "RemoveTagsFromStream",
	}
}
