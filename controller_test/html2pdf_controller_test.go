package controller_test

import (
	"testing"

	"github.com/Jeffail/tunny"
	"github.com/kataras/iris/v12/httptest"
	"github.com/lampnick/doctron/app"
	"github.com/lampnick/doctron/conf"
	"github.com/lampnick/doctron/mock"
	"github.com/lampnick/doctron/worker"
)

func init() {
	conf.LoadedConfig = conf.NewConfig()
	conf.LoadedConfig.Doctron.Uploader = conf.DoctronUploaderMock
	conf.LoadedConfig.Oss.PrivateServerDomain = "www.lampnick.com"
	worker.Pool = tunny.NewFunc(conf.LoadedConfig.Doctron.MaxConvertWorker, worker.DoctronHandler)
}

func TestHtml2Pdf(t *testing.T) {
	ts := mock.HTTPServer("text/html", "lampnick content test", false)
	defer ts.Close()

	doctron := app.NewDoctron()
	expect := httptest.New(t, doctron)
	request := expect.GET("/convert/html2pdf")
	request.WithQuery("u", "doctron")
	request.WithQuery("p", "lampnick")
	request.WithQuery("url", ts.URL)
	response := request.Expect().Status(httptest.StatusOK)
	response.Body().Length().Equal(18296)
}

func TestHtml2PdfUpload(t *testing.T) {
	ts := mock.HTTPServer("text/html", "lampnick content test", false)
	defer ts.Close()

	doctron := app.NewDoctron()
	expect := httptest.New(t, doctron)
	request := expect.GET("/convert/html2pdf")
	request.WithQuery("u", "doctron")
	request.WithQuery("p", "lampnick")
	request.WithQuery("url", ts.URL)
	request.WithQuery("uploadKey", "doctron.pdf")
	response := request.Expect().Status(httptest.StatusOK)
	expected := `{"code":0,"message":"","data":"http://www.lampnick.com/doctron.pdf"}`
	response.Body().Equal(expected)
}
