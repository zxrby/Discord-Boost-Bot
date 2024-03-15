package Utils

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/gookit/color"
)

var config, _ = LoadConfig()

func LoadConfig() (ConfigFile, error) {
	var config ConfigFile

	configFile, _ := os.Open("config.json")

	defer configFile.Close()

	_ = json.NewDecoder(configFile).Decode(&config)

	return config, nil
}

func Replacelast(token string) string {
	strLen := len(token)

	// Check if the string is longer than 20 characters
	if strLen > 20 {

		// Create a string of dots with the same length as last20Chars
		dots := strings.Repeat(".", 3)

		// Replace the last 20 characters with dots
		modifiedString := token[:strLen-41] + dots

		return modifiedString
	} else {
		// The string is less than 20 characters, no modification needed
		return ""
	}

}

func ClearScreen() {
	switch runtime.GOOS {
	case "linux", "darwin": // Unix-like systems
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows": // Windows
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		fmt.Println("Unsupported operating system. Cannot clear the screen.")
	}
}

func PrintASCII() {
	color.Println()
	color.Println()
	color.Println(`<fg=7f03fc>	   ██████╗  █████╗  █████╗  ██████╗████████╗██╗██╗     ██╗████████╗██╗   ██╗</><fg=7f03fc></>`)
	color.Println(`<fg=7f03fc>	   ██╔══██╗██╔══██╗██╔══██╗██╔════╝╚══██╔══╝██║██║     ██║╚══██╔══╝╚██╗ ██╔╝</><fg=7f03fc></>`)
	color.Println(`<fg=7f03fc>	   ██████╦╝██║  ██║██║  ██║╚█████╗    ██║   ██║██║     ██║   ██║    ╚████╔╝ </><fg=7f03fc></>`)
	color.Println(`<fg=7f03fc>	   ██╔══██╗██║  ██║██║  ██║ ╚═══██╗   ██║   ██║██║     ██║   ██║     ╚██╔╝  </><fg=7f03fc></>`)
	color.Println(`<fg=7f03fc>	   ██████╦╝╚█████╔╝╚█████╔╝██████╔╝   ██║   ██║███████╗██║   ██║      ██║   </><fg=7f03fc></>`)
	color.Println(`<fg=7f03fc>	   ╚═════╝  ╚════╝  ╚════╝ ╚═════╝    ╚═╝   ╚═╝╚══════╝╚═╝   ╚═╝      ╚═╝   </><fg=7f03fc></>`)
	color.Println()
	color.Println()
}

func SendToken(file string) string {

	// Open the file for reading and writing
	file1, err := os.OpenFile("./Data/"+file, os.O_RDWR, os.ModePerm)
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	defer file1.Close()

	// Create a scanner to read from the file
	scanner := bufio.NewScanner(file1)

	// Read the first line (if exists)
	if scanner.Scan() {
		// Get the content of the first line
		firstLine := scanner.Text()

		// Move the file cursor to the beginning
		file1.Seek(0, 0)

		// Truncate the file to remove the first line
		if err := file1.Truncate(0); err != nil {
			fmt.Println("Error:", err)
			return ""
		}

		// Write the remaining content (excluding the first line) back to the file
		for scanner.Scan() {
			_, err := fmt.Fprintln(file1, scanner.Text())
			if err != nil {
				fmt.Println("Error:", err)
				return ""
			}
		}

		// Flush changes to disk
		if err := file1.Sync(); err != nil {
			fmt.Println("Error:", err)
			return ""
		}

		return firstLine
	}

	// File is empty or an error occurred
	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}

	return ""
}

func Get3MTokens() int {

	filePath := "./Data/3 Month Tokens.txt"

	// Open the file.
	file, err := os.Open(filePath)
	if err != nil {
		LogPanic("Failed to open file", "Error", err.Error())
	}
	defer file.Close()

	// Create a scanner to read the file line by line.
	scanner := bufio.NewScanner(file)

	lineCount := 0

	// Loop through each line in the file.
	for scanner.Scan() {
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		LogPanic("Error scanning file", "Error", err.Error())
	}

	return lineCount
}

