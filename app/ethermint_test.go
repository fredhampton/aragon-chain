package app

import (
	"encoding/base32"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/codec"

	abci "github.com/tendermint/tendermint/abci/types"
)

func TestEthermintAppExport(t *testing.T) {
	db := dbm.NewMemDB()
	app := NewEthermintApp(log.NewTMLogger(log.NewSyncWriter(os.Stdout)), db, true, 0)

	genesisState := ModuleBasics.DefaultGenesis()
	stateBytes, err := codec.MarshalJSONIndent(app.cdc, genesisState)
	require.NoError(t, err)

	// Initialize the chain
	app.InitChain(
		abci.RequestInitChain{
			Validators:    []abci.ValidatorUpdate{},
			AppStateBytes: stateBytes,
		},
	)
	app.Commit()

	// Making a new app object with the db, so that initchain hasn't been called
	app2 := NewEthermintApp(log.NewTMLogger(log.NewSyncWriter(os.Stdout)), db, true, 0)
	_, _, err = app2.ExportAppStateAndValidators(false, []string{})
	require.NoError(t, err, "ExportAppStateAndValidators should not have an error")

}
func TestThis(t *testing.T) {
	testv := base32.NewEncoding("abcdefghijklmnopqrstuvwxyz234567")
	require.Equal(t, testv.WithPadding(-1).EncodeToString([]byte{0x00, 0x01, 0x2, 0x0}), []byte("0"))
}
