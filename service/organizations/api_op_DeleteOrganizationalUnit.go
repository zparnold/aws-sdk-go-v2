// Code generated by smithy-go-codegen DO NOT EDIT.

package organizations

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes an organizational unit (OU) from a root or another OU. You must first
// remove all accounts and child OUs from the OU that you want to delete. This
// operation can be called only from the organization's master account.
func (c *Client) DeleteOrganizationalUnit(ctx context.Context, params *DeleteOrganizationalUnitInput, optFns ...func(*Options)) (*DeleteOrganizationalUnitOutput, error) {
	stack := middleware.NewStack("DeleteOrganizationalUnit", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteOrganizationalUnitMiddlewares(stack)
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
	addOpDeleteOrganizationalUnitValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteOrganizationalUnit(options.Region), middleware.Before)
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
			OperationName: "DeleteOrganizationalUnit",
			Err:           err,
		}
	}
	out := result.(*DeleteOrganizationalUnitOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteOrganizationalUnitInput struct {

	// The unique identifier (ID) of the organizational unit that you want to delete.
	// You can get the ID from the ListOrganizationalUnitsForParent () operation. The
	// regex pattern (http://wikipedia.org/wiki/regex) for an organizational unit ID
	// string requires "ou-" followed by from 4 to 32 lowercase letters or digits (the
	// ID of the root that contains the OU). This string is followed by a second "-"
	// dash and from 8 to 32 additional lowercase letters or digits.
	//
	// This member is required.
	OrganizationalUnitId *string
}

type DeleteOrganizationalUnitOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteOrganizationalUnitMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteOrganizationalUnit{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteOrganizationalUnit{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteOrganizationalUnit(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "organizations",
		OperationName: "DeleteOrganizationalUnit",
	}
}
