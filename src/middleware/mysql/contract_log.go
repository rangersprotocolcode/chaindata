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

func Count(addr, chainId string) uint64 {
	sql := "select count(*) FROM chaindata WHERE (fromaddr = ? and chainid = ?);"
	rows, err := mysqlDBLog.Query(sql, addr, chainId)
	if nil != err {
		logger.Errorf("fail to count, %s, %s", addr, chainId)
		return 0
	}
	defer rows.Close()

	for rows.Next() {
		var result uint64
		err := rows.Scan(&result)
		if err != nil {
			logger.Errorf("fail to scan, %s, %s", addr, chainId)
			return 0
		}
		return result
	}

	return 0
}

func Query(addr, chainId string, page, pageSize uint64) []item {
	sql := "select height,blockhash,ts,txhash, toaddr,`value`,contract FROM chaindata WHERE (fromaddr = ? and chainid = ?) limit ?, ?;"
	rows, err := mysqlDBLog.Query(sql, addr, chainId, page*pageSize, pageSize)
	if nil != err {
		logger.Errorf("fail to count, %s, %s", addr, chainId)
		return nil
	}
	defer rows.Close()

	result := make([]item, 0)
	for rows.Next() {
		var data item
		err := rows.Scan(&data.Height, &data.Blockhash, &data.Ts, &data.Txhash, &data.Toaddr, &data.Value, &data.Contract)
		if err != nil {
			logger.Errorf("fail to scan, %s, %s", addr, chainId)
			return nil
		}

		result = append(result, data)
	}

	return result
}

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

type item struct {
	Height    string `json:"height"`
	Blockhash string `json:"blockhash"`
	Ts        string `json:"timestamp"`
	Txhash    string `json:"txhash"`
	Toaddr    string `json:"toaddr"`
	Value     string `json:"value"`
	Contract  string `json:"contract"`
}
