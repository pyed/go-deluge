package deluge

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync/atomic"
)

// Deluge represents an endpoint for Deluge RPC requests.
type Deluge struct {
	url      string
	password string

	client  *http.Client
	cookies []*http.Cookie

	id uint64
}

// New instantiates a new Deluge instance and authenticates with the
// server.
func New(url, password string) (*Deluge, error) {
	d := &Deluge{
		url,
		password,
		new(http.Client),
		nil,
		0,
	}

	err := d.authLogin()
	if err != nil {
		return nil, err
	}

	return d, err
}

func (d *Deluge) GetTorrent(hash string) (*Torrent, error) {
	response, err := d.sendJsonRequest("core.get_torrent_status", []interface{}{hash, []string{}})
	if err != nil {
		return nil, err
	}

	torrent := new(Torrent)

	data, err := json.Marshal(response["result"].(map[string]interface{}))
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, torrent)
	if err != nil {
		return nil, err
	}

	if torrent.Hash == "" {
		return torrent, fmt.Errorf("No such torrent with hash: %s", hash)
	}

	return torrent, nil
}

func (d *Deluge) GetTorrents() (Torrents, error) {
	response, err := d.sendJsonRequest("core.get_torrents_status", []interface{}{nil, []string{}})
	if err != nil {
		return nil, err
	}

	jsonMap := response["result"].(map[string]interface{})
	torrents := make(Torrents, 0, len(jsonMap))

	for _, v := range jsonMap {
		torrent := new(Torrent)
		data, err := json.Marshal(v)
		if err != nil {
			return torrents, err
		}

		err = json.Unmarshal(data, torrent)
		if err != nil {
			return torrents, err
		}

		torrents = append(torrents, torrent)
	}

	return torrents, nil
}

func (d *Deluge) AddTorrentFile(fileName, fileDump string, options map[string]interface{}) (string, error) {
	response, err := d.sendJsonRequest("core.add_torrent_file", []interface{}{fileName, fileDump, options})
	if err != nil {
		return "", err
	}

	if response["result"] == nil {
		return "", fmt.Errorf("Error adding: %s\nMaybe already added ?", fileName)
	}

	return response["result"].(string), nil
}

func (d *Deluge) AddTorrentMagnet(magnetUrl string, options map[string]interface{}) (string, error) {
	response, err := d.sendJsonRequest("core.add_torrent_magnet", []interface{}{magnetUrl, options})
	if err != nil {
		return "", err
	}

	if response["result"] == nil {
		return "", fmt.Errorf("Error adding: %s\nMaybe already added ?", magnetUrl)
	}

	return response["result"].(string), nil
}

func (d *Deluge) AddTorrentUrl(torrentUrl string, options map[string]interface{}) (string, error) {
	response, err := d.sendJsonRequest("core.add_torrent_url", []interface{}{torrentUrl, options})
	if err != nil {
		return "", err
	}
	if response["result"] == nil {
		return "", fmt.Errorf("Error adding: %s\nMaybe already added ?", torrentUrl)
	}
	return response["result"].(string), nil
}

func (d *Deluge) RemoveTorrent(hash string, removeData bool) error {
	// make sure that we have a torrent with the giving hash;
	// attempting to remove a hash that doesn't exists stalls for ever.
	if _, err := d.GetTorrent(hash); err != nil {
		return err
	}
	_, err := d.sendJsonRequest("core.remove_torrent", []interface{}{hash, removeData})
	if err != nil {
		return err
	}

	return nil
}

func (d *Deluge) authLogin() error {
	response, err := d.sendJsonRequest("auth.login", []interface{}{d.password})
	if err != nil {
		return err
	}

	if response["result"] != true {
		return fmt.Errorf("authetication failed")
	}

	return nil
}

func (d *Deluge) sendJsonRequest(method string, params []interface{}) (map[string]interface{}, error) {
	atomic.AddUint64(&(d.id), 1)
	data, err := json.Marshal(map[string]interface{}{
		"method": method,
		"id":     d.id,
		"params": params,
	})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", d.url, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	if d.cookies != nil {
		for _, cookie := range d.cookies {
			req.AddCookie(cookie)
		}
	}
	resp, err := d.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("received non-ok status to http request : %d", resp.StatusCode)
	}

	d.cookies = resp.Cookies()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result := make(map[string]interface{})
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	if result["error"] != nil {
		return nil, fmt.Errorf("json error : %v", result["error"])
	}

	return result, err
}
