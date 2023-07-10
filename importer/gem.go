package importer

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type GemConnectorError string

func (e GemConnectorError) Error() string {
	return string(e)
}

const (
	ErrNotFound    = GemConnectorError("user was not found")
	ErrInvalidData = GemConnectorError("user had invalid data")
)

type ApiResponse struct {
	Results []struct {
		Id   string `json:"id"`
		Text string `json:"text"`
	} `json:"results"`
}

// Users know their gemId, but GEM uses its internal userId that is different from the gemId.
// We need to fetch the userId from the gemId to add the user to the event.
type User struct {
	id    string
	gemId string
}

type GemConnector interface {
	Connect(username, password string) bool
	GetUsersFromEvent(eventId string) []User
	GetUserByGemId(gemId string) User
	AddUsersToEvent(eventId string, users []User) bool
}

type CookieConnector struct {
	baseUrl  string
	username string
	password string
	client   *http.Client
}

func (connector *CookieConnector) Connect() bool {
	data := url.Values{}
	data.Set("username", connector.username)
	data.Set("password", connector.password)

	csrfToken := connector.GetCsrfToken()
	data.Set("csrfmiddlewaretoken", csrfToken)

	req, err := http.NewRequest(http.MethodPost, connector.baseUrl, strings.NewReader(data.Encode()))
	if err != nil {
		log.Fatal("Rquest not created", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Origin", connector.baseUrl)
	req.Header.Add("X-CSRFToken", csrfToken)

	resp, err := connector.client.Do(req)
	if err != nil {
		log.Fatal("Error during Post: ", err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatal("Status code was not OK: ", resp.StatusCode, resp.Status)
	}
	return true
}

func (connector *CookieConnector) GetUsersFromEvent(eventId string) []User {
	resp, err := connector.client.Get(connector.baseUrl + "/gem/" + eventId + "/add/")

	if err != nil {
		log.Fatal("Error during Get: ", err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatal("Status code was not OK: ", resp.StatusCode, resp.Status)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	users := []User{}

	doc.Find("ol>li").Each(func(idx int, item *goquery.Selection) {
		splits := strings.Split(item.Text(), " (")
		if len(splits) != 2 || len(splits[1]) == 0 {
			log.Default().Println("Warning: Possible error while parsing existing user ", item.Text())
		}

		users = append(users, User{id: "", gemId: splits[1][:len(splits[1])-1]})
	})

	return users
}

func (connector *CookieConnector) GetUserByGemId(eventId, gemId string) (User, error) {
	params := url.Values{}
	params.Add("q", gemId)
	params.Add("forward", `{"id":"`+eventId+`","country_filter":"ALL"}`)

	resp, err := connector.client.Get(connector.baseUrl + "/dal-urls/player-autocomplete-add/?" + params.Encode())
	if err != nil {
		log.Fatal("Error during Get: ", err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatal("Status code was not OK: ", resp.StatusCode, resp.Status)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var test ApiResponse
	json.Unmarshal(bodyBytes, &test)

	if len(test.Results) == 0 {
		return User{}, ErrNotFound
	}

	if test.Results[0].Id == "" {
		return User{}, ErrInvalidData
	}

	return User{id: test.Results[0].Id}, nil
}

func (connector *CookieConnector) AddUsersToEvent(eventId string, users []User) bool {
	return false
}

// Regex is more than 3 times faster than using Goquery in this case according to benchmarks
func (connector *CookieConnector) GetCsrfToken() string {
	resp, err := connector.client.Get(connector.baseUrl)

	if err != nil {
		log.Fatalf("Error during Get: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatal("Status code was not OK ", resp.StatusCode, resp.Status)
	}
	defer resp.Body.Close()
	r, err := regexp.Compile(`<input type="hidden" name="csrfmiddlewaretoken" value="([^"]*)">`)
	if err != nil {
		log.Fatal("Regex not compiled: ", err)
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	match := r.FindSubmatch(bodyBytes)

	if err != nil {
		log.Fatal("Regex got an error: ", err)
	}

	if len(match) < 2 || len(match[1]) == 0 {
		log.Fatal("Csrf token not found")
	}

	return string(match[1])
}
