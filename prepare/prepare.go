package prepare

import (
	"bytes"
	"io"
	"mime/multipart"
	"os"

	"github.com/faelp22/browser"
)

func PrepareJSON(b browser.BrowserCli, payload []byte) io.Reader {
	b.SetHeader("Content-Type", "application/json; charset=utf-8")
	return bytes.NewBuffer(payload)
}

func PrepareFileUploadMultipartFormData(b browser.BrowserCli, fileFieldName, filePath string, fileName string, extraFormFields map[string]string) (buf bytes.Buffer, w *multipart.Writer, err error) {

	b.SetHeader("Content-Type", w.FormDataContentType())
	var fw io.Writer
	w = multipart.NewWriter(&buf)
	defer w.Close()

	fw, err = w.CreateFormFile(fileFieldName, fileName)
	if err != nil {
		return buf, w, err
	}

	file, err := os.Open(filePath)
	if err != nil {
		return buf, w, err
	}

	_, err = io.Copy(fw, file)
	if err != nil {
		return buf, w, err
	}

	for k, v := range extraFormFields {
		_ = w.WriteField(k, v)
	}

	return buf, w, err
}
