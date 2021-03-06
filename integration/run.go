package integration

import (
	"context"
	"encoding/hex"
	"fmt"
	"os"
	"time"

	"github.com/hdac-io/casperlabs-ee-grpc-go-util/grpc"
	"github.com/hdac-io/casperlabs-ee-grpc-go-util/protobuf/io/casperlabs/casper/consensus"
	"github.com/hdac-io/casperlabs-ee-grpc-go-util/protobuf/io/casperlabs/casper/consensus/state"
	"github.com/hdac-io/casperlabs-ee-grpc-go-util/protobuf/io/casperlabs/ipc"
	"github.com/hdac-io/casperlabs-ee-grpc-go-util/protobuf/io/casperlabs/ipc/transforms"
	"github.com/hdac-io/casperlabs-ee-grpc-go-util/storedvalue"
	"github.com/hdac-io/casperlabs-ee-grpc-go-util/util"
)

const (
	CHAIN_NAME = "hdac"

	GENESIS_ADDRESS_HEX = "d70243dd9d0d646fd6df282a8f7a8fa05a6629bec01d8024c3611eb1c1fb9f84"
	ADDRESS1_HEX        = "93236a9263d2ac6198c5ed211774c745d5dc62a910cb84276f8a7c4959208915"
	ADDRESS1_DAPP_HEX   = "0193236a9263d2ac6198c5ed211774c745d5dc62a910cb84276f8a7c4959208915"

	INITIAL_BALANCE     = "50000000000000000000000"
	INITIAL_BOND_AMOUNT = "1000000000000000000"

	BASIC_FEE = "100000000000000000"
	ADVANCED_FEE = "300000000000000000"

	DAPP_HASH_HEX = "01d70243dd9d0d646fd6df282a8f7a8fa05a6629bec01d8024c3611eb1c1fb9f84"
)

var (
	SYSTEM_ACCOUNT  = make([]byte, 32)
	GENESIS_ADDRESS = util.DecodeHexString(GENESIS_ADDRESS_HEX)
	ADDRESS1        = util.DecodeHexString(ADDRESS1_HEX)
	ADDRESS1_DAPP   = util.DecodeHexString(ADDRESS1_DAPP_HEX)
	DAPP_HASH       = util.DecodeHexString(DAPP_HASH_HEX)

	DEFAULT_GENESIS_ACCOUNT = []*ipc.ChainSpec_GenesisAccount{{
		PublicKey:    GENESIS_ADDRESS,
		Balance:      &state.BigInt{Value: INITIAL_BALANCE, BitWidth: 512},
		BondedAmount: &state.BigInt{Value: INITIAL_BOND_AMOUNT, BitWidth: 512}}}
)

func GetPaymentArgsJson(fee string) string {
	paymentArgs := []*consensus.Deploy_Arg{
		&consensus.Deploy_Arg{
			Name: "method",
			Value: &state.CLValueInstance{
				ClType: &state.CLType{Variants: &state.CLType_SimpleType{SimpleType: state.CLType_STRING}},
				Value: &state.CLValueInstance_Value{
					Value: &state.CLValueInstance_Value_StrValue{
						StrValue: "standard_payment"}}}},
		&consensus.Deploy_Arg{
			Name: "fee",
			Value: &state.CLValueInstance{
				ClType: &state.CLType{Variants: &state.CLType_SimpleType{SimpleType: state.CLType_U512}},
				Value: &state.CLValueInstance_Value{
					Value: &state.CLValueInstance_Value_U512{
						U512: &state.CLValueInstance_U512{
							Value: fee}}}}}}

	paymentArgsStr, err := util.DeployArgsToJsonString(paymentArgs)
	if err != nil {
		panic(err)
	}
	return paymentArgsStr
}

