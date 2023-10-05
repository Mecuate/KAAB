package fixtures

type FileDataType = struct {
	FileName string
	EndPoint []string
	Data     map[string]string
}

type EndpointData = map[string]FileDataType

var TablesData = EndpointData{
	"a05s": FileDataType{
		FileName: "paragraph",
		EndPoint: []string{`var vals = 5;
		var welcome = [context.welcome, "appended_data", 3];
		var goodbye = [context.goodbye, vals];
		var wiki = context.wiki;
		var contact = context.contact;
		var publik = context.publik;
		var result = [welcome, "dullabot", goodbye, wiki, contact, publik];
		return result;
		`, `return context;`},
		Data: map[string]string{
			"welcome": "Bienvenidos",
			"goodbye": "Hasta luego",
			"wiki":    "https://es.wikipedia.org/wiki/Astrophytum",
			"contact": "Buscanos en redes",
			"publik":  "este es un sitio web publico",
		},
	},
	"a02s": FileDataType{
		FileName: "academy",
		EndPoint: []string{`return context;`, `return context;`},
		Data: map[string]string{
			"academy:frontPage.title":               "Bienvenido a la Academia mecuate astrophytum",
			"academy:frontPage.description":         "this academy is made for all peoples from anywhrer in the universe.",
			"academy:workShopPage.groupName":        "Nombre de Grupo",
			"academy:workShopPage.groupType.solid":  "Grupo sólido",
			"academy:workShopPage.groupType.stable": "Grupo fijo",
			"academy:workShopPage.groupType.morph":  "Grupo Anomórfico",
		},
	},
	"a03s": FileDataType{
		FileName: "common",
		Data: map[string]string{
			"common:menu.home":     "home",
			"common:menu.mecuate":  "Mecuate",
			"common:menu.account":  "My Account",
			"common:menu.academy":  "Academy",
			"common:menu.services": "Services",
		},
	},
	"a04s": FileDataType{
		FileName: "kaab",
		Data: map[string]string{
			"kaab.common:menu.docuweb":      "DocuWeb",
			"kaab.common:menu.projects":     "Projects",
			"kaab.common:menu.publications": "Publications",
			"kaab.common:menu.contact":      "Contact",
			"kaab.common:menu.production":   "Production",
		},
	},
	"a01s": FileDataType{
		FileName: "actions",
		Data: map[string]string{
			"actions:sections.previousWorkShops": "Previous",
			"actions:sections.nextWorkShops":     "Comming next",
			"actions:sections.currentWorkShops":  "Current",
			"actions:sections.blog":              "Other media",
		},
	},
}
