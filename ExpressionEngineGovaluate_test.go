package GoMybatis

import (
	"fmt"
	"github.com/zhuxiujia/GoMybatis/example"
	"testing"
)

func TestExpress(t *testing.T) {
	var activity = example.Activity{
		Id:         "1",
		DeleteFlag: 1,
	}
	var engine = ExpressionEngineGovaluate{}
	var evaluateParameters = make(map[string]interface{})

	evaluateParameters["activity"] = &activity

	var expression = "activity.DeleteFlag == 1 || activity.DeleteFlag > 0 "
	evalExpression, err := engine.Lexer(expression)
	if err != nil {
		t.Fatal(err)
	}
	result, err := engine.Eval(evalExpression,evaluateParameters,0)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
}

func BenchmarkExpress(b *testing.B) {
	b.StopTimer()
	var activity = example.Activity{
		Id:         "1",
		DeleteFlag: 1,
	}

	var engine = ExpressionEngineGovaluate{}
	var evaluateParameters = make(map[string]interface{})

	evaluateParameters["activity"] = &activity

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		var expression = "activity.DeleteFlag == 1 || activity.DeleteFlag > 0 "
		evalExpression, err := engine.Lexer(expression)
		if err != nil {
			b.Fatal(err)
		}
		result, err := engine.Eval(evalExpression,evaluateParameters,0)
		if err != nil {
			b.Fatal(err)
		}
		if result.(bool){

		}
	}
}
