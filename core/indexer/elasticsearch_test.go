package indexer_test

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

	logger "github.com/ElrondNetwork/elrond-go-logger"
	"github.com/ElrondNetwork/elrond-go/core"
	"github.com/ElrondNetwork/elrond-go/core/check"
	"github.com/ElrondNetwork/elrond-go/core/indexer"
	"github.com/ElrondNetwork/elrond-go/core/mock"
	"github.com/ElrondNetwork/elrond-go/data/block"
	"github.com/stretchr/testify/require"
)

func newTestMetaBlock() *block.MetaBlock {
	shardData := block.ShardData{
		ShardID:               1,
		HeaderHash:            []byte{1},
		ShardMiniBlockHeaders: []block.ShardMiniBlockHeader{},
		TxCount:               100,
	}
	return &block.MetaBlock{
		Nonce:     1,
		Round:     2,
		TxCount:   100,
		ShardInfo: []block.ShardData{shardData},
	}
}

func NewElasticIndexerArguments() indexer.ElasticIndexerArgs {
	return indexer.ElasticIndexerArgs{
		Url:                "url",
		UserName:           "user",
		Password:           "password",
		Marshalizer:        &mock.MarshalizerMock{},
		Hasher:             &mock.HasherMock{},
		Options:            &indexer.Options{},
		NodesCoordinator:   &mock.NodesCoordinatorMock{},
		EpochStartNotifier: &mock.EpochStartNotifierStub{},
	}
}

func TestElasticIndexer_NewIndexerWithNilUrlShouldError(t *testing.T) {

	arguments := NewElasticIndexerArguments()
	arguments.Url = ""
	ei, err := indexer.NewElasticIndexer(arguments)

	require.Nil(t, ei)
	require.Equal(t, core.ErrNilUrl, err)
}

func TestElasticIndexer_NewIndexerWithNilMarsharlizerShouldError(t *testing.T) {
	arguments := NewElasticIndexerArguments()
	arguments.Marshalizer = nil
	ei, err := indexer.NewElasticIndexer(arguments)

	require.Nil(t, ei)
	require.Equal(t, core.ErrNilMarshalizer, err)
}

func TestElasticIndexer_NewIndexerWithNilHasherShouldError(t *testing.T) {
	arguments := NewElasticIndexerArguments()
	arguments.Hasher = nil
	ei, err := indexer.NewElasticIndexer(arguments)

	require.Nil(t, ei)
	require.Equal(t, core.ErrNilHasher, err)
}

func TestElasticIndexer_NewIndexerWithCorrectParamsShouldWork(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/blocks" {
			w.WriteHeader(http.StatusOK)
		}
		if r.URL.Path == "/transactions" {
			w.WriteHeader(http.StatusOK)
		}
		if r.URL.Path == "/tps" {
			w.WriteHeader(http.StatusOK)
		}
	}))

	arguments := NewElasticIndexerArguments()
	arguments.Url = ts.URL
	ei, err := indexer.NewElasticIndexer(arguments)
	require.Nil(t, err)
	require.False(t, check.IfNil(ei))
	require.False(t, ei.IsNilIndexer())
}

func TestNewElasticIndexerIncorrectUrl(t *testing.T) {
	url := string([]byte{1, 2, 3})

	arguments := NewElasticIndexerArguments()
	arguments.Url = url
	ind, err := indexer.NewElasticIndexer(arguments)
	require.Nil(t, ind)
	require.NotNil(t, err)
}

func TestElasticIndexer_UpdateTPS(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))

	output := &bytes.Buffer{}
	_ = logger.SetLogLevel("core/indexer:TRACE")
	_ = logger.AddLogObserver(output, &logger.PlainFormatter{})
	arguments := NewElasticIndexerArguments()
	arguments.Url = ts.URL

	defer func() {
		_ = logger.RemoveLogObserver(output)
		_ = logger.SetLogLevel("core/indexer:INFO")
	}()

	ei, err := indexer.NewElasticIndexer(arguments)
	require.Nil(t, err)

	tpsBench := mock.TpsBenchmarkMock{}
	tpsBench.Update(newTestMetaBlock())

	ei.UpdateTPS(&tpsBench)
	require.Empty(t, output.String())
}

func TestElasticIndexer_UpdateTPSNil(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))

	output := &bytes.Buffer{}
	_ = logger.SetLogLevel("core/indexer:TRACE")
	_ = logger.AddLogObserver(output, &logger.PlainFormatter{})
	arguments := NewElasticIndexerArguments()
	arguments.Url = ts.URL

	defer func() {
		_ = logger.RemoveLogObserver(output)
		_ = logger.SetLogLevel("core/indexer:INFO")
	}()

	ei, err := indexer.NewElasticIndexer(arguments)
	require.Nil(t, err)

	ei.UpdateTPS(nil)
	require.NotEmpty(t, output.String())
}

