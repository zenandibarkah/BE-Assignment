package service

import (
	"encoding/json"
	"errors"
	"net/http"
	"payment-manager/config/models"
	"payment-manager/helper"

	"github.com/go-resty/resty/v2"
)

func HttpReqTokenCheck(token, url string) (models.Response, error) {
	var resp models.Response

	client := resty.New()
	r, err := client.R().
		SetHeader("Content-type", "application/json").
		SetAuthToken(token).
		Get(url)
	if err != nil {
		helper.Log.Error("Error send in HttpReqTokenCheck :: ", err.Error(), r.Error(), r.RawResponse)
		return resp, err
	}

	if err := json.Unmarshal(r.Body(), &resp); err != nil {
		helper.Log.Error("Error send in HttpReqTokenCheck :: ", err.Error())
		return resp, err
	}

	if r.StatusCode() != http.StatusOK {
		helper.Log.Error("Error status not OK from token check")
		return resp, errors.New("invalid Token")
	}

	return resp, nil
}

func HttpReqBalanceSum(req models.ReqUpdateAccBalance, token string, url string) error {

	client := resty.New()
	r, err := client.R().
		SetHeader("Content-type", "application/json").
		SetAuthToken(token).
		SetBody(req).
		Post(url)
	if err != nil {
		helper.Log.Error("Error send in HttpReqBalanceSum --> ", err.Error(), r.Error(), r.RawResponse)
		return err
	}

	if r.StatusCode() != http.StatusOK {
		helper.Log.Error("Error status not OK")
		return errors.New("failed Sum")
	}

	return nil
}

func HttpReqBalanceReduc(req models.ReqUpdateAccBalance, token string, url string) error {

	client := resty.New()
	r, err := client.R().
		SetHeader("Content-type", "application/json").
		SetAuthToken(token).
		SetBody(req).
		Post(url)
	if err != nil {
		helper.Log.Error("Error send in HttpReqBalanceReduc --> ", err.Error(), r.Error(), r.RawResponse)
		return err
	}

	if r.StatusCode() != http.StatusOK {
		helper.Log.Error("Error status not OK")
		return errors.New("failed Reduc")
	}

	return nil
}
