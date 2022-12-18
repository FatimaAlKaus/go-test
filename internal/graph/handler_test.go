package graph

import (
	"bytes"
	"context"
	"io"
	"net"
	"net/http"
	"strconv"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

var testCases []struct {
	desc    string
	request struct {
		method   string
		endPoint string
		body     string
	}
	response struct {
		status int
		body   string
	}
} = []struct {
	desc    string
	request struct {
		method   string
		endPoint string
		body     string
	}
	response struct {
		status int
		body   string
	}
}{
	{
		desc: "vertices more then 8",
		request: struct {
			method   string
			endPoint string
			body     string
		}{"POST", "/convertMatrix", `{"data": [[1],[1],[1],[1],[1],[1],[1],[1],[1],[1],[1]]}`},
		response: struct {
			status int
			body   string
		}{
			400,
			"count of vertices must be lower then 8",
		},
	},
	{
		desc: "vertices equals 0",
		request: struct {
			method   string
			endPoint string
			body     string
		}{"POST", "/convertMatrix", `{"data": []}`},
		response: struct {
			status int
			body   string
		}{
			400,
			"passed empty matrix",
		},
	},
	{
		desc: "invalid body",
		request: struct {
			method   string
			endPoint string
			body     string
		}{"POST", "/convertMatrix", `{"dogs":"cats"}`},
		response: struct {
			status int
			body   string
		}{
			400,
			"json: unknown field \"dogs\"",
		},
	},
	{
		desc: "passed valid matrix",
		request: struct {
			method   string
			endPoint string
			body     string
		}{"POST", "/convertMatrix", `{"data": [[1],[1],[1]]}`},
		response: struct {
			status int
			body   string
		}{
			200,
			"[{\"Vertices\":[0,1,2]}]\n",
		},
	},
}

func TestServer(t *testing.T) {
	h := &Handler{}
	r := mux.NewRouter()
	r.Handle("/convertMatrix", h).Methods(http.MethodPost)
	s := http.Server{
		Addr:    net.JoinHostPort("", strconv.Itoa(8081)),
		Handler: r,
	}
	go s.ListenAndServe()
	defer s.Shutdown(context.Background())

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			req, err := http.NewRequest(tC.request.method,
				"http://localhost:8081"+tC.request.endPoint,
				bytes.NewBuffer([]byte(tC.request.body)))
			assert.Nil(t, err)

			client := http.Client{}
			resp, err := client.Do(req)
			assert.Nil(t, err)
			assert.Equal(t, tC.response.status, resp.StatusCode)
			bytes, err := io.ReadAll(resp.Body)
			assert.Nil(t, err)
			assert.Equal(t, tC.response.body, string(bytes))
			resp.Body.Close()
		})
	}

}
