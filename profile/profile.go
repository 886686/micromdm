package profile

// Profile is an Apple Configuration profile
type Profile struct {
	UUID              string `plist:"-" json:"profile_uuid,omitempty" db:"profile_uuid"`
	PayloadIdentifier string `json:"payload_identifier" db:"identifier"`
	ProfileData       string `json:"profile_data,omitempty" db:"profile_data"`
}
