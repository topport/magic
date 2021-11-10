package crypto_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/topport/magic/pkg/crypto"
)

func TestSigKeypairFromHDMnemonic(t *testing.T) {
	mnemonic := "joke basic have athlete nurse tank snow uniform busy rural depend recall dinosaur glory elegant"
	keypair, err := crypto.SigKeypairFromHDMnemonic(mnemonic, 0)
	require.NoError(t, err)
	require.Equal(t, "1c6e8d3d4e32f3c8e0bf1295a397ed5cda700888f8d289d602b15fdfd05a3f82", keypair.SigningPrivateKey.Hex())
}
