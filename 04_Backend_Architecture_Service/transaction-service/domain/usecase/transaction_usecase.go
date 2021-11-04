package usecase

import (
	"fmt"
	"time"

	"transaction-service/domain/entity"
	domain "transaction-service/domain/repository"
)

// @Author Ahmad Ridwan Mushoffa
// @Created 03/11/2021
// @Updated
type TransactionUsecase interface {
	Purchase()
	Topup(*entity.Transaction)
	GenerateReferenceNumber(string, string) string
}

// @Author Ahmad Ridwan Mushoffa
// @Created 03/11/2021
// @Updated
type transaction struct {
	r domain.TransactionRepository
}

// @Author Ahmad Ridwan Mushoffa
// @Created 03/11/2021
// @Updated
func NewTransactionUsecase(r domain.TransactionRepository) TransactionUsecase {
	return &transaction{r}
}

// @Author Ahmad Ridwan Mushoffa
// @Created 03/11/2021
// @Updated
func (u *transaction) Purchase() {

}

// @Author Ahmad Ridwan Mushoffa
// @Created 03/11/2021
// @Updated
func (u *transaction) Topup(transaction *entity.Transaction) {

}

// @Author Ahmad Ridwan Mushoffa
// @Created 03/11/2021
// @Updated
func (u *transaction) GenerateReferenceNumber(transactionType, accountNumber string) string {
	timestamp := time.Now().Unix()
	referenceNumber := fmt.Sprintf("%s-%s-%d", transactionType, accountNumber, timestamp)

	return referenceNumber
}
