package indexer

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/ElrondNetwork/elrond-go/core/statistics"
	"github.com/ElrondNetwork/elrond-go/data"
	"github.com/ElrondNetwork/elrond-go/data/block"
	"github.com/ElrondNetwork/elrond-go/hashing"
	"github.com/ElrondNetwork/elrond-go/marshal"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

// elasticSearchDatabaseArgs is struct that is used to store all parameters that are needed to create a elastic search
// database object
type elasticSearchDatabaseArgs struct {
	url         string
	userName    string
	password    string
	marshalizer marshal.Marshalizer
	hasher      hashing.Hasher
}

// elasticSearchDatabase is used to prepare and write data in elastic search database
type elasticSearchDatabase struct {
	client      databaseWriterHandler
	marshalizer marshal.Marshalizer
	hasher      hashing.Hasher
}

// newElasticSearchDatabase is method that will create a new elastic search client
func newElasticSearchDatabase(arguments elasticSearchDatabaseArgs) (*elasticSearchDatabase, error) {
	cfg := elasticsearch.Config{
		Addresses: []string{arguments.url},
		Username:  arguments.userName,
		Password:  arguments.password,
	}
	es, err := newDatabaseWriter(cfg)
	if err != nil {
		return nil, err
	}

	elasticSearchDatabase := &elasticSearchDatabase{
		client:      es,
		marshalizer: arguments.marshalizer,
		hasher:      arguments.hasher,
	}

	err = elasticSearchDatabase.createIndexes()
	if err != nil {
		return nil, err
	}

	return elasticSearchDatabase, nil
}

func (esd *elasticSearchDatabase) createIndexes() error {
	err := esd.client.checkAndCreateIndex(blockIndex, timestampMapping())
	if err != nil {
		return err
	}

	err = esd.client.checkAndCreateIndex(txIndex, timestampMapping())
	if err != nil {
		return err
	}

	err = esd.client.checkAndCreateIndex(tpsIndex, nil)
	if err != nil {
		return err
	}

	err = esd.client.checkAndCreateIndex(validatorsIndex, nil)
	if err != nil {
		return err
	}

	err = esd.client.checkAndCreateIndex(roundIndex, timestampMapping())
	if err != nil {
		return err
	}

	return nil
}

func (esd *elasticSearchDatabase) saveHeader(header data.HeaderHandler, signersIndexes []uint64) {
	var buff bytes.Buffer

	serializedBlock, headerHash := esd.getSerializedElasticBlockAndHeaderHash(header, signersIndexes)

	buff.Grow(len(serializedBlock))
	_, err := buff.Write(serializedBlock)
	if err != nil {
		log.Warn("elastic search: save header, write", "error", err.Error())
	}

	req := esapi.IndexRequest{
		Index:      blockIndex,
		DocumentID: hex.EncodeToString(headerHash),
		Body:       bytes.NewReader(buff.Bytes()),
		Refresh:    "true",
	}

	err = esd.client.doRequest(req)
	if err != nil {
		log.Warn("indexer: could not index block header", "error", err.Error())
		return
	}
}

func (esd *elasticSearchDatabase) getSerializedElasticBlockAndHeaderHash(header data.HeaderHandler, signersIndexes []uint64) ([]byte, []byte) {
	h, err := esd.marshalizer.Marshal(header)
	if err != nil {
		log.Debug("indexer: marshal", "error", "could not marshal header")
		return nil, nil
	}

	headerHash := esd.hasher.Compute(string(h))
	elasticBlock := Block{
		Nonce:         header.GetNonce(),
		Round:         header.GetRound(),
		ShardID:       header.GetShardID(),
		Hash:          hex.EncodeToString(headerHash),
		Proposer:      signersIndexes[0],
		Validators:    signersIndexes,
		PubKeyBitmap:  hex.EncodeToString(header.GetPubKeysBitmap()),
		Size:          int64(len(h)),
		Timestamp:     time.Duration(header.GetTimeStamp()),
		TxCount:       header.GetTxCount(),
		StateRootHash: hex.EncodeToString(header.GetRootHash()),
		PrevHash:      hex.EncodeToString(header.GetPrevHash()),
	}

	serializedBlock, err := json.Marshal(elasticBlock)
	if err != nil {
		log.Debug("indexer: marshal", "error", "could not marshal elastic header")
		return nil, nil
	}

	return serializedBlock, headerHash
}

func (esd *elasticSearchDatabase) saveTransactions(
	body block.Body,
	header data.HeaderHandler,
	txPool map[string]data.TransactionHandler,
	selfShardId uint32,
) {
	bulks := esd.buildTransactionBulks(body, header, txPool, selfShardId)
	for _, bulk := range bulks {
		buff := esd.serializeBulkTx(bulk)
		err := esd.client.doBulkRequest(&buff, txIndex)
		if err != nil {
			log.Warn("indexer", "error", "indexing bulk of transactions")
			continue
		}
	}
}

