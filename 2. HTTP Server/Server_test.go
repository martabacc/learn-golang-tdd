package main

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "net/http/httptest"
  "net/http"
  // "fmt"
  // "io/ioutil"
)

func TestMain(t *testing.T) {

    // sort of before block
    server := httptest.NewServer(getRouters())
    defer server.Close()


  	t.Run("return 200 OK for URL /", func(t *testing.T) {


      resp, _ := http.Get(server.URL)
      actualStatus := resp.StatusCode
      expectedStatus := 200

      assert.Equal(t, expectedStatus, actualStatus, "OK response is expected")
  	})
}