func InitalRunGenensis(genesisAccounts []*ipc.ChainSpec_GenesisAccount) (ipc.ExecutionEngineServiceClient, []byte, []byte, *state.ProtocolVersion) {
	// Init variable
	emptyStateHash := util.DecodeHexString(util.StrEmptyStateHash)
	rootStateHash := emptyStateHash

	costs := map[string]uint32{
		"regular":            1,
		"div-multiplier":     16,
		"mul-multiplier":     4,
		"mem-multiplier":     2,
		"mem-initial-pages":  4096,
		"mem-grow-per-page":  8192,
		"mem-copy-per-byte":  1,
		"max-stack-height":   65536,
		"opcodes-multiplier": 3,
		"opcodes-divisor":    8}

	protocolVersion := storedvalue.NewProtocolVersion(1, 0, 0).ToStateValue()

	// Connect to ee sock.
	socketPath := os.Getenv("HOME") + `/.casperlabs/.casper-node.sock`
	client := grpc.Connect(socketPath)

	// run genesis
	println(`RunGenesis`)
	genesisConfig, err := util.GenesisConfigMock(
		CHAIN_NAME, genesisAccounts, protocolVersion, costs,
		"./contracts/hdac_mint_install.wasm", "./contracts/pop_install.wasm", "./contracts/standard_payment_install.wasm")
	if err != nil {
		panic(err)
	}

	response, err := grpc.RunGenesis(client, genesisConfig)
	if err != nil {
		panic(err)
	}

	switch response.GetResult().(type) {
	case *ipc.GenesisResponse_Success:
		rootStateHash = response.GetSuccess().GetPoststateHash()
		// effects = response.GetSuccess().GetEffect().GetTransformMap()
	case *ipc.GenesisResponse_FailedDeploy:
		panic(response.GetFailedDeploy().GetMessage())
	}

	queryResult10, errMessage := grpc.Query(client, rootStateHash, "address", SYSTEM_ACCOUNT, []string{}, protocolVersion)
	if errMessage != "" {
		panic(errMessage)
	}
	var storedValue storedvalue.StoredValue
	storedValue.FromBytes(queryResult10)
	storedValue, err, _ = storedValue.FromBytes(queryResult10)
	if err != nil {
		panic(err)
	}

	proxyHash := storedValue.Account.NamedKeys[0].Key.Hash
	println("Proxy hash : " + util.EncodeToHexString(proxyHash))
	println(errMessage)

	return client, rootStateHash, proxyHash, protocolVersion
}

func RunCounterDefine(client ipc.ExecutionEngineServiceClient, stateHash []byte, runAddress []byte, proxyHash []byte, protocolVersion *state.ProtocolVersion) (resultStateHash []byte, bonds []*ipc.Bond) {
	counterDefineCode := util.LoadWasmFile("./contracts/counter_define.wasm")

	return RunExecute(client, stateHash, runAddress, util.WASM, counterDefineCode, "", proxyHash, BASIC_FEE, protocolVersion)
}

func RunCounterCall(client ipc.ExecutionEngineServiceClient, stateHash []byte, runAddress []byte, proxyHash []byte, protocolVersion *state.ProtocolVersion) (resultStateHash []byte, bonds []*ipc.Bond) {
	counterCallCode := util.LoadWasmFile("./contracts/counter_call.wasm")

	return RunExecute(client, stateHash, runAddress, util.WASM, counterCallCode, "", proxyHash, BASIC_FEE, protocolVersion)
}

func RunTransferToAccount(client ipc.ExecutionEngineServiceClient, stateHash []byte, runAddress []byte,
	toAddress []byte, amount string,
	proxyHash []byte, protocolVersion *state.ProtocolVersion) (resultStateHash []byte, bonds []*ipc.Bond) {
	sessionArgs := []*consensus.Deploy_Arg{
		&consensus.Deploy_Arg{
			Name: "method",
			Value: &state.CLValueInstance{
				ClType: &state.CLType{Variants: &state.CLType_SimpleType{SimpleType: state.CLType_STRING}},
				Value: &state.CLValueInstance_Value{
					Value: &state.CLValueInstance_Value_StrValue{
						StrValue: "transfer_to_account"}}}},
		&consensus.Deploy_Arg{
			Name: "address",
			Value: &state.CLValueInstance{
				ClType: &state.CLType{Variants: &state.CLType_ListType{ListType: &state.CLType_List{Inner: &state.CLType{Variants: &state.CLType_SimpleType{SimpleType: state.CLType_U8}}}}},
				Value: &state.CLValueInstance_Value{
					Value: &state.CLValueInstance_Value_BytesValue{
						BytesValue: toAddress}}}},
		&consensus.Deploy_Arg{
			Name: "amount",
			Value: &state.CLValueInstance{
				ClType: &state.CLType{Variants: &state.CLType_SimpleType{SimpleType: state.CLType_U512}},
				Value: &state.CLValueInstance_Value{
					Value: &state.CLValueInstance_Value_U512{
						U512: &state.CLValueInstance_U512{
							Value: amount}}}}}}
	sessionArgsStr, err := util.DeployArgsToJsonString(sessionArgs)
	if err != nil {
		panic(err)
	}

	return RunExecute(client, stateHash, runAddress, util.HASH, proxyHash, sessionArgsStr, proxyHash, BASIC_FEE, protocolVersion)
}

