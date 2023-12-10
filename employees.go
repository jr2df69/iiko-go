package iiko

import "github.com/google/uuid"

type OpenPersonalSessionRequest struct {
	OrganizationID  string `json:"organizationId"`
	TerminalGroupID string `json:"terminalGroupId"`
	EmployeeID      string `json:"employeeId"`
}

type OpenPersonalSessionResponse struct {
	CorrelationID uuid.UUID `json:"correlationId"`
	Error         string    `json:"error"`
}

type ClosePersonalSessionRequest struct {
	OrganizationID  string `json:"organizationId"`
	TerminalGroupID string `json:"terminalGroupId"`
	EmployeeID      string `json:"employeeId"`
}

type ClosePersonalSessionResponse struct {
	CorrelationID uuid.UUID `json:"correlationId"`
	Error         string    `json:"error"`
}

type CheckPersonalSessionOpenedRequest struct {
	OrganizationID  string `json:"organizationId"`
	TerminalGroupID string `json:"terminalGroupId"`
	EmployeeID      string `json:"employeeId"`
}

type CheckPersonalSessionOpenedResponse struct {
	CorrelationID   uuid.UUID `json:"correlationId"`
	IsSessionOpened bool      `json:"isSessionOpened"`
	Error           string    `json:"error"`
}

func (c *Client) OpenPersonalSession(req *OpenPersonalSessionRequest, opts ...Option) (*OpenPersonalSessionResponse, error) {
	var response OpenPersonalSessionResponse

	if err := c.post(true, "/api/1/employees/shift/clockin", req, &response, opts...); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) ClosePersonalSession(req *ClosePersonalSessionRequest, opts ...Option) (*ClosePersonalSessionResponse, error) {
	var response ClosePersonalSessionResponse

	if err := c.post(true, "/api/1/employees/shift/clockout", req, &response, opts...); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) CheckIfPersonalSessionOpened(req *CheckPersonalSessionOpenedRequest, opts ...Option) (*CheckPersonalSessionOpenedResponse, error) {
	var response CheckPersonalSessionOpenedResponse

	if err := c.post(true, "/api/1/employees/shift/is_open", req, &response, opts...); err != nil {
		return nil, err
	}

	return &response, nil
}
