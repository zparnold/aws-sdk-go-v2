// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves details for the specified phone number ID, such as associations,
// capabilities, and product type.
func (c *Client) GetPhoneNumber(ctx context.Context, params *GetPhoneNumberInput, optFns ...func(*Options)) (*GetPhoneNumberOutput, error) {
	stack := middleware.NewStack("GetPhoneNumber", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetPhoneNumberMiddlewares(stack)
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
	addOpGetPhoneNumberValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetPhoneNumber(options.Region), middleware.Before)

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
			OperationName: "GetPhoneNumber",
			Err:           err,
		}
	}
	out := result.(*GetPhoneNumberOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetPhoneNumberInput struct {
	// The phone number ID.
	PhoneNumberId *string
}

type GetPhoneNumberOutput struct {
	// The phone number details.
	PhoneNumber *types.PhoneNumber

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetPhoneNumberMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetPhoneNumber{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetPhoneNumber{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetPhoneNumber(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "GetPhoneNumber",
	}
}
