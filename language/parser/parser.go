package parser

import (
	"github.com/graphql-go/graphql/language/ast"
	"github.com/graphql-go/graphql/language/lexer"
	"github.com/graphql-go/graphql/language/source"
)

type parseFn func(parser *Parser) (interface{}, error)

type parseDefinitionFn func(parser *Parser) (ast.Node, error)

var tokenDefinitionFn map[string]parseDefinitionFn

func init() {
	tokenDefinitionFn = make(map[string]parseDefinitionFn)
	{
		tokenDefinitionFn[lexer.FRAGMENT] = parseFragmentDefinition
		tokenDefinitionFn[lexer.QUERY] = parseOperationDefinition
		tokenDefinitionFn[lexer.MUTATION] = parseOperationDefinition
		tokenDefinitionFn[lexer.SUBSCRIPTION] = parseOperationDefinition
		tokenDefinitionFn[lexer.SCHEMA] = parseSchemaDefinition
		tokenDefinitionFn[lexer.SCALAR] = parseScalarTypeDefinition
		tokenDefinitionFn[lexer.TYPE] = parseObjectTypeDefinition
		tokenDefinitionFn[lexer.INTERFACE] = parseInterfaceTypeDefinition
		tokenDefinitionFn[lexer.UNION] = parseUnionTypeDefinition
		tokenDefinitionFn[lexer.ENUM] = parseEnumTypeDefinition
		tokenDefinitionFn[lexer.INPUT] = parseInputObjectTypeDefinition
		tokenDefinitionFn[lexer.EXTEND] = parseTypeExtensionDefinition
		tokenDefinitionFn[lexer.DIRECTIVE] = parseDirectiveDefinition
	}
}

type ParseOptions struct {
	NoLocation bool
	NoSource   bool
}

type ParseParams struct {
	Source  interface{}
	Options ParseOptions
}

type Parser struct {
	LexToken lexer.Lexer
	Source   *source.Source
	Options  ParseOptions
	PrevEnd  int
	Token    lexer.Token
}

func Parse(p ParseParams) (*ast.Document, error) { _ = "STUB: not implemented"; return nil, nil }

func ParseValue(p ParseParams) (ast.Value, error) {
	_ = "STUB: not implemented"
	return *new(ast.Value), nil
}

func parseName(parser *Parser) (*ast.Name, error) { _ = "STUB: not implemented"; return nil, nil }

