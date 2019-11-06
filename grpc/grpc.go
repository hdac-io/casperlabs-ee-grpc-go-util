// Package grpc 는 Casperlabs의 Execution Engine의 GRPC Client 모듈을 정의한 모듈이다.
package grpc

import (
	"context"
	"time"

	"github.com/hdac-io/casperlabs-ee-grpc-go-util/protobuf/io/casperlabs/casper/consensus/state"
	"github.com/hdac-io/casperlabs-ee-grpc-go-util/protobuf/io/casperlabs/ipc"
	"github.com/hdac-io/casperlabs-ee-grpc-go-util/protobuf/io/casperlabs/ipc/transforms"

	"github.com/hdac-io/casperlabs-ee-grpc-go-util/util"

	"google.golang.org/grpc"
)

// Connect 은 Casperlabs의 Execution Engine의 unix socket으로 연결하는 함수.
func Connect(path string) ipc.ExecutionEngineServiceClient {
	path = `unix:////` + path

	conn, e := grpc.Dial(path, grpc.WithInsecure())
	if e != nil {
		panic(e)
	}

	client := ipc.NewExecutionEngineServiceClient(conn)

	return client
}

// RunGenensis 는 Execution Engine을 시작할 때 Genensis정보를 초기화하는 함수.
//
// genensis address, initial motes, timestamp, mintCode, posCode, validator Account 정보를 파라미터로 받아
// RunGenensis 후 변경될 state hash와 effects를 return 받는다.
func RunGenensis(
	client ipc.ExecutionEngineServiceClient,
	genesisAddress string,
	strInitialMotes string,
	timestamp int64,
	mintCode []byte,
	posCode []byte,
	validators map[string]string,
	protocolVersion *state.ProtocolVersion) (parentStateHash []byte, effects []*transforms.TransformEntry) {
	initialMotes := &state.BigInt{Value: strInitialMotes, BitWidth: uint32(512)}

	deployMintCode := &ipc.DeployCode{Code: mintCode}

	deployPosCode := &ipc.DeployCode{Code: posCode}

	genesisValidators := []*ipc.Bond{}

	for address, stake := range validators {
		genesisValidators = append(genesisValidators,
			&ipc.Bond{ValidatorPublicKey: util.DecodeHexString(address), Stake: &state.BigInt{Value: stake, BitWidth: 512}})
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	r, err := client.RunGenesis(
		ctx,
		&ipc.GenesisRequest{
			Address:           util.DecodeHexString(genesisAddress),
			InitialMotes:      initialMotes,
			Timestamp:         uint64(timestamp),
			MintCode:          deployMintCode,
			ProofOfStakeCode:  deployPosCode,
			GenesisValidators: genesisValidators,
			ProtocolVersion:   protocolVersion})
	if err != nil {
		panic(err)
	}

	genesisResult := r.GetSuccess()

	return genesisResult.PoststateHash, genesisResult.GetEffect().GetTransformMap()
}

// RunGenensisWithChainSpec 는 Execution Engine을 시작할 때 Genensis정보를 chain에 떄라 초기화하는 함수.
//
// chain name, timestamp, mintInstallCode, posInstallCode, validator Account, cost 정보를 파라미터로 받아
// RunGenensisWithChainSpec 후 변경될 state hash와 effects를 return 받는다.
func RunGenensisWithChainSpec(client ipc.ExecutionEngineServiceClient,
	name string,
	timestamp int64,
	protocolVersion *state.ProtocolVersion,
	mintInstallCode []byte,
	posInstallCode []byte,
	mapAccounts map[string][]string,
	mapCosts map[string]uint32) (parentStateHash []byte, effects []*transforms.TransformEntry, errMessage string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	accounts := []*ipc.ChainSpec_GenesisAccount{}

	for address, strAccount := range mapAccounts {
		accounts = append(accounts, &ipc.ChainSpec_GenesisAccount{PublicKey: util.DecodeHexString(address), Balance: &state.BigInt{Value: strAccount[0], BitWidth: 512}, BondedAmount: &state.BigInt{Value: strAccount[1], BitWidth: 512}})
	}

	costs := &ipc.ChainSpec_CostTable{
		Wasm: &ipc.ChainSpec_CostTable_WasmCosts{
			Regular:        mapCosts["regular"],
			Div:            mapCosts["div-multiplier"],
			Mul:            mapCosts["mul-multiplier"],
			Mem:            mapCosts["mem-multiplier"],
			InitialMem:     mapCosts["mem-initial-pages"],
			GrowMem:        mapCosts["mem-grow-per-page"],
			Memcpy:         mapCosts["mem-copy-per-byte"],
			MaxStackHeight: mapCosts["max-stack-height"],
			OpcodesMul:     mapCosts["opcodes-multiplier"],
			OpcodesDiv:     mapCosts["opcodes-divisor"]}}

	r, err := client.RunGenesisWithChainspec(
		ctx,
		&ipc.ChainSpec_GenesisConfig{
			Name:            name,
			Timestamp:       uint64(timestamp),
			ProtocolVersion: protocolVersion,
			MintInstaller:   mintInstallCode,
			PosInstaller:    posInstallCode,
			Accounts:        accounts,
			Costs:           costs})

	if err != nil {
		errMessage = err.Error()
	}

	switch r.GetResult().(type) {
	case *ipc.GenesisResponse_Success:
		parentStateHash = r.GetSuccess().GetPoststateHash()
		effects = r.GetSuccess().GetEffect().GetTransformMap()
	case *ipc.GenesisResponse_FailedDeploy:
		errMessage += r.GetFailedDeploy().GetMessage()
	}

	return parentStateHash, effects, errMessage
}

// Commit 은 Execute한 effects를 적용시킬 때 사용하는 함수.
//
// State Hash, Execute한 effects를 파라미터로 받아,
// Commit 후 state hash 와 현재 Bonding 된 validator의 정보를 return 받는다.
func Commit(client ipc.ExecutionEngineServiceClient,
	prestateHash []byte,
	effects []*transforms.TransformEntry,
	protocolVersion *state.ProtocolVersion) (postStateHash []byte, validators []*ipc.Bond, errMessage string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	r, err := client.Commit(
		ctx,
		&ipc.CommitRequest{
			PrestateHash:    prestateHash,
			Effects:         effects,
			ProtocolVersion: protocolVersion})
	if err != nil {
		errMessage = err.Error()
	}

	switch r.GetResult().(type) {
	case *ipc.CommitResponse_Success:
		postStateHash = r.GetSuccess().GetPoststateHash()
		validators = r.GetSuccess().GetBondedValidators()
	case *ipc.CommitResponse_MissingPrestate:
		errMessage = "Missing prestate : " + util.EncodeToHexString(r.GetMissingPrestate().GetHash())
	case *ipc.CommitResponse_KeyNotFound:
		errMessage += "Key not Found "
		var hashValue []byte
		switch r.GetKeyNotFound().GetValue().(type) {
		case *state.Key_Address_:
			errMessage += "(Address)"
			hashValue = r.GetKeyNotFound().GetAddress().GetAccount()
		case *state.Key_Hash_:
			errMessage += "(Hash)"
			hashValue = r.GetKeyNotFound().GetHash().GetHash()
		case *state.Key_Uref:
			errMessage += "(Uref)"
			hashValue = r.GetKeyNotFound().GetUref().GetUref()
		case *state.Key_Local_:
			errMessage += "(Local)"
			hashValue = r.GetKeyNotFound().GetLocal().GetHash()
		}
		errMessage += " : " + util.EncodeToHexString(hashValue)
	case *ipc.CommitResponse_TypeMismatch:
		errMessage += "Type missmatch : expected (" + r.GetTypeMismatch().GetExpected() + "), but (" + r.GetTypeMismatch().GetFound() + ")"
	case *ipc.CommitResponse_FailedTransform:
		errMessage += "Failed transform : " + r.GetFailedTransform().GetMessage()
	}

	return postStateHash, validators, errMessage
}

// Validate 는 WasmCode를 받아 정상 Wasm파일인지 확인해주는 함수.
//
// byte array의 wasm을 받아 정상 wasm 파일인지 true, false를 return 해준다.
func Validate(client ipc.ExecutionEngineServiceClient, wasmCode []byte, protocolVersion *state.ProtocolVersion) (result bool, errMessage string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := client.Validate(
		ctx,
		&ipc.ValidateRequest{
			WasmCode:        wasmCode,
			ProtocolVersion: protocolVersion})
	if err != nil {
		errMessage = err.Error()
	}

	switch r.GetResult().(type) {
	case *ipc.ValidateResponse_Success:
		result = true
	case *ipc.ValidateResponse_Failure:
		result = false
		errMessage = r.GetFailure()
	}

	return result, errMessage
}

// Query 는 특정 state 에서 해당 Key의 path에 대한 정보를 조회해주는 함수.
//
// State hash, Key type, Key Data, path를 파라미터로 받아
// Query 후 결과를 return 해준다.
func Query(client ipc.ExecutionEngineServiceClient,
	stateHash []byte,
	keyType string,
	keyData string,
	path []string,
	protocolVersion *state.ProtocolVersion) (result *state.Value, errMessage string) {

	var key *state.Key
	switch keyType {
	case "address":
		key = &state.Key{Value: &state.Key_Address_{Address: &state.Key_Address{Account: util.DecodeHexString(keyData)}}}
	case "local":
		key = &state.Key{Value: &state.Key_Local_{Local: &state.Key_Local{Hash: util.DecodeHexString(keyData)}}}
	case "uref":
		key = &state.Key{Value: &state.Key_Uref{Uref: &state.Key_URef{Uref: util.DecodeHexString(keyData)}}}
	case "hash":
		key = &state.Key{Value: &state.Key_Hash_{Hash: &state.Key_Hash{Hash: util.DecodeHexString(keyData)}}}
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	r, err := client.Query(
		ctx,
		&ipc.QueryRequest{
			StateHash:       stateHash,
			BaseKey:         key,
			Path:            path,
			ProtocolVersion: protocolVersion})
	if err != nil {
		errMessage = err.Error()
	}

	switch r.GetResult().(type) {
	case *ipc.QueryResponse_Success:
		result = r.GetSuccess()
	case *ipc.QueryResponse_Failure:
		errMessage = r.GetFailure()
	}

	return result, errMessage
}

// Execute 는 deploys를 실행할떄의 effects를 받아오는 함수.
//
// state hash, timestamp, deploys를 파라미로 받아
// Execute 후 적용하여야할 effects를 return 해준다.
func Execute(client ipc.ExecutionEngineServiceClient,
	parentStateHash []byte,
	int64timestamp int64,
	deploys []*ipc.DeployItem,
	protocolVersion *state.ProtocolVersion) (effects []*transforms.TransformEntry, errMessage string) {

	timestamp := uint64(int64timestamp)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	r, err := client.Execute(
		ctx,
		&ipc.ExecuteRequest{
			ParentStateHash: parentStateHash,
			BlockTime:       timestamp,
			Deploys:         deploys,
			ProtocolVersion: protocolVersion})
	if err != nil {
		errMessage = err.Error()
	}

	switch r.GetResult().(type) {
	case *ipc.ExecuteResponse_Success:
		for _, res := range r.GetSuccess().GetDeployResults() {
			effects = append(effects, res.GetExecutionResult().GetEffects().GetTransformMap()...)
		}
	case *ipc.ExecuteResponse_MissingParent:
		errMessage = "Missing parentstate : " + util.EncodeToHexString(r.GetMissingParent().GetHash())
	}

	return effects, errMessage
}

// Upgrade 는 Wasm 코드나 Cost를 변경하여 Protocol Version을 Upgrade할 때 활용
//
// State hash, 변경할 Insatll Wasm코드, Cost, 현재 protocol version, 다음 protocol version을 파라미터로 받으며,
// Install wasm 코드를 변경할지, Cost를 변경할지는 옵션으로 가능하며 Upgrade 를 통해 변경한 후
// 변경될 state hash, effects를 return 해준다.
func Upgrade(client ipc.ExecutionEngineServiceClient,
	parentStateHash []byte,
	wasmCode []byte,
	mapCosts map[string]uint32,
	currentProtocolVersion *state.ProtocolVersion,
	nextProtocolVersion *state.ProtocolVersion) (postStateHash []byte, effects []*transforms.TransformEntry, errMessage string) {

	costs := &ipc.ChainSpec_CostTable{
		Wasm: &ipc.ChainSpec_CostTable_WasmCosts{
			Regular:        mapCosts["regular"],
			Div:            mapCosts["div-multiplier"],
			Mul:            mapCosts["mul-multiplier"],
			Mem:            mapCosts["mem-multiplier"],
			InitialMem:     mapCosts["mem-initial-pages"],
			GrowMem:        mapCosts["mem-grow-per-page"],
			Memcpy:         mapCosts["mem-copy-per-byte"],
			MaxStackHeight: mapCosts["max-stack-height"],
			OpcodesMul:     mapCosts["opcodes-multiplier"],
			OpcodesDiv:     mapCosts["opcodes-divisor"]}}

	upgradePoint := &ipc.ChainSpec_UpgradePoint{
		ActivationPoint:  &ipc.ChainSpec_ActivationPoint{Rank: uint64(1)},
		ProtocolVersion:  nextProtocolVersion,
		UpgradeInstaller: &ipc.DeployCode{Code: wasmCode},
		NewCosts:         costs}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	r, err := client.Upgrade(
		ctx,
		&ipc.UpgradeRequest{
			ParentStateHash: parentStateHash,
			UpgradePoint:    upgradePoint,
			ProtocolVersion: currentProtocolVersion})
	if err != nil {
		errMessage = err.Error()
	}

	switch r.GetResult().(type) {
	case *ipc.UpgradeResponse_Success:
		postStateHash = r.GetSuccess().GetPostStateHash()
		effects = r.GetSuccess().GetEffect().GetTransformMap()
	case *ipc.UpgradeResponse_FailedDeploy:
		errMessage = r.GetFailedDeploy().GetMessage()
	}

	return postStateHash, effects, errMessage
}

// QueryBlanace 는 address의 balance를 조회할 때 사용하는 함수.
//
// 조회할 state hash와 address를 파라미터로 받아, key를 address로 Query한다.
// name key에서 name이 mint인 uref를 추출하여 hex string로 변환하고 purse Id를 abi로 변환한 후 hex string으로 변환하여 붙인다.
// 해당 값을 blake2b256을 하면 local bytes 값이 추출된다. 이 값을 key를 local로 하여 Query한다.
// 받아온 uref값을 Key로 하여 Query하면 BigInt 형태의 blanace를 return 해준다.
func QueryBlanace(client ipc.ExecutionEngineServiceClient,
	stateHash []byte,
	address string,
	protocolVersion *state.ProtocolVersion) (balance string, errMessage string) {

	res, errMessage := Query(client, stateHash, "address", address, []string{}, protocolVersion)
	if errMessage != "" {
		return balance, errMessage
	}

	purseID := res.GetAccount().GetPurseId().GetUref()
	namedKeys := res.GetAccount().GetNamedKeys()
	var mintUref []byte
	for _, value := range namedKeys {
		if value.GetName() == "mint" {
			mintUref = value.GetKey().GetUref().GetUref()
			break
		}
	}

	localSrc := util.EncodeToHexString(mintUref) + util.EncodeToHexString(util.AbiBytesToBytes(purseID))
	localBytes := util.Blake2b256(util.DecodeHexString(localSrc))

	res, errMessage = Query(client, stateHash, "local", util.EncodeToHexString(localBytes), []string{}, protocolVersion)
	if errMessage != "" {
		return balance, errMessage
	}

	uref := res.GetKey().GetUref().GetUref()
	res, errMessage = Query(client, stateHash, "uref", util.EncodeToHexString(uref), []string{}, protocolVersion)
	if errMessage != "" {
		return balance, errMessage
	}

	return res.GetBigInt().GetValue(), errMessage
}
