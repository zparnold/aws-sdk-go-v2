// Code generated by smithy-go-codegen DO NOT EDIT.

package devicefarm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a link to a currently running remote access session.
func (c *Client) GetRemoteAccessSession(ctx context.Context, params *GetRemoteAccessSessionInput, optFns ...func(*Options)) (*GetRemoteAccessSessionOutput, error) {
	stack := middleware.NewStack("GetRemoteAccessSession", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetRemoteAccessSessionMiddlewares(stack)
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
	addOpGetRemoteAccessSessionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetRemoteAccessSession(options.Region), middleware.Before)

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
			OperationName: "GetRemoteAccessSession",
			Err:           err,
		}
	}
	out := result.(*GetRemoteAccessSessionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to get information about the specified remote access
// session.
type GetRemoteAccessSessionInput struct {
	// The Amazon Resource Name (ARN) of the remote access session about which you want
	// to get session information.
	Arn *string
}

// Represents the response from the server that lists detailed information about
// the remote access session.
type GetRemoteAccessSessionOutput struct {
	// A container that lists detailed information about the remote access session.
	RemoteAccessSession *types.RemoteAccessSession

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetRemoteAccessSessionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetRemoteAccessSession{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetRemoteAccessSession{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetRemoteAccessSession(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devicefarm",
		OperationName: "GetRemoteAccessSession",
	}
}
