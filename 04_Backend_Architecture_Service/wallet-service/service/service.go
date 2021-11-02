package service

import (
	"context"
	"wallet-service/domain/entity"
	"wallet-service/domain/usecase"

	"github.com/mushoffa/spenmo-proto/protos"
)

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated
type wallet struct {
	protos.UnimplementedWalletServiceServer
	u usecase.WalletUsecase
}

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated
func NewWalletService(u usecase.WalletUsecase) protos.WalletServiceServer {
	return &wallet{u: u}
}

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated
func (s *wallet) CreateWallet(ctx context.Context, request *protos.CreateWalletRequest) (*protos.CreateWalletResponse, error) {
	wallet := entity.Wallet{
		UserID: request.UserId,
		Name:   request.Name,
	}

	if err := s.u.CreateWallet(&wallet); err != nil {
		return nil, err
	}

	response := protos.CreateWalletResponse{
		Wallet: &protos.Wallet{
			Id:      wallet.ID,
			Created: wallet.Created.String(),
			Name:    wallet.Name,
			Balance: wallet.Balance,
			UserId:  wallet.UserID,
		},
	}

	return &response, nil
}

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated
func (s *wallet) InquiryBalance(ctx context.Context, request *protos.InquiryBalanceRequest) (*protos.InquiryBalanceResponse, error) {
	wallet, err := s.u.InquiryBalance(request.Id)
	if err != nil {
		return nil, err
	}

	response := protos.InquiryBalanceResponse{
		Wallet: &protos.Wallet{
			Id:      wallet.ID,
			Created: wallet.Created.String(),
			Updated: wallet.Updated.String(),
			Name:    wallet.Name,
			Balance: wallet.Balance,
			UserId:  wallet.UserID,
		},
	}

	return &response, nil
}

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated
func (s *wallet) UpdateBalance(ctx context.Context, request *protos.UpdateBalanceRequest) (*protos.UpdateBalanceResponse, error) {
	wallet, err := s.u.UpdateBalance(request.Id, request.UserId, request.Amount)
	if err != nil {
		return nil, err
	}

	response := protos.UpdateBalanceResponse{
		Wallet: &protos.Wallet{
			Id:      wallet.ID,
			Created: wallet.Created.String(),
			Updated: wallet.Updated.String(),
			Name:    wallet.Name,
			Balance: wallet.Balance,
			UserId:  wallet.UserID,
		},
	}

	return &response, nil
}
