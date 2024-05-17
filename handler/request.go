package handler

import "fmt"

func errParamRequired(param string, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", param, typ)
}

// CreateOpening
type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int    `json:"salary"`
}

func (r *CreateOpeningRequest) Validate() error {
	if r.Role == "" && r.Company == "" && r.Location == "" && r.Remote == nil && r.Link == "" && r.Salary <= 0 {
		return fmt.Errorf("request json body is empty or malfomed")
	}
	if r.Role == "" {
		return errParamRequired("role", "string")
	}
	if r.Company == "" {
		return errParamRequired("company", "string")
	}
	if r.Location == "" {
		return errParamRequired("location", "string")
	}
	if r.Remote == nil {
		return errParamRequired("remote", "bool")
	}
	if r.Link == "" {
		return errParamRequired("link", "string")
	}
	if r.Salary <= 0 {
		return errParamRequired("salary", "int")
	}

	return nil
}

type UpdateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int    `json:"salary"`
}

func (r *UpdateOpeningRequest) Validate() error {
	// If any field is provided, validation is truthy
	if r.Role != "" || r.Company != "" || r.Location != "" || r.Remote != nil || r.Link != "" || r.Salary > 0 {
		return nil
	}
	// If no field is provided, return error
	return fmt.Errorf("at least one valid field must be provided to update an opening")
}
