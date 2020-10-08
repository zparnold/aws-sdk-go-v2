// Code generated by smithy-go-codegen DO NOT EDIT.

package licensemanager

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/licensemanager/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the license configurations for the specified resource.
func (c *Client) ListLicenseSpecificationsForResource(ctx context.Context, params *ListLicenseSpecificationsForResourceInput, optFns ...func(*Options)) (*ListLicenseSpecificationsForResourceOutput, error) {
	if params == nil {
		params = &ListLicenseSpecificationsForResourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListLicenseSpecificationsForResource", params, optFns, addOperationListLicenseSpecificationsForResourceMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*ListLicenseSpecificationsForResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListLicenseSpecificationsForResourceInput struct {

	// Amazon Resource Name (ARN) of a resource that has an associated license
	// configuration.
	//
	// This member is required.
	ResourceArn *string

	// Maximum number of results to return in a single call.
	MaxResults *int32

	// Token for the next set of results.
	NextToken *string
}

type ListLicenseSpecificationsForResourceOutput struct {

	// License configurations associated with a resource.
	LicenseSpecifications []*types.LicenseSpecification

	// Token for the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListLicenseSpecificationsForResourceMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListLicenseSpecificationsForResource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListLicenseSpecificationsForResource{}, middleware.After)
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
	addOpListLicenseSpecificationsForResourceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListLicenseSpecificationsForResource(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListLicenseSpecificationsForResource(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "license-manager",
		OperationName: "ListLicenseSpecificationsForResource",
	}
}
