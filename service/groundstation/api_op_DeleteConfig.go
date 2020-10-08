// Code generated by smithy-go-codegen DO NOT EDIT.

package groundstation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/groundstation/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a Config.
func (c *Client) DeleteConfig(ctx context.Context, params *DeleteConfigInput, optFns ...func(*Options)) (*DeleteConfigOutput, error) {
	if params == nil {
		params = &DeleteConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteConfig", params, optFns, addOperationDeleteConfigMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DeleteConfigInput struct {

	// UUID of a Config.
	//
	// This member is required.
	ConfigId *string

	// Type of a Config.
	//
	// This member is required.
	ConfigType types.ConfigCapabilityType
}

//
type DeleteConfigOutput struct {

	// ARN of a Config.
	ConfigArn *string

	// UUID of a Config.
	ConfigId *string

	// Type of a Config.
	ConfigType types.ConfigCapabilityType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteConfigMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteConfig{}, middleware.After)
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
	addOpDeleteConfigValidationMiddleware(stack)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}
