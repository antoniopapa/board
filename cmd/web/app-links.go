package main

import (
	"slices"
	"strings"
)

type AppLink struct {
	Title string
	Url string
	DocumentationUrl string
	Image string
	Description string
	SpecialNote string
	LDAP bool
	Tiers map[string]bool
}

func getAppLinks(tier, baseUri string) []AppLink {
	var links []AppLink

	//Get Links Based on Tier, and replace link {BASEURI} with baseUri
	for _, link := range allAppLinks {
		if link.Tiers[tier] {
			link.Url = strings.ReplaceAll(link.Url, "{BASEURI}", baseUri)
			links = append(links, link)
		}
	}

	//Sort in alphabetical order
	slices.SortFunc(links, func(a, b AppLink) int {
		return strings.Compare(a.Title, b.Title)
	})

	return links
}

var allAppLinks = []AppLink{
	{
		Title: "Panel",
		Description: "Create & Manage Users",
		Image: "/static/img/users.png",
		Url: "https://panel.{BASEURI}/log_in/",
		LDAP: true,
		DocumentationUrl: "https://documentation.federated.computer/docs/core_applications/panel/",
		Tiers: map[string]bool {
			"starter": true,
			"creator": true,
			"teams": true,
			"enterprise": true,
			"free": true,
			"good": true,
			"better": true,
			"best": true,
		},
	},
	{
		Title: "Nextcloud",
		Description: "Email, Files, Documents",
		Image: "/static/img/nextcloud.png",
		Url: "https://nextcloud.{BASEURI}",
		LDAP: true,
		DocumentationUrl: "https://docs.nextcloud.com/server/latest/user_manual/en/",
		Tiers: map[string]bool {
			"starter": true,
			"creator": true,
			"teams": true,
			"enterprise": true,
			"free": false,
			"good": true,
			"better": true,
			"best": true,
		},
	},
	{
		Title: "Vaultwarden",
		Description: "Passwords",
		Image: "/static/img/vaultwarden.png",
		Url: "https://vaultwarden.{BASEURI}",
		LDAP: false,
		SpecialNote: "Must create your own admin user at setup.",
		DocumentationUrl: "https://bitwarden.com/help/onboarding-and-succession/",
		Tiers: map[string]bool {
			"starter": true,
			"creator": true,
			"teams": true,
			"enterprise": true,
			"free": true,
			"good": true,
			"better": true,
			"best": true,
		},
	},
	{
		Title: "Power DNS",
		Description: "DNS Management",
		Image: "/static/img/powerdns.png",
		Url: "https://powerdns.{BASEURI}",
		LDAP: false,
		DocumentationUrl: "https://doc.powerdns.com/",
		Tiers: map[string]bool {
			"starter": true,
			"creator": true,
			"teams": true,
			"enterprise": true,
			"free": true,
			"good": true,
			"better": true,
			"best": true,
		},
	},
	{
		Title: "Roundcube",
		Description: "Web Mail",
		Image: "/static/img/roundcube.png",
		Url: "https://roundcube.{BASEURI}",
		LDAP: true,
		DocumentationUrl: "https://roundcube.net/support/",
		Tiers: map[string]bool {
			"starter": true,
			"creator": true,
			"teams": true,
			"enterprise": true,
			"free": false,
			"good": true,
			"better": true,
			"best": true,
		},
	},
	{
		Title: "Element",
		Description: "Team Chat",
		Image: "/static/img/element.png",
		Url: "https://element.{BASEURI}",
		LDAP: true,
		DocumentationUrl: "https://element.io/user-guide",
		Tiers: map[string]bool {
			"starter": false,
			"creator": true,
			"teams": true,
			"enterprise": true,
			"free": false,
			"good": false,
			"better": true,
			"best": true,
		},
	},
	{
		Title: "Wordpress",
		Description: "Your Website",
		Image: "/static/img/wordpress.png",
		Url: "https://{BASEURI}/wp-admin",
		LDAP: false,
		DocumentationUrl: "https://wordpress.org/documentation/",
		Tiers: map[string]bool {
			"starter": false,
			"creator": true,
			"teams": true,
			"enterprise": true,
			"free": false,
			"good": true,
			"better": true,
			"best": true,
		},
	},
	{
		Title: "Espo CRM",
		Description: "Customer Relationship Manager",
		Image: "/static/img/espo.png",
		Url: "https://espocrm.{BASEURI}",
		LDAP: true,
		DocumentationUrl: "https://docs.espocrm.com/user-guide/sales-management/",
		Tiers: map[string]bool {
			"starter": false,
			"creator": false,
			"teams": true,
			"enterprise": true,
			"free": false,
			"good": false,
			"better": true,
			"best": true,
		},
	},
	{
		Title: "FreeScout",
		Description: "Customer Help Desk",
		Image: "/static/img/freescout.png",
		Url: "https://freescout.{BASEURI}",
		LDAP: false,
		DocumentationUrl: "https://github.com/freescout-help-desk/freescout/wiki/FAQ",
		Tiers: map[string]bool {
			"starter": false,
			"creator": false,
			"teams": true,
			"enterprise": true,
			"free": false,
			"good": false,
			"better": false,
			"best": true,
		},
	},
	{
		Title: "Jitsi",
		Description: "Video Chat",
		Image: "/static/img/jitsi.png",
		Url: "https://jitsi.{BASEURI}",
		LDAP: true,
		DocumentationUrl: "https://jitsi.github.io/handbook/docs/intro/",
		Tiers: map[string]bool {
			"starter": false,
			"creator": false,
			"teams": false,
			"enterprise": true,
			"free": false,
			"good": false,
			"better": true,
			"best": true,
		},
	},
	{
		Title: "Listmonk",
		Description: "Email Marketing",
		Image: "/static/img/listmonk.png",
		Url: "https://listmonk.{BASEURI}",
		LDAP: false,
		DocumentationUrl: "https://listmonk.app/docs/concepts/",
		Tiers: map[string]bool {
			"starter": false,
			"creator": false,
			"teams": false,
			"enterprise": true,
			"free": false,
			"good": false,
			"better": false,
			"best": false,
		},
	},
	{
		Title: "Baserow",
		Description: "Visual Databases",
		Image: "/static/img/baserow.png",
		Url: "https://baserow.{BASEURI}",
		LDAP: false,
		DocumentationUrl: "https://baserow.io/user-docs",
		Tiers: map[string]bool {
			"starter": false,
			"creator": false,
			"teams": false,
			"enterprise": true,
			"free": false,
			"good": false,
			"better": false,
			"best": true,
		},
	},
	{
		Title: "Bookstack",
		Description: "Wiki Knowledgebase",
		Image: "/static/img/bookstack.png",
		Url: "https://bookstack.{BASEURI}",
		LDAP: false,
		DocumentationUrl: "https://www.bookstackapp.com/docs/",
		Tiers: map[string]bool {
			"starter": false,
			"creator": false,
			"teams": false,
			"enterprise": true,
			"free": false,
			"good": false,
			"better": false,
			"best": true,
		},
	},
	{
		Title: "Gitea",
		Description: "GIT Source Control",
		Image: "/static/img/gitea.png",
		Url: "https://gitea.{BASEURI}",
		LDAP: false,
		DocumentationUrl: "https://docs.gitea.com/category/usage",
		Tiers: map[string]bool {
			"starter": false,
			"creator": false,
			"teams": false,
			"enterprise": true,
			"free": false,
			"good": false,
			"better": false,
			"best": true,
		},
	},
	{
		Title: "Castopod",
		Description: "Podcast Distribution",
		Image: "/static/img/castopod.png",
		Url: "https://castopod.{BASEURI}/cp-auth/login",
		LDAP: false,
		Tiers: map[string]bool {
			"starter": false,
			"creator": false,
			"teams": false,
			"enterprise": true,
			"free": false,
			"good": false,
			"better": false,
			"best": false,
		},
	},	
	{
		Title: "Plane",
		Description: "Project Management",
		Image: "/static/img/plane.png",
		Url: "https://castopod.{BASEURI}/",
		LDAP: false,
		Tiers: map[string]bool {
			"starter": false,
			"creator": false,
			"teams": false,
			"enterprise": false,
			"free": false,
			"good": false,
			"better": false,
			"best": true,
		},
	},
	{
		Title: "Cal.com",
		Description: "Scheduling",
		Image: "/static/img/cal.png",
		Url: "https://calcom.{BASEURI}/",
		LDAP: false,
		SpecialNote: "Coming Soon...",
		Tiers: map[string]bool {
			"starter": false,
			"creator": false,
			"teams": false,
			"enterprise": false,
			"free": false,
			"good": true,
			"better": true,
			"best": true,
		},
	},
}