func makeParser(s *source.Source, opts ParseOptions) (*Parser, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseDocument(parser *Parser) (*ast.Document, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseOperationDefinition(parser *Parser) (ast.Node, error) {
	_ = "STUB: not implemented"
	return *new(ast.Node), nil
}

func parseOperationType(parser *Parser) (string, error) { _ = "STUB: not implemented"; return "", nil }

func parseVariableDefinitions(parser *Parser) ([]*ast.VariableDefinition, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseVariableDefinition(parser *Parser) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseVariable(parser *Parser) (*ast.Variable, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseSelectionSet(parser *Parser) (*ast.SelectionSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseSelection(parser *Parser) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseField(parser *Parser) (*ast.Field, error) { _ = "STUB: not implemented"; return nil, nil }

func parseArguments(parser *Parser) ([]*ast.Argument, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseArgument(parser *Parser) (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

func parseFragment(parser *Parser) (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

func parseFragmentDefinition(parser *Parser) (ast.Node, error) {
	_ = "STUB: not implemented"
	return *new(ast.Node), nil
}

func parseFragmentName(parser *Parser) (*ast.Name, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseValueLiteral(parser *Parser, isConst bool) (ast.Value, error) {
	_ = "STUB: not implemented"
	return *new(ast.Value), nil
}

func parseConstValue(parser *Parser) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseValueValue(parser *Parser) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseList(parser *Parser, isConst bool) (*ast.ListValue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseObject(parser *Parser, isConst bool) (*ast.ObjectValue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseObjectField(parser *Parser, isConst bool) (*ast.ObjectField, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseDirectives(parser *Parser) ([]*ast.Directive, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseDirective(parser *Parser) (*ast.Directive, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseType(parser *Parser) (ttype ast.Type, err error) {
	_ = "STUB: not implemented"
	return *new(ast.Type), nil
}

func parseNamed(parser *Parser) (*ast.Named, error) { _ = "STUB: not implemented"; return nil, nil }

func parseTypeSystemDefinition(parser *Parser) (ast.Node, error) {
	_ = "STUB: not implemented"
	return *new(ast.Node), nil
}

func parseSchemaDefinition(parser *Parser) (ast.Node, error) {
	_ = "STUB: not implemented"
	return *new(ast.Node), nil
}

func parseOperationTypeDefinition(parser *Parser) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseScalarTypeDefinition(parser *Parser) (ast.Node, error) {
	_ = "STUB: not implemented"
	return *new(ast.Node), nil
}

func parseObjectTypeDefinition(parser *Parser) (ast.Node, error) {
	_ = "STUB: not implemented"
	return *new(ast.Node), nil
}

func parseImplementsInterfaces(parser *Parser) ([]*ast.Named, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseFieldDefinition(parser *Parser) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseArgumentDefs(parser *Parser) ([]*ast.InputValueDefinition, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseInputValueDef(parser *Parser) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseInterfaceTypeDefinition(parser *Parser) (ast.Node, error) {
	_ = "STUB: not implemented"
	return *new(ast.Node), nil
}

func parseUnionTypeDefinition(parser *Parser) (ast.Node, error) {
	_ = "STUB: not implemented"
	return *new(ast.Node), nil
}

func parseUnionMembers(parser *Parser) ([]*ast.Named, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseEnumTypeDefinition(parser *Parser) (ast.Node, error) {
	_ = "STUB: not implemented"
	return *new(ast.Node), nil
}

func parseEnumValueDefinition(parser *Parser) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseInputObjectTypeDefinition(parser *Parser) (ast.Node, error) {
	_ = "STUB: not implemented"
	return *new(ast.Node), nil
}

func parseTypeExtensionDefinition(parser *Parser) (ast.Node, error) {
	_ = "STUB: not implemented"
	return *new(ast.Node), nil
}

func parseDirectiveDefinition(parser *Parser) (ast.Node, error) {
	_ = "STUB: not implemented"
	return *new(ast.Node), nil
}

func parseDirectiveLocations(parser *Parser) ([]*ast.Name, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseStringLiteral(parser *Parser) (*ast.StringValue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseDescription(parser *Parser) (*ast.StringValue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func loc(parser *Parser, start int) *ast.Location { _ = "STUB: not implemented"; return nil }

func advance(parser *Parser) error { _ = "STUB: not implemented"; return nil }

func lookahead(parser *Parser) (lexer.Token, error) {
	_ = "STUB: not implemented"
	return *new(lexer.Token), nil
}

func peek(parser *Parser, Kind lexer.TokenKind) bool { _ = "STUB: not implemented"; return false }

func peekDescription(parser *Parser) bool { _ = "STUB: not implemented"; return false }

func skip(parser *Parser, Kind lexer.TokenKind) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func expect(parser *Parser, kind lexer.TokenKind) (lexer.Token, error) {
	_ = "STUB: not implemented"
	return *new(lexer.Token), nil
}

func expectKeyWord(parser *Parser, value string) (lexer.Token, error) {
	_ = "STUB: not implemented"
	return *new(lexer.Token), nil
}

func unexpected(parser *Parser, atToken lexer.Token) error { _ = "STUB: not implemented"; return nil }

func unexpectedEmpty(parser *Parser, beginLoc int, openKind, closeKind lexer.TokenKind) error {
	_ = "STUB: not implemented"
	return nil
}

func reverse(parser *Parser, openKind lexer.TokenKind, parseFn parseFn, closeKind lexer.TokenKind, zinteger bool) ([]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
