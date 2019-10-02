// Code generated from PlayScript.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // PlayScript

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePlayScriptListener is a complete listener for a parse tree produced by PlayScriptParser.
type BasePlayScriptListener struct{}

var _ PlayScriptListener = &BasePlayScriptListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePlayScriptListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePlayScriptListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePlayScriptListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePlayScriptListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterClassOrInterfaceType is called when production classOrInterfaceType is entered.
func (s *BasePlayScriptListener) EnterClassOrInterfaceType(ctx *ClassOrInterfaceTypeContext) {}

// ExitClassOrInterfaceType is called when production classOrInterfaceType is exited.
func (s *BasePlayScriptListener) ExitClassOrInterfaceType(ctx *ClassOrInterfaceTypeContext) {}

// EnterTypeTypeOrVoid is called when production typeTypeOrVoid is entered.
func (s *BasePlayScriptListener) EnterTypeTypeOrVoid(ctx *TypeTypeOrVoidContext) {}

// ExitTypeTypeOrVoid is called when production typeTypeOrVoid is exited.
func (s *BasePlayScriptListener) ExitTypeTypeOrVoid(ctx *TypeTypeOrVoidContext) {}

// EnterFunctionType is called when production functionType is entered.
func (s *BasePlayScriptListener) EnterFunctionType(ctx *FunctionTypeContext) {}

// ExitFunctionType is called when production functionType is exited.
func (s *BasePlayScriptListener) ExitFunctionType(ctx *FunctionTypeContext) {}

// EnterPrimitiveType is called when production primitiveType is entered.
func (s *BasePlayScriptListener) EnterPrimitiveType(ctx *PrimitiveTypeContext) {}

// ExitPrimitiveType is called when production primitiveType is exited.
func (s *BasePlayScriptListener) ExitPrimitiveType(ctx *PrimitiveTypeContext) {}

// EnterTypeType is called when production typeType is entered.
func (s *BasePlayScriptListener) EnterTypeType(ctx *TypeTypeContext) {}

// ExitTypeType is called when production typeType is exited.
func (s *BasePlayScriptListener) ExitTypeType(ctx *TypeTypeContext) {}

// EnterTypeList is called when production typeList is entered.
func (s *BasePlayScriptListener) EnterTypeList(ctx *TypeListContext) {}

// ExitTypeList is called when production typeList is exited.
func (s *BasePlayScriptListener) ExitTypeList(ctx *TypeListContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasePlayScriptListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasePlayScriptListener) ExitExpression(ctx *ExpressionContext) {}

// EnterExpressionList is called when production expressionList is entered.
func (s *BasePlayScriptListener) EnterExpressionList(ctx *ExpressionListContext) {}

