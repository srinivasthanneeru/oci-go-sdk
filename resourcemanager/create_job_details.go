// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Resource Manager API
//
// API for the Resource Manager service. Use this API to install, configure, and manage resources via the "infrastructure-as-code" model. For more information, see Overview of Resource Manager (https://docs.cloud.oracle.com/iaas/Content/ResourceManager/Concepts/resourcemanager.htm).
//

package resourcemanager

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// CreateJobDetails Defines the requirements and properties of a job to create and run against the specified stack.
type CreateJobDetails struct {

	// OCID of the stack that is associated with the current job.
	StackId *string `mandatory:"true" json:"stackId"`

	// Description of the job.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Terraform-specific operation to execute.
	Operation JobOperationEnum `mandatory:"false" json:"operation,omitempty"`

	// Job details that are specific to the operation type.
	JobOperationDetails CreateJobOperationDetails `mandatory:"false" json:"jobOperationDetails"`

	// Deprecated. Use the property `executionPlanStrategy` in `jobOperationDetails` instead.
	ApplyJobPlanResolution *ApplyJobPlanResolution `mandatory:"false" json:"applyJobPlanResolution"`

	// Free-form tags associated with this resource. Each tag is a key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateJobDetails) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *CreateJobDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName            *string                           `json:"displayName"`
		Operation              JobOperationEnum                  `json:"operation"`
		JobOperationDetails    createjoboperationdetails         `json:"jobOperationDetails"`
		ApplyJobPlanResolution *ApplyJobPlanResolution           `json:"applyJobPlanResolution"`
		FreeformTags           map[string]string                 `json:"freeformTags"`
		DefinedTags            map[string]map[string]interface{} `json:"definedTags"`
		StackId                *string                           `json:"stackId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	m.DisplayName = model.DisplayName
	m.Operation = model.Operation
	nn, e := model.JobOperationDetails.UnmarshalPolymorphicJSON(model.JobOperationDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.JobOperationDetails = nn.(CreateJobOperationDetails)
	} else {
		m.JobOperationDetails = nil
	}
	m.ApplyJobPlanResolution = model.ApplyJobPlanResolution
	m.FreeformTags = model.FreeformTags
	m.DefinedTags = model.DefinedTags
	m.StackId = model.StackId
	return
}
