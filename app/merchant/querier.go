// Code generated by sqlc. DO NOT EDIT.

package merchant

import (
	"context"

	"github.com/jackc/pgtype"
)

type Querier interface {
	AddMerchant(ctx context.Context, arg AddMerchantParams) (Merchant, error)
	AddPartner(ctx context.Context, arg AddPartnerParams) (Partner, error)
	AddShop(ctx context.Context, arg AddShopParams) (Shop, error)
	AddTerminal(ctx context.Context, arg AddTerminalParams) (Terminal, error)
	Auth(ctx context.Context, arg AuthParams) (AuthRow, error)
	DeleteMerchant(ctx context.Context, id pgtype.UUID) (Merchant, error)
	DeletePartner(ctx context.Context, id pgtype.UUID) (Partner, error)
	DeleteShop(ctx context.Context, id pgtype.UUID) (Shop, error)
	DeleteTerminal(ctx context.Context, id pgtype.UUID) (Terminal, error)
}

var _ Querier = (*Queries)(nil)
