// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package forecast

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeDatasetImportJobInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the dataset import job.
	//
	// DatasetImportJobArn is a required field
	DatasetImportJobArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeDatasetImportJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDatasetImportJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeDatasetImportJobInput"}

	if s.DatasetImportJobArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("DatasetImportJobArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeDatasetImportJobOutput struct {
	_ struct{} `type:"structure"`

	// When the dataset import job was created.
	CreationTime *time.Time `type:"timestamp"`

	// The size of the dataset in gigabytes (GB) after the import job has finished.
	DataSize *float64 `type:"double"`

	// The location of the training data to import and an AWS Identity and Access
	// Management (IAM) role that Amazon Forecast can assume to access the data.
	//
	// If encryption is used, DataSource includes an AWS Key Management Service
	// (KMS) key.
	DataSource *DataSource `type:"structure"`

	// The Amazon Resource Name (ARN) of the dataset that the training data was
	// imported to.
	DatasetArn *string `type:"string"`

	// The ARN of the dataset import job.
	DatasetImportJobArn *string `type:"string"`

	// The name of the dataset import job.
	DatasetImportJobName *string `min:"1" type:"string"`

	// Statistical information about each field in the input data.
	FieldStatistics map[string]Statistics `type:"map"`

	// The last time that the dataset was modified. The time depends on the status
	// of the job, as follows:
	//
	//    * CREATE_PENDING - The same time as CreationTime.
	//
	//    * CREATE_IN_PROGRESS - The current timestamp.
	//
	//    * ACTIVE or CREATE_FAILED - When the job finished or failed.
	LastModificationTime *time.Time `type:"timestamp"`

	// If an error occurred, an informational message about the error.
	Message *string `type:"string"`

	// The status of the dataset import job. The status is reflected in the status
	// of the dataset. For example, when the import job status is CREATE_IN_PROGRESS,
	// the status of the dataset is UPDATE_IN_PROGRESS. States include:
	//
	//    * ACTIVE
	//
	//    * CREATE_PENDING, CREATE_IN_PROGRESS, CREATE_FAILED
	//
	//    * DELETE_PENDING, DELETE_IN_PROGRESS, DELETE_FAILED
	Status *string `type:"string"`

	// The format of timestamps in the dataset. The format that you specify depends
	// on the DataFrequency specified when the dataset was created. The following
	// formats are supported
	//
	//    * "yyyy-MM-dd" For the following data frequencies: Y, M, W, and D
	//
	//    * "yyyy-MM-dd HH:mm:ss" For the following data frequencies: H, 30min,
	//    15min, and 1min; and optionally, for: Y, M, W, and D
	TimestampFormat *string `type:"string"`
}

// String returns the string representation
func (s DescribeDatasetImportJobOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeDatasetImportJob = "DescribeDatasetImportJob"

// DescribeDatasetImportJobRequest returns a request value for making API operation for
// Amazon Forecast Service.
//
// Describes a dataset import job created using the CreateDatasetImportJob operation.
//
// In addition to listing the parameters provided in the CreateDatasetImportJob
// request, this operation includes the following properties:
//
//    * CreationTime
//
//    * LastModificationTime
//
//    * DataSize
//
//    * FieldStatistics
//
//    * Status
//
//    * Message - If an error occurred, information about the error.
//
//    // Example sending a request using DescribeDatasetImportJobRequest.
//    req := client.DescribeDatasetImportJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/forecast-2018-06-26/DescribeDatasetImportJob
func (c *Client) DescribeDatasetImportJobRequest(input *DescribeDatasetImportJobInput) DescribeDatasetImportJobRequest {
	op := &aws.Operation{
		Name:       opDescribeDatasetImportJob,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeDatasetImportJobInput{}
	}

	req := c.newRequest(op, input, &DescribeDatasetImportJobOutput{})
	return DescribeDatasetImportJobRequest{Request: req, Input: input, Copy: c.DescribeDatasetImportJobRequest}
}

// DescribeDatasetImportJobRequest is the request type for the
// DescribeDatasetImportJob API operation.
type DescribeDatasetImportJobRequest struct {
	*aws.Request
	Input *DescribeDatasetImportJobInput
	Copy  func(*DescribeDatasetImportJobInput) DescribeDatasetImportJobRequest
}

// Send marshals and sends the DescribeDatasetImportJob API request.
func (r DescribeDatasetImportJobRequest) Send(ctx context.Context) (*DescribeDatasetImportJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeDatasetImportJobResponse{
		DescribeDatasetImportJobOutput: r.Request.Data.(*DescribeDatasetImportJobOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeDatasetImportJobResponse is the response type for the
// DescribeDatasetImportJob API operation.
type DescribeDatasetImportJobResponse struct {
	*DescribeDatasetImportJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeDatasetImportJob request.
func (r *DescribeDatasetImportJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}