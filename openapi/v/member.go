package v

// swagger:parameters addMembersRequest
type AddMembersRequest struct {
	// in: header
	// name: organizationId
	// required: true
	OrganizationId int64 `json:"organizationId"`

	// in: body
	Body []OrganizatoinMember
}

// The request payload which is information about the members of an organization.
//
// swagger:model organizatoinMember
type OrganizatoinMember struct {
	// Name is name of the member.
	//
	// required: true
	Name string `json:"name"`

	// Mail is email of the member.
	//
	// required: true
	// swagger:strfmt email
	Mail string `json:"mail"`
}

// The response payload sent when adding memebers request.
//
// swagger:model addMembersResponse
type AddMembersResponse struct {
	// Data is data of the response of adding member.
	//
	// required: true
	Data AddMembersResponseData `json:"data"`
}

type AddMembersResponseData struct {
	Created                  []string `json:"created,omitempty"`
	EmailAlreadyExists       []string `json:"email_already_exists,omitempty"`
	PhoneMobileAlreadyExists []string `json:"phone_mobile_already_exists,omitempty"`
}
