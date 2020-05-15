// Package classification Global Region.
//
// Documentation of Global Region API.
//
//    Schemes: http
//    BasePath: /api
//    Version: 1.0.0
//    Host: 127.0.0.1:9000
//
//    Consumes:
//    - application/json
//
//    Produces:
//    - application/json
//
//    Security:
//    - oauth2
//
//    securityDefinitions:
//      OauthSecurity:
//    	type: oauth2
//    	flow: accessCode
//    	authorizationUrl: 'http://sso-dex:5556/auth'
//    	tokenUrl: 'http://sso-dex:5556/token'
//    	scopes:
//    	  user: openid email profile offline_access
//
// swagger:meta
package main

// swagger:route GET /regions region allRegions
// query all regions from the server.
// responses:
//   200: regionsResponse

// This is regions response body.
// swagger:response regionsResponse
type regionsResponseWrapper struct {
	// in:body
	Body struct {
		Regions []regionInfo `json:"regions"`
	}
}

// swagger:parameters allRegions
type regionsParamsWrapper struct {
	// No Parameters in this endpoint.
} // swagger:response regionsResponse

// swagger:route GET /user/{subject}/regions region userRegions
// query regions with the user's subject from the server.
// responses:
//   200: userRegionsResponse

// This is user's regions response body.
// swagger:response userRegionsResponse
type userRegionsResponseWrapper struct {
	// in:body
	Body struct {
		Regions []regionInfo `json:"regions"`
	}
}

// swagger:parameters userRegions
type userRegionsParamsWrapper struct {
	// subject is required, it is unique id of the user.
	// in:path
	Subject string `json:"subject"`
} // swagger:response userRegionsResponse

// swagger:route GET /user/{subject}/organizations region userOrganizations
// query organizations of the user's subject in regions from the server.
// responses:
//   200: userOrganizationsResponse

// This is user's organization in regions response body.
// swagger:response userOrganizationsResponse
type userOrganizationsResponseWrapper struct {
	// in:body
	Body struct {
		Regions []regionWithOrganizationInfo `json:"regions"`
	}
}

// swagger:parameters userOrganizations
type userOrganizationsParamsWrapper struct {
	// subject is required, it is unique id of the user.
	// in:path
	Subject string `json:"subject"`
} // swagger:response userOrganizationsResponse
