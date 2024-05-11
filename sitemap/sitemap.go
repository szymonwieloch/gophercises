package main

import (
	"encoding/xml"
	"log"
	"os"
)

/* Expected format:

<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
  <url>
    <loc>http://www.example.com/</loc>
  </url>
  <url>
    <loc>http://www.example.com/dogs</loc>
  </url>
</urlset>

*/

type Loc struct {
	XMLName xml.Name `xml:"url"`
	Loc     string   `xml:"loc"`
}

type URLSet struct {
	XMLName xml.Name `xml:"urlset"`
	XMLNS   string   `xml:"xmlns,attr"`
	Locs    []*Loc   `xml:"url"`
}

func saveSitemap(links []string, path string) {
	locs := []*Loc{}
	for _, link := range links {
		locs = append(locs, &Loc{Loc: link})
	}
	urlset := URLSet{
		XMLNS: "http://www.sitemaps.org/schemas/sitemap/0.9",
		Locs:  locs,
	}
	body, err := xml.MarshalIndent(urlset, " ", " ")
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Create(path)
	if err != nil {
		log.Fatal("Create:", err)
	}
	defer file.Close()
	_, err = file.Write([]byte(xml.Header))
	if err != nil {
		log.Fatal("Write:", err)
	}
	_, err = file.Write(body)
	if err != nil {
		log.Fatal("Write:", err)
	}
}
