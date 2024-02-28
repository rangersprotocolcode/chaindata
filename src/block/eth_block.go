package block

import (
	common2 "com.tuntun.rangers/service/chaindata/src/common"
	"com.tuntun.rangers/service/chaindata/src/middleware/mysql"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strconv"
	"strings"
	"time"
)

var big8 = big.NewInt(8)

func (self *ethModule) processBlock() {
	self.lock.Lock()
	defer func() {
		self.blockTimer.Reset(blockWait)
		self.lock.Unlock()
	}()

	client := self.getClient()
	if nil == client {
		self.logger.Errorf("fail to get client")
		return
	}

	header := self.getHeader(client)
	if nil == header {
		return
	}

	topHeight := header.Number.Int64()
	self.logger.Debugf("get header. height: %d, hash: %s", topHeight, header.Hash().String())

	last := topHeight - 12
	if last == self.lastBlockFromClient {
		self.count++
		if self.count == 6 {
			self.logger.Warnf("get max retry count, close client and reconnect")
			self.closeClient()
			return
		}
	} else {
		self.count = 0
		self.lastBlockFromClient = last
	}

	if last <= self.lastBlock {
		self.logger.Infof("no need to processBlock, remote: %d, local: %d", last, self.lastBlock)
		return
	}

	self.logger.Debugf("start getting blocks. from %d to %d", self.lastBlock, last)
	for i := self.lastBlock + 1; i < last; i++ {
		block := self.getBlock(i, client)
		if nil == block {
			return
		}

		isEvent := false
		addressList := make([]*common.Address, 0)
		blockhash := block.Hash().String()
		ts := strconv.FormatUint(block.Time(), 10)
		for _, tx := range block.Transactions() {
			if tx == nil || tx.To() == nil {
				continue
			}

			toAddr := tx.To().String()
			if _, ok := self.contracts[strings.ToLower(toAddr)]; ok {
				isEvent = true
				addressList = append(addressList, tx.To())
				continue
			}

			txHash := tx.Hash().String()

			fmt.Println(tx.Type())

			v, r, s := tx.RawSignatureValues()

			vInt64 := v.Uint64()
			var V *big.Int
			if 0 == vInt64 || 1 == vInt64 || 27 == vInt64 || 28 == vInt64 {
				V = v
			} else {
				V = new(big.Int).Sub(v, new(big.Int).Mul(tx.ChainId(), big.NewInt(2)))
				V.Sub(V, big8)
			}

			fromAddr, err := common2.RecoverPlain(tx.Hash(), r, s, V)
			if nil != err {
				self.logger.Errorf("fail to calc fromAddr, txhash: %s", txHash)
			} else {
				mysql.InsertLogs(i, self.chainId, blockhash, ts, txHash, fromAddr.String(), toAddr, tx.Value().String(), "")
			}

		}

		if isEvent {
			self.processEvent(i, addressList)
		}

		self.lastBlock = i
		self.updateHeight()
	}
}

func (self *ethModule) getBlock(i int64, client *ethclient.Client) *types.Block {
	var (
		block *types.Block
		err   error
	)
	ch := make(chan byte, 1)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go func() {
		block, err = client.BlockByNumber(ctx, big.NewInt(i))
		ch <- 0
	}()

	select {
	case <-ch:
		if err != nil {
			self.logger.Errorf("fail to get block, err: %s, close client and reconnect", err)
			self.closeClient()
			return nil
		}

		return block
	case <-ctx.Done():
		self.logger.Errorf("fail to get blockHeader, err: timeout, close client and reconnect")
		self.closeClient()
		return nil
	}
}
