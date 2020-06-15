package parse

import (
	"reflect"
	"text/template/parse"

	qspec "github.com/qiniu/goplus/exec.spec"
	"github.com/qiniu/goplus/gop"
)

// func (*parse.ActionNode).Copy() parse.Node
func execmActionNodeCopy(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*parse.ActionNode).Copy()
	p.Ret(1, ret)
}

// func (*parse.ActionNode).String() string
func execmActionNodeString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*parse.ActionNode).String()
	p.Ret(1, ret)
}

// func (*parse.BoolNode).Copy() parse.Node
func execmBoolNodeCopy(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*parse.BoolNode).Copy()
	p.Ret(1, ret)
}

// func (*parse.BoolNode).String() string
func execmBoolNodeString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*parse.BoolNode).String()
	p.Ret(1, ret)
}

// func (*parse.BranchNode).Copy() parse.Node
func execmBranchNodeCopy(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*parse.BranchNode).Copy()
	p.Ret(1, ret)
}

// func (*parse.BranchNode).String() string
func execmBranchNodeString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*parse.BranchNode).String()
	p.Ret(1, ret)
}

// func (*parse.ChainNode).Add(field string)
func execmChainNodeAdd(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	args[0].(*parse.ChainNode).Add(args[1].(string))
}

// func (*parse.ChainNode).Copy() parse.Node
func execmChainNodeCopy(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*parse.ChainNode).Copy()
	p.Ret(1, ret)
}

// func (*parse.ChainNode).String() string
func execmChainNodeString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*parse.ChainNode).String()
	p.Ret(1, ret)
}

// func (*parse.CommandNode).Copy() parse.Node
func execmCommandNodeCopy(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*parse.CommandNode).Copy()
	p.Ret(1, ret)
}

// func (*parse.CommandNode).String() string
func execmCommandNodeString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*parse.CommandNode).String()
	p.Ret(1, ret)
}

// func (*parse.DotNode).Copy() parse.Node
func execmDotNodeCopy(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*parse.DotNode).Copy()
	p.Ret(1, ret)
}

// func (*parse.DotNode).String() string
func execmDotNodeString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*parse.DotNode).String()
	p.Ret(1, ret)
}

// func (*parse.DotNode).Type() parse.NodeType
func execmDotNodeType(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*parse.DotNode).Type()
	p.Ret(1, ret)
}

// func (*parse.FieldNode).Copy() parse.Node
func execmFieldNodeCopy(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*parse.FieldNode).Copy()
	p.Ret(1, ret)
}

// func (*parse.FieldNode).String() string
func execmFieldNodeString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*parse.FieldNode).String()
	p.Ret(1, ret)
}

// func (*parse.IdentifierNode).Copy() parse.Node
func execmIdentifierNodeCopy(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*parse.IdentifierNode).Copy()
	p.Ret(1, ret)
}

// func (*parse.IdentifierNode).SetPos(pos parse.Pos) *parse.IdentifierNode
func execmIdentifierNodeSetPos(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*parse.IdentifierNode).SetPos(parse.Pos(args[1].(int)))
	p.Ret(2, ret)
}

// func (*parse.IdentifierNode).SetTree(t *parse.Tree) *parse.IdentifierNode
func execmIdentifierNodeSetTree(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret := args[0].(*parse.IdentifierNode).SetTree(args[1].(*parse.Tree))
	p.Ret(2, ret)
}

// func (*parse.IdentifierNode).String() string
func execmIdentifierNodeString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*parse.IdentifierNode).String()
	p.Ret(1, ret)
}

// func (*parse.IfNode).Copy() parse.Node
func execmIfNodeCopy(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*parse.IfNode).Copy()
	p.Ret(1, ret)
}

// func parse.IsEmptyTree(n parse.Node) bool
func execIsEmptyTree(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := parse.IsEmptyTree(args[0].(parse.Node))
	p.Ret(1, ret)
}

// func (*parse.ListNode).Copy() parse.Node
func execmListNodeCopy(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*parse.ListNode).Copy()
	p.Ret(1, ret)
}

// func (*parse.ListNode).CopyList() *parse.ListNode
func execmListNodeCopyList(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*parse.ListNode).CopyList()
	p.Ret(1, ret)
}

// func (*parse.ListNode).String() string
func execmListNodeString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*parse.ListNode).String()
	p.Ret(1, ret)
}

// func parse.New(name string, funcs ..map[string]interface{}) *parse.Tree
func execNew(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	conv := func(args []interface{}) []map[string]interface{} {
		ret := make([]map[string]interface{}, len(args))
		for i, arg := range args {
			ret[i] = arg.(map[string]interface{})
		}
		return ret
	}
	ret := parse.New(args[0].(string), conv(args[1:])...)
	p.Ret(arity, ret)
}

