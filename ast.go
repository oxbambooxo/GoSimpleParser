package main

func NewASTNode(nodeType ASTNodeType, value string) *ASTNode {
	return &ASTNode{
		NodeType: nodeType,
		value:    value,
	}
}

type ASTNode struct {
	NodeType  ASTNodeType
	value     string
	parent    *ASTNode
	childrens []*ASTNode
}

func (node *ASTNode) AddChild(child *ASTNode) {
	child.parent = node
	node.childrens = append(node.childrens, child)
}

type ASTNodeType int

const (
	NodeProgram        ASTNodeType = iota //程序入口，根节点
	NodeIntDeclaration                    //整型变量声明
	NodeExpressionStmt                    //表达式语句，即表达式后面跟个分号
	NodeAssignmentStmt                    //赋值语句
	NodePrimary                           //基础表达式
	NodeMultiply                          //乘法表达式
	NodeDivision                          //除法表达式
	NodeAdditive                          //加法表达式
	NodeSubtract                          //减法表达式
	NodeIdentifier                        //标识符
	NodeIntLiteral                        //整型字面量
)

func (t ASTNodeType) String() string {
	switch t {
	case NodeProgram:
		return "Program"
	case NodeIntDeclaration:
		return "IntDeclaration"
	case NodeExpressionStmt:
		return "ExpressionStmt"
	case NodeAssignmentStmt:
		return "AssignmentStmt"
	case NodePrimary:
		return "Primary"
	case NodeMultiply:
		return "Multiply"
	case NodeDivision:
		return "Division"
	case NodeAdditive:
		return "Additive"
	case NodeSubtract:
		return "Subtract"
	case NodeIdentifier:
		return "Identifier"
	case NodeIntLiteral:
		return "IntLiteral"
	default:
		return "unknown ast node type"
	}
}
