package kafka

import (
	"LiveDanmu/apps/rpc/danmu/dto"
	"LiveDanmu/apps/rpc/danmu/kitex_gen/danmusvr"
	"context"

	jsoniter "github.com/json-iterator/go"
)

func (r *KClient) GenDanmuKMsg(ctx context.Context, data *danmusvr.DanmuMsg) ([]byte, int, dto.Response) {
	msg, err := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(*data)
	if err != nil {
		return nil, 0, dto.ServerInternalError(err)
	}

}
