// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a display name for a customer managed customer master key (CMK). You can
// use an alias to identify a CMK in cryptographic operations
// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#cryptographic-operations),
// such as Encrypt () and GenerateDataKey (). You can change the CMK associated
// with the alias at any time. Aliases are easier to remember than key IDs. They
// can also help to simplify your applications. For example, if you use an alias in
// your code, you can change the CMK your code uses by associating a given alias
// with a different CMK. To run the same code in multiple AWS regions, use an alias
// in your code, such as alias/ApplicationKey. Then, in each AWS Region, create an
// alias/ApplicationKey alias that is associated with a CMK in that Region. When
// you run your code, it uses the alias/ApplicationKey CMK for that AWS Region
// without any Region-specific code. This operation does not return a response. To
// get the alias that you created, use the ListAliases () operation.  <p>To use
// aliases successfully, be aware of the following information.</p> <ul> <li>
// <p>Each alias points to only one CMK at a time, although a single CMK can have
// multiple aliases. The alias and its associated CMK must be in the same AWS
// account and Region. </p> </li> <li> <p>You can associate an alias with any
// customer managed CMK in the same AWS account and Region. However, you do not
// have permission to associate an alias with an <a
// href="https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-managed-cmk">AWS
// managed CMK</a> or an <a
// href="https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-owned-cmk">AWS
// owned CMK</a>. </p> </li> <li> <p>To change the CMK associated with an alias,
// use the <a>UpdateAlias</a> operation. The current CMK and the new CMK must be
// the same type (both symmetric or both asymmetric) and they must have the same
// key usage (<code>ENCRYPT_DECRYPT</code> or <code>SIGN_VERIFY</code>). This
// restriction prevents cryptographic errors in code that uses aliases.</p> </li>
// <li> <p>The alias name must begin with <code>alias/</code> followed by a name,
// such as <code>alias/ExampleAlias</code>. It can contain only alphanumeric
// characters, forward slashes (/), underscores (_), and dashes (-). The alias name
// cannot begin with <code>alias/aws/</code>. The <code>alias/aws/</code> prefix is
// reserved for <a
// href="https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-managed-cmk">AWS
// managed CMKs</a>.
// * The alias name must be unique within an AWS Region. However,
// you can use the same alias name in multiple Regions of the same AWS account.
// Each instance of the alias is associated with a CMK in its Region.
//
// * After you
// create an alias, you cannot change its alias name. However, you can use the
// DeleteAlias () operation to delete the alias and then create a new alias with
// the desired name.
//
// * You can use an alias name or alias ARN to identify a CMK in
// AWS KMS cryptographic operations
// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#cryptographic-operations)
// and in the DescribeKey () operation. However, you cannot use alias names or
// alias ARNs in API operations that manage CMKs, such as DisableKey () or
// GetKeyPolicy (). For information about the valid CMK identifiers for each AWS
// KMS API operation, see the descriptions of the KeyId parameter in the API
// operation documentation.
//  <p>Because an alias is not a property of a CMK, you
// can delete and change the aliases of a CMK without affecting the CMK. Also,
// aliases do not appear in the response from the <a>DescribeKey</a> operation. To
// get the aliases and alias ARNs of CMKs in each AWS account and Region, use the
// <a>ListAliases</a> operation.</p> <p>The CMK that you use for this operation
// must be in a compatible key state. For  details, see How Key State Affects Use
// of a Customer Master Key
// (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html) in the
// AWS Key Management Service Developer Guide.
func (c *Client) CreateAlias(ctx context.Context, params *CreateAliasInput, optFns ...func(*Options)) (*CreateAliasOutput, error) {
	if params == nil {
		params = &CreateAliasInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAlias", params, optFns, addOperationCreateAliasMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAliasOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAliasInput struct {

	// Specifies the alias name. This value must begin with alias/ followed by a name,
	// such as alias/ExampleAlias. The alias name cannot begin with alias/aws/. The
	// alias/aws/ prefix is reserved for AWS managed CMKs.
	//
	// This member is required.
	AliasName *string

	// Identifies the CMK to which the alias refers. Specify the key ID or the Amazon
	// Resource Name (ARN) of the CMK. You cannot specify another alias. For help
	// finding the key ID and ARN, see Finding the Key ID and ARN
	// (https://docs.aws.amazon.com/kms/latest/developerguide/viewing-keys.html#find-cmk-id-arn)
	// in the AWS Key Management Service Developer Guide.
	//
	// This member is required.
	TargetKeyId *string
}

type CreateAliasOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateAliasMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateAlias{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateAlias{}, middleware.After)
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
	addOpCreateAliasValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAlias(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateAlias(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kms",
		OperationName: "CreateAlias",
	}
}
