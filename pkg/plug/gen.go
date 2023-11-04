//go:build exclude
//go:generate go run gen.go
package main

import (
	"fmt"
	"go/token"
	"go/parser"
	"go/types"
	"go/ast"
	"go/importer"
)

func main() {
	fmt.Println("a")

	fset := token.NewFileSet()
    pkgs, err := parser.ParseDir(fset, "provider.go", nil, parser.ParseComments)
    if err != nil {
        fmt.Printf("error: %s\n", err.Error())
        return
    }
	// see https://junchang1031.hatenablog.com/entry/2021/12/22/193522
	pkg, _ := pkgs["plug"]

    // conf := types.Config{Importer: importer.Default()}
    // info := &types.Info{
    //     Types: make(map[ast.Expr]types.TypeAndValue),
    // }

    // _, err = conf.Check("main", fset, []*ast.File{f}, info)
    // if err != nil {
    //     fmt.Printf("error: %s\n", err.Error())
    //     return
    // }

    // interfaceType := info.Types[f.Decls[0].(*ast.GenDecl).Specs[0].(*ast.TypeSpec).Name].Type.(*types.Named)

    // methodMap := make(map[string]string)

    // for i := 0; i < interfaceType.NumMethods(); i++ {
    //     method := interfaceType.Method(i)
    //     methodMap[method.Name()] = method.Type().String()
    // }

	fmt.Println(obj)
}