// ExitExpressionList is called when production expressionList is exited.
func (s *BasePlayScriptListener) ExitExpressionList(ctx *ExpressionListContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BasePlayScriptListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BasePlayScriptListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterPrimary is called when production primary is entered.
func (s *BasePlayScriptListener) EnterPrimary(ctx *PrimaryContext) {}

// ExitPrimary is called when production primary is exited.
func (s *BasePlayScriptListener) ExitPrimary(ctx *PrimaryContext) {}

// EnterIntLiteral is called when production intLiteral is entered.
func (s *BasePlayScriptListener) EnterIntLiteral(ctx *IntLiteralContext) {}

// ExitIntLiteral is called when production intLiteral is exited.
func (s *BasePlayScriptListener) ExitIntLiteral(ctx *IntLiteralContext) {}

// EnterFloatLiteral is called when production floatLiteral is entered.
func (s *BasePlayScriptListener) EnterFloatLiteral(ctx *FloatLiteralContext) {}

// ExitFloatLiteral is called when production floatLiteral is exited.
func (s *BasePlayScriptListener) ExitFloatLiteral(ctx *FloatLiteralContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BasePlayScriptListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BasePlayScriptListener) ExitLiteral(ctx *LiteralContext) {}

// EnterBlock is called when production block is entered.
func (s *BasePlayScriptListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BasePlayScriptListener) ExitBlock(ctx *BlockContext) {}

// EnterBlockStatement is called when production blockStatement is entered.
func (s *BasePlayScriptListener) EnterBlockStatement(ctx *BlockStatementContext) {}

// ExitBlockStatement is called when production blockStatement is exited.
func (s *BasePlayScriptListener) ExitBlockStatement(ctx *BlockStatementContext) {}

// EnterBlockStatements is called when production blockStatements is entered.
func (s *BasePlayScriptListener) EnterBlockStatements(ctx *BlockStatementsContext) {}

// ExitBlockStatements is called when production blockStatements is exited.
func (s *BasePlayScriptListener) ExitBlockStatements(ctx *BlockStatementsContext) {}

// EnterFunctionBody is called when production functionBody is entered.
func (s *BasePlayScriptListener) EnterFunctionBody(ctx *FunctionBodyContext) {}

// ExitFunctionBody is called when production functionBody is exited.
func (s *BasePlayScriptListener) ExitFunctionBody(ctx *FunctionBodyContext) {}

// EnterQualifiedName is called when production qualifiedName is entered.
func (s *BasePlayScriptListener) EnterQualifiedName(ctx *QualifiedNameContext) {}

// ExitQualifiedName is called when production qualifiedName is exited.
func (s *BasePlayScriptListener) ExitQualifiedName(ctx *QualifiedNameContext) {}

// EnterQualifiedNameList is called when production qualifiedNameList is entered.
func (s *BasePlayScriptListener) EnterQualifiedNameList(ctx *QualifiedNameListContext) {}

// ExitQualifiedNameList is called when production qualifiedNameList is exited.
func (s *BasePlayScriptListener) ExitQualifiedNameList(ctx *QualifiedNameListContext) {}

// EnterFunctionDeclaration is called when production functionDeclaration is entered.
func (s *BasePlayScriptListener) EnterFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// ExitFunctionDeclaration is called when production functionDeclaration is exited.
func (s *BasePlayScriptListener) ExitFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// EnterFormalParameters is called when production formalParameters is entered.
func (s *BasePlayScriptListener) EnterFormalParameters(ctx *FormalParametersContext) {}

// ExitFormalParameters is called when production formalParameters is exited.
func (s *BasePlayScriptListener) ExitFormalParameters(ctx *FormalParametersContext) {}

// EnterFormalParameterList is called when production formalParameterList is entered.
func (s *BasePlayScriptListener) EnterFormalParameterList(ctx *FormalParameterListContext) {}

// ExitFormalParameterList is called when production formalParameterList is exited.
func (s *BasePlayScriptListener) ExitFormalParameterList(ctx *FormalParameterListContext) {}

// EnterFormalParameter is called when production formalParameter is entered.
func (s *BasePlayScriptListener) EnterFormalParameter(ctx *FormalParameterContext) {}

// ExitFormalParameter is called when production formalParameter is exited.
func (s *BasePlayScriptListener) ExitFormalParameter(ctx *FormalParameterContext) {}

// EnterLastFormalParameter is called when production lastFormalParameter is entered.
func (s *BasePlayScriptListener) EnterLastFormalParameter(ctx *LastFormalParameterContext) {}

// ExitLastFormalParameter is called when production lastFormalParameter is exited.
func (s *BasePlayScriptListener) ExitLastFormalParameter(ctx *LastFormalParameterContext) {}

// EnterVariableModifier is called when production variableModifier is entered.
func (s *BasePlayScriptListener) EnterVariableModifier(ctx *VariableModifierContext) {}

// ExitVariableModifier is called when production variableModifier is exited.
func (s *BasePlayScriptListener) ExitVariableModifier(ctx *VariableModifierContext) {}

// EnterClassDeclaration is called when production classDeclaration is entered.
func (s *BasePlayScriptListener) EnterClassDeclaration(ctx *ClassDeclarationContext) {}

// ExitClassDeclaration is called when production classDeclaration is exited.
func (s *BasePlayScriptListener) ExitClassDeclaration(ctx *ClassDeclarationContext) {}

// EnterClassBody is called when production classBody is entered.
func (s *BasePlayScriptListener) EnterClassBody(ctx *ClassBodyContext) {}

// ExitClassBody is called when production classBody is exited.
func (s *BasePlayScriptListener) ExitClassBody(ctx *ClassBodyContext) {}

// EnterClassBodyDeclaration is called when production classBodyDeclaration is entered.
func (s *BasePlayScriptListener) EnterClassBodyDeclaration(ctx *ClassBodyDeclarationContext) {}

// ExitClassBodyDeclaration is called when production classBodyDeclaration is exited.
func (s *BasePlayScriptListener) ExitClassBodyDeclaration(ctx *ClassBodyDeclarationContext) {}

// EnterMemberDecalaration is called when production memberDecalaration is entered.
func (s *BasePlayScriptListener) EnterMemberDecalaration(ctx *MemberDecalarationContext) {}

// ExitMemberDecalaration is called when production memberDecalaration is exited.
func (s *BasePlayScriptListener) ExitMemberDecalaration(ctx *MemberDecalarationContext) {}

// EnterFieldDeclaration is called when production fieldDeclaration is entered.
func (s *BasePlayScriptListener) EnterFieldDeclaration(ctx *FieldDeclarationContext) {}

// ExitFieldDeclaration is called when production fieldDeclaration is exited.
func (s *BasePlayScriptListener) ExitFieldDeclaration(ctx *FieldDeclarationContext) {}

// EnterParExpression is called when production parExpression is entered.
func (s *BasePlayScriptListener) EnterParExpression(ctx *ParExpressionContext) {}

// ExitParExpression is called when production parExpression is exited.
func (s *BasePlayScriptListener) ExitParExpression(ctx *ParExpressionContext) {}

// EnterForControl is called when production forControl is entered.
func (s *BasePlayScriptListener) EnterForControl(ctx *ForControlContext) {}

// ExitForControl is called when production forControl is exited.
func (s *BasePlayScriptListener) ExitForControl(ctx *ForControlContext) {}

// EnterForInit is called when production forInit is entered.
func (s *BasePlayScriptListener) EnterForInit(ctx *ForInitContext) {}

// ExitForInit is called when production forInit is exited.
func (s *BasePlayScriptListener) ExitForInit(ctx *ForInitContext) {}

// EnterVariableDeclarator is called when production variableDeclarator is entered.
func (s *BasePlayScriptListener) EnterVariableDeclarator(ctx *VariableDeclaratorContext) {}

// ExitVariableDeclarator is called when production variableDeclarator is exited.
func (s *BasePlayScriptListener) ExitVariableDeclarator(ctx *VariableDeclaratorContext) {}

// EnterArrayInitializer is called when production arrayInitializer is entered.
func (s *BasePlayScriptListener) EnterArrayInitializer(ctx *ArrayInitializerContext) {}

// ExitArrayInitializer is called when production arrayInitializer is exited.
func (s *BasePlayScriptListener) ExitArrayInitializer(ctx *ArrayInitializerContext) {}

// EnterVariableInitializer is called when production variableInitializer is entered.
func (s *BasePlayScriptListener) EnterVariableInitializer(ctx *VariableInitializerContext) {}

// ExitVariableInitializer is called when production variableInitializer is exited.
func (s *BasePlayScriptListener) ExitVariableInitializer(ctx *VariableInitializerContext) {}

// EnterVariableDeclarators is called when production variableDeclarators is entered.
func (s *BasePlayScriptListener) EnterVariableDeclarators(ctx *VariableDeclaratorsContext) {}

// ExitVariableDeclarators is called when production variableDeclarators is exited.
func (s *BasePlayScriptListener) ExitVariableDeclarators(ctx *VariableDeclaratorsContext) {}

// EnterEnhanceForControl is called when production enhanceForControl is entered.
func (s *BasePlayScriptListener) EnterEnhanceForControl(ctx *EnhanceForControlContext) {}

// ExitEnhanceForControl is called when production enhanceForControl is exited.
func (s *BasePlayScriptListener) ExitEnhanceForControl(ctx *EnhanceForControlContext) {}

// EnterVariableDeclaratorId is called when production variableDeclaratorId is entered.
func (s *BasePlayScriptListener) EnterVariableDeclaratorId(ctx *VariableDeclaratorIdContext) {}

// ExitVariableDeclaratorId is called when production variableDeclaratorId is exited.
func (s *BasePlayScriptListener) ExitVariableDeclaratorId(ctx *VariableDeclaratorIdContext) {}

// EnterSwitchBlockStatementGroup is called when production switchBlockStatementGroup is entered.
func (s *BasePlayScriptListener) EnterSwitchBlockStatementGroup(ctx *SwitchBlockStatementGroupContext) {
}

// ExitSwitchBlockStatementGroup is called when production switchBlockStatementGroup is exited.
func (s *BasePlayScriptListener) ExitSwitchBlockStatementGroup(ctx *SwitchBlockStatementGroupContext) {
}

// EnterSwitchLabel is called when production switchLabel is entered.
func (s *BasePlayScriptListener) EnterSwitchLabel(ctx *SwitchLabelContext) {}

// ExitSwitchLabel is called when production switchLabel is exited.
func (s *BasePlayScriptListener) ExitSwitchLabel(ctx *SwitchLabelContext) {}

// EnterStatement is called when production statement is entered.
func (s *BasePlayScriptListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BasePlayScriptListener) ExitStatement(ctx *StatementContext) {}
