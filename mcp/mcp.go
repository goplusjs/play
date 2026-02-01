package mcp

//go:generate qexp -outdir . -code github.com/goplus/mcp/...
//go:generate qexp -outdir . -code github.com/qiniu/x/test
//go:generate qexp -outdir . github.com/mark3labs/mcp-go/...
//go:generate qexp -outdir . github.com/qiniu/x/stringutil
//go:generate qexp -outdir . github.com/yosida95/uritemplate/v3
//go:generate qexp -outdir . github.com/google/uuid

import (
	_ "github.com/goplus/mcp/server"

	"github.com/goplus/ixgo/xgobuild"

	_ "github.com/goplusjs/play/mcp/github.com/qiniu/x/test"
	_ "github.com/goplusjs/play/mcp/github.com/yosida95/uritemplate/v3"

	_ "github.com/goplusjs/play/mcp/github.com/google/uuid"
	_ "github.com/goplusjs/play/mcp/github.com/goplus/mcp/mtest"
	_ "github.com/goplusjs/play/mcp/github.com/goplus/mcp/mtest/mock"
	_ "github.com/goplusjs/play/mcp/github.com/goplus/mcp/mtest/rtx"
	_ "github.com/goplusjs/play/mcp/github.com/goplus/mcp/server"
	_ "github.com/goplusjs/play/mcp/github.com/goplus/mcp/server/sse"
	_ "github.com/goplusjs/play/mcp/github.com/goplus/mcp/server/stdio"
	_ "github.com/goplusjs/play/mcp/github.com/goplus/mcp/server/svx"
	_ "github.com/goplusjs/play/mcp/github.com/mark3labs/mcp-go/client"
	_ "github.com/goplusjs/play/mcp/github.com/mark3labs/mcp-go/client/transport"
	_ "github.com/goplusjs/play/mcp/github.com/mark3labs/mcp-go/mcp"
	_ "github.com/goplusjs/play/mcp/github.com/mark3labs/mcp-go/server"
)

/*
project *_mcp.gox MCPApp github.com/goplus/mcp/server

class *_res.gox ResourceApp ResourceProto
class *_tool.gox ToolApp ToolProto
class *_prompt.gox PromptApp PromptProto

import log

project main_mtest.gox MainApp github.com/goplus/mcp/mtest github.com/qiniu/x/test

class *_mtest.gox CaseApp
*/

func init() {
	/*
		proj := &modfile.Project{
			Ext:     "_mcp.gox",
			FullExt: "*_mcp.gox",
			Class:   "MCPApp",
			Works: []*modfile.Class{
				{Ext: "_res.gox", FullExt: "*_res.gox", Class: "ResourceApp", Proto: "ResourceProto"},
				{Ext: "_tool.gox", FullExt: "*_tool.gox", Class: "ToolApp", Proto: "ToolProto"},
				{Ext: "_prompt.gox", FullExt: "*_prompt.gox", Class: "PromptApp", Proto: "PromptProto"},
			},
			Import: []*modfile.Import{
				{Name: "log", Path: "log"},
			},
			PkgPaths: []string{"github.com/goplus/mcp/server"},
		}
		proj_test := &modfile.Project{
			Ext:     "_mtest.gox",
			FullExt: "main_mtest.gox",
			Class:   "MainApp",
			Works: []*modfile.Class{
				{Ext: "_mtest.gox", FullExt: "*_mtest.gox", Class: "CaseApp"},
			},
			PkgPaths: []string{"github.com/goplus/mcp/mtest", "github.com/qiniu/x/test"},
		}
		xgobuild.RegisterProject(proj)
		xgobuild.RegisterProject(proj_test)
	*/
	xgobuild.RegisterClassFileType("_mcp.gox", "MCPApp", []*xgobuild.Class{
		{Ext: "_res.gox", Class: "ResourceApp", Proto: "ResourceProto"},
		{Ext: "_tool.gox", Class: "ToolApp", Proto: "ToolProto"},
		{Ext: "_prompt.gox", Class: "PromptApp", Proto: "PromptProto"},
	}, "github.com/goplus/mcp/server")

	xgobuild.RegisterClassFileType("main_mtest.gox", "MainApp", []*xgobuild.Class{
		{Ext: "_mtest.gox", Class: "CaseApp"},
	}, "github.com/goplus/mcp/mtest", "github.com/qiniu/x/test")
}
