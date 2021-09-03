package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path"
	"strings"
	"text/template"

	"github.com/pkg/errors"
)

const genmsg = "// Code generated by genops, which is a ops generation tool for Gorgonia. DO NOT EDIT."

const importgcontext = `import gctx "gorgonia.org/gorgonia/internal/context"`

var (
	gopath, stdopsloc string
	stubsFilename     string
	stubsFile         io.WriteCloser
)

func init() {
	gopath = os.Getenv("GOPATH")
	if gopath == "" {
		usr, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}
		gopath = path.Join(usr.HomeDir, "go")
		stat, err := os.Stat(gopath)
		if err != nil {
			log.Fatal(err)
		}
		if !stat.IsDir() {
			log.Fatal("You need to define a $GOPATH")
		}
	}
	stdopsloc = path.Join(gopath, "src/gorgonia.org/gorgonia/ops/std")
	stubsFilename = path.Join(stdopsloc, "stubs_generated.go")

	// handle stubsFile
	var err error
	if stubsFile, err = os.OpenFile(stubsFilename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644); err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(stubsFile, "package stdops\n\n%v\n\n", genmsg)
}

func goimports(filename string) error {
	cmd := exec.Command("goimports", "-w", filename)
	err := cmd.Run()
	if err != nil {
		return errors.Wrapf(err, "Unable to goimports %v", filename)
	}
	return nil
}

