package liquidstake_test

// func (s *LiquidStakePrecompileTestSuite) TestLiquidStakeEvent() {
// 	var (
// 		stDB *statedb.StateDB
// 		ctx  sdk.Context
// 	)
// 	method := s.precompile.Methods[liquidstake.LiquidStakeMethod]
// 	testCases := []struct {
// 		name        string
// 		malleate    func(delegator common.Address) []interface{}
// 		expErr      bool
// 		errContains string
// 		postCheck   func(delegator common.Address)
// 	}{
// 		{
// 			"success - LiquidStake event emitted correctly",
// 			func(delegator common.Address) []interface{} {
// 				return []interface{}{
// 					delegator,
// 					big.NewInt(1000000000000000000),
// 				}
// 			},
// 			false,
// 			"",
// 			func(delegator common.Address) {
// 				log := stDB.Logs()[0]
// 				s.Require().Equal(log.Address, s.precompile.Address())
//
// 				// Check event signature matches the one emitted
// 				event := s.precompile.ABI.Events[liquidstake.EventTypeLiquidStake]
// 				s.Require().Equal(crypto.Keccak256Hash([]byte(event.Sig)), common.HexToHash(log.Topics[0].Hex()))
// 				s.Require().Equal(log.BlockNumber, uint64(ctx.BlockHeight())) //nolint:gosec
//
// 				var liquidStakeEvent liquidstake.EventLiquidStake
// 				err := cmn.UnpackLog(s.precompile.ABI, &liquidStakeEvent, liquidstake.EventTypeLiquidStake, *log)
// 				s.Require().NoError(err)
// 				s.Require().Equal(delegator, liquidStakeEvent.DelegatorAddress)
// 				s.Require().Equal(big.NewInt(1000000000000000000), liquidStakeEvent.Amount)
// 			},
// 		},
// 	}
//
// 	for _, tc := range testCases {
// 		s.Run(tc.name, func() {
// 			s.SetupTest() // reset
// 			ctx = s.nw.GetContext()
// 			stDB = s.nw.GetStateDB()
//
// 			delegator := s.keyring.GetKey(0)
//
// 			contract := vm.NewContract(vm.AccountRef(delegator.Addr), s.precompile, common.U2560, 200000)
//
// 			_, err := s.precompile.LiquidStake(ctx, delegator.Addr, contract, stDB, &method, tc.malleate(delegator.Addr))
//
// 			if tc.expErr {
// 				s.Require().Error(err)
// 				s.Require().Contains(err.Error(), tc.errContains)
// 			} else {
// 				s.Require().NoError(err)
// 				tc.postCheck(delegator.Addr)
// 			}
// 		})
// 	}
// }
//
// func (s *LiquidStakePrecompileTestSuite) TestUpdateParamsEvent() {
// 	var (
// 		stDB *statedb.StateDB
// 		ctx  sdk.Context
// 	)
// 	method := s.precompile.Methods[liquidstake.UpdateParams]
// 	testCases := []struct {
// 		name        string
// 		malleate    func(admin common.Address) []interface{}
// 		expErr      bool
// 		errContains string
// 		postCheck   func(admin common.Address)
// 	}{
// 		{
// 			"success - UpdateParams event emitted correctly",
// 			func(admin common.Address) []interface{} {
// 				// Create test params with all required fields
// 				params := liquidstake.LiquidStakeUpdatableParams{
// 					UnstakeFeeRate:        big.NewInt(1000),
// 					LsmDisabled:           false,
// 					MinLiquidStakeAmount:  big.NewInt(1000000),
// 					CwLockedPoolAddress:   common.HexToAddress("0x1"),
// 					FeeAccountAddress:     common.HexToAddress("0x2"),
// 					AutocompoundFeeRate:   big.NewInt(500),
// 					WhitelistAdminAddress: admin,
// 				}
// 				return []interface{}{params}
// 			},
// 			false,
// 			"",
// 			func(admin common.Address) {
// 				log := stDB.Logs()[0]
// 				s.Require().Equal(log.Address, s.precompile.Address())
//
// 				// Check event signature matches the one emitted
// 				event := s.precompile.ABI.Events[liquidstake.EventTypeUpdateParams]
// 				s.Require().Equal(crypto.Keccak256Hash([]byte(event.Sig)), common.HexToHash(log.Topics[0].Hex()))
// 				s.Require().Equal(log.BlockNumber, uint64(ctx.BlockHeight())) //nolint:gosec
//
// 				var updateParamsEvent liquidstake.EventUpdateParams
// 				err := cmn.UnpackLog(s.precompile.ABI, &updateParamsEvent, liquidstake.EventTypeUpdateParams, *log)
// 				s.Require().NoError(err)
// 				s.Require().Equal(admin, updateParamsEvent.Params.WhitelistAdminAddress)
// 			},
// 		},
// 	}
//
// 	for _, tc := range testCases {
// 		s.Run(tc.name, func() {
// 			s.SetupTest() // reset
// 			ctx = s.nw.GetContext()
// 			stDB = s.nw.GetStateDB()
//
// 			contract := vm.NewContract(vm.AccountRef(s.admin.Addr), s.precompile, common.U2560, 200000)
// 			contract.CallerAddress = s.admin.Addr
//
// 			_, err := s.precompile.UpdateParams(ctx, s.admin.Addr, contract, stDB, &method, tc.malleate(s.admin.Addr))
//
// 			if tc.expErr {
// 				s.Require().Error(err)
// 				s.Require().Contains(err.Error(), tc.errContains)
// 			} else {
// 				s.Require().NoError(err)
// 				tc.postCheck(s.admin.Addr)
// 			}
// 		})
// 	}
// }
//
// func (s *LiquidStakePrecompileTestSuite) TestUpdateWhitelistedValidatorsEvent() {
// 	var (
// 		stDB *statedb.StateDB
// 		ctx  sdk.Context
// 	)
// 	method := s.precompile.Methods[liquidstake.UpdateWhitelistedValidators]
// 	testCases := []struct {
// 		name        string
// 		malleate    func(admin common.Address) []interface{}
// 		expErr      bool
// 		errContains string
// 		postCheck   func(admin common.Address)
// 	}{
// 		{
// 			"success - UpdateWhitelistedValidators event emitted correctly",
// 			func(admin common.Address) []interface{} {
// 				// Create test whitelisted validators
// 				validator1 := liquidstake.WhitelistedValidator{
// 					ValidatorAddress: s.ValidatorAddr,
// 					TargetWeight:     big.NewInt(10000),
// 				}
// 				whitelistedValidators := []liquidstake.WhitelistedValidator{validator1}
// 				return []interface{}{whitelistedValidators}
// 			},
// 			false,
// 			"",
// 			func(admin common.Address) {
// 				log := stDB.Logs()[0]
// 				s.Require().Equal(log.Address, s.precompile.Address())
//
// 				// Check event signature matches the one emitted
// 				event := s.precompile.ABI.Events[liquidstake.EventTypeUpdateWhitelistedValidator]
// 				s.Require().Equal(crypto.Keccak256Hash([]byte(event.Sig)), common.HexToHash(log.Topics[0].Hex()))
// 				s.Require().Equal(log.BlockNumber, uint64(ctx.BlockHeight())) //nolint:gosec
//
// 				var updateWhitelistEvent liquidstake.EventUpdateWhitelistedValidator
// 				err := cmn.UnpackLog(s.precompile.ABI, &updateWhitelistEvent, liquidstake.EventTypeUpdateWhitelistedValidator, *log)
// 				s.Require().NoError(err)
// 				s.Require().Equal(1, len(updateWhitelistEvent.WhitelistedValidators))
// 				s.Require().Equal(s.ValidatorAddr, updateWhitelistEvent.WhitelistedValidators[0].ValidatorAddress)
// 				s.Require().Equal(big.NewInt(10000), updateWhitelistEvent.WhitelistedValidators[0].TargetWeight)
// 			},
// 		},
// 	}
//
// 	for _, tc := range testCases {
// 		s.Run(tc.name, func() {
// 			s.SetupTest() // reset
// 			ctx = s.nw.GetContext()
// 			stDB = s.nw.GetStateDB()
//
// 			contract := vm.NewContract(vm.AccountRef(s.admin.Addr), s.precompile, common.U2560, 200000)
// 			contract.CallerAddress = s.admin.Addr
//
// 			_, err := s.precompile.UpdateWhitelistedValidators(ctx, s.admin.Addr, contract, stDB, &method, tc.malleate(s.admin.Addr))
//
// 			if tc.expErr {
// 				s.Require().Error(err)
// 				s.Require().Contains(err.Error(), tc.errContains)
// 			} else {
// 				s.Require().NoError(err)
// 				tc.postCheck(s.admin.Addr)
// 			}
// 		})
// 	}
// }
//
// func (s *LiquidStakePrecompileTestSuite) TestSetModulePausedEvent() {
// 	var (
// 		stDB *statedb.StateDB
// 		ctx  sdk.Context
// 	)
// 	method := s.precompile.Methods[liquidstake.SetModulePaused]
// 	testCases := []struct {
// 		name        string
// 		malleate    func(admin common.Address) []interface{}
// 		expErr      bool
// 		errContains string
// 		postCheck   func(admin common.Address)
// 	}{
// 		{
// 			"success - SetModulePaused event emitted correctly (pause module)",
// 			func(admin common.Address) []interface{} {
// 				return []interface{}{true}
// 			},
// 			false,
// 			"",
// 			func(admin common.Address) {
// 				log := stDB.Logs()[0]
// 				s.Require().Equal(log.Address, s.precompile.Address())
//
// 				// Check event signature matches the one emitted
// 				event := s.precompile.ABI.Events[liquidstake.EventTypeSetModulePaused]
// 				s.Require().Equal(crypto.Keccak256Hash([]byte(event.Sig)), common.HexToHash(log.Topics[0].Hex()))
// 				s.Require().Equal(log.BlockNumber, uint64(ctx.BlockHeight())) //nolint:gosec
//
// 				var setModulePausedEvent liquidstake.EventSetModulePaused
// 				err := cmn.UnpackLog(s.precompile.ABI, &setModulePausedEvent, liquidstake.EventTypeSetModulePaused, *log)
// 				s.Require().NoError(err)
// 				s.Require().Equal(true, setModulePausedEvent.IsPaused)
// 			},
// 		},
// 		{
// 			"success - SetModulePaused event emitted correctly (unpause module)",
// 			func(admin common.Address) []interface{} {
// 				return []interface{}{false}
// 			},
// 			false,
// 			"",
// 			func(admin common.Address) {
// 				log := stDB.Logs()[0]
// 				s.Require().Equal(log.Address, s.precompile.Address())
//
// 				// Check event signature matches the one emitted
// 				event := s.precompile.ABI.Events[liquidstake.EventTypeSetModulePaused]
// 				s.Require().Equal(crypto.Keccak256Hash([]byte(event.Sig)), common.HexToHash(log.Topics[0].Hex()))
// 				s.Require().Equal(log.BlockNumber, uint64(ctx.BlockHeight())) //nolint:gosec
//
// 				var setModulePausedEvent liquidstake.EventSetModulePaused
// 				err := cmn.UnpackLog(s.precompile.ABI, &setModulePausedEvent, liquidstake.EventTypeSetModulePaused, *log)
// 				s.Require().NoError(err)
// 				s.Require().Equal(false, setModulePausedEvent.IsPaused)
// 			},
// 		},
// 	}
//
// 	for _, tc := range testCases {
// 		s.Run(tc.name, func() {
// 			s.SetupTest() // reset
// 			ctx = s.nw.GetContext()
// 			stDB = s.nw.GetStateDB()
//
// 			contract := vm.NewContract(vm.AccountRef(s.admin.Addr), s.precompile, common.U2560, 200000)
// 			contract.CallerAddress = s.admin.Addr
//
// 			_, err := s.precompile.SetModulePaused(ctx, s.admin.Addr, contract, stDB, &method, tc.malleate(s.admin.Addr))
//
// 			if tc.expErr {
// 				s.Require().Error(err)
// 				s.Require().Contains(err.Error(), tc.errContains)
// 			} else {
// 				s.Require().NoError(err)
// 				tc.postCheck(s.admin.Addr)
// 			}
// 		})
// 	}
// }

