package migrations_test

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/OpenBazaar/openbazaar-go/repo/migrations"
)

var stm = `PRAGMA key = 'letmein';
				create table utxos (outpoint text primary key not null,
					value integer, height integer, scriptPubKey text,
					watchOnly integer, coin text);
				create table stxos (outpoint text primary key not null,
					value integer, height integer, scriptPubKey text,
					watchOnly integer, spendHeight integer, spendTxid text,
					coin text);
				create table txns (txid text primary key not null,
					value integer, height integer, timestamp integer,
					watchOnly integer, tx blob, coin text);`

func TestMigration030(t *testing.T) {
	var dbPath string
	if err := os.Mkdir("./datastore", os.ModePerm); err != nil {
		t.Fatal(err)
	}
	dbPath = path.Join("./", "datastore", "mainnet.db")
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := db.Exec(stm); err != nil {
		t.Fatal(err)
	}
	_, err = db.Exec("INSERT INTO utxos (outpoint, value, height, scriptPubKey, watchOnly, coin) values (?,?,?,?,?,?)", "asdf", 3, 1, "key1", 1, "TBTC")
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.Exec("INSERT INTO stxos (outpoint, value, height, scriptPubKey, watchOnly, coin) values (?,?,?,?,?,?)", "asdf", 3, 1, "key1", 1, "TBTC")
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.Exec("INSERT INTO txns (txid, value, height, timestamp, watchOnly, coin) values (?,?,?,?,?,?)", "asdf", 3, 1, 234, 1, "TBTC")
	if err != nil {
		t.Fatal(err)
	}
	var m migrations.Migration030
	err = m.Up("./", "letmein", false)
	if err != nil {
		t.Error(err)
	}

	var outpoint string
	var value string
	var height int
	var scriptPubKey string
	var watchOnlyInt int
	var value1 int

	r := db.QueryRow("select outpoint, value, height, scriptPubKey, watchOnly from utxos where coin=?", "TBTC")

	if err := r.Scan(&outpoint, &value, &height, &scriptPubKey, &watchOnlyInt); err != nil || value != "3" {
		t.Fatal(err)
	}

	r = db.QueryRow("select outpoint, value, height, scriptPubKey, watchOnly from stxos where coin=?", "TBTC")

	if err := r.Scan(&outpoint, &value, &height, &scriptPubKey, &watchOnlyInt); err != nil || value != "3" {
		t.Fatal(err)
	}

	r = db.QueryRow("select txid, value, height, watchOnly from txns where coin=?", "TBTC")

	if err := r.Scan(&outpoint, &value, &height, &watchOnlyInt); err != nil || value != "3" {
		t.Fatal(err)
	}

	repoVer, err := ioutil.ReadFile("./repover")
	if err != nil {
		t.Error(err)
	}
	if string(repoVer) != "31" {
		fmt.Println("lets see : ", string(repoVer))
		t.Error("Failed to write new repo version")
	}

	err = m.Down("./", "letmein", false)
	if err != nil {
		t.Fatal(err)
	}
	r = db.QueryRow("select outpoint, value, height, scriptPubKey, watchOnly from utxos where coin=?", "TBTC")

	if err := r.Scan(&outpoint, &value1, &height, &scriptPubKey, &watchOnlyInt); err != nil || value1 != 3 {
		t.Fatal(err)
	}

	r = db.QueryRow("select outpoint, value, height, scriptPubKey, watchOnly from stxos where coin=?", "TBTC")

	if err := r.Scan(&outpoint, &value1, &height, &scriptPubKey, &watchOnlyInt); err != nil || value1 != 3 {
		t.Fatal(err)
	}

	r = db.QueryRow("select txid, value, height, watchOnly from txns where coin=?", "TBTC")

	if err := r.Scan(&outpoint, &value1, &height, &watchOnlyInt); err != nil || value1 != 3 {
		t.Fatal(err)
	}

	repoVer, err = ioutil.ReadFile("./repover")
	if err != nil {
		t.Error(err)
	}
	if string(repoVer) != "30" {
		fmt.Println("lets see down: ", string(repoVer))
		t.Error("Failed to write new repo version")
	}
	os.RemoveAll("./datastore")
	os.RemoveAll("./repover")
}
