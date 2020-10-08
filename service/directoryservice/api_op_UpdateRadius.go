// Code generated by smithy-go-codegen DO NOT EDIT.

package directoryservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/directoryservice/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the Remote Authentication Dial In User Service (RADIUS) server
// information for an AD Connector or Microsoft AD directory.
func (c *Client) UpdateRadius(ctx context.Context, params *UpdateRadiusInput, optFns ...func(*Options)) (*UpdateRadiusOutput, error) {
	if params == nil {
		params = &UpdateRadiusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateRadius", params, optFns, addOperationUpdateRadiusMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateRadiusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the inputs for the UpdateRadius () operation.
type UpdateRadiusInput struct {

	// The identifier of the directory for which to update the RADIUS server
	// information.
	//
	// This member is required.
	DirectoryId *string

	// A RadiusSettings () object that contains information about the RADIUS server.
	//
	// This member is required.
	RadiusSettings *types.RadiusSettings
}

// Contains the results of the UpdateRadius () operation.
type UpdateRadiusOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateRadiusMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateRadius{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateRadius{}, middleware.After)
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
	addOpUpdateRadiusValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateRadius(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateRadius(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ds",
		OperationName: "UpdateRadius",
	}
}
