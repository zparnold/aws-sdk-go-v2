// Code generated by smithy-go-codegen DO NOT EDIT.

package workspaces

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Migrates a WorkSpace from one operating system or bundle type to another, while
// retaining the data on the user volume.  <p>The migration process recreates the
// WorkSpace by using a new root volume from the target bundle image and the user
// volume from the last available snapshot of the original WorkSpace. During
// migration, the original <code>D:\Users\%USERNAME%</code> user profile folder is
// renamed to <code>D:\Users\%USERNAME%MMddyyTHHmmss%.NotMigrated</code>. A new
// <code>D:\Users\%USERNAME%\</code> folder is generated by the new OS. Certain
// files in the old user profile are moved to the new user profile.</p> <p>For
// available migration scenarios, details about what happens during migration, and
// best practices, see <a
// href="https://docs.aws.amazon.com/workspaces/latest/adminguide/migrate-workspaces.html">Migrate
// a WorkSpace</a>.</p>
func (c *Client) MigrateWorkspace(ctx context.Context, params *MigrateWorkspaceInput, optFns ...func(*Options)) (*MigrateWorkspaceOutput, error) {
	if params == nil {
		params = &MigrateWorkspaceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "MigrateWorkspace", params, optFns, addOperationMigrateWorkspaceMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*MigrateWorkspaceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type MigrateWorkspaceInput struct {

	// The identifier of the target bundle type to migrate the WorkSpace to.
	//
	// This member is required.
	BundleId *string

	// The identifier of the WorkSpace to migrate from.
	//
	// This member is required.
	SourceWorkspaceId *string
}

type MigrateWorkspaceOutput struct {

	// The original identifier of the WorkSpace that is being migrated.
	SourceWorkspaceId *string

	// The new identifier of the WorkSpace that is being migrated. If the migration
	// does not succeed, the target WorkSpace ID will not be used, and the WorkSpace
	// will still have the original WorkSpace ID.
	TargetWorkspaceId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationMigrateWorkspaceMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpMigrateWorkspace{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpMigrateWorkspace{}, middleware.After)
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
	addOpMigrateWorkspaceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opMigrateWorkspace(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opMigrateWorkspace(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workspaces",
		OperationName: "MigrateWorkspace",
	}
}
