// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitosync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cognitosync/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the specific dataset. The dataset will be deleted permanently, and the
// action can't be undone. Datasets that this dataset was merged with will no
// longer report the merge. Any subsequent operation on this dataset will result in
// a ResourceNotFoundException. This API can be called with temporary user
// credentials provided by Cognito Identity or with developer credentials.
func (c *Client) DeleteDataset(ctx context.Context, params *DeleteDatasetInput, optFns ...func(*Options)) (*DeleteDatasetOutput, error) {
	if params == nil {
		params = &DeleteDatasetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteDataset", params, optFns, addOperationDeleteDatasetMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteDatasetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to delete the specific dataset.
type DeleteDatasetInput struct {

	// A string of up to 128 characters. Allowed characters are a-z, A-Z, 0-9, '_'
	// (underscore), '-' (dash), and '.' (dot).
	//
	// This member is required.
	DatasetName *string

	// A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE)
	// created by Amazon Cognito. GUID generation is unique within a region.
	//
	// This member is required.
	IdentityId *string

	// A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE)
	// created by Amazon Cognito. GUID generation is unique within a region.
	//
	// This member is required.
	IdentityPoolId *string
}

// Response to a successful DeleteDataset request.
type DeleteDatasetOutput struct {

	// A collection of data for an identity pool. An identity pool can have multiple
	// datasets. A dataset is per identity and can be general or associated with a
	// particular entity in an application (like a saved game). Datasets are
	// automatically created if they don't exist. Data is synced by dataset, and a
	// dataset can hold up to 1MB of key-value pairs.
	Dataset *types.Dataset

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteDatasetMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteDataset{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteDataset{}, middleware.After)
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
	addOpDeleteDatasetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteDataset(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteDataset(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-sync",
		OperationName: "DeleteDataset",
	}
}
