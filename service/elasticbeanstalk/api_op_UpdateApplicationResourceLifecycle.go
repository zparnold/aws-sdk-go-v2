// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticbeanstalk

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Modifies lifecycle settings for an application.
func (c *Client) UpdateApplicationResourceLifecycle(ctx context.Context, params *UpdateApplicationResourceLifecycleInput, optFns ...func(*Options)) (*UpdateApplicationResourceLifecycleOutput, error) {
	if params == nil {
		params = &UpdateApplicationResourceLifecycleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateApplicationResourceLifecycle", params, optFns, addOperationUpdateApplicationResourceLifecycleMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateApplicationResourceLifecycleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateApplicationResourceLifecycleInput struct {

	// The name of the application.
	//
	// This member is required.
	ApplicationName *string

	// The lifecycle configuration.
	//
	// This member is required.
	ResourceLifecycleConfig *types.ApplicationResourceLifecycleConfig
}

type UpdateApplicationResourceLifecycleOutput struct {

	// The name of the application.
	ApplicationName *string

	// The lifecycle configuration.
	ResourceLifecycleConfig *types.ApplicationResourceLifecycleConfig

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateApplicationResourceLifecycleMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpUpdateApplicationResourceLifecycle{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpUpdateApplicationResourceLifecycle{}, middleware.After)
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
	addOpUpdateApplicationResourceLifecycleValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateApplicationResourceLifecycle(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateApplicationResourceLifecycle(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticbeanstalk",
		OperationName: "UpdateApplicationResourceLifecycle",
	}
}
