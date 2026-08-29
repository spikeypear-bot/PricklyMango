package test

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/spikeypear-bot/PricklyMango/internal"
)

func TestServer(t *testing.T){
	testHandler:=internal.NewTestHandle()
	svr:=httptest.NewTestServer(t,testHandler)
	defer svr.Close()

	client:=svr.Client()
	resp,err:=client.Get(svr.URL)
	if err!=nil{
		t.Fatalf("Error getting response")


	}

	if resp.StatusCode!=200{

		t.Fatalf("Status code error")

	}

	body,err:=io.ReadAll(resp.Body)
	if err!=nil{
		t.Fatalf("Error reading response")

	}

	if string(body)!="Hello"{
		t.Fatalf("Expected Hello, got %v",string(body))

	}



}


