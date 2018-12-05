package main

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"

	"github.com/covrom/xml2json/xmldom"
)

func main() {
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	node := &xmldom.XMLNode{}

	if err := xml.Unmarshal(b, node); err != nil {
		log.Fatal(err)
	}

	b, err = json.MarshalIndent(node, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	
	_, err = os.Stdout.Write(b)
	if err != nil {
		log.Fatal(err)
	}
}
