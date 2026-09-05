package evaluator

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode/utf8"

	"github.com/dayattt111/lontara-lang/pkg/ast"
	"github.com/dayattt111/lontara-lang/pkg/object"
	"github.com/dayattt111/lontara-lang/pkg/token"
)

var (
	NULL  = &object.Null{}
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}

	outWriter io.Writer = os.Stdout
)

func SetOutput(w io.Writer) {
	outWriter = w
}

func GetOutput() io.Writer {
	if outWriter == nil {
		return os.Stdout
	}
	return outWriter
}

func builtinPanjang(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError(token.Token{}, "salawuk panganggara: diharapkan 1 argumen, didapat=%d", len(args))
	}
	switch arg := args[0].(type) {
	case *object.String:
		return &object.Integer{Value: int64(utf8.RuneCountInString(arg.Value))}
	default:
		return newError(token.Token{}, "argumen untuk 'panjang' tenasialai, didapat=%s", args[0].Type())
	}
}

func builtinBaca(args ...object.Object) object.Object {
	if len(args) != 0 {
		return newError(token.Token{}, "salawuk panganggara: 'baca' teppu marisi argumen")
	}
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		return &object.String{Value: scanner.Text()}
	}
	return &object.String{Value: ""}
}

func builtinTipe(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError(token.Token{}, "salawuk panganggara: diharapkan 1 argumen, didapat=%d", len(args))
	}
	return &object.String{Value: string(args[0].Type())}
}

// builtins mendaftarkan fungsi bawaan seperti paui (cetak ke konsol)
var builtins = map[string]*object.Builtin{
	"paui": {
		Fn: func(args ...object.Object) object.Object {
			out := GetOutput()
			for i, arg := range args {
				if i > 0 {
					fmt.Fprint(out, " ")
				}
				fmt.Fprint(out, arg.Inspect())
			}
			fmt.Fprintln(out)
			return NULL
		},
	},
	// Dukungan built-in aksara Lontara: ᨄᨕᨘᨕᨗ (paui)
	"ᨄᨕᨘᨕᨗ": {
		Fn: func(args ...object.Object) object.Object {
			out := GetOutput()
			for i, arg := range args {
				if i > 0 {
					fmt.Fprint(out, " ")
				}
				fmt.Fprint(out, arg.Inspect())
			}
			fmt.Fprintln(out)
			return NULL
		},
	},
	"panjang": {Fn: builtinPanjang},
	"ᨄᨍ":      {Fn: builtinPanjang},
	"baca":    {Fn: builtinBaca},
	"ᨅᨌ":      {Fn: builtinBaca},
	"tipe":    {Fn: builtinTipe},
	"ᨈᨗᨄᨙ":    {Fn: builtinTipe},
}

// Eval menelusuri AST dan mengembalikan objek hasil evaluasi
func Eval(node ast.Node, env *object.Environment) object.Object {
	switch node := node.(type) {

	// Statements
	case *ast.Program:
		return evalProgram(node, env)

	case *ast.ExpressionStatement:
		return Eval(node.Expression, env)

	case *ast.BlockStatement:
		return evalBlockStatement(node, env)

	case *ast.TaroiStatement:
		val := Eval(node.Value, env)
		if isError(val) {
			return val
		}
		env.Set(node.Name.Value, val)
		return NULL

	case *ast.LisuStatement:
		val := Eval(node.ReturnValue, env)
		if isError(val) {
			return val
		}
		return &object.ReturnValue{Value: val}

	// Expressions
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}

	case *ast.StringLiteral:
		return &object.String{Value: node.Value}

	case *ast.Boolean:
		return nativeBoolToBooleanObject(node.Value)

	case *ast.PrefixExpression:
		right := Eval(node.Right, env)
		if isError(right) {
			return right
		}
		return evalPrefixExpression(node, right)

	case *ast.InfixExpression:
		left := Eval(node.Left, env)
		if isError(left) {
			return left
		}
		right := Eval(node.Right, env)
		if isError(right) {
			return right
		}
		return evalInfixExpression(node, left, right)

	case *ast.RekkoExpression:
		return evalRekkoExpression(node, env)

	case *ast.SikiExpression:
		return evalSikiExpression(node, env)

	case *ast.Identifier:
		return evalIdentifier(node, env)

	case *ast.JamagauLiteral:
		params := node.Parameters
		body := node.Body
		return &object.Function{Parameters: params, Body: body, Env: env}

	case *ast.CallExpression:
		function := Eval(node.Function, env)
		if isError(function) {
			return function
		}

		args := evalExpressions(node.Arguments, env)
		if len(args) == 1 && isError(args[0]) {
			return args[0]
		}

		return applyFunction(node.Token, function, args)
	}

	return nil
}

