// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates a gateway's metadata, which includes the gateway's name and time zone.
// To specify which gateway to update, use the Amazon Resource Name (ARN) of the
// gateway in your request.  <note> <p>For Gateways activated after September 2,
// 2015, the gateway's ARN contains the gateway ID rather than the gateway name.
// However, changing the name of the gateway has no effect on the gateway's
// ARN.</p> </note>
func (c *Client) UpdateGatewayInformation(ctx context.Context, params *UpdateGatewayInformationInput, optFns ...func(*Options)) (*UpdateGatewayInformationOutput, error) {
	stack := middleware.NewStack("UpdateGatewayInformation", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateGatewayInformationMiddlewares(stack)
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
	addOpUpdateGatewayInformationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateGatewayInformation(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)

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
			OperationName: "UpdateGatewayInformation",
			Err:           err,
		}
	}
	out := result.(*UpdateGatewayInformationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateGatewayInformationInput struct {

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways () operation
	// to return a list of gateways for your account and AWS Region.
	//
	// This member is required.
	GatewayARN *string

	// The Amazon Resource Name (ARN) of the Amazon CloudWatch log group that you want
	// to use to monitor and log events in the gateway.  <p>For more information, see
	// <a
	// href="https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/WhatIsCloudWatchLogs.html">What
	// is Amazon CloudWatch logs?</a>.</p>
	CloudWatchLogGroupARN *string

	// The name you configured for your gateway.
	GatewayName *string

	// A value that indicates the time zone of the gateway.
	GatewayTimezone *string
}

// A JSON object containing the Amazon Resource Name (ARN) of the gateway that was
// updated.
type UpdateGatewayInformationOutput struct {

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways () operation
	// to return a list of gateways for your account and AWS Region.
	GatewayARN *string

	// The name you configured for your gateway.
	GatewayName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateGatewayInformationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateGatewayInformation{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateGatewayInformation{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateGatewayInformation(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "UpdateGatewayInformation",
	}
}
