// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the public key of an asymmetric CMK. Unlike the private key of a
// asymmetric CMK, which never leaves AWS KMS unencrypted, callers with
// kms:GetPublicKey permission can download the public key of an asymmetric CMK.
// You can share the public key to allow others to encrypt messages and verify
// signatures outside of AWS KMS. For information about symmetric and asymmetric
// CMKs, see Using Symmetric and Asymmetric CMKs
// (https://docs.aws.amazon.com/kms/latest/developerguide/symmetric-asymmetric.html)
// in the AWS Key Management Service Developer Guide. You do not need to download
// the public key. Instead, you can use the public key within AWS KMS by calling
// the Encrypt (), ReEncrypt (), or Verify () operations with the identifier of an
// asymmetric CMK. When you use the public key within AWS KMS, you benefit from the
// authentication, authorization, and logging that are part of every AWS KMS
// operation. You also reduce of risk of encrypting data that cannot be decrypted.
// These features are not effective outside of AWS KMS. For details, see Special
// Considerations for Downloading Public Keys
// (https://docs.aws.amazon.com/kms/latest/developerguide/download-public-key.html#download-public-key-considerations).
// To help you use the public key safely outside of AWS KMS, GetPublicKey returns
// important information about the public key in the response, including:  <ul>
// <li> <p> <a
// href="https://docs.aws.amazon.com/kms/latest/APIReference/API_GetPublicKey.html#KMS-GetPublicKey-response-CustomerMasterKeySpec">CustomerMasterKeySpec</a>:
// The type of key material in the public key, such as <code>RSA_4096</code> or
// <code>ECC_NIST_P521</code>.</p> </li> <li> <p> <a
// href="https://docs.aws.amazon.com/kms/latest/APIReference/API_GetPublicKey.html#KMS-GetPublicKey-response-KeyUsage">KeyUsage</a>:
// Whether the key is used for encryption or signing.</p> </li> <li> <p> <a
// href="https://docs.aws.amazon.com/kms/latest/APIReference/API_GetPublicKey.html#KMS-GetPublicKey-response-EncryptionAlgorithms">EncryptionAlgorithms</a>
// or <a
// href="https://docs.aws.amazon.com/kms/latest/APIReference/API_GetPublicKey.html#KMS-GetPublicKey-response-SigningAlgorithms">SigningAlgorithms</a>:
// A list of the encryption algorithms or the signing algorithms for the key.</p>
// </li> </ul> <p>Although AWS KMS cannot enforce these restrictions on external
// operations, it is crucial that you use this information to prevent the public
// key from being used improperly. For example, you can prevent a public signing
// key from being used encrypt data, or prevent a public key from being used with
// an encryption algorithm that is not supported by AWS KMS. You can also avoid
// errors, such as using the wrong signing algorithm in a verification
// operation.</p> <p>The CMK that you use for this operation must be in a
// compatible key state. For  details, see How Key State Affects Use of a Customer
// Master Key
// (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html) in the
// AWS Key Management Service Developer Guide.
func (c *Client) GetPublicKey(ctx context.Context, params *GetPublicKeyInput, optFns ...func(*Options)) (*GetPublicKeyOutput, error) {
	if params == nil {
		params = &GetPublicKeyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetPublicKey", params, optFns, addOperationGetPublicKeyMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*GetPublicKeyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetPublicKeyInput struct {

	// Identifies the asymmetric CMK that includes the public key.  <p>To specify a
	// CMK, use its key ID, Amazon Resource Name (ARN), alias name, or alias ARN. When
	// using an alias name, prefix it with <code>"alias/"</code>. To specify a CMK in a
	// different AWS account, you must use the key ARN or alias ARN.</p> <p>For
	// example:</p> <ul> <li> <p>Key ID:
	// <code>1234abcd-12ab-34cd-56ef-1234567890ab</code> </p> </li> <li> <p>Key ARN:
	// <code>arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab</code>
	// </p> </li> <li> <p>Alias name: <code>alias/ExampleAlias</code> </p> </li> <li>
	// <p>Alias ARN: <code>arn:aws:kms:us-east-2:111122223333:alias/ExampleAlias</code>
	// </p> </li> </ul> <p>To get the key ID and key ARN for a CMK, use <a>ListKeys</a>
	// or <a>DescribeKey</a>. To get the alias name and alias ARN, use
	// <a>ListAliases</a>.</p>
	//
	// This member is required.
	KeyId *string

	// A list of grant tokens. For more information, see Grant Tokens
	// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#grant_token)
	// in the AWS Key Management Service Developer Guide.
	GrantTokens []*string
}

type GetPublicKeyOutput struct {

	// The type of the of the public key that was downloaded.
	CustomerMasterKeySpec types.CustomerMasterKeySpec

	// The encryption algorithms that AWS KMS supports for this key. This information
	// is critical. If a public key encrypts data outside of AWS KMS by using an
	// unsupported encryption algorithm, the ciphertext cannot be decrypted. This field
	// appears in the response only when the KeyUsage of the public key is
	// ENCRYPT_DECRYPT.
	EncryptionAlgorithms []types.EncryptionAlgorithmSpec

	// The Amazon Resource Name (key ARN
	// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#key-id-key-ARN))
	// of the asymmetric CMK from which the public key was downloaded.
	KeyId *string

	// The permitted use of the public key. Valid values are ENCRYPT_DECRYPT or
	// SIGN_VERIFY. This information is critical. If a public key with SIGN_VERIFY key
	// usage encrypts data outside of AWS KMS, the ciphertext cannot be decrypted.
	KeyUsage types.KeyUsageType

	// The exported public key. The value is a DER-encoded X.509 public key, also known
	// as SubjectPublicKeyInfo (SPKI), as defined in RFC 5280
	// (https://tools.ietf.org/html/rfc5280). When you use the HTTP API or the AWS CLI,
	// the value is Base64-encoded. Otherwise, it is not Base64-encoded.
	PublicKey []byte

	// The signing algorithms that AWS KMS supports for this key. This field appears in
	// the response only when the KeyUsage of the public key is SIGN_VERIFY.
	SigningAlgorithms []types.SigningAlgorithmSpec

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetPublicKeyMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetPublicKey{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetPublicKey{}, middleware.After)
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
	addOpGetPublicKeyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetPublicKey(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetPublicKey(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kms",
		OperationName: "GetPublicKey",
	}
}
