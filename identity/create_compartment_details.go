// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity


type CreateCompartmentDetails struct {

    // The OCID of the tenancy containing the compartment.
    CompartmentId string `json:"compartmentId,omitempty"`

    // The name you assign to the compartment during creation. The name must be unique across all compartments\nin the tenancy and cannot be changed.\n
    Name string `json:"name,omitempty"`

    // The description you assign to the compartment during creation. Does not have to be unique, and it's changeable.\n
    Description string `json:"description,omitempty"`
}