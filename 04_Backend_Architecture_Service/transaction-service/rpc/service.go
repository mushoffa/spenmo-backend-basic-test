package rpc

import (
	"context"
	// "fmt"
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
// @Updated 04/11/2021
func (s *transaction) Purchase(ctx context.Context, request *protos.PurchaseRequest) (*protos.PurchaseResponse, error) {
	transaction := entity.Transaction{
		PAN: request.Pan,		
		Amount: request.Amount,		
	}

	if err := s.u.Purchase(ctx, &transaction); err != nil {
		return nil, err
	}

	response := protos.PurchaseResponse{
		Transaction: &protos.Transaction {
			Id: transaction.ID,
			Date: transaction.Date,
			Time: transaction.Time,
			Pan: transaction.PAN,
			// AccountNumber: transaction.AccountNumber,
			AccountName: transaction.AccountName,
			WalletId: transaction.WalletID,
			WalletName: transaction.WalletName,
			Amount: transaction.Amount,
			TransactionType: transaction.TransactionType,
			ReferenceNumber: transaction.ReferenceNumber,
			Status: transaction.Status,
		},
	}

	return &response, nil
}

// @Author Ahmad Ridwan Mushoffa
// @Created 03/11/2021
// @Updated 04/11/2021
func (s *transaction) Topup(ctx context.Context, request *protos.TopupRequest) (*protos.TopupResponse, error) {
	// fmt.Println("Topup => ", request)

	// userRes, err := s.user.InquiryUser(ctx, &protos.InquiryUserRequest{PhoneNumber: request.AccountNumber})
	// if err != nil {
	// 	return nil, err
	// }

	// fmt.Println("Topup => User: ", userRes.User)

	

	// walletRes, err := s.wallet.UpdateBalance(ctx, &protos.UpdateBalanceRequest{
	// 	Id:            request.WalletId,
	// 	AccountNumber: userRes.User.PhoneNumber,
	// 	Amount:        request.Amount,
	// })
	// if err != nil {
	// 	return nil, err
	// }

	// fmt.Println("Topup => Wallet: ", walletRes)


	transaction := entity.Transaction{
		AccountNumber: request.AccountNumber,
		WalletID: request.WalletId,
		Amount: request.Amount,		
	}

	if err := s.u.Topup(ctx, &transaction); err != nil {
		return nil, err
	}

	response := protos.TopupResponse{
		Transaction: &protos.Transaction {
			Id: transaction.ID,
			Date: transaction.Date,
			Time: transaction.Time,
			AccountNumber: transaction.AccountNumber,
			AccountName: transaction.AccountName,
			WalletId: transaction.WalletID,
			WalletName: transaction.WalletName,
			Amount: transaction.Amount,
			TransactionType: transaction.TransactionType,
			ReferenceNumber: transaction.ReferenceNumber,
			Status: transaction.Status,
		},
	}

	return &response, nil
}
