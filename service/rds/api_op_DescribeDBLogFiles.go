// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of DB log files for the DB instance.
func (c *Client) DescribeDBLogFiles(ctx context.Context, params *DescribeDBLogFilesInput, optFns ...func(*Options)) (*DescribeDBLogFilesOutput, error) {
	if params == nil {
		params = &DescribeDBLogFilesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDBLogFiles", params, optFns, addOperationDescribeDBLogFilesMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDBLogFilesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DescribeDBLogFilesInput struct {

	// The customer-assigned name of the DB instance that contains the log files you
	// want to list. Constraints:
	//
	//     * Must match the identifier of an existing
	// DBInstance.
	//
	// This member is required.
	DBInstanceIdentifier *string

	// Filters the available log files for files written since the specified date, in
	// POSIX timestamp format with milliseconds.
	FileLastWritten *int64

	// Filters the available log files for files larger than the specified size.
	FileSize *int64

	// Filters the available log files for log file names that contain the specified
	// string.
	FilenameContains *string

	// This parameter isn't currently supported.
	Filters []*types.Filter

	// The pagination token provided in the previous request. If this parameter is
	// specified the response includes only records beyond the marker, up to
	// MaxRecords.
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so you can retrieve the remaining results.
	MaxRecords *int32
}

// The response from a call to DescribeDBLogFiles.
type DescribeDBLogFilesOutput struct {

	// The DB log files returned.
	DescribeDBLogFiles []*types.DescribeDBLogFilesDetails

	// A pagination token that can be used in a later DescribeDBLogFiles request.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeDBLogFilesMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeDBLogFiles{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeDBLogFiles{}, middleware.After)
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
	addOpDescribeDBLogFilesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDBLogFiles(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeDBLogFiles(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DescribeDBLogFiles",
	}
}
