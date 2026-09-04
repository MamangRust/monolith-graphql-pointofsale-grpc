package role_test

import (
	"context"
	"testing"

	db "github.com/MamangRust/monolith-graphql-pointofsale-pkg/database/schema"
	role_cache "github.com/MamangRust/monolith-graphql-pointofsale-role/cache"
	role_handler "github.com/MamangRust/monolith-graphql-pointofsale-role/handler"
	"github.com/MamangRust/monolith-graphql-pointofsale-role/repository"
	"github.com/MamangRust/monolith-graphql-pointofsale-role/service"
	"github.com/MamangRust/monolith-graphql-pointofsale-shared/cache"
	"github.com/MamangRust/monolith-graphql-pointofsale-shared/observability"
	tests "github.com/MamangRust/monolith-graphql-pointofsale-test"

	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	pbrole "github.com/MamangRust/monolith-graphql-pointofsale-pb/role"
)

type RoleGapiTestSuite struct {
	tests.BaseTestSuite
	client *tests.RoleClient
}

func (s *RoleGapiTestSuite) SetupSuite() {
	s.BaseTestSuite.SetupSuite()

	// Infrastructure
	cacheMetrics, _ := observability.NewCacheMetrics("test")
	cacheStore := cache.NewCacheStore(s.RedisClient(), s.Log, cacheMetrics)
	queries := db.New(s.DBPool())

	// Role dependencies
	mencache := role_cache.NewMencache(cacheStore)
	repos := repository.NewRepositories(queries)
	svc := service.NewService(&service.Deps{
		Repositories:  repos,
		Logger:        s.Log,
		Mencache:      mencache,
		Observability: s.Obs,
	})

	// Handler
	handler := role_handler.NewHandler(&role_handler.Deps{
		Service: svc,
		Logger:  s.Log,
	})

	// Server
	server := grpc.NewServer()
	pbrole.RegisterRoleQueryServiceServer(server, handler.Role)
	pbrole.RegisterRoleCommandServiceServer(server, handler.RoleCommand)

	addr := s.RegisterServer(server)
	conn := s.GetConnection(addr)

	s.client = tests.NewRoleClient(conn)
}

func (s *RoleGapiTestSuite) TestRoleGapiLifecycle() {
	ctx := context.Background()

	// 1. Create
	createRes, err := s.client.CreateRole(ctx, &pbrole.CreateRoleRequest{
		Name: "Gapi Role",
	})
	s.Require().NoError(err)
	s.Require().NotNil(createRes)
	roleID := createRes.Data.Id

	// 2. FindById
	getRes, err := s.client.FindByIdRole(ctx, &pbrole.FindByIdRoleRequest{RoleId: roleID})
	s.Require().NoError(err)
	s.Equal("Gapi Role", getRes.Data.Name)

	// 3. FindAll
	allRes, err := s.client.FindAllRole(ctx, &pbrole.FindAllRoleRequest{Page: 1, PageSize: 10})
	s.Require().NoError(err)
	s.NotEmpty(allRes.Data)

	// 4. FindByActive
	activeRes, err := s.client.FindByActive(ctx, &pbrole.FindAllRoleRequest{Page: 1, PageSize: 10})
	s.Require().NoError(err)
	s.NotEmpty(activeRes.Data)

	// 5. Update
	updateRes, err := s.client.UpdateRole(ctx, &pbrole.UpdateRoleRequest{
		Id:   roleID,
		Name: "Gapi Role Updated",
	})
	s.Require().NoError(err)
	s.Equal("Gapi Role Updated", updateRes.Data.Name)

	// 6. Trash
	_, err = s.client.TrashedRole(ctx, &pbrole.FindByIdRoleRequest{RoleId: roleID})
	s.Require().NoError(err)

	// 7. FindByTrashed
	trashedRes, err := s.client.FindByTrashed(ctx, &pbrole.FindAllRoleRequest{Page: 1, PageSize: 10})
	s.Require().NoError(err)
	s.NotEmpty(trashedRes.Data)

	// 8. Restore
	_, err = s.client.RestoreRole(ctx, &pbrole.FindByIdRoleRequest{RoleId: roleID})
	s.Require().NoError(err)

	// 9. DeletePermanent
	_, _ = s.client.TrashedRole(ctx, &pbrole.FindByIdRoleRequest{RoleId: roleID})
	_, err = s.client.DeleteRolePermanent(ctx, &pbrole.FindByIdRoleRequest{RoleId: roleID})
	s.Require().NoError(err)

	// 10. RestoreAll
	_, err = s.client.RestoreAllRole(ctx, &emptypb.Empty{})
	s.Require().NoError(err)

	// 11. DeleteAll
	_, err = s.client.DeleteAllRolePermanent(ctx, &emptypb.Empty{})
	s.Require().NoError(err)
}

func TestRoleGapiSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(RoleGapiTestSuite))
}
