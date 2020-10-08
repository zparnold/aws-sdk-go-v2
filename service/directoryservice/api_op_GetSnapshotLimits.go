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

// Obtains the manual snapshot limits for a directory.
func (c *Client) GetSnapshotLimits(ctx context.Context, params *GetSnapshotLimitsInput, optFns ...func(*Options)) (*GetSnapshotLimitsOutput, error) {
	if params == nil {
		params = &GetSnapshotLimitsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSnapshotLimits", params, optFns, addOperationGetSnapshotLimitsMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSnapshotLimitsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the inputs for the GetSnapshotLimits () operation.
type GetSnapshotLimitsInput struct {

	// Contains the identifier of the directory to obtain the limits for.
	//
	// This member is required.
	DirectoryId *string
}

// Contains the results of the GetSnapshotLimits () operation.
type GetSnapshotLimitsOutput struct {

	// A SnapshotLimits () object that contains the manual snapshot limits for the
	// specified directory.
	SnapshotLimits *types.SnapshotLimits

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetSnapshotLimitsMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetSnapshotLimits{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetSnapshotLimits{}, middleware.After)
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
	addOpGetSnapshotLimitsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetSnapshotLimits(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetSnapshotLimits(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ds",
		OperationName: "GetSnapshotLimits",
	}
}
