package ops

import (
	"github.com/interflowrepo/interflow-wallet-api/accounts"
)

// Store defines what ops needs from the database
type Store interface {
	ListAccountsWithMissingVault(tokenName string) (*[]accounts.Account, error)
}
