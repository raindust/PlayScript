/*
 [The "BSD licence"] Copyright (c) 2013 Terence Parr, Sam Harwell Copyright (c) 2017 Ivan Kochurkin
 (upgrade to Java 8) Copyright (c) 2019 宫文学 Copyright (c) 2019 Yan Mingzhi All rights reserved.
 Redistribution and use in source and binary forms, with or
 without modification, are permitted provided that the following conditions are met: 1.
 Redistributions of source code must retain the above copyright notice, this list of conditions and
 the following disclaimer. 2. Redistributions in binary form must reproduce the above copyright
 notice, this list of conditions and the following disclaimer in the documentation and/or other
 materials provided with the distribution. 3. The name of the author may not be used to endorse or
 promote products derived from this software without specific prior written permission. THIS
 SOFTWARE IS PROVIDED BY THE AUTHOR ``AS IS'' AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT
 NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
 DISCLAIMED. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL,
 EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS
 OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
 THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR
 OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY
 OF SUCH DAMAGE.
 */

grammar PlayScript;

import CommonLexer; //import lexer

classOrInterfaceType: IDENTIFIER ('.' IDENTIFIER)*;

typeTypeOrVoid: typeType | VOID;

functionType: FUNCTION typeTypeOrVoid '(' typeList? ')';

primitiveType:
	BOOLEAN
	| CHAR
	| BYTE
	| SHORT
	| INT
	| LONG
	| FLOAT
	| DOUBLE
	| STRING;

typeType: (classOrInterfaceType | functionType | primitiveType) (
		'[' ']'
	)*;

typeList: typeType (',' typeType)*;

expression:
	primary
	| expression bop = '.' (IDENTIFIER | functionCall | THIS)
	| expression '[' expression ']'
	| functionCall
	| expression postfix = ('++' | '--')
	| prefix = ('+' | '-' | '++' | '--') expression
	| prefix = ('~' | '!') expression
	| expression bop = ('*' | '/' | '%') expression
	| expression bop = ('+' | '-') expression
	| expression ('<' '<' | '>' '>' '>' | '>' '>') expression
	| expression bop = ('<=' | '>=' | '>' | '<') expression
	| expression bop = INSTANCEOF typeType
	| expression bop = ('==' | '!=') expression
	| expression bop = '&' expression
	| expression bop = '^' expression
	| expression bop = '|' expression
	| expression bop = '&&' expression
	| expression bop = '||' expression
	| expression bop = '?' expression ':' expression
	| <assoc = right> expression bop = (
		'='
		| '+='
		| '-='
		| '*='
		| '/='
		| '&='
		| '|='
		| '^='
		| '>>='
		| '>>>='
		| '<<='
		| '%='
	) expression;

expressionList: expression (',' expression)*;

functionCall:
	IDENTIFIER '(' expressionList? ')'
	| THIS '(' expressionList? ')'
	| SUPER '(' expressionList? ')';

primary:
	'(' expression ')'
	| THIS
	| SUPER
	| literal
	| IDENTIFIER;

intLiteral:
	DECIMAL_LITERAL
	| HEX_LITERAL
	| OCT_LITERAL
	| BINARY_LITERAL;

floatLiteral: FLOAT_LITERAL | HEX_FLOAT_LITERAL;

literal:
	intLiteral
	| floatLiteral
	| CHAR_LITERAL
	| STRING_LITERAL
	| BOOL_LITERAL
	| NULL_LITERAL;

block: '{' blockStatements '}';

blockStatement:
	variableDeclarators ';'
	| statement
	| functionDeclaration
	| classDeclaration;

blockStatements: blockStatement*;

functionBody: block | ';';

qualifiedName: IDENTIFIER ('.' IDENTIFIER)*;

qualifiedNameList: qualifiedName (',' qualifiedName)*;

functionDeclaration:
	typeTypeOrVoid? IDENTIFIER formalParameters ('[' ']')* (
		THROWS qualifiedNameList
	)? functionBody;

formalParameters: '(' formalParameterList? ')';

formalParameterList:
	formalParameter (',' formalParameter)* (
		',' lastFormalParameter
	)?
	| lastFormalParameter;

formalParameter:
	variableModifier* typeType variableDeclaratorId;

lastFormalParameter:
	variableModifier* typeType '...' variableDeclaratorId;

variableModifier: FINAL;

classDeclaration:
	CLASS IDENTIFIER (EXTENDS typeType)? (IMPLEMENTS typeList)? classBody;

classBody: '{' classBodyDeclaration* '}';

classBodyDeclaration: ';' | memberDecalaration;

memberDecalaration:
	functionDeclaration
	| fieldDeclaration
	| classDeclaration;

fieldDeclaration: variableDeclarators ';';

parExpression: '(' expression ')';

forControl:
	enhanceForControl
	| forInit? ';' expression? ';' forUpdate = expressionList;

forInit: variableDeclarators | expressionList;

variableDeclarator:
	variableDeclaratorId ('=' variableInitializer);

arrayInitializer:
	'{' (variableInitializer (',' variableInitializer)* (',')?)? '}';

variableInitializer: arrayInitializer | expression;

variableDeclarators:
	typeType variableDeclarator (',' variableDeclarator)*;

enhanceForControl: typeType variableDeclaratorId ':' expression;

variableDeclaratorId: IDENTIFIER ('[' ']')*;

switchBlockStatementGroup: switchLabel+ blockStatement+;

switchLabel:
	CASE (
		constantExpression = expression
		| enumConstantName = IDENTIFIER
	) ':'
	| DEFAULT ':';

statement:
	blockLabel = block
	| IF parExpression statement (ELSE statement)?
	| FOR '(' forControl ')' statement
	| WHILE parExpression statement
	| DO statement WHILE parExpression ';'
	| SWITCH parExpression '{' switchBlockStatementGroup* switchLabel* '}'
	| RETURN expression? ';'
	| BREAK IDENTIFIER? ';'
	| CONTINUE IDENTIFIER? ';'
	| SEMI
	| statementExpression = expression
	| identifierLabel = IDENTIFIER ':' statement;