func RunBond(client ipc.ExecutionEngineServiceClient, stateHash []byte, runAddress []byte,
	amount string,
	proxyHash []byte, protocolVersion *state.ProtocolVersion) (resultStateHash []byte, bonds []*ipc.Bond) {
	sessionArgs := []*consensus.Deploy_Arg{
		&consensus.Deploy_Arg{
			Name: "method",
			Value: &state.CLValueInstance{
				ClType: &state.CLType{Variants: &state.CLType_SimpleType{SimpleType: state.CLType_STRING}},
				Value: &state.CLValueInstance_Value{
					Value: &state.CLValueInstance_Value_StrValue{
						StrValue: "bond"}}}},
		&consensus.Deploy_Arg{
			Name: "amount",
			Value: &state.CLValueInstance{
				ClType: &state.CLType{Variants: &state.CLType_SimpleType{SimpleType: state.CLType_U512}},
				Value: &state.CLValueInstance_Value{
					Value: &state.CLValueInstance_Value_U512{
						U512: &state.CLValueInstance_U512{
							Value: amount}}}}}}

	sessionArgsStr, err := util.DeployArgsToJsonString(sessionArgs)
	if err != nil {
		panic(err)
	}

	return RunExecute(client, stateHash, runAddress, util.HASH, proxyHash, sessionArgsStr, proxyHash, BASIC_FEE, protocolVersion)
}

func RunUnbond(client ipc.ExecutionEngineServiceClient, stateHash []byte, runAddress []byte,
	amount string,
	proxyHash []byte, protocolVersion *state.ProtocolVersion) (resultStateHash []byte, bonds []*ipc.Bond) {
	sessionArgs := []*consensus.Deploy_Arg{
		&consensus.Deploy_Arg{
			Name: "method",
			Value: &state.CLValueInstance{
				ClType: &state.CLType{Variants: &state.CLType_SimpleType{SimpleType: state.CLType_STRING}},
				Value: &state.CLValueInstance_Value{
					Value: &state.CLValueInstance_Value_StrValue{
						StrValue: "unbond"}}}},
		&consensus.Deploy_Arg{
			Name: "amount",
			Value: &state.CLValueInstance{
				ClType: &state.CLType{Variants: &state.CLType_OptionType{OptionType: &state.CLType_Option{Inner: &state.CLType{Variants: &state.CLType_SimpleType{SimpleType: state.CLType_U512}}}}},
				Value: &state.CLValueInstance_Value{
					Value: &state.CLValueInstance_Value_OptionValue{
						OptionValue: &state.CLValueInstance_Option{
							Value: &state.CLValueInstance_Value{
								Value: &state.CLValueInstance_Value_U512{
									U512: &state.CLValueInstance_U512{
										Value: amount}}}}}}}}}

	sessionArgsStr, err := util.DeployArgsToJsonString(sessionArgs)
	if err != nil {
		panic(err)
	}

	return RunExecute(client, stateHash, runAddress, util.HASH, proxyHash, sessionArgsStr, proxyHash, BASIC_FEE, protocolVersion)
}

