package v1

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated
type CardController interface {
}

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated
type cardController struct {
}

// @Author Ahmad Ridwan Mushoffa
// @Created 01/11/2021
// @Updated
func NewCardController() CardController {
	return &cardController{}
}
