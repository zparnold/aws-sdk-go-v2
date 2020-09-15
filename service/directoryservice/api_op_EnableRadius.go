// Code generated by smithy-go-codegen DO NOT EDIT.

package directoryservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/directoryservice/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Enables multi-factor authentication (MFA) with the Remote Authentication Dial In
// User Service (RADIUS) server for an AD Connector or Microsoft AD directory.
func (c *Client) EnableRadius(ctx context.Context, params *EnableRadiusInput, optFns ...func(*Options)) (*EnableRadiusOutput, error) {
	stack := middleware.NewStack("EnableRadius", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpEnableRadiusMiddlewares(stack)
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
	addOpEnableRadiusValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opEnableRadius(options.Region), middleware.Before)

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
			OperationName: "EnableRadius",
			Err:           err,
		}
	}
	out := result.(*EnableRadiusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the inputs for the EnableRadius () operation.
type EnableRadiusInput struct {
	// A RadiusSettings () object that contains information about the RADIUS server.
	RadiusSettings *types.RadiusSettings
	// The identifier of the directory for which to enable MFA.
	DirectoryId *string
}

// Contains the results of the EnableRadius () operation.
type EnableRadiusOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpEnableRadiusMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpEnableRadius{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpEnableRadius{}, middleware.After)
}

func newServiceMetadataMiddleware_opEnableRadius(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ds",
		OperationName: "EnableRadius",
	}
}
