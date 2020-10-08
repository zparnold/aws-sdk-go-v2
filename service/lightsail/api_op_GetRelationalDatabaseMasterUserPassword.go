// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Returns the current, previous, or pending versions of the master user password
// for a Lightsail database. The GetRelationalDatabaseMasterUserPassword operation
// supports tag-based access control via resource tags applied to the resource
// identified by relationalDatabaseName.
func (c *Client) GetRelationalDatabaseMasterUserPassword(ctx context.Context, params *GetRelationalDatabaseMasterUserPasswordInput, optFns ...func(*Options)) (*GetRelationalDatabaseMasterUserPasswordOutput, error) {
	if params == nil {
		params = &GetRelationalDatabaseMasterUserPasswordInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetRelationalDatabaseMasterUserPassword", params, optFns, addOperationGetRelationalDatabaseMasterUserPasswordMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*GetRelationalDatabaseMasterUserPasswordOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetRelationalDatabaseMasterUserPasswordInput struct {

	// The name of your database for which to get the master user password.
	//
	// This member is required.
	RelationalDatabaseName *string

	// The password version to return. Specifying CURRENT or PREVIOUS returns the
	// current or previous passwords respectively. Specifying PENDING returns the
	// newest version of the password that will rotate to CURRENT. After the PENDING
	// password rotates to CURRENT, the PENDING password is no longer available.
	// Default: CURRENT
	PasswordVersion types.RelationalDatabasePasswordVersion
}

type GetRelationalDatabaseMasterUserPasswordOutput struct {

	// The timestamp when the specified version of the master user password was
	// created.
	CreatedAt *time.Time

	// The master user password for the password version specified.
	MasterUserPassword *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetRelationalDatabaseMasterUserPasswordMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetRelationalDatabaseMasterUserPassword{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetRelationalDatabaseMasterUserPassword{}, middleware.After)
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
	addOpGetRelationalDatabaseMasterUserPasswordValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetRelationalDatabaseMasterUserPassword(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetRelationalDatabaseMasterUserPassword(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "GetRelationalDatabaseMasterUserPassword",
	}
}
