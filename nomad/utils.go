package nomad

import (
	"context"
	"time"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func convertJobSubmitTimestamp(_ context.Context, d *transform.TransformData) (interface{}, error) {
	if d.Value == nil {
		return nil, nil
	}
	var epochTime int64
	switch d.Value.(type) {
	case int64:
		epochTime = d.Value.(int64)
	case *int64:
		epochTime = *d.Value.(*int64)
	}

	unixtime := epochTime / 1e9
	unixTimestamp := time.Unix(int64(unixtime), 0)
	return unixTimestamp, nil

}
