// Code generated by smithy-go-codegen DO NOT EDIT.

package lakeformation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lakeformation/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the Lake Formation permissions for a specified table or database
// resource located at a path in Amazon S3. GetEffectivePermissionsForPath will not
// return databases and tables if the catalog is encrypted.
func (c *Client) GetEffectivePermissionsForPath(ctx context.Context, params *GetEffectivePermissionsForPathInput, optFns ...func(*Options)) (*GetEffectivePermissionsForPathOutput, error) {
	if params == nil {
		params = &GetEffectivePermissionsForPathInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetEffectivePermissionsForPath", params, optFns, addOperationGetEffectivePermissionsForPathMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*GetEffectivePermissionsForPathOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetEffectivePermissionsForPathInput struct {

	// The Amazon Resource Name (ARN) of the resource for which you want to get
	// permissions.
	//
	// This member is required.
	ResourceArn *string

	// The identifier for the Data Catalog. By default, the account ID. The Data
	// Catalog is the persistent metadata store. It contains database definitions,
	// table definitions, and other control information to manage your AWS Lake
	// Formation environment.
	CatalogId *string

	// The maximum number of results to return.
	MaxResults *int32

	// A continuation token, if this is not the first call to retrieve this list.
	NextToken *string
}

type GetEffectivePermissionsForPathOutput struct {

	// A continuation token, if this is not the first call to retrieve this list.
	NextToken *string

	// A list of the permissions for the specified table or database resource located
	// at the path in Amazon S3.
	Permissions []*types.PrincipalResourcePermissions

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetEffectivePermissionsForPathMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetEffectivePermissionsForPath{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetEffectivePermissionsForPath{}, middleware.After)
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
	addOpGetEffectivePermissionsForPathValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetEffectivePermissionsForPath(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetEffectivePermissionsForPath(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lakeformation",
		OperationName: "GetEffectivePermissionsForPath",
	}
}
