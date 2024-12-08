package utility

import (
	"fmt"
	"errors"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"io"
	"strings"
	"net/http"
)

const INPUT_URL string = "https://adventofcode.com/2024/day/{day}/input"
const USERAGENT string = "https://github.com/FailedCode/adventofcode-2024-go by xeres666@googlemail.com"

var SESSION_TOKEN string = ""

type Config map[string]string

func LoadInput(day uint, inputTypeArg ...string) []string {
	// go has no default values... use a list(?) as backup
	var inputType string = "day"
	if len(inputTypeArg) > 0 {
		inputType = inputTypeArg[0]
	}

	inputPath := fmt.Sprintf("./input/%v%02d.txt", inputType, day)
	log.Println(fmt.Sprintf("Trying to load %v", inputPath))

	content, err := os.ReadFile(inputPath)
	if os.IsNotExist(err) {
		log.Println("input not yet saved, will try to donwload")
		lines := downloadInput(day)

		f, err := os.Create(inputPath)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Fprint(f, strings.Join(lines, "\n"))
		f.Close()

		return lines
	}

	lines := strings.Split(string(content), "\n")
	return lines
}

func downloadInput(day uint) []string {
	client := &http.Client{}

	url := strings.Replace(INPUT_URL, "{day}", fmt.Sprintf("%v", day), -1)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	if SESSION_TOKEN == "" {
		err := loadConfig()
		if err != nil {
			log.Fatalln(err)
		}
	}

	req.Header.Set("User-Agent", USERAGENT)
	req.Header.Set("Cookie", "session=" + SESSION_TOKEN)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}


	return strings.Split(string(body), "\n")
}

func loadConfig() error {
	log.Println("Loading session token...")
	f, err := os.ReadFile("./config/session.yaml")
	if err != nil {
		log.Println(err)
		return err
	}

	config := Config{}
	err = yaml.Unmarshal([]byte(f), &config)
	if err != nil {
		log.Fatalf("err")
		return err
	}

	SESSION_TOKEN = config["session"]
	if SESSION_TOKEN == "" {
		log.Println("Session token wasn't set by user")
		return errors.New("session unknown")
	}
	return nil
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Sgn(x int) int {
	if x == 0 {
		return 0
	}
	if x > 0 {
		return 1
	}
	return -1
}