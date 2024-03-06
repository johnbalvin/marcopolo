package main

import (
	"fmt"
	"marcopolo/asn"
	"marcopolo/utils"
	"time"
)

var SecureState = Input{
	URL:        utils.ParseURL("https://secure.state.co.nz/car"),
	Keyworkds:  []string{"State Insurance", "secure.state.co.nz/car/favicon.ico"},
	TCPTimeout: time.Second,
	BufferSize: 2048,
	Asn: asn.Asn{
		PrioritiesNames: []string{"IAG New Zealand"},
		ForbiddenNames:  ForbidenASN,
	},
}
var Mouser = Input{
	URL:        utils.ParseURL("https://www.mouser.de"),
	TCPTimeout: time.Millisecond * 250,
	Keyworkds:  []string{"Distributor", "Deutschland"},
	BufferSize: 2048,
	Asn: asn.Asn{
		PrioritiesNames: []string{"mouser electronics", "Rogers Communications"},
		ForbiddenNames:  ForbidenASN,
	},
}
var VerifySos = Input{
	URL:        utils.ParseURL("https://verify.sos.ga.gov/verification/Search.aspx"),
	Keyworkds:  []string{"Verification", "elicense2000"},
	TCPTimeout: time.Millisecond * 250,
	BufferSize: 2048,
	Asn: asn.Asn{
		PrioritiesNames: []string{"Quality Technology Services" /*, "Comcast Cable"*/},
		ForbiddenNames:  ForbidenASN,
	},
}

var Maybank2u = Input{
	URL:        utils.ParseURL("https://www.maybank2u.com.my/home/m2u/common/login.do"),
	Keyworkds:  []string{"Maybank2u", "Malaysia"},
	TCPTimeout: time.Millisecond * 250,
	BufferSize: 2048,
	Asn: asn.Asn{
		PrioritiesNames: []string{"Binariang Berh", "Philippine Long Distan", "NTT America", "CenturyLink", "Net Onboard Sdn", "Arcnet", "TM TECHNOLOGY", "Reliance Jio Infoco"},
		ForbiddenNames:  ForbidenASN,
	},
}
var DeutscheBank = Input{
	URL:        utils.ParseURL("https://www.db.com/index?language_id=1&kid=sl.redirect-en.shortcut"),
	Keyworkds:  []string{"Deutsche Bank", "dwebcms"},
	TCPTimeout: time.Millisecond * 250,
	BufferSize: 2048,
	Asn: asn.Asn{
		PrioritiesNames: []string{"Deutsche Bank", "google"},
		ForbiddenNames:  ForbidenASN,
	},
}
var Rhbgroup = Input{
	URL:        utils.ParseURL("https://onlinebanking.rhbgroup.com/my/login"),
	Keyworkds:  []string{"Online banking"},
	TCPTimeout: time.Millisecond * 250,
	BufferSize: 2048,
	Asn: asn.Asn{
		PrioritiesNames: []string{"OSK Securities", "TM TECHNOLOGY SERVICES", "MaxisNet International Int", "TTNET-MY", "Amazon", "Binariang Berhad", "Digital Singapore", "AIMS Data Centre", "PT Cyberindo", "Forcepoint Cloud", "Telekomunikasi", "SingNet", "Gilead Sciences", "Cyberindo", "HGC Global"},
		ForbiddenNames:  ForbidenASN,
	},
}
var Kroger = Input{
	URL:        utils.ParseURL("https://www.kroger.com"),
	Keyworkds:  []string{"Kroger", "Groceries"},
	TCPTimeout: time.Millisecond * 250,
	BufferSize: 2048,
	Asn: asn.Asn{
		PrioritiesNames: []string{"The Kroger", "google" /*"Sprint"*/},
		ForbiddenNames:  ForbidenASN,
	},
}
var Nike = Input{
	URL:        utils.ParseURL("https://www.nike.com"),
	Keyworkds:  []string{"Nike delivers", "athletes"},
	TCPTimeout: time.Millisecond * 250,
	BufferSize: 2048,
	Asn: asn.Asn{
		PrioritiesNames: []string{"Verizon", "amazon", "Fastly"},
		ForbiddenNames:  ForbidenASN,
	},
}

