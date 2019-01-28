package handler

import (
	"github.com/c3systems/go-substrate/client/p2p"
	clienttypes "github.com/c3systems/go-substrate/client/types"
)

// note: ensure the struct implements the interface
var _ InterfaceHandler = (*TransactionsHandler)(nil)

// TransactionsHandler implements the transactions handler
type TransactionsHandler struct{}

// Func handles incoming transactions messages
// TODO ...
// TODO Propagate
func (t *TransactionsHandler) Func(p p2p.InterfaceP2P, pr clienttypes.InterfacePeer, msg clienttypes.InterfaceMessage) error {
	//var msgStrBytes []byte
	//if err := msg.Unmarshal(msgBytes); err != nil {
	//logger.Errorf("[handler] err unmarshalling transaction message\n%v", err)
	//return err
	//}

	//logger.Infof("%v Transaction: %v", pr.Cfg().ShortID, string(msgStrBytes))

	return nil
}

// Type returns the func enum
func (t *TransactionsHandler) Type() FuncEnum {
	return Transactions
}