//func (s *LiquidStakePrecompileTestSuite) TestStakeToLPEvent() {
//	var (
//		stDB *statedb.StateDB
//		ctx  sdk.Context
//	)
//	method := s.precompile.Methods[liquidstake.StakeToLPMethod]
//	testCases := []struct {
//		name        string
//		malleate    func(delegator common.Address) []interface{}
//		expErr      bool
//		errContains string
//		postCheck   func(delegator common.Address)
//	}{
//		{
//			"success - StakeToLP event emitted correctly",
//			func(delegator common.Address) []interface{} {
//				validator := s.validatorAdr
//				return []interface{}{
//					delegator,
//					validator,
//					big.NewInt(1000000000000000000),
//					big.NewInt(1000000000000000000),
//				}
//			},
//			false,
//			"",
//			func(delegator common.Address) {
//				s.SetupTest()
//				log := stDB.Logs()[0]
//				s.Require().Equal(log.Address, s.precompile.Address())
//
//				// Check event signature matches the one emitted
//				event := s.precompile.ABI.Events[liquidstake.EventTypeStakeToLP]
//				s.Require().Equal(crypto.Keccak256Hash([]byte(event.Sig)), common.HexToHash(log.Topics[0].Hex()))
//				s.Require().Equal(log.BlockNumber, uint64(ctx.BlockHeight())) //nolint:gosec
//
//				var stakeToLPEvent liquidstake.EventStakeToLP
//				err := cmn.UnpackLog(s.precompile.ABI, &stakeToLPEvent, liquidstake.EventTypeStakeToLP, *log)
//				s.Require().NoError(err)
//				s.Require().Equal(delegator, stakeToLPEvent.DelegatorAddress)
//				s.Require().Equal(big.NewInt(1000000000000000000), stakeToLPEvent.StakedAmount)
//				s.Require().Equal(big.NewInt(500000000000000000), stakeToLPEvent.LiquidAmount)
//			},
//		},
//	}
//
//	for _, tc := range testCases {
//		s.Run(tc.name, func() {
//			s.SetupTest() // reset
//			ctx = s.nw.GetContext()
//			stDB = s.nw.GetStateDB()
//
//			delegator := s.keyring.GetKey(0)
//
//			contract := vm.NewContract(vm.AccountRef(delegator.Addr), s.precompile, common.U2560, 200000)
//
//			_, err := s.precompile.StakeToLP(ctx, delegator.Addr, contract, stDB, &method, tc.malleate(delegator.Addr))
//
//			if tc.expErr {
//				s.Require().Error(err)
//				s.Require().Contains(err.Error(), tc.errContains)
//			} else {
//				s.Require().NoError(err)
//				tc.postCheck(delegator.Addr)
//			}
//		})
//	}
//}

