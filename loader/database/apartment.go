package database

import (
	"strings"
)

type Apartment struct {
	Found                  bool      `json:"found,omitempty" db:"found,omitempty"`
	ListingId              uint32    `json:"listingId,omitempty" db:"listingId,omitempty"`
	ListingAgencyReference string    `json:"listingAgencyReference,omitempty" db:"listingAgencyReference,omitempty"`
	IsSoldProperty         bool      `json:"isSoldProperty,omitempty" db:"isSoldProperty,omitempty"`
	Region                 string    `json:"region,omitempty" db:"region,omitempty"`
	CityName               string    `json:"cityName,omitempty" db:"cityName,omitempty"`
	Lon                    float64   `json:"lon,omitempty" db:"lon,omitempty"`
	Lat                    float64   `json:"lat,omitempty" db:"lat,omitempty"`
	Price                  int       `json:"price,omitempty" db:"price,omitempty"`
	ChargesPrice           int       `json:"chargesPrice,omitempty" db:"chargesPrice,omitempty"`
	Caution                float32   `json:"caution,omitempty" db:"caution,omitempty"`
	AgencyFee              string    `json:"agency_fee,omitempty" db:"agency_fee,omitempty"`
	PropertySubType        string    `json:"propertySubType,omitempty" db:"propertySubType,omitempty"`
	PublisherId            int       `json:"publisher_id,omitempty" db:"publisher_id,omitempty"`
	PublisherRemoteVisit   bool      `json:"publisher_remote_visit,omitempty" db:"publisher_remote_visit,omitempty"`
	PublisherPhone         string    `json:"publisher_phone,omitempty" db:"publisher_phone,omitempty"`
	PublisherName          string    `json:"publisher_name,omitempty" db:"publisher_name,omitempty"`
	PublisherAthomeId      string    `json:"publisher_athome_id,omitempty" db:"publisher_athome_id,omitempty"`
	PropertySurface        float64   `json:"propertySurface,omitempty" db:"propertySurface,omitempty"`
	BuildingYear           string    `json:"buildingYear,omitempty" db:"buildingYear,omitempty"`
	FloorNumber            string    `json:"floorNumber,omitempty" db:"floorNumber,omitempty"`
	BathroomsCount         int       `json:"bathroomsCount,omitempty" db:"bathroomsCount,omitempty"`
	BedroomsCount          int       `json:"bedroomsCount,omitempty" db:"bedroomsCount,omitempty"`
	BalconiesCount         int       `json:"balconiesCount,omitempty" db:"balconiesCount,omitempty"`
	CarparkCount           int       `json:"carparkCount,omitempty" db:"carparkCount,omitempty"`
	GaragesCount           int       `json:"garagesCount,omitempty" db:"garagesCount,omitempty"`
	HasLivingRoom          bool      `json:"hasLivingRoom,omitempty" db:"hasLivingRoom,omitempty"`
	HasKitchen             bool      `json:"hasKitchen,omitempty" db:"hasKitchen,omitempty"`
	Availability           string    `json:"availability,omitempty" db:"availability,omitempty"`
	Media                  *[]string `json:"media,omitempty" db:"media,omitempty"`
	Description            string    `json:"description,omitempty" db:"description,omitempty"`
	Link                   string    `json:"link,omitempty" db:"link,omitempty"`
	CreatedAt              string    `json:"createdAt,omitempty" db:"createdAt,omitempty"`
	UpdatedAt              string    `json:"updatedAt,omitempty" db:"updatedAt,omitempty"`
}

func (a *Apartment) MediaConcatString() string {
	var s = strings.Builder{}

	s.Write([]byte("["))
	if a.Media != nil && len(*a.Media) > 0 {
		s.WriteString(strings.Join(*a.Media, ","))
	}
	s.Write([]byte("]"))

	return s.String()
}
