// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a dashboard.
func (c *Client) DeleteDashboard(ctx context.Context, params *DeleteDashboardInput, optFns ...func(*Options)) (*DeleteDashboardOutput, error) {
	if params == nil {
		params = &DeleteDashboardInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteDashboard", params, optFns, addOperationDeleteDashboardMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteDashboardOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteDashboardInput struct {

	// The ID of the AWS account that contains the dashboard that you're deleting.
	//
	// This member is required.
	AwsAccountId *string

	// The ID for the dashboard.
	//
	// This member is required.
	DashboardId *string

	// The version number of the dashboard. If the version number property is provided,
	// only the specified version of the dashboard is deleted.
	VersionNumber *int64
}

type DeleteDashboardOutput struct {

	// The Secure Socket Layer (SSL) properties that apply for the resource.
	Arn *string

	// The ID of the dashboard.
	DashboardId *string

	// The AWS request ID for this operation.
	RequestId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteDashboardMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteDashboard{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteDashboard{}, middleware.After)
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
	addOpDeleteDashboardValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteDashboard(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteDashboard(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "DeleteDashboard",
	}
}
