package main

import "encoding/xml"

type PageData struct {
	RSSFluxArrays map[string][]Flux
	RSSData       map[string]*RSS
	RDFData       map[string]*RDF
	StringArrays  map[string][]string
	IntArrays     map[string][]int
	RSSVersion 		bool
	//	Uint16Arrays map[string][]uint16
}

type Flux struct {
	Name string
	Link string
	Version string
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

type Request struct {
	Type string
	Id   uint
	Flx  Flux
}

type RDF struct {
	XMLName xml.Name `xml:"rdf"`
	Channel RDFChannel  `xml:"channel"`
	Items   []RDFItem   `xml:"item"`
}

type RDFChannel struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Link        string `xml:"link"`
}

type RDFItem struct {
	About       string `xml:"about,attr"`
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Link        string `xml:"link"`
	Date        string `xml:"date"`
	Creator     string `xml:"creator"`
}