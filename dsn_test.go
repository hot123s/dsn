package dsn

import (
	"testing"
)

func TestParseBasicDsn(t *testing.T) {
	d, err := Parse("redis://127.0.0.1:8081/aa/bb?jj=11")
	if err != nil {
		t.Error(err)
	}
	if d.Protocol != "redis" {
		t.Error(d.Protocol)
	}
	if d.Host != "127.0.0.1:8081" {
		t.Error(d.Host)
	}
	if d.Transport != "tcp" {
		t.Error(d.Transport)
	}
	if d.Path != "aa/bb" {
		t.Error(d.Path)
	}
	if val, ok := d.Query["jj"]; !ok || val != "11" {
		if !ok {
			t.Error("jj not in map")
		} else {
			t.Error(val)
		}
	}
}
func TestParseDsnWithUDP(t *testing.T) {
	d, err := Parse("redis://udp(127.0.0.1:8081)/aa/bb?jj=11")
	if err != nil {
		t.Error(err)
	}
	if d.Protocol != "redis" {
		t.Error(d.Protocol)
	}
	if d.Host != "127.0.0.1:8081" {
		t.Error(d.Host)
	}
	if d.Transport != "udp" {
		t.Log(d.Transport)
		t.Error()
	}
	if d.Path != "aa/bb" {
		t.Error(d.Path)
	}
	if val, ok := d.Query["jj"]; !ok || val != "11" {
		if !ok {
			t.Error("jj not in map")
		} else {
			t.Error(val)
		}
	}
}
func TestParseDsnWithAuth(t *testing.T) {
	d, err := Parse("redis://root:123456@127.0.0.1:8081/aa/bb?jj=11")
	if err != nil {
		t.Error(err)
	}
	if d.Protocol != "redis" {
		t.Error(d.Protocol)
	}
	if d.Host != "127.0.0.1:8081" {
		t.Error(d.Host)
	}
	if d.Transport != "tcp" {
		t.Log(d.Transport)
		t.Error()
	}
	if d.Path != "aa/bb" {
		t.Error(d.Path)
	}
	if d.Passwd != "123456" {
		t.Error(d.Passwd)
	}
	if d.User != "root" {
		t.Error(d.User)
	}
	if val, ok := d.Query["jj"]; !ok || val != "11" {
		if !ok {
			t.Error("jj not in map")
		} else {
			t.Error(val)
		}
	}
}
func TestParseDsnFull(t *testing.T) {
	d, err := Parse("redis://root:123456@udp(127.0.0.1:8081)/aa/bb?jj=11")
	if err != nil {
		t.Error(err)
	}
	if d.Protocol != "redis" {
		t.Error(d.Protocol)
	}
	if d.Host != "127.0.0.1:8081" {
		t.Error(d.Host)
	}
	if d.Transport != "udp" {
		t.Log(d.Transport)
		t.Error()
	}
	if d.Path != "aa/bb" {
		t.Error(d.Path)
	}
	if d.Passwd != "123456" {
		t.Error(d.Passwd)
	}
	if d.User != "root" {
		t.Error(d.User)
	}
	if val, ok := d.Query["jj"]; !ok || val != "11" {
		if !ok {
			t.Error("jj not in map")
		} else {
			t.Error(val)
		}
	}
}
