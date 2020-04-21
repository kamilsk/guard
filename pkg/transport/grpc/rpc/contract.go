package rpc

import (
	"context"

	domain "github.com/kamilsk/guard/pkg/service/types"
	"github.com/kamilsk/guard/pkg/service/types/request"
	"github.com/kamilsk/guard/pkg/service/types/response"
	"github.com/kamilsk/guard/pkg/storage/query"
	repository "github.com/kamilsk/guard/pkg/storage/types"
)

// Maintenance TODO issue#docs
type Maintenance interface {
	// Install TODO issue#docs
	Install(context.Context, request.Install) response.Install
}

// Storage TODO issue#docs
type Storage interface {
	// RegisterLicense TODO issue#docs
	RegisterLicense(context.Context, domain.Token, query.RegisterLicense) (repository.License, error)
	// CreateLicense TODO issue#docs
	CreateLicense(context.Context, domain.Token, query.CreateLicense) (repository.License, error)
	// ReadLicense TODO issue#docs
	ReadLicense(context.Context, domain.Token, query.ReadLicense) (repository.License, error)
	// UpdateLicense TODO issue#docs
	UpdateLicense(context.Context, domain.Token, query.UpdateLicense) (repository.License, error)
	// DeleteLicense TODO issue#docs
	DeleteLicense(context.Context, domain.Token, query.DeleteLicense) (repository.License, error)
	// RestoreLicense TODO issue#docs
	RestoreLicense(context.Context, domain.Token, query.RestoreLicense) (repository.License, error)

	// TODO issue#draft {

	// AddEmployee TODO issue#docs
	AddEmployee(context.Context, domain.Token, query.LicenseEmployee) error
	// DeleteEmployee TODO issue#docs
	DeleteEmployee(context.Context, domain.Token, query.LicenseEmployee) error
	// LicenseEmployees TODO issue#docs
	LicenseEmployees(context.Context, domain.Token, query.EmployeeList) ([]repository.Employee, error)
	// AddWorkplace TODO issue#docs
	AddWorkplace(context.Context, domain.Token, query.LicenseWorkplace) error
	// DeleteWorkplace TODO issue#docs
	DeleteWorkplace(context.Context, domain.Token, query.LicenseWorkplace) error
	// PushWorkplace TODO issue#docs
	PushWorkplace(context.Context, domain.Token, query.LicenseWorkplace) error
	// LicenseWorkplaces TODO issue#docs
	LicenseWorkplaces(context.Context, domain.Token, query.WorkplaceList) ([]repository.Workplace, error)

	// issue#draft }
}