func RunDelegate(client ipc.ExecutionEngineServiceClient, stateHash []byte, runAddress []byte,
	validator []byte, amount string,
	proxyHash []byte, protocolVersion *state.ProtocolVersion) (resultStateHash []byte, bonds []*ipc.Bond) {
	sessionArgs := []*consensus.Deploy_Arg{
		&consensus.Deploy_Arg{
			Name: "method",
			Value: &state.CLValueInstance{
				ClType: &state.CLType{Variants: &state.CLType_SimpleType{SimpleType: state.CLType_STRING}},
				Value: &state.CLValueInstance_Value{
					Value: &state.CLValueInstance_Value_StrValue{
						StrValue: "delegate"}}}},
		&consensus.Deploy_Arg{
			Name: "validator",
			Value: &state.CLValueInstance{
				ClType: &state.CLType{Variants: &state.CLType_ListType{ListType: &state.CLType_List{Inner: &state.CLType{Variants: &state.CLType_SimpleType{SimpleType: state.CLType_U8}}}}},
				Value: &state.CLValueInstance_Value{
					Value: &state.CLValueInstance_Value_BytesValue{
						BytesValue: validator}}}},
		&consensus.Deploy_Arg{
			Name: "amount",
			Value: &state.CLValueInstance{
				ClType: &state.CLType{Variants: &state.CLType_SimpleType{SimpleType: state.CLType_U512}},
				Value: &state.CLValueInstance_Value{
					Value: &state.CLValueInstance_Value_U512{
						U512: &state.CLValueInstance_U512{
							Value: amount}}}}}}

	sessionArgsStr, err := util.DeployArgsToJsonString(sessionArgs)
	if err != nil {
		panic(err)
	}

	return RunExecute(client, stateHash, runAddress, util.HASH, proxyHash, sessionArgsStr, proxyHash, ADVANCED_FEE, protocolVersion)
}

func RunUndelegate(client ipc.ExecutionEngineServiceClient, stateHash []byte, runAddress []byte,
	validator []byte, amount string,
	proxyHash []byte, protocolVersion *state.ProtocolVersion) (resultStateHash []byte, bonds []*ipc.Bond) {
	sessionArgs := []*consensus.Deploy_Arg{
		&consensus.Deploy_Arg{
			Name: "method",
			Value: &state.CLValueInstance{
				ClType: &state.CLType{Variants: &state.CLType_SimpleType{SimpleType: state.CLType_STRING}},
				Value: &state.CLValueInstance_Value{
					Value: &state.CLValueInstance_Value_StrValue{
						StrValue: "undelegate"}}}},
		&consensus.Deploy_Arg{
			Name: "validator",
			Value: &state.CLValueInstance{
				ClType: &state.CLType{Variants: &state.CLType_ListType{ListType: &state.CLType_List{Inner: &state.CLType{Variants: &state.CLType_SimpleType{SimpleType: state.CLType_U8}}}}},
				Value: &state.CLValueInstance_Value{
					Value: &state.CLValueInstance_Value_BytesValue{
						BytesValue: validator}}}},
		&consensus.Deploy_Arg{
			Name: "amount",
			Value: &state.CLValueInstance{
				ClType: &state.CLType{Variants: &state.CLType_OptionType{OptionType: &state.CLType_Option{Inner: &state.CLType{Variants: &state.CLType_SimpleType{SimpleType: state.CLType_U512}}}}},
				Value: &state.CLValueInstance_Value{
					Value: &state.CLValueInstance_Value_OptionValue{
						OptionValue: &state.CLValueInstance_Option{
							Value: &state.CLValueInstance_Value{
								Value: &state.CLValueInstance_Value_U512{
									U512: &state.CLValueInstance_U512{
										Value: amount}}}}}}}}}

	sessionArgsStr, err := util.DeployArgsToJsonString(sessionArgs)
	if err != nil {
		panic(err)
	}

	return RunExecute(client, stateHash, runAddress, util.HASH, proxyHash, sessionArgsStr, proxyHash, ADVANCED_FEE, protocolVersion)
}

