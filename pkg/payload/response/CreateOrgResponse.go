package response

import "time"

type CreateOrgResponse struct {
	Name      string
	CreatedBy int
	CreatedAt time.Time
	UpdatedAt time.Time
}
