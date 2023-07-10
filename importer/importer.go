package importer

import (
	"fmt"
	"log"
	"net/http"
	"net/http/cookiejar"
	"time"
)

func fetchUser(connector CookieConnector, eventId string, csvUser CsvUser, userChannel chan User) {
	newUser, err := connector.GetUserByGemId(eventId, csvUser.GemId)
	if err != nil {
		log.Default().Println("ImportUsers - Warning! Unable to find user on gem: ", csvUser, err)
	}
	fmt.Println("Preparing to add user", csvUser.Name)
	userChannel <- newUser
}

func ImportUsers(username, password, eventId, csvPath string) {
	userFromCsv := Read(csvPath)

	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatal("ImportUsers - Got error while creating cookie jar: ", err)
	}
	client := http.Client{Timeout: time.Duration(60) * time.Second, Jar: jar}
	connector := CookieConnector{
		client:   &client,
		baseUrl:  "https://gem.fabtcg.com",
		username: username,
		password: password,
	}

	if !connector.Connect() {
		log.Fatal("ImportUsers - Unable to log in to GEM server")
	}

	existingUsers := connector.GetUsersFromEvent(eventId)
	usersToAdd := []User{}
	userChannel := make(chan User)
	goRoutineCreated := 0

	for _, csvUser := range userFromCsv {
		found := false
		for _, gemUser := range existingUsers {
			if gemUser.gemId == csvUser.GemId {
				found = true
				break
			}
		}

		if !found {
			go fetchUser(connector, eventId, csvUser, userChannel)
			goRoutineCreated++
		}
	}

	for i := 0; i < goRoutineCreated; i++ {
		newUser := <-userChannel
		usersToAdd = append(usersToAdd, newUser)
	}

	fmt.Println("ImportUsers - Trying to add users: ", usersToAdd)
}
