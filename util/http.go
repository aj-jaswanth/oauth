package util

import (
	"io"
	"io/ioutil"
)

func ResponseBodyInBytes(respBody io.ReadCloser) []byte {
	respInBytes, err := ioutil.ReadAll(respBody)
	HandleErr(err)
	respBody.Close()
	return respInBytes
}
