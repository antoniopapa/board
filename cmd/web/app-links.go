package main

type AppLink struct {
	Title string
	Url string
	Image string
	Description string
}

func getAppLinks(tier, baseUri string) []AppLink {
	if tier == "enterprise" {
		return getEnterpriseLinks(baseUri)
	} else if tier == "creator" {
		return getCreatorLinks(baseUri)
	} else if tier == "teams" {
		return getTeamsLinks(baseUri)
	}

	return getStarterLinks(baseUri)
}

func getStarterLinks(baseUri string) []AppLink {
	return []AppLink {
		{
			Title: "User Management",
			Description: "Create users and manage their access",
			Image: "/img/panel.jpg",
			Url: "https://panel." + baseUri,
		},
		{
			Title: "Nextcloud",
			Description: "Email, Files, Documents",
			Image: "/img/nextcloud.jpg",
			Url: "https://nextcloud." + baseUri,
		},
		{
			Title: "Vaultwarden",
			Description: "Password Management",
			Image: "/img/vaultwarden.jpg",
			Url: "https://vaultwarden." + baseUri,
		},
	}
}

func getCreatorLinks(baseUri string) []AppLink {
	creatorLinks := []AppLink {
		{
			Title: "Element and Matrix",
			Description: "Team Chat",
			Image: "/img/element.jpg",
			Url: "https://element." + baseUri,
		},
		{
			Title: "Wordpress",
			Description: "Your website",
			Image: "/img/wordpress.jpg",
			Url: "https://" + baseUri,
		},
	}

	return append(creatorLinks, getStarterLinks(baseUri)...)
}

func getTeamsLinks(baseUri string) []AppLink {
	teamsLinks := []AppLink {
		{
			Title: "Espo CRM",
			Description: "Customer relationship manager",
			Image: "/img/espo.jpg",
			Url: "https://espocrm." + baseUri,
		},
		{
			Title: "FreeScout",
			Description: "Customer Help Desk",
			Image: "/img/freescout.jpg",
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
			Image: "/img/jitsi.jpg",
			Url: "https://jitsi." + baseUri,
		},
		{
			Title: "Listmonk",
			Description: "Email Marketing",
			Image: "/img/listmonk.jpg",
			Url: "https://listmonk." + baseUri,
		},
		{
			Title: "Baserow",
			Description: "Visual Databases",
			Image: "/img/baserow.jpg",
			Url: "https://baserow." + baseUri,
		},
		{
			Title: "Bookstack",
			Description: "Wiki Knowledgebase",
			Image: "/img/bookstack.jpg",
			Url: "https://bookstack." + baseUri,
		},
		{
			Title: "Gitea",
			Description: "GIT Source Control",
			Image: "/img/gitea.jpg",
			Url: "https://gitea." + baseUri,
		},
		{
			Title: "Castopod",
			Description: "Podcast Distribution",
			Image: "/img/castopod.jpg",
			Url: "https://castopod." + baseUri,
		},	
	}

	return append(enterpriseLinks, getTeamsLinks(baseUri)...)
}