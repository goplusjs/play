package mcp

//go:generate qexp -outdir . -code github.com/goplus/mcp/...
//go:generate qexp -outdir . -code github.com/qiniu/x/test
//go:generate qexp -outdir . github.com/mark3labs/mcp-go/...
//go:generate qexp -outdir . github.com/qiniu/x/stringutil
//go:generate qexp -outdir . github.com/yosida95/uritemplate/v3
//go:generate qexp -outdir . github.com/google/uuid

import (
	_ "github.com/goplus/mcp/server"

	"github.com/goplus/igop/gopbuild"

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
project _mcp.gox MCPApp github.com/goplus/mcp/server

class _res.gox ResourceApp ResourceProto
class _tool.gox ToolApp ToolProto
class _prompt.gox PromptApp PromptProto

import log

project _mtest.gox MainApp github.com/goplus/mcp/mtest github.com/qiniu/x/test

class _mtest.gox CaseApp

*/

func init() {
	gopbuild.RegisterClassFileType("_mcp.gox", "MCPApp", []*gopbuild.Class{
		{Ext: "_res.gox", Class: "ResourceApp", Proto: "ResourceProto"},
		{Ext: "_tool.gox", Class: "ToolApp", Proto: "ToolProto"},
		{Ext: "_prompt.gox", Class: "PromptApp", Proto: "PromptProto"},
	}, "github.com/goplus/mcp/server")

	gopbuild.RegisterClassFileType("_mtest.gox", "MainApp", []*gopbuild.Class{
		{Ext: "_mtest.gox", Class: "CaseApp"},
	}, "github.com/goplus/mcp/mtest", "github.com/qiniu/x/test")
}