func generateBinOp(ops []op, tmpl *template.Template) error {
	for _, op := range ops {
		filename := strings.ToLower(op.Name) + "_generated.go"
		p := path.Join(stdopsloc, filename)
		f, err := os.OpenFile(p, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		fmt.Fprintf(f, "package stdops\n\n%v\n\n %v\n\n", genmsg, importgcontext)
		if err := tmpl.Execute(f, op); err != nil {
			return errors.Wrapf(err, "Unable to execute binopTmpl for %v", op.Name)
		}
		if err := f.Close(); err != nil {
			return errors.Wrapf(err, "Unable to close %v", p)
		}
		if err := goimports(p); err != nil {
			return err
		}

		// extra: write symdiff to stubs
		if err := binSymDiffTmpl.Execute(stubsFile, op); err != nil {
			return errors.Wrapf(err, "Unable to add %v SymDiff stubs", op.Name)
		}
	}
	return nil
}

func generateBinOpTest(ops []op, input binopTestInput, results []binopTestResult, isCmp bool, tmpl *template.Template) error {
	for i, op := range ops {
		opTest := binopTest{op: op, binopTestInput: input, binopTestResult: results[i]}
		filename := strings.ToLower(op.Name) + "_generated_test.go"
		p := path.Join(stdopsloc, filename)
		f, err := os.OpenFile(p, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		fmt.Fprintf(f, "package stdops\n\n%v\n\n", genmsg)
		if err := tmpl.Execute(f, opTest); err != nil {
			return errors.Wrapf(err, "Unable to execute binopTmpl for %v", op.Name)
		}
		// for cmp
		if isCmp {
			opTest.IsCmpRetTrue = true
			opTest.binopTestInput = cmpTestInputSame
			opTest.binopTestResult = cmpTestResultsSame[i]
			if err := tmpl.Execute(f, opTest); err != nil {
				return errors.Wrapf(err, "Unable to execute binopTmpl for %v", op.Name)
			}
		}
		if err := f.Close(); err != nil {
			return errors.Wrapf(err, "Unable to close %v", p)
		}
		if err := goimports(p); err != nil {
			return err
		}
	}
	return nil
}

func generateAriths() error {
	if err := generateBinOp(ariths, arithOpTmpl); err != nil {
		return errors.Wrap(err, "generateAriths.generateBinOp")
	}
	if err := generateBinOpTest(ariths, arithTestInput, arithTestResults, false, arithOpTestTmpl); err != nil {
		return errors.Wrap(err, "generateAriths.generateBinOpTests")
	}

	return nil
}

func generateCmps() error {
	if err := generateBinOp(cmps, cmpOpTmpl); err != nil {
		return errors.Wrap(err, "generateCmps.generateBinOp")
	}
	if err := generateBinOpTest(cmps, cmpTestInputBool, cmpTestResultsBool, true, arithOpTestTmpl); err != nil {
		return errors.Wrap(err, "generateCmps.generateBinOpTests")
	}
	return nil
}

func generateUnOps() error {
	tmpl := unopTmpl
	for _, op := range unops {
		filename := strings.ToLower(op.Name) + "_generated.go"
		p := path.Join(stdopsloc, filename)
		f, err := os.OpenFile(p, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		fmt.Fprintf(f, "package stdops\n\n%v\n\n %v\n\n", genmsg, importgcontext)
		if err := tmpl.Execute(f, op); err != nil {
			return errors.Wrapf(err, "Unable to execute unopTmpl for %v", op.Name)
		}
		if err := f.Close(); err != nil {
			return errors.Wrapf(err, "Unable to close %v", p)
		}
		if err := goimports(p); err != nil {
			return err
		}

		// extra: write symdiff to stubs
		if err := binSymDiffTmpl.Execute(stubsFile, op); err != nil {
			return errors.Wrapf(err, "Unable to add %v SymDiff stubs", op.Name)
		}
	}

	tmpl = unopTestTmpl
	for i, op := range unops {
		filename := strings.ToLower(op.Name) + "_generated_test.go"
		p := path.Join(stdopsloc, filename)
		f, err := os.OpenFile(p, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		fmt.Fprintf(f, "package stdops\n\n%v\n\n", genmsg)

		o := unoptestWithOp{op, unopTests[i]}
		if err := tmpl.Execute(f, o); err != nil {
			return errors.Wrapf(err, "Unable to execute unopTmpl for %v", op.Name)
		}
		if err := f.Close(); err != nil {
			return errors.Wrapf(err, "Unable to close %v", p)
		}
		if err := goimports(p); err != nil {
			return err
		}
	}
	return nil
}

func generateBinOpAPI() (err error) {

	type apiwrap struct {
		op
		IsCmp bool
	}

	filename := "api_generated.go"
	filenameTest := "api_generated_test.go"
	p := path.Join(stdopsloc, filename)
	pt := path.Join(stdopsloc, filenameTest)
	f, err := os.OpenFile(p, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	g, err := os.OpenFile(pt, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	fmt.Fprintf(f, "package stdops\n\n%v\n\n", genmsg)
	fmt.Fprintf(g, "package stdops\n\n%v\n\n", genmsg)
	for _, o := range ariths {
		if err := binopAPITmpl.Execute(f, apiwrap{o, false}); err != nil {
			return errors.Wrapf(err, "Unable to execute binopAPITmpl for %v", o.Name)
		}

		if err := binopAPITestTmpl.Execute(g, apiwrap{o, false}); err != nil {
			return errors.Wrapf(err, "Unable to execute binopAPITestTmpl for %v", o.Name)
		}
	}
	for _, o := range cmps {
		if err := binopAPITmpl.Execute(f, apiwrap{o, true}); err != nil {
			return errors.Wrapf(err, "Unable to execute binopAPITmpl for %v", o.Name)
		}

		if err := binopAPITestTmpl.Execute(g, apiwrap{o, true}); err != nil {
			return errors.Wrapf(err, "Unable to execute binopAPITestTmpl for %v", o.Name)
		}
	}

	if err := f.Close(); err != nil {
		return errors.Wrapf(err, "Unable to close %v", p)
	}
	if err := g.Close(); err != nil {
		return errors.Wrapf(err, "Unable to close %v", pt)
	}

	if err := goimports(p); err != nil {
		return errors.Wrapf(err, "Unable to goimports %v", p)
	}
	return goimports(pt)
}

func finishStubs() error {
	if err := stubsFile.Close(); err != nil {
		return err
	}
	return goimports(stubsFilename)
}

func main() {
	defer finishStubs()
	if err := generateAriths(); err != nil {
		log.Fatal(err)
	}
	if err := generateCmps(); err != nil {
		log.Fatal(err)
	}
	if err := generateUnOps(); err != nil {
		log.Fatal(err)
	}
	if err := generateBinOpAPI(); err != nil {
		log.Fatal(err)
	}
}
