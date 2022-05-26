package genesis

import (
	"github.com/zenon-network/go-zenon/common/types"
	"github.com/zenon-network/go-zenon/vm/constants"
	"github.com/zenon-network/go-zenon/vm/embedded/definition"
	"math/big"
)

const (
	// 26 May 2022 00:00:00 UTC
	genesisTimestamp = 1653523200

	// 10 ^ 8
	Zexp = 100000000
)

var (
	sporkAddress = types.ParseAddressPanic("z1qrs8l37648snuqsay93zmsy2m4v6tkwl52j7jp")

	pillar1Owner    = types.ParseAddressPanic("z1qr5ntdwmnw59zwpmh4cadc3sk6ln63ts8dag8j")
	pillar1Producer = types.ParseAddressPanic("z1qquwlzvmcrqayw4d2hl0uurz56t0jz0n4jhyfd")
	pillar1Name     = "alien-valley.io"

	embeddedGenesis = &GenesisConfig{
		ChainIdentifier:     2,
		ExtraData:           "custom dev-net",
		GenesisTimestampSec: genesisTimestamp,
		SporkAddress:        &sporkAddress,
		PillarConfig: &PillarContractConfig{
			Pillars: []*definition.PillarInfo{
				{

					Name:                         pillar1Name,
					BlockProducingAddress:        pillar1Producer,
					StakeAddress:                 pillar1Owner,
					RewardWithdrawAddress:        pillar1Owner,
					Amount:                       new(big.Int).Set(constants.PillarStakeAmount),
					RegistrationTime:             genesisTimestamp,
					RevokeTime:                   0,
					GiveBlockRewardPercentage:    0,
					GiveDelegateRewardPercentage: 100,
					PillarType:                   definition.LegacyPillarType,
				},
			},
			Delegations:   []*definition.DelegationInfo{},
			LegacyEntries: []*definition.LegacyPillarEntry{},
		},
		TokenConfig: &TokenContractConfig{
			Tokens: []*definition.TokenInfo{
				{

					Owner:       types.PillarContract,
					TokenName:   "Zenon Coin",
					TokenSymbol: "ZNN",
					TokenDomain: "zenon.network",
					// The total supply needs to be computed manually
					TotalSupply:   big.NewInt(201 * 15000 * Zexp),
					MaxSupply:     big.NewInt(4611686018427387903),
					Decimals:      8,
					IsMintable:    true,
					IsBurnable:    true,
					IsUtility:     true,
					TokenStandard: types.ZnnTokenStandard,
				},
				{

					Owner:       types.StakeContract,
					TokenName:   "QuasarCoin",
					TokenSymbol: "QSR",
					TokenDomain: "zenon.network",
					// The total supply needs to be computed manually
					TotalSupply:   big.NewInt(Zexp * (2*10000 + 200*150000)),
					MaxSupply:     big.NewInt(4611686018427387903),
					Decimals:      8,
					IsMintable:    true,
					IsBurnable:    true,
					IsUtility:     true,
					TokenStandard: types.QsrTokenStandard,
				},
			},
		},
		PlasmaConfig: &PlasmaContractConfig{
			Fusions: []*definition.FusionInfo{
				{

					Owner:            pillar1Owner,
					Amount:           big.NewInt(10000 * Zexp),
					ExpirationHeight: 0,
					Beneficiary:      pillar1Producer,
				},
				{

					Owner:            pillar1Owner,
					Amount:           big.NewInt(10000 * Zexp),
					ExpirationHeight: 0,
					Beneficiary:      pillar1Owner,
				},
			},
		},
		GenesisBlocks: &GenesisBlocksConfig{
			Blocks: []*GenesisBlockConfig{
				{
					Address: types.PillarContract,
					BalanceList: map[types.ZenonTokenStandard]*big.Int{
						// 1 pillar
						types.ZnnTokenStandard: big.NewInt(1 * 15000 * Zexp),
					},
				},
				{
					Address: types.PlasmaContract,
					BalanceList: map[types.ZenonTokenStandard]*big.Int{
						// 2 fuse entries
						types.QsrTokenStandard: big.NewInt(2 * 10000 * Zexp),
					},
				},
				{
					Address: pillar1Owner,
					BalanceList: map[types.ZenonTokenStandard]*big.Int{
						// 200 pillars worth of ZNN
						types.ZnnTokenStandard: big.NewInt(200 * 15000 * Zexp),
						// 200 pillars worth of QSR
						types.QsrTokenStandard: big.NewInt(200 * 150000 * Zexp),
					},
				},
			},
		},
		SwapConfig: &SwapContractConfig{},
		SporkConfig: &SporkConfig{
			Sporks: []*definition.Spork{
				// deploy the genesis with the accelerator spork already activated, to make sure it has the same id,
				// for simplicity only
				{
					Id:                types.AcceleratorSpork.SporkId,
					Name:              "Pre-activates accelerator spork",
					Description:       "Pre-activates accelerator spork",
					Activated:         true,
					EnforcementHeight: 1,
				},
			},
		},
	}
)