func evalProgram(program *ast.Program, env *object.Environment) object.Object {
	var result object.Object

	for _, statement := range program.Statements {
		result = Eval(statement, env)

		switch result := result.(type) {
		case *object.ReturnValue:
			return result.Value
		case *object.Error:
			return result
		}
	}

	return result
}

func evalBlockStatement(block *ast.BlockStatement, env *object.Environment) object.Object {
	var result object.Object

	for _, statement := range block.Statements {
		result = Eval(statement, env)

		if result != nil {
			rt := result.Type()
			if rt == object.RETURN_VALUE_OBJ || rt == object.ERROR_OBJ {
				return result
			}
		}
	}

	return result
}

func nativeBoolToBooleanObject(input bool) *object.Boolean {
	if input {
		return TRUE
	}
	return FALSE
}

func evalPrefixExpression(node *ast.PrefixExpression, right object.Object) object.Object {
	switch node.Operator {
	case "!":
		return evalBangOperatorExpression(right)
	case "-":
		return evalMinusPrefixOperatorExpression(node, right)
	default:
		return newError(node.Token, "operator tenriisseng: %s%s", node.Operator, right.Type())
	}
}

func evalBangOperatorExpression(right object.Object) object.Object {
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return TRUE
	case nil:
		return TRUE
	default:
		return FALSE
	}
}

func evalMinusPrefixOperatorExpression(node *ast.PrefixExpression, right object.Object) object.Object {
	if right == nil {
		return newError(node.Token, "ekspresi kanan bernilai nil")
	}

	if right.Type() != object.INTEGER_OBJ {
		return newError(node.Token, "operator tenriisseng: -%s", right.Type())
	}

	value := right.(*object.Integer).Value
	return &object.Integer{Value: -value}
}

func evalInfixExpression(node *ast.InfixExpression, left, right object.Object) object.Object {
	switch {
	case left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ:
		return evalIntegerInfixExpression(node, left, right)
	case left.Type() == object.STRING_OBJ && right.Type() == object.STRING_OBJ:
		return evalStringInfixExpression(node, left, right)
	case node.Operator == "==":
		return nativeBoolToBooleanObject(left == right)
	case node.Operator == "!=":
		return nativeBoolToBooleanObject(left != right)
	case left.Type() != right.Type():
		return newError(node.Token, "tipe data tenasialai: %s %s %s", left.Type(), node.Operator, right.Type())
	default:
		return newError(node.Token, "operator tenriisseng: %s %s %s", left.Type(), node.Operator, right.Type())
	}
}

func evalIntegerInfixExpression(node *ast.InfixExpression, left, right object.Object) object.Object {
	leftVal := left.(*object.Integer).Value
	rightVal := right.(*object.Integer).Value

	switch node.Operator {
	case "+":
		return &object.Integer{Value: leftVal + rightVal}
	case "-":
		return &object.Integer{Value: leftVal - rightVal}
	case "*":
		return &object.Integer{Value: leftVal * rightVal}
	case "/":
		if rightVal == 0 {
			return newError(node.Token, "pambagean nol: tekkulle nibage nol")
		}
		return &object.Integer{Value: leftVal / rightVal}
	case "<":
		return nativeBoolToBooleanObject(leftVal < rightVal)
	case ">":
		return nativeBoolToBooleanObject(leftVal > rightVal)
	case "<=":
		return nativeBoolToBooleanObject(leftVal <= rightVal)
	case ">=":
		return nativeBoolToBooleanObject(leftVal >= rightVal)
	case "==":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObject(leftVal != rightVal)
	default:
		return newError(node.Token, "operator tenriisseng: %s %s %s", left.Type(), node.Operator, right.Type())
	}
}

