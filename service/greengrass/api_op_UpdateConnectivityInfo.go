// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrass

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/greengrass/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the connectivity information for the core. Any devices that belong to
// the group which has this core will receive this information in order to find the
// location of the core and connect to it.
func (c *Client) UpdateConnectivityInfo(ctx context.Context, params *UpdateConnectivityInfoInput, optFns ...func(*Options)) (*UpdateConnectivityInfoOutput, error) {
	if params == nil {
		params = &UpdateConnectivityInfoInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateConnectivityInfo", params, optFns, addOperationUpdateConnectivityInfoMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateConnectivityInfoOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Connectivity information.
type UpdateConnectivityInfoInput struct {

	// The thing name.
	//
	// This member is required.
	ThingName *string

	// A list of connectivity info.
	ConnectivityInfo []*types.ConnectivityInfo
}

type UpdateConnectivityInfoOutput struct {

	// A message about the connectivity info update request.
	Message *string

	// The new version of the connectivity info.
	Version *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateConnectivityInfoMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateConnectivityInfo{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateConnectivityInfo{}, middleware.After)
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
	addOpUpdateConnectivityInfoValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateConnectivityInfo(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateConnectivityInfo(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "greengrass",
		OperationName: "UpdateConnectivityInfo",
	}
}
