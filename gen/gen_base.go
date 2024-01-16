package gen

import (
	"bytes"
	"fmt"
	"github.com/doytowin/goooqo/core"
	"go/ast"
	"strings"
)

type Generator interface {
	appendPackage(string2 string)
	appendImports()
	appendBuildMethod(ts *ast.TypeSpec)
	String() string
}

type generator struct {
	*bytes.Buffer
	key        string
	imports    []string
	bodyFormat string
	ifFormat   string
	intent     string
	replaceIns func(string) string
}

func newGenerator(key string, imports []string, bodyFormat string) *generator {
	return &generator{
		Buffer:     bytes.NewBuffer(make([]byte, 0, 1024)),
		key:        key,
		imports:    imports,
		bodyFormat: bodyFormat,
		ifFormat:   "if q.%s%s {",
		replaceIns: keep,
	}
}

func keep(ins string) string {
	return ins
}

func (g *generator) appendPackage(pkg string) {
	g.WriteString("package " + pkg)
	g.WriteString(NewLine)
	g.WriteString(NewLine)
}

func (g *generator) appendImports() {
	for _, s := range g.imports {
		g.WriteString("import " + s)
		g.WriteString(NewLine)
	}
	g.WriteString(NewLine)
}

func (g *generator) appendIfEnd() {
	g.WriteString(g.intent)
	g.WriteString("}")
	g.WriteString(NewLine)
}

func (g *generator) appendIfStart(structName string, cond string) {
	g.writeInstruction(g.ifFormat, structName, cond)
}

func (g *generator) appendIfStartNil(fieldName string) {
	g.writeInstruction(g.ifFormat, fieldName, " != nil")
}

func (g *generator) appendIfBody(ins string, args ...any) {
	if ins == "" {
		ins = g.bodyFormat
	}
	ins = g.replaceIns(ins)
	g.WriteString("\t")
	g.writeInstruction(ins, args...)
}

func (g *generator) writeInstruction(ins string, args ...any) {
	g.WriteString(g.intent)
	g.WriteString(fmt.Sprintf(ins, args...))
	g.WriteString(NewLine)
}

func (g *generator) suffixMatch(fieldName string) (string, operator) {
	if match := suffixRgx.FindStringSubmatch(fieldName); len(match) > 0 {
		op := opMap[g.key][match[1]]
		column := strings.TrimSuffix(fieldName, match[1])
		column = core.ConvertToColumnCase(column)
		return column, op
	}
	return core.ConvertToColumnCase(fieldName), opMap[g.key]["Eq"]
}
