package udm

import "time"

type Attribute struct {
	Cloud          *Cloud      `json:"cloud,omitempty"`
	CreationTime   *time.Time  `json:"creation_time,omitempty"`
	Labels         []*Label    `json:"labels,omitempty"`
	LastUpdateTime *time.Time  `json:"last_update_time,omitempty"`
	Permissions    *Permission `json:"permissions,omitempty"`
	Roles          []*Role     `json:"roles,omitempty"`
}
