// Code generated by smithy-go-codegen DO NOT EDIT.

package fms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Permanently deletes an AWS Firewall Manager applications list.
func (c *Client) DeleteAppsList(ctx context.Context, params *DeleteAppsListInput, optFns ...func(*Options)) (*DeleteAppsListOutput, error) {
	if params == nil {
		params = &DeleteAppsListInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteAppsList", params, optFns, addOperationDeleteAppsListMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteAppsListOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteAppsListInput struct {

	// The ID of the applications list that you want to delete. You can retrieve this
	// ID from PutAppsList, ListAppsLists, and GetAppsList.
	//
	// This member is required.
	ListId *string
}

type DeleteAppsListOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteAppsListMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteAppsList{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteAppsList{}, middleware.After)
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
	addOpDeleteAppsListValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteAppsList(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteAppsList(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "fms",
		OperationName: "DeleteAppsList",
	}
}
