package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

func main(){
	file := "./go/ast/a1.go"//os.Getenv("GOFILE")
	// 保存注释信息
	var comments = make(map[string]string)

	// 解析代码源文件，获取常量和注释之间的关系
	fset := token.NewFileSet()
	f, _ := parser.ParseFile(fset, file, nil, parser.ParseComments)

	// Create an ast.CommentMap from the ast.File's comments.
	// This helps keeping the association between comments
	// and AST nodes.
	cmap := ast.NewCommentMap(fset, f, f.Comments)

	for node := range cmap {
		// 仅支持一条声明语句，一个常量的情况
		if spec, ok := node.(*ast.ValueSpec); ok && len(spec.Names) == 1 {
			// 仅提取常量的注释
			ident := spec.Names[0]
			if ident.Obj.Kind == ast.Con {
				// 获取注释信息
				comments[ident.Name] = getComment(ident.Name, spec.Doc)
			}
		}
	}

	fmt.Println(comments)
	//getComment("abcdefg",comments)
}

// getComment 获取注释信息，来自AST标准库的summary方法
func getComment(name string, group *ast.CommentGroup) string {
	var buf bytes.Buffer


	for _, comment := range group.List {
		// 注释信息会以 // 参数名，开始，我们实际使用时不需要，去掉
		text := strings.TrimSpace(strings.TrimPrefix(comment.Text, fmt.Sprintf("// %s", name)))
		buf.WriteString(text)
	}

	// replace any invisibles with blanks
	bytes := buf.Bytes()
	for i, b := range bytes {
		switch b {
		case '\t', '\n', '\r':
			bytes[i] = ' '
		}
	}

	return string(bytes)
}