func Get1mTokens() int {

	filePath := "./Data/1 Month Tokens.txt"

	// Open the file.
	file, err := os.Open(filePath)
	if err != nil {
		LogPanic("Failed to open file", "Error", err.Error())
	}
	defer file.Close()

	// Create a scanner to read the file line by line.
	scanner := bufio.NewScanner(file)

	lineCount := 0

	// Loop through each line in the file.
	for scanner.Scan() {
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		LogPanic("Error scanning file", "Error", err.Error())
	}

	return lineCount
}

func Proxy() string {
	var proxy string
	file, err := os.Open("./Data/proxies.txt")
	stat, _ := os.Stat("./Data/proxies.txt")
	if err != nil {
		LogError("Failed Opening Proxies File", "Error", err.Error())
		err1 := errors.New("Failed Opening Proxies File")
		return err1.Error()
	}
	defer file.Close()

	if stat.Size() == 0 {
		LogError("No Proxies in Proxies.txt File, Failed to Boost!", "Error", err.Error())
		err2 := errors.New("No Proxies in Proxies.txt File, Failed to Boost!")
		return err2.Error()
	} else {
		scanner := bufio.NewScanner(file)
		lines := []string{}

		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}

		rand.Seed(time.Now().UnixNano())
		randomIndex := rand.Intn(len(lines))

		proxy = lines[randomIndex]

	}
	return proxy
}

func ImageToB64(path string) string {
	file, err := os.Open(path)

	if strings.Contains(path, "https") {
		LogError("No URL's Allowed in Config File for Banner or PFP, Must Enter Name of the File in the Avatars or Banners Folder", "Error", err.Error())
		err1 := errors.New("No URL's Allowed in Config File for Banner or PFP, Must Enter Name of the File in the Avatars or Banners Folder")
		return err1.Error()
	}
	if err != nil {
		LogError("Failed to Open Image, Check Name in Config File!", "Error", err.Error())
		err1 := errors.New("Failed to Open Image, Check Name in Config File!")
		return err1.Error()
	}

	defer file.Close()

	data, err := io.ReadAll(file)

	if err != nil {
		LogPanic(err.Error(), "", "")
	}

	return "data:image/png;base64," + base64.StdEncoding.EncodeToString(data)
}

