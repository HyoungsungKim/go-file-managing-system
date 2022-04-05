package controller

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"net/http"
	"net/http/httptest"
	"os/exec"

	"testing"
)

func TestPingR(t *testing.T) {
	ts := httptest.NewServer(SetupRouter())
	defer ts.Close()

	response, err := http.Get(fmt.Sprintf("%s/ping", ts.URL))

	if err != nil {
		t.Fatalf("Exprexted no error, got %v", err)
	}

	if response.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", response.StatusCode)
	} else {
		t.Logf("Expected status code 200, got %v", response.StatusCode)
	}

	val, ok := response.Header["Content-Type"]

	if !ok {
		t.Fatalf("Expected Content-Type header to be set")
	}

	if val[0] != "application/json; charset=utf-8" {
		t.Fatalf("Expected \"application/json; charset=utf-8\", got %s", val[0])
	} else {
		t.Logf("Expected \"application/json; charset=utf-8\", got %s", val[0])
	}
}

func TestUpload(t *testing.T) {
	ts := httptest.NewServer(SetupRouter())
	defer ts.Close()

	os.Mkdir("storage", os.ModePerm)
	tempDir, _ := ioutil.TempDir("./", "temp_dir")

	defer os.RemoveAll(tempDir)
	defer os.RemoveAll("storage")

	txtFile, _ := ioutil.TempFile(tempDir, "pre-")
	txtFile.Write([]byte("Hello world"))

	t.Log(txtFile.Name())
	t.Log(ts.URL + "/upload/kim")

	//exec.Command(curlCommand).Run()
	c := exec.Command("curl", "-X", "POST", ts.URL+"/upload/"+tempDir, "-F", "file=@"+txtFile.Name(), "-H", "Content-Type: multipart/form-data")
	_, err := c.Output()
	if err != nil {
		t.Fatal("curl error. Check command again")
	}

	srcFile, err := ioutil.ReadFile(txtFile.Name())
	if err != nil {
		t.Fatal("srcFile does not exists")
	}

	dstFile, err := ioutil.ReadFile("./storage/" + txtFile.Name())
	if err != nil {
		t.Fatal("dstFile does not exists")
	}

	if !bytes.Equal(srcFile, dstFile) {
		t.Fail()
		t.Logf("srcFile and dstFile are different")
	}

}
