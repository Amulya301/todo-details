package utils

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jmoiron/sqlx"
)

type TestServer struct {
	*httptest.Server
}

func (ts *TestServer) Post(t *testing.T, urlPath string, reqBody string, contentType string) (int, http.Header, []byte) {
	req, err := http.NewRequest("POST", ts.URL+urlPath, bytes.NewBuffer([]byte(reqBody)))

	if err != nil {
		t.Fatal(err)
	}

	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}

	rs, err := ts.Client().Do(req)
	if err != nil {
		t.Fatal(err)
	}

	defer rs.Body.Close()
	body, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}

	return rs.StatusCode, rs.Header, body
}

func (ts *TestServer) Get(t *testing.T, urlPath string) (int, http.Header, []byte) {
	req, err := http.NewRequest("GET", ts.URL+urlPath, nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rs, err := ts.Client().Do(req)
	if err != nil {
		t.Fatal(err)
	}

	defer rs.Body.Close()
	body, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}

	return rs.StatusCode, rs.Header, body
}

func (ts *TestServer) Delete(t *testing.T, urlPath string, contentType string) (int, http.Header, []byte) {
	req, err := http.NewRequest("DELETE", ts.URL+urlPath, nil)
	if err != nil {
		t.Fatal(err)
	}

	rs, err := ts.Client().Do(req)
	if err != nil {
		t.Fatal(err)
	}

	defer rs.Body.Close()
	body, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}

	return rs.StatusCode, rs.Header, body
}

func (ts *TestServer) Put(t *testing.T, urlPath string, reqBody string) (int, http.Header, []byte) {
	req, err := http.NewRequest("PUT", ts.URL+urlPath, bytes.NewBuffer([]byte(reqBody)))
	if err != nil {
		t.Fatal(err)
	}

	rs, err := ts.Client().Do(req)
	if err != nil {
		t.Fatal(err)
	}

	defer rs.Body.Close()
	body, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}

	return rs.StatusCode, rs.Header, body
}

func ClearTestDatabase(db *sqlx.DB) {
	_, _ = db.Exec(`TRUNCATE TABLE todos;`)
}

func SeedTodos(db *sqlx.DB) error {
	query := `INSERT INTO todos(name, completed, created_at) VALUES 
                                                	("title1", "Not done", CURRENT_TIMESTAMP),
                                                   ("title2", "In Progress", CURRENT_TIMESTAMP);
                                                   `

	_, err := db.Exec(query)

	return err
}
