package model

import "github.com/google/uuid"

// Implements the Unmarshaler interface of the yaml pkg.
func (r *RelatedActivity) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var myuuid uuid.UUID
	err := unmarshal(&myuuid)
	if err != nil {
		return err
	}

	// make sure to dereference before assignment,
	// otherwise only the local variable will be overwritten
	// and not the value the pointer actually points to
	r.ID = myuuid

	return nil
}
