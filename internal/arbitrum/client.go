package arbitrum

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const apiUrlMaskForTokens = `
https://api.arbiscan.io/api
?module=account
&action=tokentx
&contractaddress=%s
&address=%s
&offset=%d
&page=%d
&sort=asc
&apikey=%s`

// https://api.arbiscan.io/api?module=account&action=txlist&address=0x09197c3dd57E86Cb8b02A7ca2c315a7e59dE9383&page=1&offset=200&startblock=0&endblock=99999999&sort=asc&apikey=

type Client struct {
	client       http.Client
	c            *Config
	contractAddr map[string]string
}

type TxList struct {
	Hash           string
	Incomes        []Income
	TxTotal        int
	TxOffsetIncrBy int
}

type Income struct {
	Hash     string
	Received int
	Time     time.Time
}

type rawResp struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  []struct {
		BlockNumber       string `json:"blockNumber"`
		TimeStamp         string `json:"timeStamp"`
		Hash              string `json:"hash"`
		Nonce             string `json:"nonce"`
		BlockHash         string `json:"blockHash"`
		From              string `json:"from"`
		ContractAddress   string `json:"contractAddress"`
		To                string `json:"to"`
		Value             string `json:"value"`
		TokenName         string `json:"tokenName"`
		TokenSymbol       string `json:"tokenSymbol"`
		TokenDecimal      string `json:"tokenDecimal"`
		TransactionIndex  string `json:"transactionIndex"`
		Gas               string `json:"gas"`
		GasPrice          string `json:"gasPrice"`
		GasUsed           string `json:"gasUsed"`
		CumulativeGasUsed string `json:"cumulativeGasUsed"`
		Input             string `json:"input"`
		Confirmations     string `json:"confirmations"`
	} `json:"result"`
}

type Config struct {
	Token       string
	MetamaskAdr string
}

func New(c *Config) *Client {

	var tokenTxByNet = map[string]string{
		"ARBI_USDT": "0xfd086bc7cd5c481dcc9c85ebe478a1c0b69fcbb9",
	}

	return &Client{
		client:       http.Client{},
		c:            c,
		contractAddr: tokenTxByNet,
	}
}

func (c *Client) GetListTx(ctx context.Context, net, addr string, offset int) (*TxList, error) {
	reqUrl := fmt.Sprintf(apiUrlMaskForTokens, c.contractAddr[net], addr, 1, offset+1, c.c.Token)
	reqUrl = strings.ReplaceAll(reqUrl, "\n", "")

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqUrl, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	parsedBody := rawResp{}
	err = json.Unmarshal(body, &parsedBody)
	if err != nil {
		return nil, err
	}

	txList, err := parsedBody.toTxList(addr)
	if err != nil {
		return nil, err
	}

	return txList, nil
}

func (r *rawResp) toTxList(addr string) (*TxList, error) {
	res := &TxList{
		TxTotal:        0,
		TxOffsetIncrBy: len(r.Result),
	}
	var incomes []Income

	for i := range r.Result {
		tx := r.Result[i]

		if tx.To == strings.ToLower(addr) {
			ts, err := strconv.ParseInt(tx.TimeStamp, 10, 64)
			if err != nil {
				return nil, err
			}

			value, err := strconv.ParseInt(tx.Value, 10, 32)
			if err != nil {
				return nil, err
			}

			incomes = append(incomes, Income{
				Time:     time.Unix(ts, 0),
				Hash:     tx.Hash,
				Received: int(value),
			})
		}
	}
	res.Incomes = incomes

	return res, nil
}
