package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
	"flag"
	"github.com/robfig/cron/v3"
	"github.com/rodaine/table"
)


func forever() {
	for {
		time.Sleep(time.Hour)
	}
}

func startCron(config *Config, cron *cron.Cron, httpClient *http.Client) {
	// setup cron

	for _, record := range config.Cloudflare.Records {
		fmt.Print(record.Schedule)
		cron.AddFunc(record.Schedule, func() {
			fmt.Println("[Job] Updating ", record.Name)
			// update record
			success, err := updateRecord(config, httpClient, record.ZoneID, record.Id, record.Name)
			if err != nil {
				fmt.Printf("Error: %+v\n", err)
			}
			if success {
				fmt.Println("[Job OK] Updated for ", record.Name)
			}
		})
	}
	fmt.Printf("%+v\n", config)

	fmt.Printf("Cron started at %v+\n", time.Now())

	cron.Start()
	fmt.Printf("%+v\n", cron.Entries())
}

func printHelp() {
	helpString := `
Usage: 

	ddns [options] [command]

Commands:

	help				- Get this help text
	ip				- Get your current external IP address
	list zones			- List cloudflare zones
	list records [zoneID]		- List cloudflare records
	list jobs			- List your scheduled jobs in config
	start				- Start the cron job
	hammer				- Force update your jobs with current IP, use with a hammer
	
Options:

	-f [path]      - Provide config file path (default config.yaml)
	`
	fmt.Print(helpString)
	// fmt.Println("ddns help            - Get this help text")
	// fmt.Println("ddns ip              - Get your current external IP address")
	// fmt.Println("ddns list zones      - List cloudflare zones")
	// fmt.Println("ddns list records [zoneID] - List cloudflare records")
	// fmt.Println("ddns list jobs       - List your scheduled jobs in config")
	// fmt.Println("ddns start           - Start the cron job")
	// fmt.Println("ddns hammer          - Force update your jobs with current IP, use with a hammer")
	// fmt.Println("\n Arguments:")
	// fmt.Println("-f [path]            - Provide config file path (default config.yaml)")
}

func main() {
	// read args
	wd, _ := os.Getwd()
	var configPath string
	flag.StringVar(&configPath, "f", path.Join(wd, "config.yaml"), "Path to config file")
	flag.Parse()
	
	// fmt.Println(flag.Args())
	args := flag.Args()

	if len(args) == 0 || args[0] == "help" {
		printHelp()
		return
	}

	// setup cron
	cron := cron.New()

	// setup HTTP client
	client := &http.Client{}

	// Read config file
	config := getConfig(configPath)

	if args[0] == "ip" {
		ip := getIp(config)
		fmt.Println(ip)
	} else if args[0] == "help" {
		printHelp()
	} else if args[0] == "start" {
		startCron(config, cron, client)

		go forever()
		select {}
	} else if args[0] == "list" {

		if len(args) < 2 {
			fmt.Println("Specify resources to list, available resources are: zones, records")
			return

		}
		if args[1] == "zones" {
			// list zones
			zones, _ := listZones(config, client)
			tbl := table.New("Name", "ID")
			for _, zone := range zones {
				tbl.AddRow(zone.Name, zone.ZoneID)
			}
			tbl.Print()
			return

		} else if args[1] == "records" {
			if len(args) < 3 {
				fmt.Println("zone id required")
				return
			}
			// list records
			zoneId := args[2]
			records, _ := listRecords(config, client, zoneId)
			tbl := table.New("Name", "Content", "ID")
			for _, record := range records {
				tbl.AddRow(record.Name, record.Content, record.RecordID)
			}
			tbl.Print()
			return
		} else if args[1] == "jobs" {
			// list cron jobs
			fmt.Printf("%+v\n", cron.Entries())
		} else {
			fmt.Println("Specify resources to list, available resources are: zones, records")
		}
	} else if args[0] == "hammer" {
		for _, record := range config.Cloudflare.Records {
			fmt.Println("updating", record.Name)
			// update record
			success, err := updateRecord(config, client, record.ZoneID, record.Id, record.Name)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(success)
		}
	} else {
		fmt.Println("Unknown commands, run `ddns help` for help")
	}
	// ip := getIp()
	// fmt.Println(ip)
}

func updateRecord(config *Config, client *http.Client, zoneId string, recordId string, name string) (bool, error) {
	ip := getIp(config)
	fmt.Println("Updating record " + recordId)
	jsonData, err := json.Marshal(CloudflarePatchDNSBody{
		Content: ip,
		Name:    name,
	})
	if err != nil {
		log.Println(err)
	}

	req, err := http.NewRequest("PATCH", "https://api.cloudflare.com/client/v4/zones/"+zoneId+"/dns_records/"+recordId, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return false, err
	}
	req.Header.Add("Authorization", "Bearer "+config.Cloudflare.ApiKey)

	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
		return false, err
	}
	defer response.Body.Close()

	result := CloudflarePatchDNSResponse{}

	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		fmt.Println("Error decoding JSON response:", err)
		return false, err
	}

	if !result.Success {
		fmt.Printf("Error: %+v\n", result)
		return false, err
	}

	return true, nil
}

func listZones(config *Config, httpc *http.Client) ([]CloudflareZone, error) {
	req, err := http.NewRequest("GET", "https://api.cloudflare.com/client/v4/zones", nil)
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+config.Cloudflare.ApiKey)
	client := httpc

	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
		return nil, err
	}
	defer response.Body.Close()

	result := CloudflareZoneResponse{}

	json.NewDecoder(response.Body).Decode(&result)

	zones := []CloudflareZone{}


	for _, value := range result.Result {
		zones = append(zones, CloudflareZone{ZoneID: value.ID, Name: value.Name})
		// fmt.Println(index, value.Name, value.ID)
	}

	


	return zones, nil
}

func listRecords(config *Config, httpc *http.Client, zoneId string) ([]CloudflareRecords, error) {
	req, err := http.NewRequest("GET", "https://api.cloudflare.com/client/v4/zones/"+zoneId+"/dns_records", nil)
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+config.Cloudflare.ApiKey)

	response, err := httpc.Do(req)
	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
		return nil, err
	}
	defer response.Body.Close()

	result := CloudflareRecordsResponse{}

	json.NewDecoder(response.Body).Decode(&result)

	records := []CloudflareRecords{}

	for _, value := range result.Result {
		records = append(records, CloudflareRecords{RecordID: value.ID, Name: value.Name, Content: value.Content})
	}
	return records, nil
}

func getIp(config *Config) string {
	resp, err := http.Get(config.IpProvider)
	if err != nil {
		return getIp(config)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return getIp(config)
	}

	return strings.TrimSpace(string(body))
}
