package handler

import (
	"LiveDanmu/apps/gateway/danmu_gateway/core"
	"LiveDanmu/apps/gateway/danmu_gateway/core/dto"
	"LiveDanmu/apps/public/models/dao"
	"LiveDanmu/apps/public/response"
	"LiveDanmu/apps/public/union_var"
	"context"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func PubDanmuHandleFunc() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 弹幕请求
		var danmuData dao.DanmuData
		// 获取上下文中的claims
		claims, _ := c.Get(union_var.JWT_CONTEXT_KEY)
		// 类型断言
		claim := claims.(*dao.MainClaims)
		// 提取请求体中的弹幕信息
		err := c.BindAndValidate(&danmuData)
		if err != nil {
			c.JSON(consts.StatusOK, response.ValidateRequestFail)
			return
		}
		// 填充danmuData
		danmuData.UserId = claim.Uid
		// 转换结构体
		pubReq := dto.GenPubReq(danmuData)
		// 调用PubDanmu微服务
		_, err = core.DanmuSvr.PubDanmu(ctx, pubReq)
		if err != nil {
			c.JSON(consts.StatusOK, response.InternalError(err))
			return
		}

		c.JSON(consts.StatusOK, response.OperationSuccess)
	}
}

func PubLiveDanmuHandleFunc() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 弹幕请求
		var danmuData dao.DanmuData
		// 获取上下文中的claims
		claims, _ := c.Get(union_var.JWT_CONTEXT_KEY)
		// 类型断言
		claim := claims.(*dao.MainClaims)
		// 提取请求体中的弹幕信息
		err := c.BindAndValidate(&danmuData)
		if err != nil {
			c.JSON(consts.StatusOK, response.ValidateRequestFail)
			return
		}
		// 填充danmuData
		danmuData.UserId = claim.Uid
		// 转换结构体
		pubReq := dto.GenPubLiveReq(danmuData)
		// 调用PubDanmu微服务
		_, err = core.DanmuSvr.PubLiveDanmu(ctx, pubReq)
		if err != nil {
			c.JSON(consts.StatusOK, response.InternalError(err))
			return
		}

		c.JSON(consts.StatusOK, response.OperationSuccess)
	}
}

func GetHotDanmuHandleFunc() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// danmu/hot/:rvid
		// 从路由中提取rvid
		rvid := c.Param("rvid")
		if rvid == "" {
			c.JSON(consts.StatusOK, response.EmptyRVID)
			return
		}
		// 将string转为int64
		num, err := strconv.ParseInt(rvid, 10, 64)
		if err != nil {
			c.JSON(consts.StatusOK, response.InternalError(err))
			return
		}
		// 生成GetTopReq
		getTopReq := dto.GenGetTopReq(num)
		// 调用GetTop
		resp, err := core.DanmuSvr.GetTop(ctx, getTopReq)
		if err != nil {
			c.JSON(consts.StatusOK, response.InternalError(err))
			return
		}
		finalResp := dto.GenFinalResponseForGetTopReq(resp)
		c.JSON(consts.StatusOK, finalResp)
	}
}

func GetFullDanmuHandleFunc() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// danmu/full/:rvid
		// 从路由中提取rvid
		rvid := c.Param("rvid")
		if rvid == "" {
			c.JSON(consts.StatusOK, response.EmptyRVID)
			return
		}
		// 将string转为int64
		num, err := strconv.ParseInt(rvid, 10, 64)
		if err != nil {
			c.JSON(consts.StatusOK, response.InternalError(err))
			return
		}
		// 生成GetTopReq
		getReq := dto.GenGetDanmuReq(num)
		// 调用GetDanmu
		resp, err := core.DanmuSvr.GetDanmu(ctx, getReq)
		if err != nil {
			c.JSON(consts.StatusOK, response.InternalError(err))
			return
		}
		finalResp := dto.GenFinalResponseForGetDanmuReq(resp)
		c.JSON(consts.StatusOK, finalResp)
	}
}