var Disney = Input{
	URL:        utils.ParseURL("https://disneyworld.disney.go.com"),
	Keyworkds:  []string{"50th_anniversary_countdown_clock"},
	TCPTimeout: time.Millisecond * 250,
	BufferSize: 2048,
	Asn: asn.Asn{
		PrioritiesNames: []string{"Rackspace Hosting", "Disney Worldwide", "Google"},
		ForbiddenNames:  ForbidenASN,
	},
}
var Walmart = Input{
	URL:        utils.ParseURL("https://www.walmart.com"),
	Keyworkds:  []string{"walmartimages", "Save Money", "free delivery"},
	TCPTimeout: time.Millisecond * 250,
	BufferSize: 7048,
	Asn: asn.Asn{
		PrioritiesNames: []string{"Cox Communications", "Rackspace", "Wal-Mart", "CyrusOne", "google", "amazon", "NTT America", "AT&T Services", "Steadfast"},
		ForbiddenNames:  ForbidenASN,
	},
}
var Adidas = Input{
	URL:        utils.ParseURL("https://www.adidas.com/us"),
	Keyworkds:  []string{"adidas", "Official", "running"},
	TCPTimeout: time.Millisecond * 250,
	BufferSize: 2048,
	Asn: asn.Asn{
		PrioritiesNames: []string{"Globe Telecoms", "noris network", "CenturyLink Communications", "amazon"},
		ForbiddenNames:  ForbidenASN,
	},
}

var Adobe = Input{
	URL:        utils.ParseURL("https://www.adobe.com"),
	Keyworkds:  []string{"Adobe", "marketing and document"},
	TCPTimeout: time.Millisecond * 250,
	BufferSize: 2048,
	Asn: asn.Asn{
		PrioritiesNames: []string{"Globe Telecom", "StarHub", "Adobe", "Fastly", "Amazon"},
		ForbiddenNames:  ForbidenASN,
	},
}
var Selesforce = Input{
	URL:        utils.ParseURL("https://www.salesforce.com"),
	Keyworkds:  []string{"Salesforce", "Customer"},
	TCPTimeout: time.Millisecond * 250,
	BufferSize: 2048,
	Asn: asn.Asn{
		PrioritiesNames: []string{"Salesforce"},
		ForbiddenNames:  ForbidenASN,
	},
}
var Microsoft = Input{
	URL:        utils.ParseURL("https://www.microsoft.com"),
	Keyworkds:  []string{"Microsoft", "Cloud", "Computers"},
	TCPTimeout: time.Millisecond * 250,
	BufferSize: 2048,
	Asn: asn.Asn{
		PrioritiesNames: []string{"GTT Communications", "Microsoft", "Edgecast"},
		ForbiddenNames:  ForbidenASN,
	},
}
var BestBuy = Input{
	URL:        utils.ParseURL("https://www.bestbuy.com"),
	Keyworkds:  []string{"Best Buy", "Shop Now"},
	TCPTimeout: time.Millisecond * 250,
	BufferSize: 2048,
	Asn: asn.Asn{
		PrioritiesNames: []string{"Rogers Communicatio", "Rackspace", "Amazon", "Best Buy", "GTT Communications"},
		ForbiddenNames:  ForbidenASN,
	},
}

