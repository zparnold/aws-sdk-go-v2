// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a paginated list of self-service actions associated with the specified
// Product ID and Provisioning Artifact ID.
func (c *Client) ListServiceActionsForProvisioningArtifact(ctx context.Context, params *ListServiceActionsForProvisioningArtifactInput, optFns ...func(*Options)) (*ListServiceActionsForProvisioningArtifactOutput, error) {
	stack := middleware.NewStack("ListServiceActionsForProvisioningArtifact", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListServiceActionsForProvisioningArtifactMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpListServiceActionsForProvisioningArtifactValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListServiceActionsForProvisioningArtifact(options.Region), middleware.Before)

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
			OperationName: "ListServiceActionsForProvisioningArtifact",
			Err:           err,
		}
	}
	out := result.(*ListServiceActionsForProvisioningArtifactOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListServiceActionsForProvisioningArtifactInput struct {
	// The identifier of the provisioning artifact. For example, pa-4abcdjnxjj6ne.
	ProvisioningArtifactId *string
	// The page token for the next set of results. To retrieve the first set of
	// results, use null.
	PageToken *string
	// The product identifier. For example, prod-abcdzk7xy33qa.
	ProductId *string
	// The language code.
	//
	//     * en - English (default)
	//
	//     * jp - Japanese
	//
	//     * zh
	// - Chinese
	AcceptLanguage *string
	// The maximum number of items to return with this call.
	PageSize *int32
}

type ListServiceActionsForProvisioningArtifactOutput struct {
	// The page token to use to retrieve the next set of results. If there are no
	// additional results, this value is null.
	NextPageToken *string
	// An object containing information about the self-service actions associated with
	// the provisioning artifact.
	ServiceActionSummaries []*types.ServiceActionSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListServiceActionsForProvisioningArtifactMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListServiceActionsForProvisioningArtifact{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListServiceActionsForProvisioningArtifact{}, middleware.After)
}

func newServiceMetadataMiddleware_opListServiceActionsForProvisioningArtifact(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "ListServiceActionsForProvisioningArtifact",
	}
}
