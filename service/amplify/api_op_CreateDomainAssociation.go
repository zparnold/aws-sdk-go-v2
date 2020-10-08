// Code generated by smithy-go-codegen DO NOT EDIT.

package amplify

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/amplify/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new domain association for an Amplify app. This action associates a
// custom domain with the Amplify app
func (c *Client) CreateDomainAssociation(ctx context.Context, params *CreateDomainAssociationInput, optFns ...func(*Options)) (*CreateDomainAssociationOutput, error) {
	if params == nil {
		params = &CreateDomainAssociationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDomainAssociation", params, optFns, addOperationCreateDomainAssociationMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDomainAssociationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request structure for the create domain association request.
type CreateDomainAssociationInput struct {

	// The unique ID for an Amplify app.
	//
	// This member is required.
	AppId *string

	// The domain name for the domain association.
	//
	// This member is required.
	DomainName *string

	// The setting for the subdomain.
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

// The result structure for the create domain association request.
type CreateDomainAssociationOutput struct {

	// Describes the structure of a domain association, which associates a custom
	// domain with an Amplify app.
	//
	// This member is required.
	DomainAssociation *types.DomainAssociation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateDomainAssociationMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateDomainAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateDomainAssociation{}, middleware.After)
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
	addOpCreateDomainAssociationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDomainAssociation(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateDomainAssociation(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "amplify",
		OperationName: "CreateDomainAssociation",
	}
}