var InstantCart = Input{
	URL:        utils.ParseURL("https://www.instacart.com"),
	Keyworkds:  []string{"delivery", "pickup", "grocers"},
	TCPTimeout: time.Millisecond * 250,
	BufferSize: 2048,
	Asn: asn.Asn{
		PrioritiesNames: []string{"Amazon", "google"},
		ForbiddenNames:  ForbidenASN,
	},
}
var Iberia = Input{
	URL:        utils.ParseURL("https://www.iberia.com/us/"),
	Keyworkds:  []string{"IBERIA.COM", "best prices"},
	TCPTimeout: time.Millisecond * 250,
	BufferSize: 2048,
	Asn: asn.Asn{
		PrioritiesNames: []string{"IBERIA LINEAS AEREAS", "Amazon"},
		ForbiddenNames:  ForbidenASN,
	},
}
var Cars = Input{
	URL:        utils.ParseURL("https://www.cars.com"),
	Keyworkds:  []string{"New Cars", "Dealers"},
	TCPTimeout: time.Millisecond * 250,
	BufferSize: 2048,
	Asn: asn.Asn{
		PrioritiesNames: []string{"Classified Venture", "Amazon"},
		ForbiddenNames:  ForbidenASN,
	},
}
var starngage = Input{
	URL:        utils.ParseURL("https://starngage.com/plus/en-us"),
	Keyworkds:  []string{"Influencer", "Agencies"},
	TCPTimeout: time.Millisecond * 250,
	BufferSize: 2048,
	Asn: asn.Asn{
		PrioritiesNames: []string{"Amazon", "GoDaddy"},
		ForbiddenNames:  ForbidenASN,
	},
}
var Fansale = Input{
	URL:        utils.ParseURL("https://www.fansale.de/fansale/"),
	Keyworkds:  []string{"fanSALE", "Tickets"},
	TCPTimeout: time.Millisecond * 250,
	BufferSize: 2048,
	Asn: asn.Asn{
		PrioritiesNames: []string{"Mobile Telecommunication Company Saudi Arabia", "CTS Eventim"},
		ForbiddenNames:  ForbidenASN,
	},
}
var smiles = Input{
	URL:        utils.ParseURL("https://www.smiles.com.br"),
	Keyworkds:  []string{"Milhas para quem"},
	TCPTimeout: time.Millisecond * 250,
	BufferSize: 2048,
	Asn: asn.Asn{
		PrioritiesNames: []string{"Cox", "amazon", "Cisco", "Microsoft"},
		ForbiddenNames:  ForbidenASN,
	},
}
var RealtorCa = Input{
	URL:        utils.ParseURL("https://www.realtor.ca"),
	Keyworkds:  []string{"Real Estate", "realtor"},
	TCPTimeout: time.Millisecond * 250,
	BufferSize: 2048,
	Asn: asn.Asn{
		PrioritiesNames: []string{"Microsoft", "DigitalOcean", "Rogers Communications"},
		ForbiddenNames:  ForbidenASN,
	},
}
var Crunchbase = Input{
	URL:        utils.ParseURL("https://www.crunchbase.com"),
	Keyworkds:  []string{"Crunchbase: Discover innovative"},
	TCPTimeout: time.Millisecond * 250,
	BufferSize: 2048,
	Asn: asn.Asn{
		PrioritiesNames: []string{"amazon", "Google"},
		ForbiddenNames:  ForbidenASN,
	},
}

// Hetzner be aware of Hetzner
var platesmania = Input{
	URL:        utils.ParseURL("https://platesmania.com"),
	Keyworkds:  []string{"Photos of vehicles"},
	TCPTimeout: time.Millisecond * 250,
	Asn: asn.Asn{
		PrioritiesNames: []string{"GoDaddy", "Hetzner", "OVH"},
		ForbiddenNames:  ForbidenASN,
	},
}

func (input *Input) setInputs() {
	path := input.URL.Path
	if path == "" {
		path = "/"
	}
	input.request = []byte(fmt.Sprintf("GET %s HTTP/1.1\r\nHost: %s\r\n", path, input.URL.Host))
	for key, values := range headers {
		input.request = append(input.request, []byte(fmt.Sprintf("%s: %s\r\n", key, values[0]))...)
	}
	input.request = append(input.request, []byte("\r\n\r\n")...)

	for _, keyword := range input.Keyworkds {
		input.keyworkds = append(input.keyworkds, []byte(keyword))
	}
}
