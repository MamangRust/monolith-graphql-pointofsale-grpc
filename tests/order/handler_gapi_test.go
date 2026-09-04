package order_test

import (
	"context"
	"testing"

	order_cache "github.com/MamangRust/monolith-graphql-pointofsale-order/cache"
	order_handler "github.com/MamangRust/monolith-graphql-pointofsale-order/handler"
	order_repo "github.com/MamangRust/monolith-graphql-pointofsale-order/repository"
	order_service "github.com/MamangRust/monolith-graphql-pointofsale-order/service"
	db "github.com/MamangRust/monolith-graphql-pointofsale-pkg/database/schema"
	"github.com/MamangRust/monolith-graphql-pointofsale-shared/cache"
	"github.com/MamangRust/monolith-graphql-pointofsale-shared/observability"
	tests "github.com/MamangRust/monolith-graphql-pointofsale-test"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	pborder "github.com/MamangRust/monolith-graphql-pointofsale-pb/order"
	pb "github.com/MamangRust/monolith-graphql-pointofsale-pb"
)

type OrderGapiTestSuite struct {
	tests.BaseTestSuite
	client *tests.OrderClient
}

func (s *OrderGapiTestSuite) SetupSuite() {
	s.BaseTestSuite.SetupSuite()

	// Setup dependencies
	s.SetupUserService()
	s.SetupCategoryService()
	s.SetupMerchantService()
	s.SetupProductService()
	s.SetupOrderItemService()
	s.SetupCashierService()
	s.SetupTransactionService()

	// Infrastructure
	cacheMetrics, _ := observability.NewCacheMetrics("test")
	cacheStore := cache.NewCacheStore(s.RedisClient(), s.Log, cacheMetrics)
	queries := db.New(s.DBPool())

	// Order dependencies
	mencache := order_cache.NewMencache(cacheStore)
	repos := order_repo.NewRepositories(
		queries,
		tests.NewCashierClient(s.Conns["cashier"]),
		tests.NewMerchantClient(s.Conns["merchant"]),
		tests.NewProductClient(s.Conns["product"]),
		pb.NewOrderItemServiceClient(s.Conns["order-item"]),
	)
	svc := order_service.NewService(&order_service.Deps{
		Mencache:      mencache,
		Repositories:  repos,
		Logger:        s.Log,
		Observability: s.Obs,
	})

	// Handler
	handler := order_handler.NewHandler(&order_handler.Deps{
		Service: svc,
		Logger:  s.Log,
	})

	// Server
	server := grpc.NewServer()
	pborder.RegisterOrderQueryServiceServer(server, handler.Order)
	pborder.RegisterOrderCommandServiceServer(server, handler.OrderCommand)
	pborder.RegisterOrderStatsServiceServer(server, handler.OrderStats)

	addr := s.RegisterServer(server)
	conn := s.GetConnection(addr)

	s.client = tests.NewOrderClient(conn)
}

func (s *OrderGapiTestSuite) TestOrderGapiLifecycle() {
	ctx := context.Background()

	// 1. Seed dependencies
	userID := s.SeedUser(ctx)
	catID := s.SeedCategory(ctx)
	merchID := s.SeedMerchant(ctx, userID)
	prodID := s.SeedProduct(ctx, merchID, catID)

	// Seed a cashier (orders.cashier_id references cashiers.cashier_id)
	var cashierID int
	err := s.DBPool().QueryRow(ctx,
		`INSERT INTO cashiers (merchant_id, user_id, name) VALUES ($1, $2, 'Order Gapi Cashier') RETURNING cashier_id`,
		merchID, userID,
	).Scan(&cashierID)
	s.Require().NoError(err)

	// 2. Create
	createRes, err := s.client.Create(ctx, &pborder.CreateOrderRequest{
		MerchantId: int32(merchID),
		CashierId:  int32(cashierID),
		Items: []*pborder.CreateOrderItemRequest{
			{
				ProductId: int32(prodID),
				Quantity:  1,
			},
		},
	})
	s.Require().NoError(err)
	s.Require().NotNil(createRes)
	orderID := createRes.Data.Id

	// 3. FindById
	getRes, err := s.client.FindById(ctx, &pborder.FindByIdOrderRequest{Id: orderID})
	s.Require().NoError(err)
	s.Equal(int32(userID), getRes.Data.CashierId)

	// 4. FindAll
	allRes, err := s.client.FindAll(ctx, &pborder.FindAllOrderRequest{Page: 1, PageSize: 10})
	s.Require().NoError(err)
	s.NotEmpty(allRes.Data)

	// 5. FindByActive
	activeRes, err := s.client.FindByActive(ctx, &pborder.FindAllOrderRequest{Page: 1, PageSize: 10})
	s.Require().NoError(err)
	s.NotEmpty(activeRes.Data)

	// 6. Update
	// Fetch order items first
	itemClient := pb.NewOrderItemServiceClient(s.Conns["order-item"])
	itemsRes, err := itemClient.FindOrderItemByOrder(ctx, &pb.FindByIdOrderItemRequest{Id: orderID})
	s.Require().NoError(err)
	s.NotEmpty(itemsRes.Data)
	orderItemID := itemsRes.Data[0].Id

	_, err = s.client.Update(ctx, &pborder.UpdateOrderRequest{
		OrderId: orderID,
		Items: []*pborder.UpdateOrderItemRequest{
			{
				OrderItemId: orderItemID,
				ProductId:   int32(prodID),
				Quantity:    1,
			},
		},
	})
	s.Require().NoError(err)

	// 7. Trash
	_, err = s.client.TrashedOrder(ctx, &pborder.FindByIdOrderRequest{Id: orderID})
	s.Require().NoError(err)

	// 8. FindByTrashed
	trashedRes, err := s.client.FindByTrashed(ctx, &pborder.FindAllOrderRequest{Page: 1, PageSize: 10})
	s.Require().NoError(err)
	s.NotEmpty(trashedRes.Data)

	// 9. Restore
	_, err = s.client.RestoreOrder(ctx, &pborder.FindByIdOrderRequest{Id: orderID})
	s.Require().NoError(err)

	// 10. DeletePermanent
	_, _ = s.client.TrashedOrder(ctx, &pborder.FindByIdOrderRequest{Id: orderID})
	_, err = s.client.DeleteOrderPermanent(ctx, &pborder.FindByIdOrderRequest{Id: orderID})
	s.Require().NoError(err)

	// 11. RestoreAll — no trashed orders remain after permanent delete, expect error.
	_, err = s.client.RestoreAllOrder(ctx, &emptypb.Empty{})
	s.Require().Error(err)

	// 12. DeleteAll — no trashed orders remain, expect error.
	_, err = s.client.DeleteAllOrderPermanent(ctx, &emptypb.Empty{})
	s.Require().Error(err)
}

func TestOrderGapiSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(OrderGapiTestSuite))
}
