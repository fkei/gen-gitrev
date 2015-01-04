package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"text/template"
)

func main() {

	///
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	///
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `
Usage of %s:
   %s [OPTIONS] ARGS...

Options`, os.Args[0], os.Args[0])

		flag.PrintDefaults()
	}

	pkgname := flag.String("pkgname", "main", "package name")
	out := flag.String("out", cwd+"/gitrev_gen.go", "output file path")

	flag.Parse()

	///
	cmd := exec.Command("git", "rev-parse", "HEAD")
	cmd.Stdout = new(bytes.Buffer)
	cmd.Stderr = os.Stderr

	err = cmd.Run()

	if err != nil {
		panic(err)
		return
	}

	revision := fmt.Sprint(cmd.Stdout)

	///
	tmpl := `package {{.pkgname}}

func GetRevision() (string){
	return "{{.revision}}"
}
`

	t := template.New("t")
	template.Must(t.Parse(tmpl))

	dest := new(bytes.Buffer)
	t.Execute(dest, map[string]string{
		"revision": revision[:len(revision)-1],
		"pkgname":  *pkgname,
	})

	ioutil.WriteFile(*out, dest.Bytes(), 0644)
	fmt.Println("Output generate file:", *out)
	fmt.Println("\tpacakge name:", *pkgname)
	fmt.Println("\trevision:", revision)
}
