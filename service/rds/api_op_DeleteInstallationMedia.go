// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the installation medium for a DB engine that requires an on-premises
// customer provided license, such as Microsoft SQL Server.
func (c *Client) DeleteInstallationMedia(ctx context.Context, params *DeleteInstallationMediaInput, optFns ...func(*Options)) (*DeleteInstallationMediaOutput, error) {
	if params == nil {
		params = &DeleteInstallationMediaInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteInstallationMedia", params, optFns, addOperationDeleteInstallationMediaMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteInstallationMediaOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteInstallationMediaInput struct {

	// The installation medium ID.
	//
	// This member is required.
	InstallationMediaId *string
}

// Contains the installation media for a DB engine that requires an on-premises
// customer provided license, such as Microsoft SQL Server.
type DeleteInstallationMediaOutput struct {

	// The custom Availability Zone (AZ) that contains the installation media.
	CustomAvailabilityZoneId *string

	// The DB engine.
	Engine *string

	// The path to the installation medium for the DB engine.
	EngineInstallationMediaPath *string

	// The engine version of the DB engine.
	EngineVersion *string

	// If an installation media failure occurred, the cause of the failure.
	FailureCause *types.InstallationMediaFailureCause

	// The installation medium ID.
	InstallationMediaId *string

	// The path to the installation medium for the operating system associated with the
	// DB engine.
	OSInstallationMediaPath *string

	// The status of the installation medium.
	Status *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteInstallationMediaMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeleteInstallationMedia{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteInstallationMedia{}, middleware.After)
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
	addOpDeleteInstallationMediaValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteInstallationMedia(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteInstallationMedia(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DeleteInstallationMedia",
	}
}
