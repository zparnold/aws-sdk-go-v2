// Code generated by smithy-go-codegen DO NOT EDIT.

package sms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes existing replication configuration for an application.
func (c *Client) DeleteAppReplicationConfiguration(ctx context.Context, params *DeleteAppReplicationConfigurationInput, optFns ...func(*Options)) (*DeleteAppReplicationConfigurationOutput, error) {
	if params == nil {
		params = &DeleteAppReplicationConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteAppReplicationConfiguration", params, optFns, addOperationDeleteAppReplicationConfigurationMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteAppReplicationConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteAppReplicationConfigurationInput struct {

	// ID of the application associated with the replication configuration.
	AppId *string
}

type DeleteAppReplicationConfigurationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteAppReplicationConfigurationMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteAppReplicationConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteAppReplicationConfiguration{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteAppReplicationConfiguration(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteAppReplicationConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sms",
		OperationName: "DeleteAppReplicationConfiguration",
	}
}