func evalStringInfixExpression(node *ast.InfixExpression, left, right object.Object) object.Object {
	if node.Operator != "+" {
		return newError(node.Token, "operator tenriisseng: %s %s %s", left.Type(), node.Operator, right.Type())
	}

	leftVal := left.(*object.String).Value
	rightVal := right.(*object.String).Value
	return &object.String{Value: leftVal + rightVal}
}

func evalRekkoExpression(ie *ast.RekkoExpression, env *object.Environment) object.Object {
	condition := Eval(ie.Condition, env)
	if isError(condition) {
		return condition
	}

	if isTruthy(condition) {
		return Eval(ie.Consequence, env)
	} else if ie.Alternative != nil {
		return Eval(ie.Alternative, env)
	} else {
		return NULL
	}
}

func evalSikiExpression(se *ast.SikiExpression, env *object.Environment) object.Object {
	var result object.Object = NULL

	for {
		condition := Eval(se.Condition, env)
		if isError(condition) {
			return condition
		}

		if !isTruthy(condition) {
			break
		}

		result = Eval(se.Body, env)

		if result != nil {
			rt := result.Type()
			if rt == object.RETURN_VALUE_OBJ || rt == object.ERROR_OBJ {
				return result
			}
		}
	}

	return result
}

func evalIdentifier(node *ast.Identifier, env *object.Environment) object.Object {
	if val, ok := env.Get(node.Value); ok {
		return val
	}

	if builtin, ok := builtins[node.Value]; ok {
		return builtin
	}

	return newError(node.Token, "variabel tekkedeteng (tenriisseng): "+node.Value)
}

func evalExpressions(exps []ast.Expression, env *object.Environment) []object.Object {
	var result []object.Object

	for _, e := range exps {
		evaluated := Eval(e, env)
		if isError(evaluated) {
			return []object.Object{evaluated}
		}
		result = append(result, evaluated)
	}

	return result
}

func applyFunction(tok token.Token, fn object.Object, args []object.Object) object.Object {
	switch fn := fn.(type) {
	case *object.Function:
		extendedEnv := extendFunctionEnv(fn, args)
		evaluated := Eval(fn.Body, extendedEnv)
		return unwrapReturnValue(evaluated)
	case *object.Builtin:
		return fn.Fn(args...)
	default:
		return newError(tok, "bannang fungsi: %s", fn.Type())
	}
}

func extendFunctionEnv(fn *object.Function, args []object.Object) *object.Environment {
	env := object.NewEnclosedEnvironment(fn.Env)

	for paramIdx, param := range fn.Parameters {
		env.Set(param.Value, args[paramIdx])
	}

	return env
}

func unwrapReturnValue(obj object.Object) object.Object {
	if returnValue, ok := obj.(*object.ReturnValue); ok {
		return returnValue.Value
	}
	return obj
}

func isTruthy(obj object.Object) bool {
	switch obj {
	case NULL:
		return false
	case TRUE:
		return true
	case FALSE:
		return false
	default:
		return true
	}
}

func newError(tok token.Token, format string, a ...interface{}) *object.Error {
	msg := fmt.Sprintf(format, a...)
	if tok.Line > 0 {
		return &object.Error{Message: fmt.Sprintf("[Baris %d, Kolom %d] %s", tok.Line, tok.Column, msg)}
	}
	return &object.Error{Message: msg}
}

func isError(obj object.Object) bool {
	if obj != nil {
		return obj.Type() == object.ERROR_OBJ
	}
	return false
}
