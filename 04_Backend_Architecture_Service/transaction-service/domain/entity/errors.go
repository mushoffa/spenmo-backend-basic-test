package entity

import (
	"errors"
)

// @Author Ahmad Ridwan Mushoffa
// @Created 04/11/2021
// @Updated
var (
	ErrInsufficientBalance = errors.New("51 - Insufficient balance")

	ErrExceedsDailyLimit = errors.New("65 - Declined - Customer Exceeds Activity Limit")
)