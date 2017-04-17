/* 
 * Lyft API
 *
 * Drive your app to success with Lyft's API
 *
 * OpenAPI spec version: 1.0.0
 * Contact: api-support@lyft.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package lyft

// Represents a requested, ongoing, or finished Lyft ride
type Ride struct {

	RideType RideTypeEnum `json:"ride_type,omitempty"`

	// The *requested* location for passenger pickup
	Origin interface{} `json:"origin,omitempty"`

	// The *requested* location for passenger drop off
	Destination interface{} `json:"destination,omitempty"`

	// A token that confirms the user has accepted current primetime charges (Deprecated)
	PrimetimeConfirmationToken string `json:"primetime_confirmation_token,omitempty"`

	// A token that confirms the user has accepted current Prime Time and/or fixed price charges
	CostToken string `json:"cost_token,omitempty"`
}
