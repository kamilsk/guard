package types

import (
	"time"

	domain "go.octolab.org/ecosystem/guard/internal/service/types"
)

// Action TODO issue#docs
type Action string

// Create, Update, Delete defines Action set.
const (
	Create  Action = "create"
	Update  Action = "update"
	Delete  Action = "delete"
	Restore Action = "restore"
)

// LicenseAudit TODO issue#docs
type LicenseAudit struct {
	ID        uint64          `db:"id"`
	LicenseID domain.ID       `db:"license_id"`
	Contract  domain.Contract `db:"contract"`
	What      Action          `db:"what"`
	Who       domain.ID       `db:"who"`
	When      time.Time       `db:"when"`
	With      domain.Token    `db:"with"`
}
