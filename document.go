package esssdk

import (
	"bytes"
	"encoding/base64"
	"fmt"
)

// DocType resprents doc type.
type DocType int

// DocType enum
const (
	PDFDocType DocType = 2
	OFDDocType DocType = 3
)

// Doc represents an ess document.
type Doc struct {
	DocName          string  `json:"docName"`
	DocType          DocType `json:"docType"`
	DocContentBase64 string  `json:"docContentBase64"`
	UserID           string  `json:"userId"`
	UUID             string  `json:"uuid"`
}

// CreateDoc creates an ess document.
func (c *Client) CreateDoc(doc *Doc) (docURL string, err error) {
	url := fmt.Sprintf("http://%s/ess/api/user/doc/create", c.opts.Addr)

	formData := map[string]string{
		"docName": doc.DocName,
		"docType": fmt.Sprintf("%v", doc.DocType),
		"userId":  doc.UserID,
		"uuid":    doc.UUID,
	}

	var fileName string
	switch doc.DocType {
	case PDFDocType:
		fileName = fmt.Sprintf("%s.pdf", doc.DocName)
	case OFDDocType:
		fileName = fmt.Sprintf("%s.ofd", doc.DocName)
	default:
		fileName = doc.DocName
	}

	fileContent, err := base64.StdEncoding.DecodeString(doc.DocContentBase64)
	if err != nil {
		err = fmt.Errorf("文档内容格式无效: %v", err)
		return
	}
	fileReader := bytes.NewReader(fileContent)

	data, err := c.postFileForm(url, formData, "file", fileName, fileReader)
	if err != nil {
		return
	}

	v, ok := data.(map[string]interface{})
	if !ok {
		err = fmt.Errorf("response data invalid: %v", data)
		return
	}

	// Parse doc id param
	tv, found := v["id"]
	if !found {
		err = fmt.Errorf("response data invalid: %v", v)
		return
	}
	docID, ok := tv.(string)
	if !ok {
		err = fmt.Errorf("response data invalid: %v", tv)
		return
	}
	if docID == "" {
		err = fmt.Errorf("response data invalid: id is empty")
		return
	}
	doc.UUID = docID

	// Parse doc url param
	tv, found = v["url"]
	if !found {
		err = fmt.Errorf("response data invalid: %v", v)
		return
	}
	docURL, ok = tv.(string)
	if !ok {
		err = fmt.Errorf("response data invalid: %v", tv)
		return
	}
	if docURL == "" {
		err = fmt.Errorf("response data invalid: url is empty")
		return
	}

	return
}