// func parse.NewIdentifier(ident string) *parse.IdentifierNode
func execNewIdentifier(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := parse.NewIdentifier(args[0].(string))
	p.Ret(1, ret)
}

// func (*parse.NilNode).Copy() parse.Node
func execmNilNodeCopy(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*parse.NilNode).Copy()
	p.Ret(1, ret)
}

// func (*parse.NilNode).String() string
func execmNilNodeString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*parse.NilNode).String()
	p.Ret(1, ret)
}

// func (*parse.NilNode).Type() parse.NodeType
func execmNilNodeType(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*parse.NilNode).Type()
	p.Ret(1, ret)
}

// func (parse.NodeType).Type() parse.NodeType
func execmNodeTypeType(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(parse.NodeType).Type()
	p.Ret(1, ret)
}

// func (*parse.NumberNode).Copy() parse.Node
func execmNumberNodeCopy(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*parse.NumberNode).Copy()
	p.Ret(1, ret)
}

// func (*parse.NumberNode).String() string
func execmNumberNodeString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*parse.NumberNode).String()
	p.Ret(1, ret)
}

// func parse.Parse(name string, text string, leftDelim string, rightDelim string, funcs ..map[string]interface{}) (map[string]*parse.Tree, error)
func execParse(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	conv := func(args []interface{}) []map[string]interface{} {
		ret := make([]map[string]interface{}, len(args))
		for i, arg := range args {
			ret[i] = arg.(map[string]interface{})
		}
		return ret
	}
	ret, ret1 := parse.Parse(args[0].(string), args[1].(string), args[2].(string), args[3].(string), conv(args[4:])...)
	p.Ret(arity, ret, ret1)
}

// func (*parse.PipeNode).Copy() parse.Node
func execmPipeNodeCopy(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*parse.PipeNode).Copy()
	p.Ret(1, ret)
}

// func (*parse.PipeNode).CopyPipe() *parse.PipeNode
func execmPipeNodeCopyPipe(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*parse.PipeNode).CopyPipe()
	p.Ret(1, ret)
}

// func (*parse.PipeNode).String() string
func execmPipeNodeString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*parse.PipeNode).String()
	p.Ret(1, ret)
}

// func (parse.Pos).Position() parse.Pos
func execmPosPosition(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(parse.Pos).Position()
	p.Ret(1, ret)
}

// func (*parse.RangeNode).Copy() parse.Node
func execmRangeNodeCopy(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*parse.RangeNode).Copy()
	p.Ret(1, ret)
}

// func (*parse.StringNode).Copy() parse.Node
func execmStringNodeCopy(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*parse.StringNode).Copy()
	p.Ret(1, ret)
}

// func (*parse.StringNode).String() string
func execmStringNodeString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*parse.StringNode).String()
	p.Ret(1, ret)
}

// func (*parse.TemplateNode).Copy() parse.Node
func execmTemplateNodeCopy(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*parse.TemplateNode).Copy()
	p.Ret(1, ret)
}

// func (*parse.TemplateNode).String() string
func execmTemplateNodeString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*parse.TemplateNode).String()
	p.Ret(1, ret)
}

// func (*parse.TextNode).Copy() parse.Node
func execmTextNodeCopy(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*parse.TextNode).Copy()
	p.Ret(1, ret)
}

// func (*parse.TextNode).String() string
func execmTextNodeString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*parse.TextNode).String()
	p.Ret(1, ret)
}

// func (*parse.Tree).Copy() *parse.Tree
func execmTreeCopy(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*parse.Tree).Copy()
	p.Ret(1, ret)
}

// func (*parse.Tree).ErrorContext(n parse.Node) (location string, context string)
func execmTreeErrorContext(_ int, p *gop.Context) {
	args := p.GetArgs(2)
	ret, ret1 := args[0].(*parse.Tree).ErrorContext(args[1].(parse.Node))
	p.Ret(2, ret, ret1)
}

// func (*parse.Tree).Parse(text string, leftDelim string, rightDelim string, treeSet map[string]*parse.Tree, funcs ..map[string]interface{}) (tree *parse.Tree, err error)
func execmTreeParse(arity int, p *gop.Context) {
	args := p.GetArgs(arity)
	conv := func(args []interface{}) []map[string]interface{} {
		ret := make([]map[string]interface{}, len(args))
		for i, arg := range args {
			ret[i] = arg.(map[string]interface{})
		}
		return ret
	}
	ret, ret1 := args[0].(*parse.Tree).Parse(args[1].(string), args[2].(string), args[3].(string), args[4].(map[string]*parse.Tree), conv(args[5:])...)
	p.Ret(arity, ret, ret1)
}

