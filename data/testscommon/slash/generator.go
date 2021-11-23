package slash

import (
	"fmt"
	"testing"

	"github.com/ElrondNetwork/elrond-go-core/core"
	"github.com/ElrondNetwork/elrond-go-core/core/mock"
	"github.com/ElrondNetwork/elrond-go-core/data"
	"github.com/ElrondNetwork/elrond-go-core/data/block"
	dataMock "github.com/ElrondNetwork/elrond-go-core/data/mock"
	"github.com/ElrondNetwork/elrond-go-core/data/slash"
	"github.com/ElrondNetwork/elrond-go-core/marshal"
	"github.com/stretchr/testify/require"
)

func GenerateSlashResults(b *testing.B, noOfPubKeys uint32, noOfHeaders uint32) map[string]slash.SlashingResult {
	hasher := &mock.HasherMock{}
	marshaller := &marshal.GogoProtoMarshalizer{}

	headers := make([]data.HeaderInfoHandler, 0, noOfHeaders)
	for i := 0; i < int(noOfHeaders); i++ {
		header := &block.HeaderV2{Header: &block.Header{Round: 1, TimeStamp: uint64(i)}}
		hash, err := core.CalculateHash(marshaller, hasher, header)
		require.Nil(b, err)

		headerInfo := &dataMock.HeaderInfoStub{Header: header, Hash: hash}
		headers = append(headers, headerInfo)
	}

	threatLevel := slash.Zero
	if noOfHeaders == slash.MinSlashableNoOfHeaders {
		threatLevel = slash.Medium
	} else if noOfHeaders >= slash.MinSlashableNoOfHeaders {
		threatLevel = slash.High
	}
	slashRes := make(map[string]slash.SlashingResult, noOfPubKeys)
	for i := 0; i < int(noOfPubKeys); i++ {
		tmp := fmt.Sprintf("pubKey%v", i)
		pubKey := hasher.Compute(tmp)
		slashRes[string(pubKey)] = slash.SlashingResult{
			Headers:       headers,
			SlashingLevel: threatLevel,
		}
	}

	return slashRes
}