func RunRedelegate(client ipc.ExecutionEngineServiceClient, stateHash []byte, runAddress []byte,
	srcValidator []byte, destValidator []byte, amount string,
	proxyHash []byte, protocolVersion *state.ProtocolVersion) (resultStateHash []byte, bonds []*ipc.Bond) {
	sessionArgs := []*consensus.Deploy_Arg{
		&consensus.Deploy_Arg{
			Name: "method",
			Value: &state.CLValueInstance{
				ClType: &state.CLType{Variants: &state.CLType_SimpleType{SimpleType: state.CLType_STRING}},
				Value: &state.CLValueInstance_Value{
					Value: &state.CLValueInstance_Value_StrValue{
						StrValue: "redelegate"}}}},
		&consensus.Deploy_Arg{
			Name: "src",
			Value: &state.CLValueInstance{
				ClType: &state.CLType{Variants: &state.CLType_ListType{ListType: &state.CLType_List{Inner: &state.CLType{Variants: &state.CLType_SimpleType{SimpleType: state.CLType_U8}}}}},
				Value: &state.CLValueInstance_Value{
					Value: &state.CLValueInstance_Value_BytesValue{
						BytesValue: srcValidator}}}},
		&consensus.Deploy_Arg{
			Name: "dest",
			Value: &state.CLValueInstance{
				ClType: &state.CLType{Variants: &state.CLType_ListType{ListType: &state.CLType_List{Inner: &state.CLType{Variants: &state.CLType_SimpleType{SimpleType: state.CLType_U8}}}}},
				Value: &state.CLValueInstance_Value{
					Value: &state.CLValueInstance_Value_BytesValue{
						BytesValue: destValidator}}}},
		&consensus.Deploy_Arg{
			Name: "amount",
			Value: &state.CLValueInstance{
				ClType: &state.CLType{Variants: &state.CLType_OptionType{OptionType: &state.CLType_Option{Inner: &state.CLType{Variants: &state.CLType_SimpleType{SimpleType: state.CLType_U512}}}}},
				Value: &state.CLValueInstance_Value{
					Value: &state.CLValueInstance_Value_OptionValue{
						OptionValue: &state.CLValueInstance_Option{
							Value: &state.CLValueInstance_Value{
								Value: &state.CLValueInstance_Value_U512{
									U512: &state.CLValueInstance_U512{
										Value: amount}}}}}}}}}
	sessionArgsStr, err := util.DeployArgsToJsonString(sessionArgs)
	if err != nil {
		panic(err)
	}

	return RunExecute(client, stateHash, runAddress, util.HASH, proxyHash, sessionArgsStr, proxyHash, ADVANCED_FEE, protocolVersion)
}

func RunVote(client ipc.ExecutionEngineServiceClient, stateHash []byte, runAddress []byte,
	hash []byte, amount string,
	proxyHash []byte, protocolVersion *state.ProtocolVersion) (resultStateHash []byte, bonds []*ipc.Bond) {
	sessionArgs := []*consensus.Deploy_Arg{
		&consensus.Deploy_Arg{
			Name: "method",
			Value: &state.CLValueInstance{
				ClType: &state.CLType{Variants: &state.CLType_SimpleType{SimpleType: state.CLType_STRING}},
				Value: &state.CLValueInstance_Value{
					Value: &state.CLValueInstance_Value_StrValue{
						StrValue: "vote"}}}},
		&consensus.Deploy_Arg{
			Name: "hash",
			Value: &state.CLValueInstance{
				ClType: &state.CLType{Variants: &state.CLType_SimpleType{SimpleType: state.CLType_KEY}},
				Value: &state.CLValueInstance_Value{
					Value: &state.CLValueInstance_Value_Key{
						Key: &state.Key{Value: &state.Key_Hash_{Hash: &state.Key_Hash{Hash: hash}}}}}}},
		&consensus.Deploy_Arg{
			Name: "amount",
			Value: &state.CLValueInstance{
				ClType: &state.CLType{Variants: &state.CLType_SimpleType{SimpleType: state.CLType_U512}},
				Value: &state.CLValueInstance_Value{
					Value: &state.CLValueInstance_Value_U512{
						U512: &state.CLValueInstance_U512{
							Value: amount}}}}}}

	sessionArgsStr, err := util.DeployArgsToJsonString(sessionArgs)
	if err != nil {
		panic(err)
	}

	return RunExecute(client, stateHash, runAddress, util.HASH, proxyHash, sessionArgsStr, proxyHash, ADVANCED_FEE, protocolVersion)
}

