package service

import (
	"context"
	"fmt"
	"transaction-service/domain/entity"
	"transaction-service/domain/usecase"

	"github.com/mushoffa/spenmo-proto/protos"
)

// @Author Ahmad Ridwan Mushoffa
// @Created 03/11/2021
// @Updated
type transaction struct {
	protos.UnimplementedTransactionServiceServer
	u      usecase.TransactionUsecase
	user   protos.UserServiceClient
	wallet protos.WalletServiceClient
	card   protos.CardServiceClient
}

// @Author Ahmad Ridwan Mushoffa
// @Created 03/11/2021
// @Updated
func NewTransactionService(u usecase.TransactionUsecase, user protos.UserServiceClient, wallet protos.WalletServiceClient, card protos.CardServiceClient) protos.TransactionServiceServer {
	return &transaction{
		u:      u,
		user:   user,
		wallet: wallet,
		card:   card,
	}
}

// @Author Ahmad Ridwan Mushoffa
// @Created 03/11/2021
// @Updated
func (s *transaction) Purchase(ctx context.Context, request *protos.PurchaseRequest) (*protos.PurchaseResponse, error) {

	s.u.Purchase()

	response := protos.PurchaseResponse{}

	return &response, nil
}

// @Author Ahmad Ridwan Mushoffa
// @Created 03/11/2021
// @Updated
func (s *transaction) Topup(ctx context.Context, request *protos.TopupRequest) (*protos.TopupResponse, error) {
	// fmt.Println("Topup => ", request)

	userRes, err := s.user.InquiryUser(ctx, &protos.InquiryUserRequest{PhoneNumber: request.AccountNumber})
	if err != nil {
		return nil, err
	}

	// fmt.Println("Topup => User: ", userRes.User)

	wallet, err := s.wallet.UpdateBalance(ctx, &protos.UpdateBalanceRequest{
		Id:            request.WalletId,
		AccountNumber: userRes.User.PhoneNumber,
		Amount:        request.Amount,
	})
	if err != nil {
		return nil, err
	}

	// wallet, err := s.wallet.InquiryBalance(ctx, &protos.InquiryBalanceRequest{Id: request.WalletId})
	// if err != nil {
	// 	return nil, err
	// }

	fmt.Println("Topup => Wallet: ", wallet)
	transaction := entity.Transaction{}

	s.u.Topup(&transaction)

	response := protos.TopupResponse{}

	return &response, nil
}
