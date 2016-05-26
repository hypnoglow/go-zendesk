package zendesk

import (
	"fmt"
	"time"
)

// Organization represents a Zendesk organization.
//
// Zendesk Core API docs: https://developer.zendesk.com/rest_api/docs/core/organizations
type Organization struct {
	ID                 *int64                 `json:"id,omitempty"`
	URL                *string                `json:"url,omitempty"`
	ExternalID         *string                `json:"external_id,omitempty"`
	Name               *string                `json:"name,omitempty"`
	CreatedAt          *time.Time             `json:"created_at,omitempty"`
	UpdatedAt          *time.Time             `json:"updated_at,omitempty"`
	DomainNames        *[]string              `json:"domain_names,omitempty"`
	Details            *string                `json:"details,omitempty"`
	Notes              *string                `json:"notes,omitempty"`
	GroupID            *int64                 `json:"group_id,omitempty"`
	SharedTickets      *bool                  `json:"shared_tickets,omitempty"`
	SharedComments     *bool                  `json:"shared_comments,omitempty"`
	OrganizationFields map[string]interface{} `json:"organization_fields,omitempty"`
}

// ShowOrganization fetches an organization by its ID.
//
// Zendesk Core API docs: https://developer.zendesk.com/rest_api/docs/core/organizations#show-organization
func (c *client) ShowOrganization(id int64) (*Organization, error) {
	out := new(APIPayload)
	err := c.get(fmt.Sprintf("/api/v2/organizations/%d.json", id), out)
	return out.Organization, err
}

// CreateOrganization creates an organization.
//
// Zendesk Core API docs: https://developer.zendesk.com/rest_api/docs/core/organizations#create-organization
func (c *client) CreateOrganization(org *Organization) (*Organization, error) {
	in := &APIPayload{Organization: org}
	out := new(APIPayload)
	err := c.post("/api/v2/organizations.json", in, out)
	return out.Organization, err
}