func RunUnvote(client ipc.ExecutionEngineServiceClient, stateHash []byte, runAddress []byte,
	hash []byte, amount string,
	proxyHash []byte, protocolVersion *state.ProtocolVersion) (resultStateHash []byte, bonds []*ipc.Bond) {
	sessionArgs := []*consensus.Deploy_Arg{
		&consensus.Deploy_Arg{
			Name: "method",
			Value: &state.CLValueInstance{
				ClType: &state.CLType{Variants: &state.CLType_SimpleType{SimpleType: state.CLType_STRING}},
				Value: &state.CLValueInstance_Value{
					Value: &state.CLValueInstance_Value_StrValue{
						StrValue: "unvote"}}}},
		&consensus.Deploy_Arg{
			Name: "hash",
			Value: &state.CLValueInstance{
				ClType: &state.CLType{Variants: &state.CLType_SimpleType{SimpleType: state.CLType_KEY}},
				Value: &state.CLValueInstance_Value{
					Value: &state.CLValueInstance_Value_Key{
						Key: &state.Key{Value: &state.Key_Hash_{Hash: &state.Key_Hash{Hash: hash}}}}}}},
		&consensus.Deploy_Arg{
			Name: "amount",
			Value: &state.CLValueInstance{
				ClType: &state.CLType{Variants: &state.CLType_OptionType{OptionType: &state.CLType_Option{Inner: &state.CLType{Variants: &state.CLType_SimpleType{SimpleType: state.CLType_U512}}}}},
				Value: &state.CLValueInstance_Value{
					Value: &state.CLValueInstance_Value_OptionValue{
						OptionValue: &state.CLValueInstance_Option{
							Value: &state.CLValueInstance_Value{
								Value: &state.CLValueInstance_Value_U512{
									U512: &state.CLValueInstance_U512{
										Value: amount}}}}}}}}}

	sessionArgsStr, err := util.DeployArgsToJsonString(sessionArgs)
	if err != nil {
		panic(err)
	}

	return RunExecute(client, stateHash, runAddress, util.HASH, proxyHash, sessionArgsStr, proxyHash, ADVANCED_FEE, protocolVersion)
}

func RunStep(client ipc.ExecutionEngineServiceClient, stateHash []byte, runAddress []byte,
	proxyHash []byte, protocolVersion *state.ProtocolVersion) (resultStateHash []byte, bonds []*ipc.Bond) {

	res, err := client.Step(
		context.TODO(),
		&ipc.StepRequest{
			ParentStateHash: stateHash,
			BlockTime:       uint64(time.Now().Unix()),
			ProtocolVersion: protocolVersion,
		},
	)
	if err != nil {
		panic(err)
	}

	stateHash, bonds, errMessage := grpc.Commit(client, res.GetSuccess().GetPostStateHash(), res.GetSuccess().GetEffect().TransformMap, protocolVersion)
	if errMessage != "" {
		panic(errMessage)
	}
	printCommitResult(stateHash, bonds)

	return stateHash, bonds
}

func RunClaimCommission(client ipc.ExecutionEngineServiceClient, stateHash []byte, runAddress []byte,
	proxyHash []byte, protocolVersion *state.ProtocolVersion) (resultStateHash []byte, bonds []*ipc.Bond) {
	sessionArgs := []*consensus.Deploy_Arg{
		&consensus.Deploy_Arg{
			Name: "method",
			Value: &state.CLValueInstance{
				ClType: &state.CLType{Variants: &state.CLType_SimpleType{SimpleType: state.CLType_STRING}},
				Value: &state.CLValueInstance_Value{
					Value: &state.CLValueInstance_Value_StrValue{
						StrValue: "claim_commission"}}}}}

	sessionArgsStr, err := util.DeployArgsToJsonString(sessionArgs)
	if err != nil {
		panic(err)
	}

	return RunExecute(client, stateHash, runAddress, util.HASH, proxyHash, sessionArgsStr, proxyHash, ADVANCED_FEE, protocolVersion)
}

