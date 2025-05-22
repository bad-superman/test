package process

import (
	"context"
	"time"

	"github.com/bad-superman/test/logging"
	"github.com/bad-superman/test/sdk/thegraph"
)

const (
	_subgraphId = "ANk8smWo9Y1FCBY6EBU5mG2ArWzYJ1iex7iQb6SCB41X"
)

func (d *DataCron) SyncAllTheGraphIndexer() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Hour*2)
	defer cancel()
	d.syncAllTheGraphIndexerWorker(ctx)
}

func (d *DataCron) syncAllTheGraphIndexerWorker(ctx context.Context) {
	page := 50
	offset := 0
	for {
		select {
		case <-ctx.Done():
			logging.Warnf("syncAllTheGraphIndexerWorker stopped: %v", ctx.Err())
			return
		default:
		}
		result := &thegraph.QueryIndexersResponse{}
		err := d.thegraphClient.QueryIndexers(_subgraphId, page, offset, result)
		if err != nil {
			logging.Errorf("query indexers failed: %v", err)
			time.Sleep(time.Second * 10)
			continue
		}
		offset += page
		logging.Infof("sync indexers success, offset: %d len: %d", offset, len(result.Data.Indexers))
		if len(result.Data.Indexers) == 0 {
			return
		}
		for _, indexer := range result.Data.Indexers {
			indexer.UpdatedAt = time.Now()
			err := d.dao.UpsertIndexer(&indexer)
			if err != nil {
				logging.Errorf("upsert indexer failed: %v", err)
			}
		}
	}
}
