package main

import (
	"time"
)

type CloudflareZone struct {
	ZoneID string
	Name   string
}

type CloudflareRecords struct {
	RecordID string
	Name     string
	Content  string
}
type CloudflareRecordsResponse struct {
	Errors   []interface{} `json:"errors"`
	Messages []interface{} `json:"messages"`
	Result   []struct {
		Content   string    `json:"content"`
		Name      string    `json:"name"`
		Proxied   bool      `json:"proxied"`
		Type      string    `json:"type"`
		Comment   string    `json:"comment"`
		CreatedOn time.Time `json:"created_on"`
		ID        string    `json:"id"`
		Locked    bool      `json:"locked"`
		Meta      struct {
			AutoAdded bool   `json:"auto_added"`
			Source    string `json:"source"`
		} `json:"meta"`
		ModifiedOn time.Time `json:"modified_on"`
		Proxiable  bool      `json:"proxiable"`
		Tags       []string  `json:"tags"`
		TTL        int       `json:"ttl"`
		ZoneID     string    `json:"zone_id"`
		ZoneName   string    `json:"zone_name"`
	} `json:"result"`
	Success    bool `json:"success"`
	ResultInfo struct {
		Count      int `json:"count"`
		Page       int `json:"page"`
		PerPage    int `json:"per_page"`
		TotalCount int `json:"total_count"`
	} `json:"result_info"`
}

type CloudflareZoneResponse struct {
	Errors     []interface{} `json:"errors"`
	Messages   []interface{} `json:"messages"`
	Success    bool          `json:"success"`
	ResultInfo struct {
		Count      int `json:"count"`
		Page       int `json:"page"`
		PerPage    int `json:"per_page"`
		TotalCount int `json:"total_count"`
	} `json:"result_info"`
	Result []struct {
		Account struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"account"`
		ActivatedOn     time.Time `json:"activated_on"`
		CreatedOn       time.Time `json:"created_on"`
		DevelopmentMode int       `json:"development_mode"`
		ID              string    `json:"id"`
		Meta            struct {
			CdnOnly                bool `json:"cdn_only"`
			CustomCertificateQuota int  `json:"custom_certificate_quota"`
			DNSOnly                bool `json:"dns_only"`
			FoundationDNS          bool `json:"foundation_dns"`
			PageRuleQuota          int  `json:"page_rule_quota"`
			PhishingDetected       bool `json:"phishing_detected"`
			Step                   int  `json:"step"`
		} `json:"meta"`
		ModifiedOn          time.Time `json:"modified_on"`
		Name                string    `json:"name"`
		OriginalDnshost     string    `json:"original_dnshost"`
		OriginalNameServers []string  `json:"original_name_servers"`
		OriginalRegistrar   string    `json:"original_registrar"`
		Owner               struct {
			ID   string `json:"id"`
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"owner"`
		VanityNameServers []string `json:"vanity_name_servers"`
	} `json:"result"`
}
