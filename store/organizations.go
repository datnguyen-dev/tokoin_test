// Author: dat.nguyen

package store

import "time"

// Organization -
type Organization struct {
	ID            int64     `json:"_id"`
	URL           string    `json:"url"`
	ExternalID    int       `json:"external_id"`
	Name          string    `json:"name"`
	DomainNames   []string  `json:"domain_names"`
	CreateAt      time.Time `json:"created_at"`
	Details       string    `json:"details"`
	SharedTickets bool      `json:"shared_tickets"`
	Tags          []string  `json:"tags"`
}
