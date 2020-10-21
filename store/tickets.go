// Author: dat.nguyen

package store

import "time"

// Ticket -
type Ticket struct {
	ID             int64     `json:"_id"`
	URL            string    `json:"url"`
	ExternalID     int       `json:"external_id"`
	CreateAt       time.Time `json:"created_at"`
	Type           string    `json:"type"`
	Subject        string    `json:"subject"`
	Description    string    `json:"description"`
	Priority       string    `json:"priority"`
	Status         string    `json:"status"`
	SubmitterID    int       `json:"submitter_id"`
	AssigneeID     int       `json:"assignee_id"`
	OrganizationID int       `json:"organization_id"`
	Tags           []string  `json:"tags"`
	HasIncidents   bool      `json:"has_incidents"`
	DueAt          time.Time `json:"due_at"`
	Via            string    `json:"via"`
}
