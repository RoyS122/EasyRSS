package main

type PageData struct {
	RSSFluxArrays map[string][]Flux
	RSSData       map[string]*RSS
	StringArrays  map[string][]string
	IntArrays     map[string][]int
	//	Uint16Arrays map[string][]uint16
}

type Flux struct {
	Name string
	Link string
}
type RSS struct {
	Name    string
	Channel Channel `xml:"channel"`
}

type Channel struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Items       []Item `xml:"item"`
}

type Item struct {
	Title       string    `xml:"title"`
	Description string    `xml:"description"`
	Link        string    `xml:"link"`
	PubDate     string    `xml:"pubDate"`
	Enclosure   Enclosure `xml:"enclosure"`
}

type Enclosure struct {
	URL    string `xml:"url,attr"`
	Length string `xml:"length,attr"`
	Type   string `xml:"type,attr"`
}

type Categorie struct {
	Title       string
	Description string
	Color       string
}