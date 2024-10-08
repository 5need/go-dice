// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"go-dice/models"
	"strconv"
)

func Main() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var2 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
			templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
			if !templ_7745c5c3_IsBuffer {
				defer func() {
					templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
					if templ_7745c5c3_Err == nil {
						templ_7745c5c3_Err = templ_7745c5c3_BufErr
					}
				}()
			}
			ctx = templ.InitializeContext(ctx)
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div _=\"\n\t\t\tdef animate()\n\t\t\t\tif the length of .sheen &gt; 0\n\t\t\t\t\tthen set x to random .sheen\n\t\t\t\t\tthen take .animate-diceShine for x\n\t\t\t\t\tthen add .animate-diceShine to x\n\t\t\t\tend\n\t\t\tend\n\t\t\tinit wait 1s\n\t\t\tthen repeat forever\n\t\t\t\tanimate()\n\t\t\t\tthen wait 3s\n\t\t\tend\n\t\t\t\" class=\"animate-diceShine hidden\"></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = Form(models.NewRollStats()).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = Base("Go-Dice").Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func Form(rollStats models.RollStats) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var3 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var3 == nil {
			templ_7745c5c3_Var3 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<form hx-post=\"/\" hx-swap=\"outerHTML\" class=\"[&amp;_input]:bg-transparent flex flex-col\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = DiceBox(rollStats).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex flex-wrap gap-1\"><button type=\"button\" _=\"on click set the value of #input to &#39;-1+&#39; then send submit to the closest &lt;form/&gt;\" class=\"w-24 h-16 leading-none bg-ctp-surface0 rounded\">clear</button>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = AddDiceButton("1").Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = AddDiceButton("2").Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = AddDiceButton("5").Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = AddDiceButton("10").Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = AddDiceButton("25").Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<button type=\"button\" _=\"\n\t\t\t\ton click\n\t\t\t\t\trepeat in &lt;input[name=selection]/&gt;\n\t\t\t\t\t\tjs(it) it.checked = !it.checked; end\n\t\t\t\t\tend\n\t\t\t\tend\n\t\t\t\t\" class=\"w-24 h-16 leading-none bg-ctp-surface0 rounded\">invert selection</button></div><input id=\"input\" name=\"input\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(rollStats.PreviousInput)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/main.templ`, Line: 57, Col: 64}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"> <input id=\"currentResult\" name=\"currentResult\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var5 string
		templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(rollStats.CurrentRoll)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/main.templ`, Line: 58, Col: 78}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"> <button>submit</button></form>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func AddDiceButton(num string) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var6 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var6 == nil {
			templ_7745c5c3_Var6 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<button type=\"button\" _=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var7 string
		templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs("on click set the value of #input to '+" + num + "d6' then send submit to the closest <form/>")
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/main.templ`, Line: 66, Col: 100}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"w-24 h-16 leading-none bg-ctp-surface0 rounded\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var8 string
		templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(num)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/main.templ`, Line: 69, Col: 7}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("x <img src=\"/images/dice-d6.svg\" class=\"w-4 invert inline\"></button>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func DiceBox(rollStats models.RollStats) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var9 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var9 == nil {
			templ_7745c5c3_Var9 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<ul id=\"diceBox\" class=\"flex flex-col overflow-x-auto select-none\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for i := 5; i >= 0; i-- {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li class=\"flex h-20 items-center min-w-max border-b border-ctp-surface0 last:border-none\" _=\"\n\t\t\t\ton touchmove(touch)\n\t\t\t\t\tset the innerHTML of &lt;div.info/&gt; in me to touch.touches[0].clientX\n\t\t\t\tend\n\t\t\t\t\"><div class=\"info\"></div><button class=\"font-mono h-12 w-20 border-2 border-ctp-blue/10 flex items-center justify-center\"><img src=\"images/16/solid/arrow-path.svg\" class=\"inline-block invert\"> ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if i+1 != 1 {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<img src=\"images/16/solid/arrow-right.svg\" class=\"inline-block invert opacity-50 mx-0.5\"> ")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			if i+1 == 6 {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("6s")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else if i+1 == 1 {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("all")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				var templ_7745c5c3_Var10 string
				templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.Itoa(i + 1))
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/main.templ`, Line: 96, Col: 25}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("+")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</button><div class=\"self-center w-8 text-center font-mono\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if rollStats.RollAmounts[i] > 0 {
				var templ_7745c5c3_Var11 string
				templ_7745c5c3_Var11, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.Itoa(rollStats.RollAmounts[i]))
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/main.templ`, Line: 101, Col: 46}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var11))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("x")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("-")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			for j := 0; j < rollStats.RollAmounts[i]; j++ {
				templ_7745c5c3_Err = Dice(i+1, j, Some(rollStats.Selection, func(n models.DiceIdentifier) bool {
					return n == models.DiceIdentifier{RollValue: i + 1, FromTheLeft: j}
				})).Render(ctx, templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</li>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</ul>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func Some[T any](slice []T, predicate func(T) bool) bool {
	for _, v := range slice {
		if predicate(v) {
			return true
		}
	}
	return false
}

func Dice(rollValue int, fromTheLeft int, checked bool) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var12 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var12 == nil {
			templ_7745c5c3_Var12 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<label data-rollValue=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var13 string
		templ_7745c5c3_Var13, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.Itoa(rollValue))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/main.templ`, Line: 127, Col: 42}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var13))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"\n\t\tgroup\n\t\tdata-[rollValue=&#39;1&#39;]:[animation-delay:0ms]\n\t\tdata-[rollValue=&#39;2&#39;]:[animation-delay:25ms]\n\t\tdata-[rollValue=&#39;3&#39;]:[animation-delay:50ms]\n\t\tdata-[rollValue=&#39;4&#39;]:[animation-delay:75ms]\n\t\tdata-[rollValue=&#39;5&#39;]:[animation-delay:100ms]\n\t\tdata-[rollValue=&#39;6&#39;]:[animation-delay:125ms]\n\t\tborder-4\n\t\tborder-transparent\n\t\th-20 flex items-center justify-center\n\t\trelative\n\t\t//[&amp;_div.counter]:last-of-type:has-[input:checked]:flex\n\t\t//[&amp;_div.counter]:[&amp;:has(+_label_input:not(:checked)):has(input:checked)]:flex\n\t\t\" _=\"\n\t\ton dblclick \n\t\t\trepeat in &lt;input/&gt; in the closest &lt;li/&gt;\n\t\t\t\tset it&#39;s checked to true\n\t\t\tend\n\t\tend\n\t\ton touchmove\n\t\t\tsend lol to &lt;input/&gt; in me\n\t\tend\n\t\t\"><input data-rollValue=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var14 string
		templ_7745c5c3_Var14, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.Itoa(rollValue))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/main.templ`, Line: 155, Col: 43}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var14))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" data-fromTheLeft=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var15 string
		templ_7745c5c3_Var15, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.Itoa(fromTheLeft))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/main.templ`, Line: 156, Col: 47}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var15))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" name=\"selection\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if checked {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" checked")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var16 string
		templ_7745c5c3_Var16, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.Itoa(rollValue) + "@" + strconv.Itoa(fromTheLeft))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/main.templ`, Line: 159, Col: 68}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var16))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" type=\"checkbox\" class=\"hidden peer\" _=\"\n\t\t\tdef fillLeftClearRight()\n\t\t\t\tset y to my @data-fromTheLeft\n\t\t\t\tthen set y to parseInt(y)\n\t\t\t\tthen repeat in &lt;input/&gt; in the closest &lt;li/&gt;\n\t\t\t\t\tset x to it&#39;s @data-fromTheLeft\n\t\t\t\t\tthen set x to parseInt(x)\n\t\t\t\t\tthen if x &lt; y\n\t\t\t\t\t\tjs(it) it.checked = true; end\n\t\t\t\t\tend\n\t\t\t\t\tthen if x &gt; y\n\t\t\t\t\t\tjs(it) it.checked = false; end\n\t\t\t\t\tend\n\t\t\t\tend\n\t\t\tend\n\n\t\t\tdef clearAll()\n\t\t\t\trepeat in &lt;input/&gt; in the closest &lt;li/&gt;\n\t\t\t\t\tjs(it) it.checked = false; end\n\t\t\t\tend\n\t\t\tend\n\n\t\t\ton click\n\t\t\t\tif my checked == true\n\t\t\t\t\tfillLeftClearRight()\n\t\t\t\telse\n\t\t\t\t\tset isLast to true\n\t\t\t\t\tthen repeat in &lt;input:checked/&gt; in the closest &lt;li/&gt;\n\t\t\t\t\t\tset x to it&#39;s @data-fromTheLeft\n\t\t\t\t\t\tthen set x to parseInt(x)\n\t\t\t\t\t\tthen set y to my @data-fromTheLeft\n\t\t\t\t\t\tthen set y to parseInt(y)\n\t\t\t\t\t\tthen if x &gt; y\n\t\t\t\t\t\t\tset isLast to false\n\t\t\t\t\t\tend\n\t\t\t\t\tend\n\t\t\t\t\tthen if isLast\n\t\t\t\t\t\tclearAll()\n\t\t\t\t\telse\n\t\t\t\t\t\thalt the event\n\t\t\t\t\t\tthen fillLeftClearRight()\n\t\t\t\t\tend\n\t\t\t\tend\n\t\t\tend\n\n\t\t\ton lol\n\t\t\t\tfillLeftClearRight()\n\t\t\tend\n\n\t\t\t\"><div class=\"[&amp;_img]:w-full [&amp;_img]:h-full [&amp;_img]:opacity-75 peer-checked:[&amp;_img]:opacity-10\n\t\t\tw-12 h-12 rounded-lg pointer-events-none\n\t\t\tgroup-data-[rollValue=&#39;1&#39;]:bg-ctp-subtext0\n\t\t\tgroup-data-[rollValue=&#39;2&#39;]:bg-ctp-text\n\t\t\tgroup-data-[rollValue=&#39;3&#39;]:bg-ctp-green\n\t\t\tgroup-data-[rollValue=&#39;4&#39;]:bg-ctp-blue\n\t\t\tgroup-data-[rollValue=&#39;5&#39;]:bg-ctp-mauve\n\t\t\tgroup-data-[rollValue=&#39;6&#39;]:bg-ctp-yellow\n\t\t\trelative\n\t\t\t\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = DiceSheen().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = DiceSelectionCurrentCount(fromTheLeft+1).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"\n\t\t\t\thidden\n\t\t\t\tgroup-[:not(:first-of-type):has(input:checked)]:block\n\t\t\t\tgroup-data-[rollValue=&#39;1&#39;]:bg-ctp-subtext0\n\t\t\t\tgroup-data-[rollValue=&#39;2&#39;]:bg-ctp-text\n\t\t\t\tgroup-data-[rollValue=&#39;3&#39;]:bg-ctp-green\n\t\t\t\tgroup-data-[rollValue=&#39;4&#39;]:bg-ctp-blue\n\t\t\t\tgroup-data-[rollValue=&#39;5&#39;]:bg-ctp-mauve\n\t\t\t\tgroup-data-[rollValue=&#39;6&#39;]:bg-ctp-yellow\n\t\t\t\tabsolute inset-0 -z-10 -translate-x-1/2 bg-ctp-blue\n\t\t\t\tpointer-events-none\"></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		switch rollValue {
		case 1:
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<img src=\"/images/dice-one.svg\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		case 2:
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<img src=\"/images/dice-two.svg\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		case 3:
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<img src=\"/images/dice-three.svg\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		case 4:
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<img src=\"/images/dice-four.svg\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		case 5:
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<img src=\"/images/dice-five.svg\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		case 6:
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<img src=\"/images/dice-six.svg\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></label>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func DiceSelectionCurrentCount(num int) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var17 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var17 == nil {
			templ_7745c5c3_Var17 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"\n\t\tcounter pointer-events-none rounded-lg absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 font-mono\n\t\tw-12 h-12 z-10 hidden items-center justify-center text-ctp-crust\n\t\tgroup-hover:bg-white\n\t\tgroup-hover:flex\n\t\ttext-xl\n\t\tgroup-[:last-of-type:has(input:checked)]:text-2xl\n\t\tgroup-has-[input:checked]:group-has-[+_label_input:not(:checked)]:text-2xl\n\t\tgroup-[:last-of-type:has(input:checked)]:flex\n\t\tgroup-has-[input:checked]:group-has-[+_label_input:not(:checked)]:flex\n\t\t\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var18 string
		templ_7745c5c3_Var18, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.Itoa(num))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/main.templ`, Line: 272, Col: 21}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var18))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("x</div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func DiceSheen() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var19 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var19 == nil {
			templ_7745c5c3_Var19 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"absolute inset-0.5 rounded overflow-hidden\"><div class=\"sheen pointer-events-none w-full h-[150%] absolute left-0 -top-3 opacity-0 z-10\"><div class=\"\n\t\t\t\t\tbg-gradient-to-r from-transparent to-transparent\n\t\t\t\t\tgroup-data-[rollValue=&#39;1&#39;]:via-ctp-subtext0\n\t\t\t\t\tgroup-data-[rollValue=&#39;2&#39;]:via-ctp-text\n\t\t\t\t\tgroup-data-[rollValue=&#39;3&#39;]:via-ctp-green\n\t\t\t\t\tgroup-data-[rollValue=&#39;4&#39;]:via-ctp-blue\n\t\t\t\t\tgroup-data-[rollValue=&#39;5&#39;]:via-ctp-mauve\n\t\t\t\t\tgroup-data-[rollValue=&#39;6&#39;]:via-ctp-yellow\n\t\t\t\t\tgroup-data-[rollValue=&#39;1&#39;]:w-1\n\t\t\t\t\tgroup-data-[rollValue=&#39;2&#39;]:w-2\n\t\t\t\t\tgroup-data-[rollValue=&#39;3&#39;]:w-4\n\t\t\t\t\tgroup-data-[rollValue=&#39;4&#39;]:w-8\n\t\t\t\t\tgroup-data-[rollValue=&#39;5&#39;]:w-12\n\t\t\t\t\tgroup-data-[rollValue=&#39;6&#39;]:w-16\n\t\t\t\t\th-full absolute left-1/2 -translate-x-1/2 -rotate-12\"></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
