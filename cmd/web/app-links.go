package main

import (
	"slices"
	"strings"
)

type AppLink struct {
	Title string
	Url string
	Image string
	Description string
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
			Title: "User Management",
			Description: "Create and Manage Users",
			Image: "/static/img/users.png",
			Url: "https://panel." + baseUri + "/log_in/",
		},
		{
			Title: "Nextcloud",
			Description: "Email, Files, Documents",
			Image: "/static/img/nextcloud.png",
			Url: "https://nextcloud." + baseUri,
		},
		{
			Title: "Vaultwarden",
			Description: "Password Management",
			Image: "/static/img/vaultwarden.png",
			Url: "https://vaultwarden." + baseUri,
		},
		{
			Title: "Power DNS",
			Description: "DNS Management",
			Image: "/static/img/powerdns.png",
			Url: "https://powerdns." + baseUri,
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
		},
		{
			Title: "Wordpress",
			Description: "Your website",
			Image: "/static/img/wordpress.png",
			Url: "https://" + baseUri + "/wp-admin",
		},
	}

	return append(creatorLinks, getStarterLinks(baseUri)...)
}

func getTeamsLinks(baseUri string) []AppLink {
	teamsLinks := []AppLink {
		{
			Title: "Espo CRM",
			Description: "Customer relationship manager",
			Image: "/static/img/espo.png",
			Url: "https://espocrm." + baseUri,
		},
		{
			Title: "FreeScout",
			Description: "Customer Help Desk",
			Image: "/static/img/freescout.png",
			Url: "https://freescout." + baseUri,
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
		},
		{
			Title: "Listmonk",
			Description: "Email Marketing",
			Image: "/static/img/listmonk.png",
			Url: "https://listmonk." + baseUri,
		},
		{
			Title: "Baserow",
			Description: "Visual Databases",
			Image: "/static/img/baserow.png",
			Url: "https://baserow." + baseUri,
		},
		{
			Title: "Bookstack",
			Description: "Wiki Knowledgebase",
			Image: "/static/img/bookstack.png",
			Url: "https://bookstack." + baseUri,
		},
		{
			Title: "Gitea",
			Description: "GIT Source Control",
			Image: "/static/img/gitea.png",
			Url: "https://gitea." + baseUri,
		},
		{
			Title: "Castopod",
			Description: "Podcast Distribution",
			Image: "/static/img/castopod.png",
			Url: "https://castopod." + baseUri + "/cp-auth/login",
		},	
	}

	return append(enterpriseLinks, getTeamsLinks(baseUri)...)
}