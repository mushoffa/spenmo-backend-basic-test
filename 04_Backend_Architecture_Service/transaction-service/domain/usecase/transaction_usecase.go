package usecase

import (
	"context"

	"transaction-service/domain/entity"
	domain "transaction-service/domain/repository"

	"github.com/mushoffa/spenmo-proto/protos"
)

// @Author Ahmad Ridwan Mushoffa
// @Created 03/11/2021
// @Updated 04/11/2021
type TransactionUsecase interface {
	Purchase(context.Context, *entity.Transaction) error
	Topup(context.Context, *entity.Transaction) error
	// GenerateUniqueID(*entity.Transaction)
	// GenerateDateTime(*entity.Transaction)
	// GenerateReferenceNumber(*entity.Transaction)
}

// @Author Ahmad Ridwan Mushoffa
// @Created 03/11/2021
// @Updated 04/11/2021
type transaction struct {
	r domain.TransactionRepository
	user   protos.UserServiceClient
	card   protos.CardServiceClient
	wallet protos.WalletServiceClient
}

// @Author Ahmad Ridwan Mushoffa
// @Created 03/11/2021
// @Updated
func NewTransactionUsecase(r domain.TransactionRepository, user protos.UserServiceClient, card protos.CardServiceClient, wallet protos.WalletServiceClient) TransactionUsecase {
	return &transaction{r, user, card, wallet}
}

// @Author Ahmad Ridwan Mushoffa
// @Created 03/11/2021
// @Updated
func (u *transaction) Purchase(ctx context.Context, transaction *entity.Transaction) error {
	
	// Inquiry card to ensure the card belongs to our system (ON US)
	cardResponse, err := u.card.InquiryCard(ctx, &protos.InquiryCardRequest{Id: transaction.PAN})
	if err != nil {
		return err
	}

	// Inquiry purchase limit: WIP
	if transaction.Amount > cardResponse.Card.LimitDaily {
		return entity.ErrExceedsDailyLimit
	}

	// Inquiry wallet to ensure the balance is sufficient to process payment
	walletResponse, err := u.wallet.InquiryBalance(ctx, &protos.InquiryBalanceRequest{Id: cardResponse.Card.WalletId})
	if err != nil {
		return err
	}

	if walletResponse.Wallet.Balance < transaction.Amount {
		return entity.ErrInsufficientBalance
	}

	transaction.AccountName = cardResponse.Card.Name
	transaction.WalletID = walletResponse.Wallet.Id
	transaction.WalletName = walletResponse.Wallet.Name
	transaction.TransactionType = "DEBIT"
	transaction.Status = "SUCCESS"
	transaction.GenerateID()
	transaction.GenerateDateTime()
	transaction.GenerateReferenceNumber(transaction.PAN)


	return nil
}

// @Author Ahmad Ridwan Mushoffa
// @Created 03/11/2021
// @Updated 04/11/2021
func (u *transaction) Topup(ctx context.Context, transaction *entity.Transaction) error {
	
	// Inquiry user to ensure it is exist and valid
	userResponse, err := u.user.InquiryUser(ctx, &protos.InquiryUserRequest{PhoneNumber: transaction.AccountNumber})
	if err != nil {
		return err
	}

	// Update wallet balance, the validation is occured within Wallet Service
	walletResponse, err := u.wallet.UpdateBalance(ctx, &protos.UpdateBalanceRequest{
		Id:            transaction.WalletID,
		AccountNumber: transaction.AccountNumber,
		Amount:        transaction.Amount,
	})
	if err != nil {
		return err
	}

	transaction.AccountName = userResponse.User.Name
	transaction.WalletName = walletResponse.Wallet.Name
	transaction.TransactionType = "TOPUP"
	transaction.Status = "SUCCESS"
	transaction.GenerateID()
	transaction.GenerateDateTime()
	transaction.GenerateReferenceNumber(transaction.AccountNumber)

	return nil
}

// @Author Ahmad Ridwan Mushoffa
// @Created 04/11/2021
// @Updated
// func (u *transaction) GenerateUniqueID(transaction *entity.Transaction) {
// 	transaction.ID = uuid.New().String()
// }

// @Author Ahmad Ridwan Mushoffa
// @Created 04/11/2021
// @Updated
// func (u *transaction) GenerateDateTime(transaction *entity.Transaction) {
// 	today := time.Now()

// 	dateLayout := "02-01-2006"
// 	timeLayout := "15:04:05"

// 	transaction.Date = today.Format(dateLayout)
// 	transaction.Time = today.Format(timeLayout)
// }

// @Author Ahmad Ridwan Mushoffa
// @Created 03/11/2021
// @Updated 04/11/2021
// func (u *transaction) GenerateReferenceNumber(transactionType, accountNumber string) string {
// func (u *transaction) GenerateReferenceNumber(transaction *entity.Transaction)  {
// 	timestamp := time.Now().Unix()
// 	transaction.ReferenceNumber = fmt.Sprintf("%s-%s-%d", transaction.TransactionType, transaction.AccountNumber, timestamp)
// }