// buildTransactionBulks creates bulks of maximum txBulkSize transactions to be indexed together
//  using the elastic search bulk API
func (esd *elasticSearchDatabase) buildTransactionBulks(
	body block.Body,
	header data.HeaderHandler,
	txPool map[string]data.TransactionHandler,
	selfShardId uint32,
) [][]*Transaction {
	processedTxCount := 0
	bulks := make([][]*Transaction, (header.GetTxCount()/txBulkSize)+1)
	blockMarshal, _ := esd.marshalizer.Marshal(header)
	blockHash := esd.hasher.Compute(string(blockMarshal))

	for _, mb := range body {
		mbMarshal, err := esd.marshalizer.Marshal(mb)
		if err != nil {
			log.Debug("indexer: marshal", "error", "could not marshal miniblock")
			continue
		}
		mbHash := esd.hasher.Compute(string(mbMarshal))

		mbTxStatus := "Pending"
		if selfShardId == mb.ReceiverShardID {
			mbTxStatus = "Success"
		}

		for _, txHash := range mb.TxHashes {
			processedTxCount++

			currentBulk := processedTxCount / txBulkSize
			currentTxHandler, ok := txPool[string(txHash)]
			if !ok {
				log.Debug("indexer: elasticsearch could not find tx hash in pool")
				continue
			}

			currentTx := getTransactionByType(currentTxHandler, txHash, mbHash, blockHash, mb, header, mbTxStatus)
			if currentTx == nil {
				log.Debug("indexer: elasticsearch found tx in pool but of wrong type")
				continue
			}

			bulks[currentBulk] = append(bulks[currentBulk], currentTx)
		}
	}

	return bulks
}

func (esd *elasticSearchDatabase) serializeBulkTx(bulk []*Transaction) bytes.Buffer {
	var buff bytes.Buffer
	for _, tx := range bulk {
		meta := []byte(fmt.Sprintf(`{ "index" : { "_id" : "%s", "_type" : "%s" } }%s`, tx.Hash, "_doc", "\n"))
		serializedTx, err := json.Marshal(tx)
		if err != nil {
			log.Debug("indexer: marshal",
				"error", "could not serialize transaction, will skip indexing",
				"tx hash", tx.Hash)
			continue
		}
		// append a newline foreach element
		serializedTx = append(serializedTx, "\n"...)

		buff.Grow(len(meta) + len(serializedTx))
		_, err = buff.Write(meta)
		if err != nil {
			log.Warn("elastic search: serialize bulk tx, write meta", "error", err.Error())
		}
		_, err = buff.Write(serializedTx)
		if err != nil {
			log.Warn("elastic search: serialize bulk tx, write serialized tx", "error", err.Error())
		}
	}

	return buff
}

func (esd *elasticSearchDatabase) saveRoundInfo(info RoundInfo) {
	var buff bytes.Buffer

	marshalizedRoundInfo, err := esd.marshalizer.Marshal(info)
	if err != nil {
		log.Debug("indexer: marshal", "error", "could not marshal signers indexes")
		return
	}

	buff.Grow(len(marshalizedRoundInfo))
	_, err = buff.Write(marshalizedRoundInfo)
	if err != nil {
		log.Warn("elastic search: save round info, write", "error", err.Error())
		return
	}

	req := esapi.IndexRequest{
		Index:      roundIndex,
		DocumentID: strconv.FormatUint(uint64(info.ShardId), 10) + "_" + strconv.FormatUint(info.Index, 10),
		Body:       bytes.NewReader(buff.Bytes()),
		Refresh:    "true",
	}

	err = esd.client.doRequest(req)
	if err != nil {
		log.Warn("indexer: can not index round info", "error", err.Error())
		return
	}
}

func (esd *elasticSearchDatabase) saveShardValidatorsPubKeys(shardId uint32, shardValidatorsPubKeys []string) {
	var buff bytes.Buffer

	shardValPubKeys := ValidatorsPublicKeys{PublicKeys: shardValidatorsPubKeys}
	marshalizedValidatorPubKeys, err := esd.marshalizer.Marshal(shardValPubKeys)
	if err != nil {
		log.Debug("indexer: marshal", "error", "could not marshal validators public keys")
		return
	}

	buff.Grow(len(marshalizedValidatorPubKeys))
	_, err = buff.Write(marshalizedValidatorPubKeys)
	if err != nil {
		log.Warn("elastic search: save shard validators pub keys, write", "error", err.Error())
	}

	req := esapi.IndexRequest{
		Index:      validatorsIndex,
		DocumentID: strconv.FormatUint(uint64(shardId), 10),
		Body:       bytes.NewReader(buff.Bytes()),
		Refresh:    "true",
	}

	err = esd.client.doRequest(req)
	if err != nil {
		log.Warn("indexer: can not index validators pubkey", "error", err.Error())
		return
	}
}

func (esd *elasticSearchDatabase) saveShardStatistics(tpsBenchmark statistics.TPSBenchmark) {
	buff := prepareGeneralInfo(tpsBenchmark)

	for _, shardInfo := range tpsBenchmark.ShardStatistics() {
		serializedShardInfo, serializedMetaInfo := serializeShardInfo(shardInfo)
		if serializedShardInfo == nil {
			continue
		}

		buff.Grow(len(serializedMetaInfo) + len(serializedShardInfo))
		_, err := buff.Write(serializedMetaInfo)
		if err != nil {
			log.Warn("elastic search: update TPS write meta", "error", err.Error())
		}
		_, err = buff.Write(serializedShardInfo)
		if err != nil {
			log.Warn("elastic search: update TPS write serialized data", "error", err.Error())
		}

		err = esd.client.doBulkRequest(&buff, tpsIndex)
		if err != nil {
			log.Warn("indexer: error indexing tps information", "error", err.Error())
			continue
		}
	}
}