// func (s *LiquidStakePrecompileTestSuite) TestLiquidUnstakeEvent() {
// 	var (
// 		stDB *statedb.StateDB
// 		ctx  sdk.Context
// 	)
// 	method := s.precompile.Methods[liquidstake.LiquidUnstakeMethod]
// 	testCases := []struct {
// 		name        string
// 		malleate    func(delegator common.Address) []interface{}
// 		expErr      bool
// 		errContains string
// 		postCheck   func(delegator common.Address)
// 	}{
// 		{
// 			"success - LiquidUnstake event emitted correctly",
// 			func(delegator common.Address) []interface{} {
// 				_, err := s.nw.App.LiquidStakeKeeper.LiquidStake(ctx, liquidstaketypes.LiquidStakeProxyAcc, sdk.AccAddress(delegator.Bytes()), sdk.NewCoin(s.bondDenom, sdkmath.NewInt(1000000000000000000)))
// 				s.Require().NoError(err)
//
// 				s.Require().NoError(err, "failed to pack input")
//
// 				return []interface{}{
// 					delegator,
// 					big.NewInt(1000000000000000000),
// 				}
// 			},
// 			false,
// 			"",
// 			func(delegator common.Address) {
// 				log := stDB.Logs()[0]
// 				s.Require().Equal(log.Address, s.precompile.Address())
//
// 				// Check event signature matches the one emitted
// 				event := s.precompile.ABI.Events[liquidstake.EventTypeLiquidUnstake]
// 				s.Require().Equal(crypto.Keccak256Hash([]byte(event.Sig)), common.HexToHash(log.Topics[0].Hex()))
// 				s.Require().Equal(log.BlockNumber, uint64(ctx.BlockHeight())) //nolint:gosec
//
// 				var liquidUnstakeEvent liquidstake.EventLiquidUnstake
// 				err := cmn.UnpackLog(s.precompile.ABI, &liquidUnstakeEvent, liquidstake.EventTypeLiquidUnstake, *log)
// 				s.Require().NoError(err)
// 				s.Require().Equal(delegator, liquidUnstakeEvent.DelegatorAddress)
// 				s.Require().Equal(big.NewInt(1000000000000000000), liquidUnstakeEvent.Amount)
// 			},
// 		},
// 	}
//
// 	for _, tc := range testCases {
// 		s.Run(tc.name, func() {
// 			s.SetupTest() // reset
// 			ctx = s.nw.GetContext()
// 			stDB = s.nw.GetStateDB()
//
// 			delegator := s.keyring.GetKey(0)
//
// 			contract := vm.NewContract(vm.AccountRef(delegator.Addr), s.precompile, common.U2560, 200000)
//
// 			_, err := s.precompile.LiquidUnstake(ctx, delegator.Addr, contract, stDB, &method, tc.malleate(delegator.Addr))
//
// 			if tc.expErr {
// 				s.Require().Error(err)
// 				s.Require().Contains(err.Error(), tc.errContains)
// 			} else {
// 				s.Require().NoError(err)
// 				tc.postCheck(delegator.Addr)
// 			}
// 		})
// 	}
// }
