package main

type Whoami struct {
	StatusCode int           `json:"StatusCode,omitempty"`
	ErrorCode  int           `json:"ErrorCode,omitempty"`
	Message    string        `json:"Message,omitempty"`
	Details    WhoamiDetails `json:"Details,omitempty"`
}

type WhoamiDetails struct {
	Id           int    `json:"Id,omitempty"`
	Username     string `json:"Username,omitempty"`
	Type         int    `json:"Type,omitempty"`
	Status       int    `json:"Status,omitempty"`
	Email        string `json:"Email,omitempty"`
	FirstName    string `json:"FirstName,omitempty"`
	LastName     string `json:"LastName,omitempty"`
	Organization string `json:"Organization,omitempty"`
}

type DomainResult struct {
	StatusCode int                `json:"StatusCode,omitempty"`
	ErrorCode  int                `json:"ErrorCode,omitempty"`
	Message    string             `json:"Message,omitempty"`
	Details    DomainResultDetail `json:"Details,omitempty"`
}

type DomainResultDetail struct {
	Code          int
	Message       string
	Id            int
	Name          string
	Username      string
	Password      string
	DomainUser    bool
	ModuleResults []DomainOperationModuleResult
}

type DomainOperationModuleResult struct {
	Status bool
	Msg    string
	Name   string
	PType  string
}
