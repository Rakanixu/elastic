package elastic

import (
	"errors"
	"fmt"
	elasticsearch "github.com/Rakanixu/elasticsearch/srv/proto/elasticsearch"
	lib "github.com/mattbaird/elastigo/lib"
)

var (
	ErrNotFound = errors.New("flag not found")
	Hosts       []string
	conn        *lib.Conn
)

// Init ES connection
func Init() {
	conn = lib.NewConn()
	conn.SetHosts(Hosts)
}

// Create record
func Create(cr *elasticsearch.CreateRequest) error {
	_, err := conn.Index(cr.Index, cr.Type, cr.Id, nil, cr.Data)

	return err
}

// Read record
func Read(rr *elasticsearch.ReadRequest) (string, error) {
	r, err := conn.Get(rr.Index, rr.Type, rr.Id, nil)
	if err != nil {
		return "", err
	}

	data, _ := r.Source.MarshalJSON()

	return string(data), nil
}

// Update record
func Update(ur *elasticsearch.UpdateRequest) error {
	_, err := conn.Index(ur.Index, ur.Type, ur.Id, nil, ur.Data)

	return err
}

// Delete record
func Delete(dr *elasticsearch.DeleteRequest) error {
	_, err := conn.Delete(dr.Index, dr.Type, dr.Id, nil)

	return err
}

// Search ES index
func Search(sr *elasticsearch.SearchRequest) (string, error) {
	if len(sr.Query) <= 0 {
		sr.Query = "*"
	}
	size := fmt.Sprintf("%d", sr.Limit)
	from := fmt.Sprintf("%d", sr.Offset)

	out, err := lib.Search(sr.Index).Type(sr.Type).Size(size).From(from).Search(sr.Query).Result(conn)
	if err != nil {
		return "", err
	}

	return string(out.RawJSON), nil
}

// Query DSL ES
func Query(sr *elasticsearch.QueryRequest) (string, error) {
	result, err := conn.Search(sr.Index, sr.Type, nil, sr.Query)

	if err != nil {
		return "", err
	}

	return string(result.RawJSON), nil
}
