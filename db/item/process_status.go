package item

import (
	"github.com/jchavannes/jgo/jerr"
	"github.com/jchavannes/jgo/jutil"
	"github.com/memocash/index/db/client"
	"github.com/memocash/index/db/item/db"
	"github.com/memocash/index/ref/config"
)

const (
	ProcessStatusTopicBlocks       = "blocks"
	ProcessStatusTopicBlockHeights = "block-heights"
	ProcessStatusTopicBlockTxes    = "block-txes"
)

type ProcessStatus struct {
	Name   string
	Height int64
	Shard  uint
}

func (s ProcessStatus) GetUid() []byte {
	return []byte(s.Name)
}

func (s ProcessStatus) GetShard() uint {
	return s.Shard
}

func (s ProcessStatus) GetTopic() string {
	return db.TopicProcessStatus
}

func (s ProcessStatus) Serialize() []byte {
	return jutil.GetInt64DataBig(s.Height)
}

func (s *ProcessStatus) SetUid(uid []byte) {
	s.Name = string(uid)
}

func (s *ProcessStatus) Deserialize(data []byte) {
	s.Height = jutil.GetInt64Big(data)
}

func (s *ProcessStatus) Save() error {
	if err := db.Save([]db.Object{s}); err != nil {
		return jerr.Get("error saving process status", err)
	}
	return nil
}

func NewProcessStatus(shard uint, name string) *ProcessStatus {
	return &ProcessStatus{
		Name:  name,
		Shard: shard,
	}
}

func GetProcessStatus(shard uint, name string) (*ProcessStatus, error) {
	shardConfig := config.GetShardConfig(uint32(shard), config.GetQueueShards())
	dbClient := client.NewClient(shardConfig.GetHost())
	err := dbClient.GetSingle(db.TopicProcessStatus, []byte(name))
	if err != nil {
		return nil, jerr.Get("error getting db message process status", err)
	}
	if len(dbClient.Messages) == 0 || len(dbClient.Messages[0].Uid) == 0 {
		return nil, jerr.Get("error status not found", client.MessageNotSetError)
	}
	var processStatus = new(ProcessStatus)
	db.Set(processStatus, dbClient.Messages[0])
	processStatus.Shard = shard
	return processStatus, nil
}
