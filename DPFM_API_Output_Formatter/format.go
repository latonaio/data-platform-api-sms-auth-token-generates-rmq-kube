package dpfm_api_output_formatter

import (
	"data-platform-api-sms-auth-token-generates-rmq-kube/sub_func_complementer"
	"encoding/json"

	"golang.org/x/xerrors"
)

func ConvertToSMSAuthTokenCreates(subfuncSDC *sub_func_complementer.SDC) (*[]SMSAuthToken, error) {
	sMSAuthTokens := make([]SMSAuthToken, 0)

	for _, data := range *subfuncSDC.Message.SMSAuthToken {
		sMSAuthToken, err := TypeConverter[*SMSAuthToken](data)
		if err != nil {
			return nil, err
		}

		sMSAuthTokens = append(sMSAuthTokens, *sMSAuthToken)
	}

	return &sMSAuthTokens, nil
}

//func ConvertToAuthCreates(subfuncSDC *sub_func_complementer.SDC) (*[]Auth, error) {
//	auths := make([]Auth, 0)
//
//	for _, data := range *subfuncSDC.Message.Auth {
//		auth, err := TypeConverter[*Auth](data)
//		if err != nil {
//			return nil, err
//		}

//		auths = append(auths, *auth)
//	}

//	return &auths, nil
//}

func TypeConverter[T any](data interface{}) (T, error) {
	var dist T
	b, err := json.Marshal(data)
	if err != nil {
		return dist, xerrors.Errorf("Marshal error: %w", err)
	}
	err = json.Unmarshal(b, &dist)
	if err != nil {
		return dist, xerrors.Errorf("Unmarshal error: %w", err)
	}
	return dist, nil
}
