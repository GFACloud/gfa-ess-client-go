package client_test

import (
	"encoding/base64"
	"io/ioutil"
	"strings"
	"testing"

	ecv2 "github.com/gfacloud/gfa-ess-client-go/v2"
)

const (
	// For 145
	APP_KEY_TEST    = "ff80808185a616240185aa352dc409a6"
	APP_SECRET_TEST = "2d291657812b0c56a8a197c307f1dfa2116f501d"
	ADDR_TEST       = "localhost:8080"

	// For 109
	// APP_KEY_TEST    = "8ae58ded77c7a5a90177c7a63d260002"
	// APP_SECRET_TEST = "e09ea9ae4927005e747763e014a78fe84d75024c"
	// ADDR_TEST       = "10.101.13.109:8800"
)

func TestNewClient(t *testing.T) {
	opts := &ecv2.Options{
		AppKey:    APP_KEY_TEST,
		AppSecret: APP_SECRET_TEST,
		Addr:      ADDR_TEST,
	}
	c, err := ecv2.NewClient(opts)
	if err != nil {
		t.Fatalf("NewClient failed: %v", err)
	}
	if c == nil {
		t.Fatalf("client should not be nil")
	}
}

func readFile(filename string) (content []byte, fileType string, err error) {
	// Read file content
	content, err = ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	// Parse file type
	items := strings.Split(filename, ".")
	if len(items) >= 2 {
		fileType = items[len(items)-1]
	}

	return
}

func readFileBase64(filename string) (contentBase64 string, fileType string, err error) {
	// Read file content, and encode it with Base64
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	contentBase64 = base64.StdEncoding.EncodeToString(content)

	// Parse file type
	items := strings.Split(filename, ".")
	if len(items) >= 2 {
		fileType = items[len(items)-1]
	}

	return
}
