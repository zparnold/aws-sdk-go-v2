// Code generated by smithy-go-codegen DO NOT EDIT.

package amplify

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/amplify/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new domain association for an Amplify app.
func (c *Client) UpdateDomainAssociation(ctx context.Context, params *UpdateDomainAssociationInput, optFns ...func(*Options)) (*UpdateDomainAssociationOutput, error) {
	stack := middleware.NewStack("UpdateDomainAssociation", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateDomainAssociationMiddlewares(stack)
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
	addOpUpdateDomainAssociationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDomainAssociation(options.Region), middleware.Before)
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
			OperationName: "UpdateDomainAssociation",
			Err:           err,
		}
	}
	out := result.(*UpdateDomainAssociationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request structure for the update domain association request.
type UpdateDomainAssociationInput struct {

	// The unique ID for an Amplify app.
	//
	// This member is required.
	AppId *string

	// The name of the domain.
	//
	// This member is required.
	DomainName *string

	// Describes the settings for the subdomain.
	//
	// This member is required.
	SubDomainSettings []*types.SubDomainSetting

	// Sets the branch patterns for automatic subdomain creation.
	AutoSubDomainCreationPatterns []*string

	// The required AWS Identity and Access Management (IAM) service role for the
	// Amazon Resource Name (ARN) for automatically creating subdomains.
	AutoSubDomainIAMRole *string

	// Enables the automated creation of subdomains for branches.
	EnableAutoSubDomain *bool
}

// The result structure for the update domain association request.
type UpdateDomainAssociationOutput struct {

	// Describes a domain association, which associates a custom domain with an Amplify
	// app.
	//
	// This member is required.
	DomainAssociation *types.DomainAssociation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateDomainAssociationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateDomainAssociation{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateDomainAssociation{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateDomainAssociation(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "amplify",
		OperationName: "UpdateDomainAssociation",
	}
}
