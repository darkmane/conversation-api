package conversationapi

type Location struct {
	Coordinates      *LatLng        `json:"coordinates,omitempty"`
	FormattedAddress string         `json:"formattedAddress"`
	ZipCode          string         `json:"zipCode"`
	City             string         `json:"city"`
	PostalAddress    *PostalAddress `json:"postalAddress,omitempty"`
	Name             string         `json:"name"`
	PhoneNumber      string         `json:"phoneNumber"`
	Notes            string         `json:"notes"`
}
