package faaschain

import (
	"testing"
)

var (
	data = []byte("<svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 100 100\">" +
		"<path d=\"M30,1h40l29,29v40l-29,29h-40l-29-29v-40z\" stroke=\"#000\" fill=\"none\"/> " +
		"<path d=\"M31,3h38l28,28v38l-28,28h-38l-28-28v-38z\" fill=\"#a23\"/> " +
		"<text x=\"50\" y=\"68\" font-size=\"48\" fill=\"#FFF\" text-anchor=\"middle\"> " +
		"<![CDATA[410]]> " +
		"</text></svg>")
)

func TestChainCreate(t *testing.T) {
	chain := NewFaaschain("http://127.0.0.1:8080", "mychain")
	if chain == nil {
		t.Errorf("Creating faas chain: got %v", chain)
		t.Fail()
	}
}

func TestApply(t *testing.T) {
	chain := NewFaaschain("http://127.0.0.1:8080", "mychain")
	chain.Apply("compress", Header("Method", "Post"), Sync).Apply("upload", Header("Method", "Post"), Query("URL", "my.file.storage/s8sg"), Sync)
}

func TestApplyAsync(t *testing.T) {
	chain := NewFaaschain("http://127.0.0.1:8080", "mychain")
	chain.Apply("compress", Header("Method", "Post")).Apply("upload", Header("Method", "Post"), Query("URL", "my.file.storage/s8sg"))
}