func RunClaimReward(client ipc.ExecutionEngineServiceClient, stateHash []byte, runAddress []byte,
	proxyHash []byte, protocolVersion *state.ProtocolVersion) (resultStateHash []byte, bonds []*ipc.Bond) {
	sessionArgs := []*consensus.Deploy_Arg{
		&consensus.Deploy_Arg{
			Name: "method",
			Value: &state.CLValueInstance{
				ClType: &state.CLType{Variants: &state.CLType_SimpleType{SimpleType: state.CLType_STRING}},
				Value: &state.CLValueInstance_Value{
					Value: &state.CLValueInstance_Value_StrValue{
						StrValue: "claim_reward"}}}}}

	sessionArgsStr, err := util.DeployArgsToJsonString(sessionArgs)
	if err != nil {
		panic(err)
	}

	return RunExecute(client, stateHash, runAddress, util.HASH, proxyHash, sessionArgsStr, proxyHash, ADVANCED_FEE, protocolVersion)
}

func RunExecute(client ipc.ExecutionEngineServiceClient, stateHash []byte,
	fromAddress []byte,
	sessionType util.ContractType, sessionData []byte, sessionArgsStr string,
	proxyHash []byte, fee string,
	protocolVersion *state.ProtocolVersion) (resultStateHash []byte, bonds []*ipc.Bond) {
	timestamp := time.Now().Unix()

	paymentArgsStr := GetPaymentArgsJson(fee)

	deploy, _ := util.MakeDeploy(fromAddress, sessionType, sessionData, sessionArgsStr, util.HASH, proxyHash, paymentArgsStr, uint64(10), timestamp, CHAIN_NAME)
	deploys := util.MakeInitDeploys()
	deploys = util.AddDeploy(deploys, deploy)

	res, err := grpc.Execute(client, stateHash, timestamp, deploys, protocolVersion)
	effect, err := executeErrorHandler(res)
	if err != nil {
		panic(err)
	}

	stateHash, bonds, errMessage := grpc.Commit(client, stateHash, effect, protocolVersion)
	if errMessage != "" {
		panic(errMessage)
	}
	printCommitResult(stateHash, bonds)

	return stateHash, bonds
}

func RunQuery(client ipc.ExecutionEngineServiceClient, stateHash []byte, types string, value []byte, path []string, protocolVersion *state.ProtocolVersion) storedvalue.StoredValue {
	var storedValue storedvalue.StoredValue
	queryResult, errMessage := grpc.Query(client, stateHash, types, value, path, protocolVersion)
	if errMessage != "" {
		panic(errMessage)
	}
	storedValue, err, _ := storedValue.FromBytes(queryResult)
	if err != nil {
		panic(err)
	}

	return storedValue
}

func executeErrorHandler(r *ipc.ExecuteResponse) (effects []*transforms.TransformEntry, err error) {
	switch r.GetResult().(type) {
	case *ipc.ExecuteResponse_Success:
		for _, res := range r.GetSuccess().GetDeployResults() {
			switch res.GetExecutionResult().GetError().GetValue().(type) {
			case *ipc.DeployError_GasError:
				err = fmt.Errorf("DeployError_GasError")
			case *ipc.DeployError_ExecError:
				err = fmt.Errorf("DeployError_ExecError : %s", res.GetExecutionResult().GetError().GetExecError().GetMessage())
			default:
				effects = append(effects, res.GetExecutionResult().GetEffects().GetTransformMap()...)
			}

		}
	case *ipc.ExecuteResponse_MissingParent:
		err = fmt.Errorf("Missing parentstate : %s", util.EncodeToHexString(r.GetMissingParent().GetHash()))
	default:
		err = fmt.Errorf("Unknown result : %s", r.String())
	}

	return effects, err
}

func printCommitResult(stateHash []byte, bonds []*ipc.Bond) {
	println("State hash : " + hex.EncodeToString(stateHash))
	for _, bond := range bonds {
		println(hex.EncodeToString(bond.ValidatorPublicKey) + " : " + bond.GetStake().GetValue())
	}
	println()
}
