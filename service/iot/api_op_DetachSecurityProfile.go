// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Disassociates a Device Defender security profile from a thing group or from this
// account.
func (c *Client) DetachSecurityProfile(ctx context.Context, params *DetachSecurityProfileInput, optFns ...func(*Options)) (*DetachSecurityProfileOutput, error) {
	stack := middleware.NewStack("DetachSecurityProfile", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDetachSecurityProfileMiddlewares(stack)
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
	addOpDetachSecurityProfileValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDetachSecurityProfile(options.Region), middleware.Before)

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
			OperationName: "DetachSecurityProfile",
			Err:           err,
		}
	}
	out := result.(*DetachSecurityProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DetachSecurityProfileInput struct {
	// The security profile that is detached.
	SecurityProfileName *string
	// The ARN of the thing group from which the security profile is detached.
	SecurityProfileTargetArn *string
}

type DetachSecurityProfileOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDetachSecurityProfileMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDetachSecurityProfile{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDetachSecurityProfile{}, middleware.After)
}

func newServiceMetadataMiddleware_opDetachSecurityProfile(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "DetachSecurityProfile",
	}
}
