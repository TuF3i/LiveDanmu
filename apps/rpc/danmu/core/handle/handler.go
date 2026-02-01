package handle

import (
	"LiveDanmu/apps/public/dto"
	"LiveDanmu/apps/rpc/danmu/kitex_gen/danmusvr"
	"context"
)

// DanmuSvrImpl implements the last handle interface defined in the IDL.
type DanmuSvrImpl struct{}

// PubDanmu implements the DanmuSvrImpl interface.
func (s *DanmuSvrImpl) PubDanmu(ctx context.Context, req *danmusvr.PubReq) (resp *danmusvr.PubResp, err error) {
	rawResp := PubVideoDanmu(ctx, req)
	return dto.GenFinalRespForPubDanMu(rawResp), nil
}

// PubLiveDanmu implements the DanmuSvrImpl interface.
func (s *DanmuSvrImpl) PubLiveDanmu(ctx context.Context, req *danmusvr.PubLiveReq) (resp *danmusvr.PubLiveResp, err error) {
	rawResp := PubLiveDanmu(ctx, req)
	return dto.GenFinalRespForPubLiveDanMu(rawResp), nil
}

// GetDanmu implements the DanmuSvrImpl interface.
func (s *DanmuSvrImpl) GetDanmu(ctx context.Context, req *danmusvr.GetReq) (resp *danmusvr.GetResp, err error) {
	data, rawResp := GetFullDanmu(ctx, req)
	return dto.GenFinalRespForGetDanMu(rawResp, data), nil
}

// GetTop implements the DanmuSvrImpl interface.
func (s *DanmuSvrImpl) GetTop(ctx context.Context, req *danmusvr.GetTopReq) (resp *danmusvr.GetTopResp, err error) {
	data, rawResp := GetHotDanmu(ctx, req)
	return dto.GenFinalRespForGetHotDanMu(rawResp, data), nil
}
