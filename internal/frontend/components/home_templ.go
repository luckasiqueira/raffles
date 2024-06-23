// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.648
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "raffles/internal/frontend/components/parts"

func Home() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = parts.Header().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"grid md:grid-cols-[80px_1fr] min-h-screen w-full\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = parts.Menu().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex flex-col items-center justify-center py-12\"><div class=\"mx-auto w-full max-w-[800px] space-y-6\"><div class=\"space-y-2 text-center\"><h1 class=\"text-3xl font-bold\">Painel de gestão</h1><p class=\"text-gray-500 dark:text-gray-400\">Acompanhe aqui os detalhes da sua rifa.</p></div><div class=\"grid grid-cols-2 md:grid-cols-4 gap-4\"><div class=\"rounded-lg border bg-card text-card-foreground shadow-sm\" data-v0-t=\"card\"><div class=\"flex flex-col space-y-1.5 p-6\"><h3 class=\"whitespace-nowrap text-xl font-semibold leading-none tracking-tight\">Participantes</h3></div><div class=\"p-6\"><div class=\"text-6xl font-bold\">45</div></div></div><div class=\"rounded-lg border bg-card text-card-foreground shadow-sm\" data-v0-t=\"card\"><div class=\"flex flex-col space-y-1.5 p-6\"><h3 class=\"whitespace-nowrap text-xl font-semibold leading-none tracking-tight\">Disponíveis</h3></div><div class=\"p-6\"><div class=\"text-6xl font-bold\">36</div></div></div><div class=\"rounded-lg border bg-card text-card-foreground shadow-sm\" data-v0-t=\"card\"><div class=\"flex flex-col space-y-1.5 p-6\"><h3 class=\"whitespace-nowrap text-xl font-semibold leading-none tracking-tight\">Comprados</h3></div><div class=\"p-6\"><div class=\"text-6xl font-bold\">92</div></div></div><div class=\"rounded-lg border bg-card text-card-foreground shadow-sm\" data-v0-t=\"card\"><div class=\"flex flex-col space-y-1.5 p-6\"><h3 class=\"whitespace-nowrap text-xl font-semibold leading-none tracking-tight\">Finalizado?</h3></div><div class=\"p-6\"><div class=\"text-6xl font-bold\">NÃO</div></div></div></div><div class=\"space-y-4\"><div class=\"border rounded-lg p-4 text-center\"><p class=\"text-2xl font-bold\">Sorteio não realizado</p><p class=\"text-4xl font-bold mt-4\">00:00:00s </p><div class=\"flex justify-center mt-4\"><button class=\"inline-flex items-center justify-center whitespace-nowrap rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 bg-primary text-primary-foreground hover:bg-primary/90 h-10 px-4 py-2\">contagem regressiva para o sorteio\r</button></div></div></div></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = parts.Footer().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