func copyFiles() error {
	// Read the content of the first file
	content1, err := ioutil.ReadFile("./Data/3 Month Tokens.txt")
	if err != nil {
		return err
	}

	// Read the content of the second file
	content2, err := ioutil.ReadFile("./Data/1 Month Tokens.txt")
	if err != nil {
		return err
	}

	// Open the output file for writing, creating it if it doesn't exist
	output, err := os.OpenFile("./Data/Onliner Tokens.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer output.Close()

	// Write the contents of both files to the output file
	if _, err := output.Write(content1); err != nil {
		return err
	}
	if _, err := output.Write(content2); err != nil {
		return err
	}

	return nil
}

func CheckPermissions(userID string) bool {
	owners := config.DiscordSettings.Owners

	for _, owner := range owners {
		if userID == owner {
			return true
		}
	}

	return false
}

func readLinesFromFile(filename string) ([]string, error) {
	var lines []string

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Fields(line)
		if len(tokens) > 0 {
			lines = append(lines, tokens...)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func OnlinerTokens() []string {

	filePaths := []string{"./Data/1 Month Tokens.txt", "./Data/3 Month Tokens.txt", "./Data/used.txt"}

	var allTokens []string

	for _, filePath := range filePaths {
		tokens, err := readLinesFromFile(filePath)
		if err != nil {
			fmt.Printf("Error reading %s: %v\n", filePath, err)
			continue
		}
		allTokens = append(allTokens, tokens...)
	}

	return allTokens
}

func ExtractLinesAndSave(inputFilePath int, numLinesToExtract int) error {
	var file string

	if inputFilePath == 1 {
		file = "1 Month Tokens.txt"
	} else if inputFilePath == 3 {
		file = "3 Month Tokens.txt"
	}
	outputFile, err := os.Create("./Data/tokens.txt")
	if err != nil {
		return err
	}

	linesExtracted := 0

	for {
		line := SendToken(file)

		// Write the line to the output file
		_, err := outputFile.WriteString(line + "\n")
		if err != nil {
			return err
		}

		linesExtracted++

		// Check if we have extracted the desired number of lines
		if linesExtracted >= numLinesToExtract {
			break
		}
	}

	outputFile.Close()
	return nil
}

type Cycle struct {
	Mutex  *sync.Mutex
	Locked []string
	List   []string
	I      int

	WaitTime time.Duration
}

func New(List *[]string) *Cycle {
	rand.Seed(time.Now().UnixNano())

	return &Cycle{
		WaitTime: 50 * time.Millisecond,
		Mutex:    &sync.Mutex{},
		Locked:   []string{},
		List:     *List,
		I:        0,
	}
}

func NewFromFile(Path string) (*Cycle, error) {
	file, err := os.Open(fmt.Sprintf("./Data/%v", Path))
	if err != nil {
		return nil, err
	}
	var lines []string

	defer file.Close()
	defer func() {
		lines = nil
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return New(&lines), nil
}

func (c *Cycle) RandomiseIndex() {
	c.I = rand.Intn(len(c.List)-1) + 1
}

func (c *Cycle) IsInList(Element string) bool {
	for _, v := range c.List {
		if Element == v {
			return true
		}
	}
	return false
}

func (c *Cycle) IsLocked(Element string) bool {
	for _, v := range c.Locked {
		if Element == v {
			return true
		}
	}
	return false
}

func isInList(List *[]string, Element *string) bool {
	for _, v := range *List {
		if *Element == v {
			return true
		}
	}
	return false
}

func (c *Cycle) Next() string {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	for {
		c.I++
		if c.I >= len(c.List) {
			c.I = 0
		}

		if !c.IsLocked(c.List[c.I]) {
			return c.List[c.I]
		}

		time.Sleep(c.WaitTime)
	}
}

func (c *Cycle) Lock(Element string) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	if c.IsInList(Element) {
		c.Locked = append(c.Locked, Element)
	}
}

func (c *Cycle) Unlock(Element string) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	for i, v := range c.Locked {
		if Element == v {
			c.Locked = append(c.Locked[:i], c.Locked[i+1:]...)
		}
	}
}

func (c *Cycle) ClearDuplicates() int {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	removed := 0
	var list []string
	for _, v := range c.List {
		if !isInList(&list, &v) {
			list = append(list, v)
		} else {
			removed++
		}
	}
	c.List = list
	list = nil

	return removed
}

func (c *Cycle) Remove(Element string) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	for i, v := range c.List {
		if Element == v {
			c.List = append(c.List[:i], c.List[i+1:]...)
		}
	}

	for i, v := range c.Locked {
		if Element == v {
			c.Locked = append(c.Locked[:i], c.Locked[i+1:]...)
		}
	}
}

func (c *Cycle) LockByTimeout(Element string, Timeout time.Duration) {
	defer c.Unlock(Element)

	c.Lock(Element)
	time.Sleep(Timeout)
}

func AppendTextToFile(text, file string, extra ...string) {
	f, err := os.OpenFile(fmt.Sprintf("./Data/%v", file), os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		LogError("Failed to write to a file!", "Error", err.Error())
	}
	var t string
	if len(extra) != 0 {
		t = extra[0]
	}
	_, err = f.Write([]byte(t + text))
	if err != nil {
		AppendTextToFile(text, file)
		return
	}
	f.Close()
	return
}

func RemoveToken(text, fileName string) {
	file, err := os.Open(fmt.Sprintf("./Data/%v", fileName))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if line == text {
			continue
		}

		lines = append(lines, line)
	}

	err = file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	file, err = os.Create(fmt.Sprintf("./Data/%v", fileName))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	for _, line := range lines {
		_, err = fmt.Fprintln(file, line)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	err = file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func FormatToken(token string) string {
	if strings.Contains(token, ":") {
		split := strings.Split(token, ":")
		if len(split) == 3 {
			return split[2]
		} else if len(split) == 2 {
			return split[1]
		}
	}
	return token
}