func TestElasticIndexer_SaveBlockNilHeaderHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))

	output := &bytes.Buffer{}
	_ = logger.SetLogLevel("core/indexer:TRACE")
	_ = logger.AddLogObserver(output, &logger.PlainFormatter{})
	arguments := NewElasticIndexerArguments()
	arguments.Url = ts.URL
	ei, _ := indexer.NewElasticIndexer(arguments)

	defer func() {
		_ = logger.RemoveLogObserver(output)
		_ = logger.SetLogLevel("core/indexer:INFO")
	}()

	ei.SaveBlock(&block.Body{MiniBlocks: []*block.MiniBlock{}}, nil, nil, nil, nil)
	require.True(t, strings.Contains(output.String(), indexer.ErrNoHeader.Error()))
}

func TestElasticIndexer_SaveBlockNilBodyHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))

	output := &bytes.Buffer{}
	_ = logger.SetLogLevel("core/indexer:TRACE")
	_ = logger.AddLogObserver(output, &logger.PlainFormatter{})
	arguments := NewElasticIndexerArguments()
	arguments.Url = ts.URL
	ei, _ := indexer.NewElasticIndexer(arguments)

	defer func() {
		_ = logger.RemoveLogObserver(output)
		_ = logger.SetLogLevel("core/indexer:INFO")
	}()

	ei.SaveBlock(nil, nil, nil, nil, nil)
	require.True(t, strings.Contains(output.String(), indexer.ErrBodyTypeAssertion.Error()))
}

func TestElasticIndexer_SaveBlockNoMiniBlocks(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))

	output := &bytes.Buffer{}
	_ = logger.SetLogLevel("core/indexer:TRACE")
	_ = logger.AddLogObserver(output, &logger.PlainFormatter{})
	arguments := NewElasticIndexerArguments()
	arguments.Url = ts.URL
	ei, _ := indexer.NewElasticIndexer(arguments)

	defer func() {
		_ = logger.RemoveLogObserver(output)
		_ = logger.SetLogLevel("core/indexer:INFO")
	}()

	header := &block.Header{}
	body := &block.Body{}
	ei.SaveBlock(body, header, nil, []uint64{0}, nil)
	require.True(t, strings.Contains(output.String(), indexer.ErrNoMiniblocks.Error()))
}

func TestElasticIndexer_SaveRoundInfo(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))

	output := &bytes.Buffer{}
	_ = logger.SetLogLevel("core/indexer:TRACE")
	_ = logger.AddLogObserver(output, &logger.PlainFormatter{})
	arguments := NewElasticIndexerArguments()
	arguments.Marshalizer = &mock.MarshalizerMock{Fail: true}
	arguments.Url = ts.URL
	ei, _ := indexer.NewElasticIndexer(arguments)

	defer func() {
		_ = logger.RemoveLogObserver(output)
		_ = logger.SetLogLevel("core/indexer:INFO")
	}()

	ei.SaveRoundInfo(indexer.RoundInfo{})
	require.Empty(t, output.String())
}

func TestElasticIndexer_SaveValidatorsPubKeys(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))

	output := &bytes.Buffer{}
	_ = logger.SetLogLevel("core/indexer:TRACE")
	_ = logger.AddLogObserver(output, &logger.PlainFormatter{})
	arguments := NewElasticIndexerArguments()
	arguments.Url = ts.URL
	arguments.Marshalizer = &mock.MarshalizerMock{Fail: true}
	ei, _ := indexer.NewElasticIndexer(arguments)

	valPubKey := make(map[uint32][][]byte)

	keys := [][]byte{[]byte("key")}
	valPubKey[0] = keys
	epoch := uint32(0)
	ei.SaveValidatorsPubKeys(valPubKey, epoch)

	defer func() {
		_ = logger.RemoveLogObserver(output)
		_ = logger.SetLogLevel("core/indexer:INFO")
	}()

	require.Empty(t, output.String())
}

func TestElasticIndexer_EpochChange(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))

	getEligibleValidatorsCalled := false

	output := &bytes.Buffer{}
	_ = logger.SetLogLevel("core/indexer:TRACE")
	_ = logger.AddLogObserver(output, &logger.PlainFormatter{})
	arguments := NewElasticIndexerArguments()
	arguments.Url = ts.URL
	arguments.Marshalizer = &mock.MarshalizerMock{Fail: true}
	arguments.ShardId = core.MetachainShardId
	epochChangeNotifier := &mock.EpochStartNotifierStub{}
	arguments.EpochStartNotifier = epochChangeNotifier

	var wg sync.WaitGroup
	wg.Add(1)

	arguments.NodesCoordinator = &mock.NodesCoordinatorMock{
		GetAllEligibleValidatorsPublicKeysCalled: func() (m map[uint32][][]byte, err error) {
			defer wg.Done()
			getEligibleValidatorsCalled = true
			return nil, nil
		},
	}

	ei, _ := indexer.NewElasticIndexer(arguments)
	assert.NotNil(t, ei)

	epochChangeNotifier.NotifyAll(&block.Header{Nonce: 1})
	wg.Wait()

	assert.True(t, getEligibleValidatorsCalled)
}
