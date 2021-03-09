package service

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

// UploadXML загрузка каталога
func (s *PimService) UploadXML(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	ctx := r.Context()

	cg, err := getCatalogFromRequest(r)
	if err != nil {
		http.Error(w, "failed to parse multipart message", http.StatusBadRequest)
		return
	}

	err = s.SaveCatalog(ctx, cg)
	if err != nil {
		http.Error(w, "failed to save catalog", http.StatusInternalServerError)
	}
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