package test

import (
	tronWallet "github.com/ranjbar-dev/tron-wallet"
	"testing"
)

// GenerateTronWallet test
func TestGenerateWallet(t *testing.T) {
	w := tronWallet.GenerateTronWallet(node)
	if w == nil {
		t.Errorf("GenerateTronWallet res was incorect, got: %q, want: %q.", w, "*tronWallet")
	}
	if len(w.PrivateKey) == 0 {
		t.Errorf("GenerateTronWallet PrivateKey was incorect, got: %q, want: %q.", w.PrivateKey, "valid PrivateKey")
	}
	if len(w.PublicKey) == 0 {
		t.Errorf("GenerateTronWallet PublicKey was incorect, got: %q, want: %q.", w.PublicKey, "valid PublicKey")
	}
	if len(w.Address) == 0 {
		t.Errorf("GenerateTronWallet Address was incorect, got: %q, want: %q.", w.Address, "valid Address")
	}
	if len(w.AddressBase58) == 0 {
		t.Errorf("GenerateTronWallet AddressBase58 was incorect, got: %q, want: %q.", w.AddressBase58, "valid AddressBase58")
	}
}

// CreateTronWallet test
func TestCreateWallet(t *testing.T) {
	_, err := tronWallet.CreateTronWallet(node, invalidPrivateKey)
	if err == nil {
		t.Errorf("CreateTronWallet error was incorect, got: %q, want: %q.", err, "not nil")
	}

	w, err := tronWallet.CreateTronWallet(node, validPrivateKey)
	if err != nil {
		t.Errorf("CreateTronWallet error was incorect, got: %q, want: %q.", err, "nil")
	}
	if len(w.PrivateKey) == 0 {
		t.Errorf("CreateTronWallet PrivateKey was incorect, got: %q, want: %q.", w.PrivateKey, "valid PrivateKey")
	}
	if len(w.PublicKey) == 0 {
		t.Errorf("CreateTronWallet PublicKey was incorect, got: %q, want: %q.", w.PublicKey, "valid PublicKey")
	}
	if len(w.Address) == 0 {
		t.Errorf("CreateTronWallet Address was incorect, got: %q, want: %q.", w.Address, "valid Address")
	}
	if len(w.AddressBase58) == 0 {
		t.Errorf("CreateTronWallet AddressBase58 was incorect, got: %q, want: %q.", w.AddressBase58, "valid AddressBase58")
	}
}

// PrivateKeyRCDSA test
func TestPrivateKeyRCDSA(t *testing.T) {
	w := wallet()

	_, err := w.PrivateKeyRCDSA()
	if err != nil {
		t.Errorf("PrivateKeyRCDSA error was incorect, got: %q, want: %q.", err, "nil")
	}
}

// PrivateKeyBytes test
func TestPrivateKeyBytes(t *testing.T) {
	w := wallet()

	bytes, err := w.PrivateKeyBytes()
	if err != nil {
		t.Errorf("PrivateKeyBytes error was incorect, got: %q, want: %q.", err, "nil")
	}
	if len(bytes) == 0 {
		t.Errorf("PrivateKeyBytes bytes len was incorect, got: %q, want: %q.", len(bytes), "more than 0")
	}
}

// Balance test
func TestBalance(t *testing.T) {
	w := wallet()

	_, err := w.Balance()
	if err != nil {
		t.Errorf("Balance error was incorect, got: %q, want: %q.", err, "nil")
	}
}

// BalanceTRC20 test
func TestBalanceTRC20(t *testing.T) {
	w := wallet()

	_, err := w.BalanceTRC20(token())
	if err != nil {
		t.Errorf("BalanceTRC20 error was incorect, got: %q, want: %q.", err, "nil")
	}
}

// Transfer test
func TestTransfer(t *testing.T) {
	w := wallet()

	_, err := w.Transfer(invalidToAddress, trxAmount)
	if err == nil {
		t.Errorf("Transfer error was incorect, got: %q, want: %q.", err, "not nil becuase to address is invalid")
	}

	txId, err := w.Transfer(validToAddress, trxAmount)
	if err != nil {
		t.Errorf("Transfer error was incorect, got: %q, want: %q.", err, "nil")
	}
	if len(txId) == 0 {
		t.Errorf("Transfer txId was incorect, got: %q, want: %q.", txId, "not nil")
	}
}

// TransferTRC20 test
func TestTransferTRC20(t *testing.T) {
	w := wallet()
	_t := token()

	_, err := w.TransferTRC20(_t, invalidToAddress, trc20Amount)
	if err == nil {
		t.Errorf("TestTransferTRC20 error was incorect, got: %q, want: %q.", err, "not nil becuase to address is invalid")
	}

	txId, err := w.TransferTRC20(_t, validToAddress, trc20Amount)
	if err != nil {
		t.Errorf("Transfer error was incorect, got: %q, want: %q.", err, "nil")
	}
	if len(txId) == 0 {
		t.Errorf("Transfer txId was incorect, got: %q, want: %q.", txId, "not nil")
	}
}
