// Code generated by smithy-go-codegen DO NOT EDIT.

package forecast

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/forecast/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an Amazon Forecast dataset. The information about the dataset that you
// provide helps Forecast understand how to consume the data for model training.
// This includes the following:
//
//     * DataFrequency - How frequently your
// historical time-series data is collected.
//
//     * Domain and DatasetType - Each
// dataset has an associated dataset domain and a type within the domain. Amazon
// Forecast provides a list of predefined domains and types within each domain. For
// each unique dataset domain and type within the domain, Amazon Forecast requires
// your data to include a minimum set of predefined fields.
//
//     * Schema - A
// schema specifies the fields in the dataset, including the field name and data
// type.
//
// After creating a dataset, you import your training data into it and add
// the dataset to a dataset group. You use the dataset group to create a predictor.
// For more information, see howitworks-datasets-groups (). To get a list of all
// your datasets, use the ListDatasets () operation. For example Forecast datasets,
// see the Amazon Forecast Sample GitHub repository
// (https://github.com/aws-samples/amazon-forecast-samples). The Status of a
// dataset must be ACTIVE before you can import training data. Use the
// DescribeDataset () operation to get the status.
func (c *Client) CreateDataset(ctx context.Context, params *CreateDatasetInput, optFns ...func(*Options)) (*CreateDatasetOutput, error) {
	if params == nil {
		params = &CreateDatasetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDataset", params, optFns, addOperationCreateDatasetMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDatasetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDatasetInput struct {

	// A name for the dataset.
	//
	// This member is required.
	DatasetName *string

	// The dataset type. Valid values depend on the chosen Domain.
	//
	// This member is required.
	DatasetType types.DatasetType

	// The domain associated with the dataset. When you add a dataset to a dataset
	// group, this value and the value specified for the Domain parameter of the
	// CreateDatasetGroup () operation must match. The Domain and DatasetType that you
	// choose determine the fields that must be present in the training data that you
	// import to the dataset. For example, if you choose the RETAIL domain and
	// TARGET_TIME_SERIES as the DatasetType, Amazon Forecast requires item_id,
	// timestamp, and demand fields to be present in your data. For more information,
	// see howitworks-datasets-groups ().
	//
	// This member is required.
	Domain types.Domain

	// The schema for the dataset. The schema attributes and their order must match the
	// fields in your data. The dataset Domain and DatasetType that you choose
	// determine the minimum required fields in your training data. For information
	// about the required fields for a specific dataset domain and type, see
	// howitworks-domains-ds-types ().
	//
	// This member is required.
	Schema *types.Schema

	// The frequency of data collection. This parameter is required for
	// RELATED_TIME_SERIES datasets. Valid intervals are Y (Year), M (Month), W (Week),
	// D (Day), H (Hour), 30min (30 minutes), 15min (15 minutes), 10min (10 minutes),
	// 5min (5 minutes), and 1min (1 minute). For example, "D" indicates every day and
	// "15min" indicates every 15 minutes.
	DataFrequency *string

	// An AWS Key Management Service (KMS) key and the AWS Identity and Access
	// Management (IAM) role that Amazon Forecast can assume to access the key.
	EncryptionConfig *types.EncryptionConfig

	// The optional metadata that you apply to the dataset to help you categorize and
	// organize them. Each tag consists of a key and an optional value, both of which
	// you define. The following basic restrictions apply to tags:
	//
	//     * Maximum
	// number of tags per resource - 50.
	//
	//     * For each resource, each tag key must be
	// unique, and each tag key can have only one value.
	//
	//     * Maximum key length -
	// 128 Unicode characters in UTF-8.
	//
	//     * Maximum value length - 256 Unicode
	// characters in UTF-8.
	//
	//     * If your tagging schema is used across multiple
	// services and resources, remember that other services may have restrictions on
	// allowed characters. Generally allowed characters are: letters, numbers, and
	// spaces representable in UTF-8, and the following characters: + - = . _ : / @.
	//
	//
	// * Tag keys and values are case sensitive.
	//
	//     * Do not use aws:, AWS:, or any
	// upper or lowercase combination of such as a prefix for keys as it is reserved
	// for AWS use. You cannot edit or delete tag keys with this prefix. Values can
	// have this prefix. If a tag value has aws as its prefix but the key does not,
	// then Forecast considers it to be a user tag and will count against the limit of
	// 50 tags. Tags with only the key prefix of aws do not count against your tags per
	// resource limit.
	Tags []*types.Tag
}

type CreateDatasetOutput struct {

	// The Amazon Resource Name (ARN) of the dataset.
	DatasetArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateDatasetMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateDataset{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateDataset{}, middleware.After)
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
	addOpCreateDatasetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDataset(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateDataset(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "forecast",
		OperationName: "CreateDataset",
	}
}
