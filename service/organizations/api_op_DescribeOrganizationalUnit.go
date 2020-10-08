// Code generated by smithy-go-codegen DO NOT EDIT.

package organizations

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves information about an organizational unit (OU). This operation can be
// called only from the organization's master account or by a member account that
// is a delegated administrator for an AWS service.
func (c *Client) DescribeOrganizationalUnit(ctx context.Context, params *DescribeOrganizationalUnitInput, optFns ...func(*Options)) (*DescribeOrganizationalUnitOutput, error) {
	if params == nil {
		params = &DescribeOrganizationalUnitInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeOrganizationalUnit", params, optFns, addOperationDescribeOrganizationalUnitMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeOrganizationalUnitOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeOrganizationalUnitInput struct {

	// The unique identifier (ID) of the organizational unit that you want details
	// about. You can get the ID from the ListOrganizationalUnitsForParent ()
	// operation. The regex pattern (http://wikipedia.org/wiki/regex) for an
	// organizational unit ID string requires "ou-" followed by from 4 to 32 lowercase
	// letters or digits (the ID of the root that contains the OU). This string is
	// followed by a second "-" dash and from 8 to 32 additional lowercase letters or
	// digits.
	//
	// This member is required.
	OrganizationalUnitId *string
}

type DescribeOrganizationalUnitOutput struct {

	// A structure that contains details about the specified OU.
	OrganizationalUnit *types.OrganizationalUnit

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeOrganizationalUnitMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeOrganizationalUnit{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeOrganizationalUnit{}, middleware.After)
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
	addOpDescribeOrganizationalUnitValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeOrganizationalUnit(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeOrganizationalUnit(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "organizations",
		OperationName: "DescribeOrganizationalUnit",
	}
}
