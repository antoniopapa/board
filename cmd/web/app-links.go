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
}

func getAppLinks(tier, baseUri string) []AppLink {
	var output []AppLink

	//Get Links Based on Tier
	if tier == "enterprise" {
		output = getEnterpriseLinks(baseUri)
	} else if tier == "creator" {
		output = getCreatorLinks(baseUri)
	} else if tier == "teams" {
		output = getTeamsLinks(baseUri)
	} else {
		output = getStarterLinks(baseUri)
	}

	//Sort in alphabetical order
	slices.SortFunc(output, func(a, b AppLink) int {
		return strings.Compare(a.Title, b.Title)
	})

	return output
}

func getStarterLinks(baseUri string) []AppLink {
	return []AppLink {
		{
			Title: "Panel",
			Description: "Create & Manage Users",
			Image: "/static/img/users.png",
			Url: "https://panel." + baseUri + "/log_in/",
			LDAP: true,
			DocumentationUrl: "https://documentation.federated.computer/docs/core_applications/panel/",
		},
		{
			Title: "Nextcloud",
			Description: "Email, Files, Documents",
			Image: "/static/img/nextcloud.png",
			Url: "https://nextcloud." + baseUri,
			LDAP: true,
			DocumentationUrl: "https://docs.nextcloud.com/server/latest/user_manual/en/",
		},
		{
			Title: "Vaultwarden",
			Description: "Passwords",
			Image: "/static/img/vaultwarden.png",
			Url: "https://vaultwarden." + baseUri,
			LDAP: false,
			SpecialNote: "Must create your own admin user at setup.",
			DocumentationUrl: "https://bitwarden.com/help/onboarding-and-succession/",
		},
		{
			Title: "Power DNS",
			Description: "DNS Management",
			Image: "/static/img/powerdns.png",
			Url: "https://powerdns." + baseUri,
			LDAP: false,
			DocumentationUrl: "https://doc.powerdns.com/",
		},
		{
			Title: "Web Mail",
			Description: "Email",
			Image: "/static/img/roundcube.png",
			Url: "https://roundcube." + baseUri,
			LDAP: false,
			DocumentationUrl: "https://roundcube.net/support/",
		},
	}
}

func getCreatorLinks(baseUri string) []AppLink {
	creatorLinks := []AppLink {
		{
			Title: "Element",
			Description: "Team Chat",
			Image: "/static/img/element.png",
			Url: "https://element." + baseUri,
			LDAP: true,
			DocumentationUrl: "https://element.io/user-guide",
		},
		{
			Title: "Wordpress",
			Description: "Your Website",
			Image: "/static/img/wordpress.png",
			Url: "https://" + baseUri + "/wp-admin",
			LDAP: false,
			DocumentationUrl: "https://wordpress.org/documentation/",
		},
	}

	return append(creatorLinks, getStarterLinks(baseUri)...)
}

func getTeamsLinks(baseUri string) []AppLink {
	teamsLinks := []AppLink {
		{
			Title: "Espo CRM",
			Description: "Customer Relationship Manager",
			Image: "/static/img/espo.png",
			Url: "https://espocrm." + baseUri,
			LDAP: true,
			DocumentationUrl: "https://docs.espocrm.com/user-guide/sales-management/",
		},
		{
			Title: "FreeScout",
			Description: "Customer Help Desk",
			Image: "/static/img/freescout.png",
			Url: "https://freescout." + baseUri,
			LDAP: false,
			DocumentationUrl: "https://github.com/freescout-help-desk/freescout/wiki/FAQ",
		},
	}

	return append(teamsLinks, getCreatorLinks(baseUri)...)
}

func getEnterpriseLinks(baseUri string) []AppLink {
	enterpriseLinks := []AppLink{
		{
			Title: "Jitsi",
			Description: "Video Chat",
			Image: "/static/img/jitsi.png",
			Url: "https://jitsi." + baseUri,
			LDAP: true,
			DocumentationUrl: "https://jitsi.github.io/handbook/docs/intro/",
		},
		{
			Title: "Listmonk",
			Description: "Email Marketing",
			Image: "/static/img/listmonk.png",
			Url: "https://listmonk." + baseUri,
			LDAP: false,
			DocumentationUrl: "https://listmonk.app/docs/concepts/",
		},
		{
			Title: "Baserow",
			Description: "Visual Databases",
			Image: "/static/img/baserow.png",
			Url: "https://baserow." + baseUri,
			LDAP: false,
			DocumentationUrl: "https://baserow.io/user-docs",
		},
		{
			Title: "Bookstack",
			Description: "Wiki Knowledgebase",
			Image: "/static/img/bookstack.png",
			Url: "https://bookstack." + baseUri,
			LDAP: false,
			DocumentationUrl: "https://www.bookstackapp.com/docs/",
		},
		{
			Title: "Gitea",
			Description: "GIT Source Control",
			Image: "/static/img/gitea.png",
			Url: "https://gitea." + baseUri,
			LDAP: false,
			DocumentationUrl: "https://docs.gitea.com/category/usage",
		},
		{
			Title: "Castopod",
			Description: "Podcast Distribution",
			Image: "/static/img/castopod.png",
			Url: "https://castopod." + baseUri + "/cp-auth/login",
			LDAP: false,
		},	
	}

	return append(enterpriseLinks, getTeamsLinks(baseUri)...)
}