package dsn

import (
	"errors"
	"strings"
)

var (
	errInvalidDSNUnescaped       = errors.New("invalid DSN: did you forget to escape a param value?")
	errInvalidDSNAddr            = errors.New("invalid DSN: network address not terminated (missing closing brace)")
	errInvalidDSNNoSlash         = errors.New("invalid DSN: missing the slash separating the database name")
	errInvalidDSNUnsafeCollation = errors.New("invalid DSN: interpolateParams can not be used with unsafe collations")
)

const formatTemplate = "%s://%s%s/" //protocol,auth,host,path,query
type DSN struct {
	Protocol  string
	User      string            // Username
	Passwd    string            // Password (requires User)
	Transport string            // Network type
	Host      string            // Network address (requires Transport)
	Path      string            // Database name
	Query     map[string]string // Connection parameters
}

func (dsn *DSN) normalize() error {
	if dsn.Transport == "" {
		dsn.Transport = "tcp"
	}
	// Set default address if empty
	if dsn.Host == "" {
		return errors.New("default addr for network '" + dsn.Transport + "' unknown")

	}
	return nil
}

//func (dsn *DSN) String() string {
//	return ""
//}

func Parse(dsn string) (*DSN, error) {
	d := new(DSN)
	tmp := strings.SplitN(dsn, "://", 2)
	if len(tmp) != 2 {
		return nil, errInvalidDSNAddr
	}
	d.Protocol = tmp[0]
	dsn = tmp[1]
	foundSlash := false

	for i := len(dsn) - 1; i >= 0; i-- {
		if dsn[i] == '/' {
			foundSlash = true
			var j, k int
			if i > 0 {
				//parse auth
				for j = i; j >= 0; j-- {
					if dsn[j] == '@' {
						for k = 0; k < j; k++ {
							if dsn[k] == ':' {
								d.Passwd = dsn[k+1 : j]
								break
							}
						}
						d.User = dsn[:k]

						break
					} else if dsn[j] == '/' {
						i = j
					}
				}
				//parse transport
				for k = j + 1; ; k++ {
					if !(k < i) {
						d.Host = dsn[j+1 : k]
						break
					}
					if dsn[k] == '(' {
						if dsn[i-1] != ')' {
							if strings.ContainsRune(dsn[k+1:i], ')') {
								return nil, errInvalidDSNUnescaped
							}
							return nil, errInvalidDSNAddr
						}

						d.Host = dsn[k+1 : i-1]
						d.Transport = dsn[j+1 : k]
						break
					}
				}

			}
			// parse query
			for j = i + 1; j < len(dsn); j++ {
				if dsn[j] == '?' {
					for _, v := range strings.Split(dsn[j+1:], "&") {
						param := strings.SplitN(v, "=", 2)
						if len(param) != 2 {
							continue
						}
						if d.Query == nil {
							d.Query = make(map[string]string)
						}
						d.Query[param[0]] = param[1]
					}
					break
				}
			}
			d.Path = dsn[i+1 : j]
			break
		}
	}

	if !foundSlash && len(dsn) > 0 {
		return nil, errInvalidDSNNoSlash
	}

	if err := d.normalize(); err != nil {
		return nil, err
	}

	return d, nil
}
