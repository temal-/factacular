/* factacular.go */
package main

import (
    "fmt"
    "os"
    "github.com/temal-/go-puppetdb"
    "github.com/codegangsta/cli"
)

func main() {
    app := cli.NewApp()
    app.Name = "factacular"
    app.Version = "0.2"
    app.Usage = "Get facts and informations from PuppetDB."
    app.Flags = []cli.Flag {
      cli.StringFlag{
        Name: "puppetdb, p",
        Value: "http://localhost:8080",
        Usage: "PuppetDB host.",
        EnvVar: "PUPPETDB_HOST",
      },
    }
    app.Commands = []cli.Command{
        {
            Name:      "list-facts",
            ShortName: "lf",
            Usage:     "List all available facts",
            Action: func(c *cli.Context) {
                fmt.Println("PuppetDB host: " + c.GlobalString("puppetdb"))
                client := puppetdb.NewClient(c.GlobalString("puppetdb"))
                resp, err := client.FactNames()
                if err != nil {
                    fmt.Println(err)
                }
                fmt.Println("Facts: ")
                for _, element := range resp {
                    fmt.Printf("%v\n", element)
                }
            },
        },
        {
            Name:      "list-nodes",
            ShortName: "ln",
            Usage:     "List all available nodes",
            Action: func(c *cli.Context) {
                fmt.Println("PuppetDB host: " + c.GlobalString("puppetdb"))
                client := puppetdb.NewClient(c.GlobalString("puppetdb"))
                resp, err := client.Nodes()
                if err != nil {
                    fmt.Println(err)
                }
                fmt.Println("Nodes: ")
                for _, element := range resp {
                    fmt.Printf("%v\n", element.Name)
                }
            },
        },
        {
            Name:      "node-facts",
            ShortName: "nf",
            Usage:     "List all facts for a specific node.",
            Action: func(c *cli.Context) {
                if(c.Args().First() == "") {
                    fmt.Println("Please provide the FQDN of a node.")
                    return
                }
                fmt.Println("PuppetDB host: " + c.GlobalString("puppetdb"))
                client := puppetdb.NewClient(c.GlobalString("puppetdb"))
                resp, err := client.NodeFacts(c.Args().First())
                if err != nil {
                    fmt.Println(err)
                }
                fmt.Println("Node-facts: ")
                for _, element := range resp {
                    fmt.Printf("%v - %v\n", c.Args().First(), element.Name)
                    fmt.Printf("%v\n", element.Value)
                }
            },
        },
        {
            Name:      "fact",
            ShortName: "f",
            Usage:     "List fact for all nodes.",
            Action: func(c *cli.Context) {
                if(c.Args().First() == "") {
                    fmt.Println("Please provide a fact.")
                    return
                }
                fmt.Println("PuppetDB host: " + c.GlobalString("puppetdb"))
                client := puppetdb.NewClient(c.GlobalString("puppetdb"))
                resp, err := client.FactPerNode(c.Args().First())
                if err != nil {
                    fmt.Println(err)
                }
                fmt.Println("Fact per node: ")
                for _, element := range resp {
                    fmt.Printf("%v - %v - %v\n", element.CertName, element.Name, element.Value)
                }
            },
        },
    }
    app.Action = func(c *cli.Context) {
        fmt.Println("Please provide a command to do stuff. 'h' brings up the help.")
    }
    app.Run(os.Args)
}
