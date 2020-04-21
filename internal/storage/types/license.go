package types

import (
	"time"

	domain "go.octolab.org/ecosystem/guard/internal/service/types"
)

// License TODO issue#docs
type License struct {
	ID        domain.ID       `db:"id"`
	AccountID domain.ID       `db:"account_id"`
	Contract  domain.Contract `db:"contract"`
	CreatedAt time.Time       `db:"created_at"`
	UpdatedAt *time.Time      `db:"updated_at"`
	DeletedAt *time.Time      `db:"deleted_at"`
	Account   *Account        `db:"-"`
}

// Employee TODO issue#docs
type Employee struct {
	ID        domain.ID `db:"employee"`
	License   domain.ID `db:"license"`
	CreatedAt time.Time `db:"created_at"`
}

// Workplace TODO issue#docs
type Workplace struct {
	ID        domain.ID  `db:"workplace"`
	License   domain.ID  `db:"license"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
}
