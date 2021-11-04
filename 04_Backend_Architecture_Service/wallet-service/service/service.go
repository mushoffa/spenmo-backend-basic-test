package service

import (
	"context"
	"wallet-service/domain/entity"
	"wallet-service/domain/usecase"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/mushoffa/spenmo-proto/protos"
)

// @Author Ahmad Ridwan Mushoffa
// @Created 03/11/2021
// @Updated
type wallet struct {
	protos.UnimplementedWalletServiceServer
	u    usecase.WalletUsecase
	user protos.UserServiceClient
}

// @Author Ahmad Ridwan Mushoffa
// @Created 03/11/2021
// @Updated
func NewWalletService(u usecase.WalletUsecase, user protos.UserServiceClient) protos.WalletServiceServer {
	return &wallet{
		u:    u,
		user: user,
	}
}

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated 03/11/2021
func (s *wallet) CreateWallet(ctx context.Context, request *protos.CreateWalletRequest) (*protos.CreateWalletResponse, error) {
	_, err := s.user.InquiryUser(ctx, &protos.InquiryUserRequest{PhoneNumber: request.AccountNumber})
	if err != nil {
		return nil, err
	}

	wallet := entity.Wallet{
		AccountNumber: request.AccountNumber,
		Name:          request.Name,
	}

	if err := s.u.CreateWallet(&wallet); err != nil {
		return nil, err
	}

	response := protos.CreateWalletResponse{
		Wallet: &protos.Wallet{
			Id:            wallet.ID,
			Created:       wallet.Created.String(),
			Name:          wallet.Name,
			Balance:       wallet.Balance,
			AccountNumber: wallet.AccountNumber,
		},
	}

	return &response, nil
}

// @Author Ahmad Ridwan Mushoffa
// @Created 03/11/2021
// @Updated
func (s *wallet) GetAllWallets(ctx context.Context, request *empty.Empty) (*protos.GetAllWalletsResponse, error) {
	_wallets, err := s.u.GetAllWallets()
	if err != nil {
		return nil, err
	}

	wallets := []*protos.Wallet{}

	for _, _wallet := range _wallets {

		wallet := &protos.Wallet{
			Id:            _wallet.ID,
			Created:       _wallet.Created.String(),
			Updated:       _wallet.Updated.String(),
			Name:          _wallet.Name,
			Balance:       _wallet.Balance,
			AccountNumber: _wallet.AccountNumber,
		}

		wallets = append(wallets, wallet)
	}

	response := protos.GetAllWalletsResponse{
		Wallets: wallets,
	}

	return &response, nil
}

// @Author Ahmad Ridwan Mushoffa
// @Created 03/11/2021
// @Updated
func (s *wallet) GetByUserID(ctx context.Context, request *protos.GetByUserIDRequest) (*protos.GetByUserIDResponse, error) {
	_wallets, err := s.u.GetByAccountNumber(request.AccountNumber)
	if err != nil {
		return nil, err
	}

	wallets := []*protos.Wallet{}

	for _, _wallet := range _wallets {
		wallet := &protos.Wallet{
			Id:            _wallet.ID,
			Created:       _wallet.Created.String(),
			Updated:       _wallet.Updated.String(),
			Name:          _wallet.Name,
			Balance:       _wallet.Balance,
			AccountNumber: _wallet.AccountNumber,
		}

		wallets = append(wallets, wallet)
	}

	response := protos.GetByUserIDResponse{
		Wallets: wallets,
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
			Id:            wallet.ID,
			Created:       wallet.Created.String(),
			Updated:       wallet.Updated.String(),
			Name:          wallet.Name,
			Balance:       wallet.Balance,
			AccountNumber: wallet.AccountNumber,
		},
	}

	return &response, nil
}

// @Author Ahmad Ridwan Mushoffa
// @Created 02/11/2021
// @Updated
func (s *wallet) UpdateBalance(ctx context.Context, request *protos.UpdateBalanceRequest) (*protos.UpdateBalanceResponse, error) {
	_, err := s.user.InquiryUser(ctx, &protos.InquiryUserRequest{PhoneNumber: request.AccountNumber})
	if err != nil {
		return nil, err
	}

	wallet, err := s.u.UpdateBalance(request.Id, request.AccountNumber, request.Amount)
	if err != nil {
		return nil, err
	}

	response := protos.UpdateBalanceResponse{
		Wallet: &protos.Wallet{
			Id:            wallet.ID,
			Created:       wallet.Created.String(),
			Updated:       wallet.Updated.String(),
			Name:          wallet.Name,
			Balance:       wallet.Balance,
			AccountNumber: wallet.AccountNumber,
		},
	}

	return &response, nil
}
