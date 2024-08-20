// Copyright 2022 Evmos Foundation
// This file is part of the Ethermint Network packages.
//
// Ethermint is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The Ethermint packages are distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the Ethermint packages. If not, see https://github.com/xpladev/ethermint/blob/main/LICENSE

package transfer

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/types/module"

	ibctransfer "github.com/cosmos/ibc-go/v8/modules/apps/transfer"
	ibctransferkeeper "github.com/cosmos/ibc-go/v8/modules/apps/transfer/keeper"
	"github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	"github.com/xpladev/ethermint/x/ibc/transfer/keeper"
)

var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
)

// AppModuleBasic embeds the IBC Transfer AppModuleBasic
type AppModuleBasic struct {
	*ibctransfer.AppModuleBasic
}

// AppModule represents the AppModule for this module
type AppModule struct {
	*ibctransfer.AppModule
	keeper keeper.Keeper
}

// NewAppModule creates a new 20-transfer module
func NewAppModule(k keeper.Keeper) AppModule {
	am := ibctransfer.NewAppModule(*k.Keeper)
	return AppModule{
		AppModule: &am,
		keeper:    k,
	}
}

// RegisterServices registers module services.
func (am AppModule) RegisterServices(cfg module.Configurator) {
	// Override Transfer Msg Server
	types.RegisterMsgServer(cfg.MsgServer(), am.keeper)
	types.RegisterQueryServer(cfg.QueryServer(), am.keeper)

	m := ibctransferkeeper.NewMigrator(*am.keeper.Keeper)
	if err := cfg.RegisterMigration(types.ModuleName, 1, m.MigrateTraces); err != nil {
		panic(fmt.Errorf("failed to migrate transfer app from version 1 to 2 (denom trace format migration): %v", err))
	}

	if err := cfg.RegisterMigration(types.ModuleName, 2, m.MigrateTotalEscrowForDenom); err != nil {
		panic(fmt.Errorf("failed to migrate transfer app from version 2 to 3 (total escrow entry migration): %v", err))
	}

	if err := cfg.RegisterMigration(types.ModuleName, 3, m.MigrateParams); err != nil {
		panic(fmt.Errorf("failed to migrate transfer app version 3 to 4 (self-managed params migration): %v", err))
	}

	if err := cfg.RegisterMigration(types.ModuleName, 4, m.MigrateDenomMetadata); err != nil {
		panic(fmt.Errorf("failed to migrate transfer app from version 4 to 5 (set denom metadata migration): %v", err))
	}
}
