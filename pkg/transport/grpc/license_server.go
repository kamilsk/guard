package grpc

import (
	"context"

	domain "github.com/kamilsk/guard/pkg/service/types"

	"github.com/kamilsk/guard/pkg/storage/query"
	"github.com/kamilsk/guard/pkg/transport/grpc/middleware"
	"github.com/kamilsk/guard/pkg/transport/grpc/protobuf"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// NewLicenseServer TODO issue#docs
func NewLicenseServer(storage ProtectedStorage) protobuf.LicenseServer {
	return &licenseServer{storage}
}

type licenseServer struct {
	storage ProtectedStorage
}

// Register TODO issue#docs
func (server *licenseServer) Register(ctx context.Context, req *protobuf.RegisterLicenseRequest) (*protobuf.RegisterLicenseResponse, error) {
	token, authErr := middleware.TokenExtractor(ctx)
	if authErr != nil {
		return nil, authErr
	}
	license, registerErr := server.storage.RegisterLicense(ctx, token, query.RegisterLicense{
		ID:       domain.ID(req.Id),
		Contract: convertToDomainContract(req.Contract),
	})
	if registerErr != nil {
		return nil, status.Errorf(codes.Internal, "something happen: %v", registerErr) // TODO issue#6
	}
	return &protobuf.RegisterLicenseResponse{Id: license.ID.String()}, nil
}

// Create TODO issue#docs
func (server *licenseServer) Create(ctx context.Context, req *protobuf.CreateLicenseRequest) (*protobuf.CreateLicenseResponse, error) {
	token, authErr := middleware.TokenExtractor(ctx)
	if authErr != nil {
		return nil, authErr
	}
	license, createErr := server.storage.CreateLicense(ctx, token, query.CreateLicense{
		ID:       ptrToID(req.Id),
		Contract: convertToDomainContract(req.Contract),
	})
	if createErr != nil {
		return nil, status.Errorf(codes.Internal, "something happen: %v", createErr) // TODO issue#6
	}
	return &protobuf.CreateLicenseResponse{Id: license.ID.String(), CreatedAt: Timestamp(&license.CreatedAt)}, nil
}

// Read TODO issue#docs
func (server *licenseServer) Read(ctx context.Context, req *protobuf.ReadLicenseRequest) (*protobuf.ReadLicenseResponse, error) {
	token, authErr := middleware.TokenExtractor(ctx)
	if authErr != nil {
		return nil, authErr
	}
	license, readErr := server.storage.ReadLicense(ctx, token, query.ReadLicense{ID: domain.ID(req.Id)})
	if readErr != nil {
		return nil, status.Errorf(codes.Internal, "something happen: %v", readErr) // TODO issue#6
	}
	return &protobuf.ReadLicenseResponse{
		Id:        license.ID.String(),
		Contract:  convertFromDomainContract(license.Contract),
		CreatedAt: Timestamp(&license.CreatedAt),
		UpdatedAt: Timestamp(license.UpdatedAt),
		DeletedAt: Timestamp(license.DeletedAt),
	}, nil
}

// Update TODO issue#docs
func (server *licenseServer) Update(ctx context.Context, req *protobuf.UpdateLicenseRequest) (*protobuf.UpdateLicenseResponse, error) {
	token, authErr := middleware.TokenExtractor(ctx)
	if authErr != nil {
		return nil, authErr
	}
	license, updateErr := server.storage.UpdateLicense(ctx, token, query.UpdateLicense{
		ID:       domain.ID(req.Id),
		Contract: convertToDomainContract(req.Contract),
	})
	if updateErr != nil {
		return nil, status.Errorf(codes.Internal, "something happen: %v", updateErr) // TODO issue#6
	}
	return &protobuf.UpdateLicenseResponse{Id: license.ID.String(), UpdatedAt: Timestamp(license.UpdatedAt)}, nil
}

// Delete TODO issue#docs
func (server *licenseServer) Delete(ctx context.Context, req *protobuf.DeleteLicenseRequest) (*protobuf.DeleteLicenseResponse, error) {
	token, authErr := middleware.TokenExtractor(ctx)
	if authErr != nil {
		return nil, authErr
	}
	license, deleteErr := server.storage.DeleteLicense(ctx, token, query.DeleteLicense{ID: domain.ID(req.Id)})
	if deleteErr != nil {
		return nil, status.Errorf(codes.Internal, "something happen: %v", deleteErr) // TODO issue#6
	}
	return &protobuf.DeleteLicenseResponse{Id: license.ID.String(), DeletedAt: Timestamp(license.DeletedAt)}, nil
}

// Restore TODO issue#docs
func (server *licenseServer) Restore(ctx context.Context, req *protobuf.RestoreLicenseRequest) (*protobuf.RestoreLicenseResponse, error) {
	token, authErr := middleware.TokenExtractor(ctx)
	if authErr != nil {
		return nil, authErr
	}
	license, restoreErr := server.storage.RestoreLicense(ctx, token, query.RestoreLicense{ID: domain.ID(req.Id)})
	if restoreErr != nil {
		return nil, status.Errorf(codes.Internal, "something happen: %v", restoreErr) // TODO issue#6
	}
	return &protobuf.RestoreLicenseResponse{Id: license.ID.String(), UpdatedAt: Timestamp(license.UpdatedAt)}, nil
}

// TODO issue#draft {

// AddEmployee TODO issue#docs
func (server *licenseServer) AddEmployee(ctx context.Context, req *protobuf.AddEmployeeRequest) (*protobuf.EmptyResponse, error) {
	token, authErr := middleware.TokenExtractor(ctx)
	if authErr != nil {
		return nil, authErr
	}
	if addErr := server.storage.AddEmployee(ctx, token, query.LicenseEmployee{
		ID:       domain.ID(req.Id),
		Employee: domain.ID(req.Employee),
	}); addErr != nil {
		return nil, status.Errorf(codes.Internal, "something happen: %v", addErr) // TODO issue#6
	}
	return new(protobuf.EmptyResponse), nil
}

// DeleteEmployee TODO issue#docs
func (server *licenseServer) DeleteEmployee(ctx context.Context, req *protobuf.DeleteEmployeeRequest) (*protobuf.EmptyResponse, error) {
	token, authErr := middleware.TokenExtractor(ctx)
	if authErr != nil {
		return nil, authErr
	}
	if deleteErr := server.storage.DeleteEmployee(ctx, token, query.LicenseEmployee{
		ID:       domain.ID(req.Id),
		Employee: domain.ID(req.Employee),
	}); deleteErr != nil {
		return nil, status.Errorf(codes.Internal, "something happen: %v", deleteErr) // TODO issue#6
	}
	return new(protobuf.EmptyResponse), nil
}

// AddWorkplace TODO issue#docs
func (server *licenseServer) AddWorkplace(ctx context.Context, req *protobuf.AddWorkplaceRequest) (*protobuf.EmptyResponse, error) {
	token, authErr := middleware.TokenExtractor(ctx)
	if authErr != nil {
		return nil, authErr
	}
	if addErr := server.storage.AddWorkplace(ctx, token, query.LicenseWorkplace{
		ID:        domain.ID(req.Id),
		Workplace: domain.ID(req.Workplace),
	}); addErr != nil {
		return nil, status.Errorf(codes.Internal, "something happen: %v", addErr) // TODO issue#6
	}
	return new(protobuf.EmptyResponse), nil
}

// DeleteWorkplace TODO issue#docs
func (server *licenseServer) DeleteWorkplace(ctx context.Context, req *protobuf.DeleteWorkplaceRequest) (*protobuf.EmptyResponse, error) {
	token, authErr := middleware.TokenExtractor(ctx)
	if authErr != nil {
		return nil, authErr
	}
	if deleteErr := server.storage.DeleteWorkplace(ctx, token, query.LicenseWorkplace{
		ID:        domain.ID(req.Id),
		Workplace: domain.ID(req.Workplace),
	}); deleteErr != nil {
		return nil, status.Errorf(codes.Internal, "something happen: %v", deleteErr) // TODO issue#6
	}
	return new(protobuf.EmptyResponse), nil
}

// PushWorkplace TODO issue#docs
func (server *licenseServer) PushWorkplace(ctx context.Context, req *protobuf.PushWorkplaceRequest) (*protobuf.EmptyResponse, error) {
	token, authErr := middleware.TokenExtractor(ctx)
	if authErr != nil {
		return nil, authErr
	}
	if pushErr := server.storage.PushWorkplace(ctx, token, query.LicenseWorkplace{
		ID:        domain.ID(req.Id),
		Workplace: domain.ID(req.Workplace),
	}); pushErr != nil {
		return nil, status.Errorf(codes.Internal, "something happen: %v", pushErr) // TODO issue#6
	}
	return new(protobuf.EmptyResponse), nil
}

// issue#draft }
