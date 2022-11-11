package ws

import (
	"apirest/internal/env"
	"apirest/internal/logger"
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
)

func ConsumeSOAP(bodyRequest string) ([]byte, int, error) {
	e := env.NewConfiguration()

	var req http.Request

	response, err := http.NewRequest("POST", e.SoapS.BaseUrl, bytes.NewBuffer([]byte(bodyRequest)))

	if err != nil {
		logger.Error.Printf("no se  puedo obtener respuesta: %v  -- log: ", err)
		return nil, 1, err
	}

	req = *response
	req.Header.Set("Content-Type", "text/xml; charset=UTF-8")
	//req.SetBasicAuth(username, password)

	client := &http.Client{}

	responseClient, err := client.Do(&req)

	if err != nil {
		logger.Error.Printf("no se  puedo enviar la petición: %v  -- log: ", err)
		return nil, responseClient.StatusCode, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logger.Error.Printf("no se pudo ejecutar defer body close: %v  -- log: ", err)
		}
	}(responseClient.Body)

	responseBody, err := ioutil.ReadAll(responseClient.Body)
	if err != nil {
		logger.Error.Printf("no se  puedo obtener respuesta: %v  -- log: ", err)
		return responseBody, responseClient.StatusCode, err
	}

	return responseBody, responseClient.StatusCode, nil
}
