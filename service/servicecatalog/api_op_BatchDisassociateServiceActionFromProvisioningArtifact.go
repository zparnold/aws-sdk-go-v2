// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Disassociates a batch of self-service actions from the specified provisioning
// artifact.
func (c *Client) BatchDisassociateServiceActionFromProvisioningArtifact(ctx context.Context, params *BatchDisassociateServiceActionFromProvisioningArtifactInput, optFns ...func(*Options)) (*BatchDisassociateServiceActionFromProvisioningArtifactOutput, error) {
	stack := middleware.NewStack("BatchDisassociateServiceActionFromProvisioningArtifact", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpBatchDisassociateServiceActionFromProvisioningArtifactMiddlewares(stack)
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
	addOpBatchDisassociateServiceActionFromProvisioningArtifactValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opBatchDisassociateServiceActionFromProvisioningArtifact(options.Region), middleware.Before)
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
			OperationName: "BatchDisassociateServiceActionFromProvisioningArtifact",
			Err:           err,
		}
	}
	out := result.(*BatchDisassociateServiceActionFromProvisioningArtifactOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchDisassociateServiceActionFromProvisioningArtifactInput struct {

	// One or more associations, each consisting of the Action ID, the Product ID, and
	// the Provisioning Artifact ID.
	//
	// This member is required.
	ServiceActionAssociations []*types.ServiceActionAssociation

	// The language code.
	//
	//     * en - English (default)
	//
	//     * jp - Japanese
	//
	//     * zh
	// - Chinese
	AcceptLanguage *string
}

type BatchDisassociateServiceActionFromProvisioningArtifactOutput struct {

	// An object that contains a list of errors, along with information to help you
	// identify the self-service action.
	FailedServiceActionAssociations []*types.FailedServiceActionAssociation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpBatchDisassociateServiceActionFromProvisioningArtifactMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpBatchDisassociateServiceActionFromProvisioningArtifact{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchDisassociateServiceActionFromProvisioningArtifact{}, middleware.After)
}

func newServiceMetadataMiddleware_opBatchDisassociateServiceActionFromProvisioningArtifact(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "BatchDisassociateServiceActionFromProvisioningArtifact",
	}
}
