package dsn

import (
	"strings"
	"testing"
)

type _TestTable struct {
	dsn           string
	wantProtocol  string
	wantUser      string
	wantPasswd    string
	wantHost      string
	wantTransport string
	wantPath      string
	wantParam     string
}

func TestParse(t *testing.T) {
	table := []_TestTable{
		{
			dsn:           "redis://127.0.0.1:8081/aa/bb?redis=redis://127.0.0.1",
			wantProtocol:  "redis",
			wantUser:      "",
			wantPasswd:    "",
			wantHost:      "127.0.0.1:8081",
			wantTransport: "tcp",
			wantPath:      "aa/bb",
			wantParam:     "redis=redis://127.0.0.1",
		},
		{
			dsn:           "redis://root:123456@127.0.0.1:8081/aa/bb?jj=11",
			wantProtocol:  "redis",
			wantUser:      "root",
			wantPasswd:    "123456",
			wantHost:      "127.0.0.1:8081",
			wantTransport: "tcp",
			wantPath:      "aa/bb",
			wantParam:     "jj=11",
		},
		{
			dsn:           "redis://udp(127.0.0.1:8081)/aa/bb?jj=11",
			wantProtocol:  "redis",
			wantUser:      "",
			wantPasswd:    "",
			wantHost:      "127.0.0.1:8081",
			wantTransport: "udp",
			wantPath:      "aa/bb",
			wantParam:     "jj=11",
		},
		{
			dsn:           "okex://127.0.0.1:8081/aa/bb?redis=redis://root:123456@127.0.0.1",
			wantProtocol:  "okex",
			wantUser:      "",
			wantPasswd:    "",
			wantHost:      "127.0.0.1:8081",
			wantTransport: "tcp",
			wantPath:      "aa/bb",
			wantParam:     "redis=redis://root:123456@127.0.0.1",
		},
	}
	for i := range table {
		tmp, err := Parse(table[i].dsn)
		if err != nil {
			t.Error(err)
		}
		if tmp.Host != table[i].wantHost {
			t.Error(tmp.Host, table[i].wantHost)
		}
		if tmp.Protocol != table[i].wantProtocol {
			t.Error(tmp.Protocol, table[i].wantProtocol)
		}
		if tmp.User != table[i].wantUser {
			t.Error(tmp.User, table[i].wantUser)
		}
		if tmp.Passwd != table[i].wantPasswd {
			t.Error(tmp.Passwd, table[i].wantPasswd)
		}
		if tmp.Transport != table[i].wantTransport {
			t.Error(tmp.Transport, table[i].wantTransport)
		}
		if tmp.Path != table[i].wantPath {
			t.Error(tmp.Path, table[i].wantPath)
		}
		params := strings.Split(table[i].wantParam, "&")
		for j := range params {
			kv := strings.SplitN(params[j], "=", 2)
			if tmp.Query[kv[0]] == kv[1] {
				continue
			}
			t.Error(kv[0], tmp.Query[kv[0]], kv[1])
		}
	}
}

