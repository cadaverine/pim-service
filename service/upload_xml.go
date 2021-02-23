package service

import (
	"context"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

// UploadXML загрузка каталога
func (s *PimService) UploadXML(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	cg, err := getCatalogFromRequest(r)
	if err != nil {
		http.Error(w, "failed to parse multipart message", http.StatusBadRequest)
		return
	}

	fmt.Printf("%+v", cg)
}

func (s *PimService) SaveCatalog(ctx context.Context, cg *catalog) error {
	return nil
}

func getCatalogFromRequest(req *http.Request) (*catalog, error) {
	err := req.ParseMultipartForm(32 << 20)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse multipart message")
	}

	file, _, err := req.FormFile("file")
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse multipart message")
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)

	var cg catalog

	err = xml.Unmarshal(bytes, &cg)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse multipart message")
	}

	return &cg, nil
}