// func (*parse.VariableNode).Copy() parse.Node
func execmVariableNodeCopy(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*parse.VariableNode).Copy()
	p.Ret(1, ret)
}

// func (*parse.VariableNode).String() string
func execmVariableNodeString(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*parse.VariableNode).String()
	p.Ret(1, ret)
}

// func (*parse.WithNode).Copy() parse.Node
func execmWithNodeCopy(_ int, p *gop.Context) {
	args := p.GetArgs(1)
	ret := args[0].(*parse.WithNode).Copy()
	p.Ret(1, ret)
}

// I is a Go package instance.
var I = gop.NewGoPackage("text/template/parse")

func init() {
	I.RegisterConsts(
		I.Const("NodeAction", reflect.Int, parse.NodeAction),
		I.Const("NodeBool", reflect.Int, parse.NodeBool),
		I.Const("NodeChain", reflect.Int, parse.NodeChain),
		I.Const("NodeCommand", reflect.Int, parse.NodeCommand),
		I.Const("NodeDot", reflect.Int, parse.NodeDot),
		I.Const("NodeField", reflect.Int, parse.NodeField),
		I.Const("NodeIdentifier", reflect.Int, parse.NodeIdentifier),
		I.Const("NodeIf", reflect.Int, parse.NodeIf),
		I.Const("NodeList", reflect.Int, parse.NodeList),
		I.Const("NodeNil", reflect.Int, parse.NodeNil),
		I.Const("NodeNumber", reflect.Int, parse.NodeNumber),
		I.Const("NodePipe", reflect.Int, parse.NodePipe),
		I.Const("NodeRange", reflect.Int, parse.NodeRange),
		I.Const("NodeString", reflect.Int, parse.NodeString),
		I.Const("NodeTemplate", reflect.Int, parse.NodeTemplate),
		I.Const("NodeText", reflect.Int, parse.NodeText),
		I.Const("NodeVariable", reflect.Int, parse.NodeVariable),
		I.Const("NodeWith", reflect.Int, parse.NodeWith),
	)
	I.RegisterTypes(
		I.Rtype(reflect.TypeOf((*parse.ActionNode)(nil))),
		I.Rtype(reflect.TypeOf((*parse.BoolNode)(nil))),
		I.Rtype(reflect.TypeOf((*parse.BranchNode)(nil))),
		I.Rtype(reflect.TypeOf((*parse.ChainNode)(nil))),
		I.Rtype(reflect.TypeOf((*parse.CommandNode)(nil))),
		I.Rtype(reflect.TypeOf((*parse.DotNode)(nil))),
		I.Rtype(reflect.TypeOf((*parse.FieldNode)(nil))),
		I.Rtype(reflect.TypeOf((*parse.IdentifierNode)(nil))),
		I.Rtype(reflect.TypeOf((*parse.IfNode)(nil))),
		I.Rtype(reflect.TypeOf((*parse.ListNode)(nil))),
		I.Rtype(reflect.TypeOf((*parse.NilNode)(nil))),
		I.Type("NodeType", qspec.TyInt),
		I.Rtype(reflect.TypeOf((*parse.NumberNode)(nil))),
		I.Rtype(reflect.TypeOf((*parse.PipeNode)(nil))),
		I.Type("Pos", qspec.TyInt),
		I.Rtype(reflect.TypeOf((*parse.RangeNode)(nil))),
		I.Rtype(reflect.TypeOf((*parse.StringNode)(nil))),
		I.Rtype(reflect.TypeOf((*parse.TemplateNode)(nil))),
		I.Rtype(reflect.TypeOf((*parse.TextNode)(nil))),
		I.Rtype(reflect.TypeOf((*parse.Tree)(nil))),
		I.Rtype(reflect.TypeOf((*parse.VariableNode)(nil))),
		I.Rtype(reflect.TypeOf((*parse.WithNode)(nil))),
	)
	I.RegisterFuncs(
		I.Func("(*ActionNode).Copy", (*parse.ActionNode).Copy, execmActionNodeCopy),
		I.Func("(*ActionNode).String", (*parse.ActionNode).String, execmActionNodeString),
		I.Func("(*BoolNode).Copy", (*parse.BoolNode).Copy, execmBoolNodeCopy),
		I.Func("(*BoolNode).String", (*parse.BoolNode).String, execmBoolNodeString),
		I.Func("(*BranchNode).Copy", (*parse.BranchNode).Copy, execmBranchNodeCopy),
		I.Func("(*BranchNode).String", (*parse.BranchNode).String, execmBranchNodeString),
		I.Func("(*ChainNode).Add", (*parse.ChainNode).Add, execmChainNodeAdd),
		I.Func("(*ChainNode).Copy", (*parse.ChainNode).Copy, execmChainNodeCopy),
		I.Func("(*ChainNode).String", (*parse.ChainNode).String, execmChainNodeString),
		I.Func("(*CommandNode).Copy", (*parse.CommandNode).Copy, execmCommandNodeCopy),
		I.Func("(*CommandNode).String", (*parse.CommandNode).String, execmCommandNodeString),
		I.Func("(*DotNode).Copy", (*parse.DotNode).Copy, execmDotNodeCopy),
		I.Func("(*DotNode).String", (*parse.DotNode).String, execmDotNodeString),
		I.Func("(*DotNode).Type", (*parse.DotNode).Type, execmDotNodeType),
		I.Func("(*FieldNode).Copy", (*parse.FieldNode).Copy, execmFieldNodeCopy),
		I.Func("(*FieldNode).String", (*parse.FieldNode).String, execmFieldNodeString),
		I.Func("(*IdentifierNode).Copy", (*parse.IdentifierNode).Copy, execmIdentifierNodeCopy),
		I.Func("(*IdentifierNode).SetPos", (*parse.IdentifierNode).SetPos, execmIdentifierNodeSetPos),
		I.Func("(*IdentifierNode).SetTree", (*parse.IdentifierNode).SetTree, execmIdentifierNodeSetTree),
		I.Func("(*IdentifierNode).String", (*parse.IdentifierNode).String, execmIdentifierNodeString),
		I.Func("(*IfNode).Copy", (*parse.IfNode).Copy, execmIfNodeCopy),
		I.Func("IsEmptyTree", parse.IsEmptyTree, execIsEmptyTree),
		I.Func("(*ListNode).Copy", (*parse.ListNode).Copy, execmListNodeCopy),
		I.Func("(*ListNode).CopyList", (*parse.ListNode).CopyList, execmListNodeCopyList),
		I.Func("(*ListNode).String", (*parse.ListNode).String, execmListNodeString),
		I.Func("NewIdentifier", parse.NewIdentifier, execNewIdentifier),
		I.Func("(*NilNode).Copy", (*parse.NilNode).Copy, execmNilNodeCopy),
		I.Func("(*NilNode).String", (*parse.NilNode).String, execmNilNodeString),
		I.Func("(*NilNode).Type", (*parse.NilNode).Type, execmNilNodeType),
		I.Func("(NodeType).Type", (parse.NodeType).Type, execmNodeTypeType),
		I.Func("(*NumberNode).Copy", (*parse.NumberNode).Copy, execmNumberNodeCopy),
		I.Func("(*NumberNode).String", (*parse.NumberNode).String, execmNumberNodeString),
		I.Func("(*PipeNode).Copy", (*parse.PipeNode).Copy, execmPipeNodeCopy),
		I.Func("(*PipeNode).CopyPipe", (*parse.PipeNode).CopyPipe, execmPipeNodeCopyPipe),
		I.Func("(*PipeNode).String", (*parse.PipeNode).String, execmPipeNodeString),
		I.Func("(Pos).Position", (parse.Pos).Position, execmPosPosition),
		I.Func("(*RangeNode).Copy", (*parse.RangeNode).Copy, execmRangeNodeCopy),
		I.Func("(*StringNode).Copy", (*parse.StringNode).Copy, execmStringNodeCopy),
		I.Func("(*StringNode).String", (*parse.StringNode).String, execmStringNodeString),
		I.Func("(*TemplateNode).Copy", (*parse.TemplateNode).Copy, execmTemplateNodeCopy),
		I.Func("(*TemplateNode).String", (*parse.TemplateNode).String, execmTemplateNodeString),
		I.Func("(*TextNode).Copy", (*parse.TextNode).Copy, execmTextNodeCopy),
		I.Func("(*TextNode).String", (*parse.TextNode).String, execmTextNodeString),
		I.Func("(*Tree).Copy", (*parse.Tree).Copy, execmTreeCopy),
		I.Func("(*Tree).ErrorContext", (*parse.Tree).ErrorContext, execmTreeErrorContext),
		I.Func("(*VariableNode).Copy", (*parse.VariableNode).Copy, execmVariableNodeCopy),
		I.Func("(*VariableNode).String", (*parse.VariableNode).String, execmVariableNodeString),
		I.Func("(*WithNode).Copy", (*parse.WithNode).Copy, execmWithNodeCopy),
	)
	I.RegisterFuncvs(
		I.Funcv("New", parse.New, execNew),
		I.Funcv("Parse", parse.Parse, execParse),
		I.Funcv("(*Tree).Parse", (*parse.Tree).Parse, execmTreeParse),
	)
}
