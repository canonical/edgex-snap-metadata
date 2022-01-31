package main

import (
	"bufio"
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	name := flag.String("name", "", "name of the snap")
	flag.Parse()

	if *name == "" {
		fmt.Println("Missing input.\nUsage:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	var component struct {
		XMLName     xml.Name `xml:"component"`
		Summary     string   `xml:"summary"`
		Description struct {
			// use CDATA to make the value human-readable but
			// not interpreted as XML markup
			Body string `xml:",cdata"`
		} `xml:"description"`
	}

	file, err := os.Open("metadata/" + *name + ".md")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 0
	var description, summary string
	for scanner.Scan() {
		line := scanner.Text()
		if lineNum == 0 {
			// line 0 is the summary
			summary = line
		} else if lineNum == 1 {
			// line 1 is the seperator
			if line != "---" {
				log.Fatalln("Missing three-dashes (---) separator on line 1 between summary and description")
			}
			lineNum++
			continue
		} else {
			// lines 2 to EOF are the description
			description += line + "\n"
		}
		lineNum++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("Summary:\n%s\n\n", summary)
	// fmt.Printf("Description:\n%s\n\n", description)

	component.Summary = summary
	component.Description.Body = "\n" + description

	output, err := xml.MarshalIndent(&component, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n\n", output)

	// output = bytes.ReplaceAll(output, []byte("\n"), []byte("  \n"))
	err = os.MkdirAll("snap", 0755)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(*name+".metainfo.xml", output, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
