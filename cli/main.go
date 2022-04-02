package main

import (
	"log"
	"os"

	droneGitClone "github.com/beanjs-pipeline/drone-git-clone"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Version = "1.0.1"
	app.Name = "Drone git clone"
	app.Usage = "clone a repository"
	app.Copyright = "@ 2022 beanjs"
	app.Authors = []cli.Author{
		{Name: "beanjs", Email: "502554248@qq.com"},
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "remote",
			Usage:  "remote",
			EnvVar: "PLUGIN_REMOTE",
		},
		cli.BoolFlag{
			Name:   "recursive",
			Usage:  "recursive",
			EnvVar: "PLUGIN_RECURSIVE",
		},
		cli.StringFlag{
			Name:   "cache",
			Usage:  "cache",
			EnvVar: "PLUGIN_CACHE",
			Value:  "./cache",
		},
		cli.StringFlag{
			Name:   "branch",
			Usage:  "branch",
			EnvVar: "PLUGIN_BRANCH",
			Value:  "master",
		},
		cli.StringFlag{
			Name:   "username",
			Usage:  "username",
			EnvVar: "PLUGIN_USERNAME",
		},
		cli.StringFlag{
			Name:   "password",
			Usage:  "password",
			EnvVar: "PLUGIN_PASSWORD",
		},
	}

	app.Action = run

	if err := app.Run(os.Args); err != nil {
		log.Fatalf("Fatal: %v", err)
	}
}

func run(c *cli.Context) {
	p := droneGitClone.Plugin{
		Remote:    c.String("remote"),
		Cache:     c.String("cache"),
		Branch:    c.String("branch"),
		Username:  c.String("username"),
		Password:  c.String("password"),
		Recursive: c.Bool("recursive"),
	}

	if err := p.Exec(); err != nil {
		log.Fatalf("Run Error: %v", err)
		return
	}

	log.Printf("Clone Success")
}
