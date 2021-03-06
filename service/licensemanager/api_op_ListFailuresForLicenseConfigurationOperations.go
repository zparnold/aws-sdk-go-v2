// Code generated by smithy-go-codegen DO NOT EDIT.

package licensemanager

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/licensemanager/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the license configuration operations that failed.
func (c *Client) ListFailuresForLicenseConfigurationOperations(ctx context.Context, params *ListFailuresForLicenseConfigurationOperationsInput, optFns ...func(*Options)) (*ListFailuresForLicenseConfigurationOperationsOutput, error) {
	stack := middleware.NewStack("ListFailuresForLicenseConfigurationOperations", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListFailuresForLicenseConfigurationOperationsMiddlewares(stack)
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
	addOpListFailuresForLicenseConfigurationOperationsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListFailuresForLicenseConfigurationOperations(options.Region), middleware.Before)
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
			OperationName: "ListFailuresForLicenseConfigurationOperations",
			Err:           err,
		}
	}
	out := result.(*ListFailuresForLicenseConfigurationOperationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListFailuresForLicenseConfigurationOperationsInput struct {

	// Amazon Resource Name of the license configuration.
	//
	// This member is required.
	LicenseConfigurationArn *string

	// Maximum number of results to return in a single call.
	MaxResults *int32

	// Token for the next set of results.
	NextToken *string
}

type ListFailuresForLicenseConfigurationOperationsOutput struct {

	// License configuration operations that failed.
	LicenseOperationFailureList []*types.LicenseOperationFailure

	// Token for the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListFailuresForLicenseConfigurationOperationsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListFailuresForLicenseConfigurationOperations{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListFailuresForLicenseConfigurationOperations{}, middleware.After)
}

func newServiceMetadataMiddleware_opListFailuresForLicenseConfigurationOperations(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "license-manager",
		OperationName: "ListFailuresForLicenseConfigurationOperations",
	}
}
