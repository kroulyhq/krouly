package main

import (
	"fmt"
	"os"

	"github.com/gocolly/colly"
	"gopkg.in/yaml.v2"
)

type Task struct {
	Name      string `yaml:"name"`
	Connector string `yaml:"connector"`
	Params    struct {
		URL string `yaml:"url"`
	} `yaml:"parameters"`
}

type Playbook struct {
	Name  string `yaml:"playbook"`
	Tasks []Task `yaml:"tasks"`
}

type KroulyConnector struct {
	URL string // URL for source
}

func NewKroulyConnector(url string) *KroulyConnector {
	// pre-init and config
	return &KroulyConnector{
		URL: url,
	}
}

func (c *KroulyConnector) ExtractData(collector *colly.Collector) error {
	// extract data from the specified URL using Colly
	collector.OnHTML("tr.simpTblRow", func(e *colly.HTMLElement) {
		// Find all rows with the class "simpTblRow"
		symbol := e.ChildText("a[data-test='quoteLink']")                      // Extract symbol from the link with data-test attribute "quoteLink"
		price := e.ChildText("td[aria-label='Price (Intraday)'] fin-streamer") // Extract price from the specified cell

		// Print the extracted data
		fmt.Printf("Symbol: %s, Price: %s\n", symbol, price)
	})

	// start extraction
	err := collector.Visit(c.URL)
	if err != nil {
		return fmt.Errorf("error visiting URL: %v", err)
	}

	return nil
}

func loadPlaybook(filename string) (*Playbook, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening playbook file: %v", err)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	var playbook Playbook
	err = decoder.Decode(&playbook)
	if err != nil {
		return nil, fmt.Errorf("error parsing playbook YAML: %v", err)
	}

	return &playbook, nil
}

func main() {
	// load yaml playbook from playbooks folder
	playbookFile := "../playbooks/krouly.sample.yaml"
	playbook, err := loadPlaybook(playbookFile)
	if err != nil {
		fmt.Println("Error loading playbook:", err)
		return
	}

	// create a new colly collector
	collector := colly.NewCollector()

	// iterate through all playbooks tasks
	for _, task := range playbook.Tasks {
		fmt.Println("Executing task:", task.Name)

		switch task.Connector {
		case "KroulyYahooCryptoConnector":
			connector := NewKroulyConnector(task.Params.URL)
			err := connector.ExtractData(collector)
			if err != nil {
				fmt.Printf("Error executing task %s: %v\n", task.Name, err)
			}
		default:
			fmt.Printf("Unknown connector type for task %s\n", task.Name)
		}
	}

	fmt.Println("Data extraction successful!")
}
