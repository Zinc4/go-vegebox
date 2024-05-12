package payment

import (
	"mini-project/user"
	"os"
	"strconv"

	"github.com/veritrans/go-midtrans"
)

type Usecase interface {
	GetPaymentURL(transaction Transaction, user user.User) (string, error)
}

type usecase struct {
}

func NewUsecase() *usecase {
	return &usecase{}
}

func (u *usecase) GetPaymentURL(transaction Transaction, user user.User) (string, error) {
	midClient := midtrans.NewClient()
	server := os.Getenv("MSERVER")
	client := os.Getenv("MCLIENT")
	midClient.ServerKey = server
	midClient.ClientKey = client
	midClient.APIEnvType = midtrans.Sandbox
	orderID := strconv.Itoa(transaction.ID)
	snapGateway := midtrans.SnapGateway{
		Client: midClient,
	}
	snapReq := &midtrans.SnapReq{
		CustomerDetail: &midtrans.CustDetail{
			Email: user.Email,
			FName: user.Name,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderID,
			GrossAmt: int64(transaction.Amount),
		},
	}

	snapTokenRes, err := snapGateway.GetToken(snapReq)
	if err != nil {
		return "", err
	}

	return snapTokenRes.RedirectURL, nil
}
