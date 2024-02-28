// Copyright 2020 The RangersProtocol Authors
// This file is part of the RocketProtocol library.
//
// The RangersProtocol library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The RangersProtocol library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the RangersProtocol library. If not, see <http://www.gnu.org/licenses/>.

package mysql

//func SelectLogs(from, to uint64, contractAddresses []common.Address) []*types.Log {
//	if nil == mysqlDBLog {
//		return nil
//	}
//
//	sql := "select height,logindex, blockhash,txhash,contractaddress,topic,data FROM contractlogs WHERE (height>=? and height<=?) "
//	if 0 != len(contractAddresses) {
//		sql += "and( "
//		for _, contractAddress := range contractAddresses {
//			sql += " contractaddress = \"" + contractAddress.GetHexString() + "\" " + "or"
//		}
//		sql = sql[:len(sql)-2] + ")"
//	}
//
//	rows, err := mysqlDBLog.Query(sql, from, to)
//	if err != nil {
//		return nil
//	}
//	defer rows.Close()
//
//	result := make([]*types.Log, 0)
//	for rows.Next() {
//		var (
//			height, index                                   uint64
//			blockhash, txhash, contractaddress, topic, data string
//		)
//		err := rows.Scan(&height, &index, &blockhash, &txhash, &contractaddress, &topic, &data)
//		if err != nil {
//			logger.Errorf("scan failed, err: %v", err)
//			return nil
//		}
//
//		log := types.Log{
//			Address:     common.HexToAddress(contractaddress),
//			Data:        common.FromHex(data),
//			TxHash:      common.HexToHash(txhash),
//			BlockHash:   common.HexToHash(blockhash),
//			BlockNumber: height,
//			Index:       uint(index),
//		}
//
//		json.Unmarshal(utility.StrToBytes(topic), &log.Topics)
//		result = append(result, &log)
//	}
//
//	return result
//}
//
//func SelectLogsByHash(blockhash common.Hash, contractAddresses []common.Address) []*types.Log {
//	if nil == mysqlDBLog {
//		return nil
//	}
//
//	sql := "select height,logindex, blockhash,txhash,contractaddress,topic,data FROM contractlogs WHERE blockhash = ? "
//	if 0 != len(contractAddresses) {
//		sql += "and( "
//		for _, contractAddress := range contractAddresses {
//			sql += " contractaddress = \"" + contractAddress.GetHexString() + "\" " + "or"
//		}
//		sql = sql[:len(sql)-2] + ")"
//	}
//
//	rows, err := mysqlDBLog.Query(sql, blockhash.Hex())
//	if err != nil {
//		return nil
//	}
//	defer rows.Close()
//
//	result := make([]*types.Log, 0)
//	for rows.Next() {
//		var (
//			height, index                                   uint64
//			blockhash, txhash, contractaddress, topic, data string
//		)
//		err := rows.Scan(&height, &index, &blockhash, &txhash, &contractaddress, &topic, &data)
//		if err != nil {
//			logger.Errorf("scan failed, err: %v", err)
//			return nil
//		}
//
//		log := types.Log{
//			Address:     common.HexToAddress(contractaddress),
//			Data:        common.FromHex(data),
//			TxHash:      common.HexToHash(txhash),
//			BlockHash:   common.HexToHash(blockhash),
//			BlockNumber: height,
//			Index:       uint(index),
//		}
//		json.Unmarshal(utility.StrToBytes(topic), &log.Topics)
//		result = append(result, &log)
//	}
//
//	return result
//}
//

func InsertLogs(height int64, chainid, blockhash, ts, txhash, fromaddr, toaddr, value, contract string) {
	stmt, err := mysqlDBLog.Prepare("replace INTO chaindata(chainid,height,blockhash,ts,txhash,fromaddr,toaddr,`value`, contract) values(?,?,?,?,?,?,?,?,?)")
	if err != nil {
		logger.Errorf("fail to prepare. chainId: %s, height: %d, blockhash: %s, txhash: %s", chainid, height, blockhash, txhash)
		return
	}
	defer stmt.Close()

	result, err := stmt.Exec(chainid, height, blockhash, ts, txhash, fromaddr, toaddr, value, contract)
	if err != nil {
		logger.Errorf("fail to exec. chainId: %s, height: %d, blockhash: %s, txhash: %s", chainid, height, blockhash, txhash)
		return
	}

	row, _ := result.RowsAffected()
	lastId, _ := result.LastInsertId()
	logger.Infof("inserted. chainId: %s, height: %d, blockhash: %s,  txhash: %s, lines: %d, lastId: %d", chainid, height, blockhash, txhash, row, lastId)
}
