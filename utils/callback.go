package utils

import (
	"context"
	"net/http"
	"time"

	"shared-modules/common"
	"shared-modules/logger"
)

func CallbackPlain(ctx context.Context, req string, url string) (oriRes []byte, code int, success bool, err error) {

	cctx, cancel := context.WithTimeout(
		ctx,
		time.Duration(Config.Notify.CallbackTimeoutSeconds*time.Second),
	)
	resChan, errChan := make(chan struct{}, 1), make(chan error, 1)
	reqID, _ := GetMDCValue("reqId")
	go func(cctx context.Context) {
		SetMDCValue("reqId", reqID)

		header := make(http.Header)
		header.Add("accept", "application/json")
		header.Add("content-type", "application/json")

		oriRes, _, code, err = HttpPost(cctx, url, req, header)
		if err != nil {
			logger.Warn("post failed", err)
			err = NewBusinessError(cctx, common.CODE_EXTERNAL_API_ERROR, err.Error())
			errChan <- err
			return
		}
		logger.Infof("callback response body: %s", string(oriRes))
		if code < http.StatusOK || code >= http.StatusMultipleChoices {
			logger.Warnf("post failed, resp body: [%s], code: [%d]", string(oriRes), code)
			err = NewBusinessError(cctx, common.CODE_EXTERNAL_API_ERROR)
			errChan <- err
			return
		}
		success = true
		resChan <- struct{}{}

	}(cctx)
	defer cancel()

	select {
	case err = <-errChan:
		if err != nil {
			logger.Warn("notify failed", err)
			return
		}
	case <-resChan:
		return
	case <-cctx.Done():
		switch cctx.Err() {
		case context.DeadlineExceeded, context.Canceled:
			return nil, 0, false, NewBusinessError(ctx, common.CODE_EXTERNAL_API_TIMEOUT)
		}
		return nil, 0, false, NewBusinessError(ctx, common.CODE_UNKOWN_ERROR)
	case <-time.After(Config.Notify.CallbackTimeoutSeconds * time.Second):
		return nil, 0, false, NewBusinessError(ctx, common.CODE_EXTERNAL_API_TIMEOUT)
	}
	return nil, 0, false, nil
}

func CallbackEncrypted(ctx context.Context, req string, url string, pubKey string) (oriRes []byte, decRes []byte, code int, success bool, err error) {
	err = NewBusinessError(ctx, common.CODE_UNIMPLEMENTED)
	return
}
