package store

type ShardingDirStrategyFn func(string) string

func defaultShardingDirStrategyFn(id string) string {
	if len(id) < 2 {
		return ""
	}

	return id[0:2]
}
