package cmd

import (
	"html/template"
	"log"
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

const slnTpl = `package {{.Pkg}}`
const testTpl = `package {{.Pkg}}
import (
	"reflect"
	"testing"
)

type Args []interface{}
type TestCase struct {
	input  Args
	expect interface{}
}

func TestSolution(t *testing.T) {
	cases := []*TestCase{
		&TestCase{Args{}, xxx},
	}
	for _, c := range cases {
		res := XXX(c.input[0].(xxx), c.input[1].(xxx))
		pass := reflect.DeepEqual(res, c.expect.(xxx))
		if !pass {
			t.Errorf("\n Case %v fail, Expect %v Got %v\n", c.input, c.expect, res)
		}
	}
}
`

var rootCmd = &cobra.Command{
	Use:   "lc",
	Short: "Leetcode helper command",
}

var createCmd = &cobra.Command{
	Use:   "create [name]",
	Short: "Create a question (format: lcNN-the-question-name)",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Println("Question name must be provided!")
			os.Exit(1)
		}

		name := args[0]
		if match, _ := regexp.MatchString(`^lc\d`, name); !match {
			log.Println("Question name must started with \"lc\"")
			os.Exit(1)
		}

		re, err := regexp.Compile(`lc\d*`)
		if err != nil {
			log.Fatalln(err)
		}

		pkg := string(re.Find([]byte(name)))
		name = strings.ReplaceAll(name, "-", "_")

		log.Println("Creating dirtory ...")
		os.Mkdir("./algorithm/"+name, os.ModePerm)

		slnT, err := template.New("solution").Parse(slnTpl)
		if err != nil {
			log.Fatalln(err)
		}
		slnTestT, err := template.New("solution_test").Parse(testTpl)
		if err != nil {
			log.Fatalln(err)
		}
		data := struct{ Pkg string }{pkg}

		slnFile, err := os.OpenFile(path.Join("./algorithm", name, "solution.go"), os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			log.Fatalln(err)
		}
		slnTestFile, err := os.OpenFile(path.Join("./algorithm", name, "solution_test.go"), os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			log.Fatalln(err)
		}

		log.Println("Generating solution.go ...")
		slnT.Execute(slnFile, data)
		log.Println("Generating solution_test.go ...")
		slnTestT.Execute(slnTestFile, data)

		log.Println("Done!")
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
