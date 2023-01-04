package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/tabwriter"
)

type command struct {
	name    string
	flagSet func() *flag.FlagSet
	exec    func(*flag.FlagSet) error
}

func main() {
	allCommands := []command{
		profileListCmd(),
	}
	if len(os.Args) >= 2 {
		for _, cmd := range allCommands {
			if cmd.name == os.Args[1] {
				fs := cmd.flagSet()
				if err := fs.Parse(os.Args[2:]); err != nil {
					log.Fatal(err)
				}
				if err := cmd.exec(fs); err != nil {
					log.Fatal(err)
				}
				return
			}
		}
	}
	var cmdNames []string
	for _, cmd := range allCommands {
		cmdNames = append(cmdNames, cmd.name)
	}
	fmt.Printf("expected %s subcommands\n", strings.Join(cmdNames, ", "))
	os.Exit(1)
}

func profileListCmd() command {
	name := "profile-list"
	return command{
		name: name,
		flagSet: func() *flag.FlagSet {
			fs := flag.NewFlagSet(name, flag.ExitOnError)
			return fs
		},
		exec: func(*flag.FlagSet) error {
			return profileList()
		},
	}
}

func profileList() error {
	buf := bytes.Buffer{}
	cmd := exec.Command("nix", "profile", "list")
	cmd.Stdout = &buf
	if err := cmd.Run(); err != nil {
		return err
	}

	lines := strings.Split(buf.String(), "\n")
	home, _ := os.UserHomeDir()

	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	defer w.Flush()

	for _, l := range lines {
		l = strings.TrimSpace(l)
		if l == "" {
			continue
		}
		parts := strings.Split(l, " ")
		if len(parts) < 2 {
			continue
		}

		var pkg, ver string
		{
			b := filepath.Base(parts[3])
			b = strings.TrimSuffix(b, "-bin")
			i := strings.Index(b, "-")
			j := strings.LastIndex(b, "-")
			if i >= j {
				pkg = b[i+1:]
			} else {
				pkg = b[i+1 : j]
				ver = b[j+1:]
			}
		}

		var repo string
		{
			u, _ := url.Parse(parts[1])
			if u.Scheme == "git+file" {
				repo = strings.ReplaceAll(u.Path, home, "~")
			} else if u.Scheme == "github" {
				repo = fmt.Sprintf("github:%v", u.Opaque)
			}
		}

		var rev string
		{
			u, _ := url.Parse(parts[2])
			rev = u.Query().Get("rev")
			if rev != "" {
				rev = rev[0:12]
			}
		}

		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n", parts[0], pkg, ver, repo, rev)
	}

	return nil
}
