package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"strings"

	bolt "go.etcd.io/bbolt"
)

const (
	path        = "urlshortener.db"
	shortURLLen = 4
)

func main() {

	http.HandleFunc("/", home)
	http.HandleFunc("/shorten/", shortener)
	http.HandleFunc("/redirect/", redirect)

	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "To shorten a url, go to %s<br><br>", fmt.Sprintf("http://%s/shorten?url=", r.Host))
	exampleURL := fmt.Sprintf("%s/shorten?url=%s", r.Host, "https://github.com/etcd-io/bbolt")
	fmt.Fprintf(w, "Example: <a href=\"%[1]s\">%[1]s</a><br>", exampleURL)
}

func redirect(w http.ResponseWriter, r *http.Request) {
	shortCode := strings.Split(r.URL.Path, "/")[2]

	db := dbHandler()
	defer db.Close()

	valueFromDB, err := readFromDB(db, shortCode)
	if err != nil {
		panic(err)
	}
	if valueFromDB == nil {
		fmt.Fprintf(w, "Short URL not generated for short code: %s\n", shortCode)
	}

	http.Redirect(w, r, string(valueFromDB), http.StatusFound)
}

func shortener(w http.ResponseWriter, r *http.Request) {

	urlToShorten := strings.Split(r.URL.RawQuery, "=")[1]
	urlHash := hex.EncodeToString(hashInput(urlToShorten))[:shortURLLen]

	db := dbHandler()
	defer db.Close()

	valueFromDB, err := readFromDB(db, urlHash)
	if err != nil {
		panic(err)
	}
	if valueFromDB == nil {
		err := insertToDB(db, urlHash, urlToShorten)
		if err != nil {
			panic(err)
		}
	}

	shortURL := fmt.Sprintf("%s/redirect/%s", r.Host, urlHash)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "Short URL: <a href=\"%[1]s\">%[1]s</a><br>", shortURL)
}

func hashInput(input string) []byte {
	h := sha256.New()
	h.Write([]byte(input))
	hash := h.Sum(nil)
	return hash
}

func dbHandler() (db *bolt.DB) {
	db, err := bolt.Open(path, 0666, nil)
	if err != nil {
		panic(fmt.Errorf("unable to open db: %v", err))
	}

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("urls"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})

	return db
}

func insertToDB(db *bolt.DB, key, value string) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("urls"))
		err := b.Put([]byte(key), []byte(value))
		return err
	})
}

func readFromDB(db *bolt.DB, key string) ([]byte, error) {
	tx, err := db.Begin(true)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	b := tx.Bucket([]byte("urls"))
	value := b.Get([]byte(key))

	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return value, nil
}
