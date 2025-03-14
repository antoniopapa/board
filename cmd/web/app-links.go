package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type AppLink struct {
	Title            string
	Url              string
	DocumentationUrl string
	Image            string
	Description      string
	SpecialNote      string
	LDAP             bool
	Tiers            map[string]bool
}

func getAppLinks(tier, baseUri string) []AppLink {
	folders, err := os.ReadDir("/services")
	if err != nil {
		log.Fatal(err)
	}

	var links []AppLink

	for _, folder := range folders {
		envFile, _ := godotenv.Read(fmt.Sprintf("./services/%s/service.env", folder.Name()))
		dashboard := envFile["DASHBOARD"]
		if len(dashboard) == 0 {
			dashboard = "yes"
		}

		if dashboard == "no" {
			continue
		}

		ldap, _ := strconv.ParseBool(envFile["LDAP"])
		starter, _ := strconv.ParseBool(envFile["TIER_STARTER"])
		creator, _ := strconv.ParseBool(envFile["TIER_CREATOR"])
		teams, _ := strconv.ParseBool(envFile["TIER_TEAMS"])
		enterprise, _ := strconv.ParseBool(envFile["TIER_ENTERPRISE"])
		free, _ := strconv.ParseBool(envFile["TIER_FREE"])
		good, _ := strconv.ParseBool(envFile["TIER_GOOD"])
		better, _ := strconv.ParseBool(envFile["TIER_BETTER"])
		best, _ := strconv.ParseBool(envFile["TIER_BEST"])

		links = append(links, AppLink{
			Title:            envFile["TITLE"],
			Description:      envFile["DESCRIPTION"],
			Url:              strings.ReplaceAll(envFile["URL"], "{BASEURI}", baseUri),
			DocumentationUrl: envFile["DOCUMENTATION_URL"],
			Image:            envFile["IMAGE"],
			SpecialNote:      envFile["SPECIAL_NOTE"],
			LDAP:             ldap,
			Tiers: map[string]bool{
				"starter":    starter,
				"creator":    creator,
				"teams":      teams,
				"enterprise": enterprise,
				"free":       free,
				"good":       good,
				"better":     better,
				"best":       best,
			},
		})
	}

	//Sort in alphabetical order
	slices.SortFunc(links, func(a, b AppLink) int {
		return strings.Compare(a.Title, b.Title)
	})

	return links
}
