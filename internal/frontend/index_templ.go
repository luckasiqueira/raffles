// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.648
package frontend

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Index() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\"><head><meta charset=\"UTF-8\"><title>Title</title><script src=\"https://cdn.tailwindcss.com\"></script></head><body><div class=\"flex flex-col min-h-[100dvh]\"><header class=\"px-4 lg:px-6 h-14 flex items-center\"><a class=\"flex items-center justify-center\" href=\"#\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\" class=\"h-6 w-6\"><rect width=\"18\" height=\"18\" x=\"3\" y=\"3\" rx=\"2\" ry=\"2\"></rect> <path d=\"M12 12h.01\"></path></svg> <span class=\"sr-only\">Rifa de Números</span></a><nav class=\"ml-auto flex gap-4 sm:gap-6\"><a class=\"text-sm font-medium hover:underline underline-offset-4\" href=\"#\">Números\r</a> <a class=\"text-sm font-medium hover:underline underline-offset-4\" href=\"#\">Sobre\r</a> <a class=\"text-sm font-medium hover:underline underline-offset-4\" href=\"#\">Contato\r</a></nav></header><main class=\"flex-1\"><section class=\"w-full py-6 sm:py-12 md:py-24 lg:py-32 xl:py-48\"><div class=\"container px-4 md:px-6\"><div class=\"grid gap-6 lg:grid-cols-[1fr_400px] lg:gap-12 xl:grid-cols-[1fr_600px]\"><img src=\"/assets/img/perfume.png\" width=\"550\" height=\"550\" alt=\"Hero\" class=\"mx-auto aspect-video overflow-hidden rounded-xl object-cover sm:w-full lg:order-last lg:aspect-square\"><div class=\"flex flex-col justify-center space-y-4\"><div class=\"space-y-2\"><h1 class=\"text-3xl font-bold tracking-tighter sm:text-5xl xl:text-6xl/none\">Rifa de Números</h1><p class=\"max-w-[600px] text-muted-foreground md:text-xl\">Participe da nossa rifa e concorra ao incrível prêmio de um XXXX XXX avaliado no valor de R$XXX,XX.\r</p></div><div class=\"flex flex-col gap-2 min-[400px]:flex-row\"><a class=\"inline-flex h-10 items-center justify-center rounded-md bg-primary px-8 text-sm font-medium text-primary-foreground shadow transition-colors hover:bg-primary/90 focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring disabled:pointer-events-none disabled:opacity-50\" href=\"#\">Comprar Números\r</a> <a class=\"inline-flex h-10 items-center justify-center rounded-md border border-input bg-background px-8 text-sm font-medium shadow-sm transition-colors hover:bg-accent hover:text-accent-foreground focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring disabled:pointer-events-none disabled:opacity-50\" href=\"#\">Sobre a Rifa\r</a></div></div></div></div></section><section class=\"w-full py-12 md:py-24 lg:py-32 bg-muted\"><div class=\"container px-4 md:px-6\"><div class=\"flex flex-col items-center justify-center space-y-4 text-center\"><div class=\"space-y-2\"><div class=\"inline-block rounded-lg bg-muted px-3 py-1 text-sm\">Números Disponíveis</div><h2 class=\"text-3xl font-bold tracking-tighter sm:text-5xl\">Confira os números</h2><p class=\"max-w-[900px] text-muted-foreground md:text-xl/relaxed lg:text-base/relaxed xl:text-xl/relaxed\">Cada número custa apenas R$5,00. Não perca a sua chance de ganhar!\r</p></div></div><div class=\"mx-auto grid max-w-5xl items-center gap-6 py-12 lg:grid-cols-5 lg:gap-4\"><div class=\"bg-background border border-input rounded-full w-12 h-12 flex items-center justify-center text-lg font-medium\">1\r</div><div class=\"bg-background border border-input rounded-full w-12 h-12 flex items-center justify-center text-lg font-medium\">2\r</div><div class=\"bg-background border border-input rounded-full w-12 h-12 flex items-center justify-center text-lg font-medium\">3\r</div><div class=\"bg-background border border-input rounded-full w-12 h-12 flex items-center justify-center text-lg font-medium\">4\r</div><div class=\"bg-background border border-input rounded-full w-12 h-12 flex items-center justify-center text-lg font-medium\">5\r</div><div class=\"bg-background border border-input rounded-full w-12 h-12 flex items-center justify-center text-lg font-medium\">6\r</div><div class=\"bg-background border border-input rounded-full w-12 h-12 flex items-center justify-center text-lg font-medium\">7\r</div><div class=\"bg-background border border-input rounded-full w-12 h-12 flex items-center justify-center text-lg font-medium\">8\r</div><div class=\"bg-background border border-input rounded-full w-12 h-12 flex items-center justify-center text-lg font-medium\">9\r</div><div class=\"bg-background border border-input rounded-full w-12 h-12 flex items-center justify-center text-lg font-medium\">10\r</div><div class=\"bg-background border border-input rounded-full w-12 h-12 flex items-center justify-center text-lg font-medium\">11\r</div><div class=\"bg-background border border-input rounded-full w-12 h-12 flex items-center justify-center text-lg font-medium\">12\r</div><div class=\"bg-background border border-input rounded-full w-12 h-12 flex items-center justify-center text-lg font-medium\">13\r</div><div class=\"bg-background border border-input rounded-full w-12 h-12 flex items-center justify-center text-lg font-medium\">14\r</div><div class=\"bg-background border border-input rounded-full w-12 h-12 flex items-center justify-center text-lg font-medium\">15\r</div><div class=\"bg-background border border-input rounded-full w-12 h-12 flex items-center justify-center text-lg font-medium\">16\r</div><div class=\"bg-background border border-input rounded-full w-12 h-12 flex items-center justify-center text-lg font-medium\">17\r</div><div class=\"bg-background border border-input rounded-full w-12 h-12 flex items-center justify-center text-lg font-medium\">18\r</div><div class=\"bg-background border border-input rounded-full w-12 h-12 flex items-center justify-center text-lg font-medium\">19\r</div><div class=\"bg-background border border-input rounded-full w-12 h-12 flex items-center justify-center text-lg font-medium\">20\r</div><div class=\"bg-background border border-input rounded-full w-12 h-12 flex items-center justify-center text-lg font-medium\">21\r</div><div class=\"bg-background border border-input rounded-full w-12 h-12 flex items-center justify-center text-lg font-medium\">22\r</div><div class=\"bg-background border border-input rounded-full w-12 h-12 flex items-center justify-center text-lg font-medium\">23\r</div><div class=\"bg-background border border-input rounded-full w-12 h-12 flex items-center justify-center text-lg font-medium\">24\r</div><div class=\"bg-background border border-input rounded-full w-12 h-12 flex items-center justify-center text-lg font-medium\">25\r</div><div class=\"bg-background border border-input rounded-full w-12 h-12 flex items-center justify-center text-lg font-medium\">26\r</div><div class=\"bg-background border border-input rounded-full w-12 h-12 flex items-center justify-center text-lg font-medium\">27\r</div><div class=\"bg-background border border-input rounded-full w-12 h-12 flex items-center justify-center text-lg font-medium\">28\r</div><div class=\"bg-background border border-input rounded-full w-12 h-12 flex items-center justify-center text-lg font-medium\">29\r</div><div class=\"bg-background border border-input rounded-full w-12 h-12 flex items-center justify-center text-lg font-medium\">30\r</div><div class=\"bg-background border border-input rounded-full w-12 h-12 flex items-center justify-center text-lg font-medium\">31\r</div><div class=\"bg-background border border-input rounded-full w-12 h-12 flex items-center justify-center text-lg font-medium\">32\r</div><div class=\"bg-background border border-input rounded-full w-12 h-12 flex items-center justify-center text-lg font-medium\">33\r</div><div class=\"bg-background border border-input rounded-full w-12 h-12 flex items-center justify-center text-lg font-medium\">34\r</div><div class=\"bg-background border border-input rounded-full w-12 h-12 flex items-center justify-center text-lg font-medium\">35\r</div><div class=\"bg-background border border-input rounded-full w-12 h-12 flex items-center justify-center text-lg font-medium\">36\r</div><div class=\"bg-background border border-input rounded-full w-12 h-12 flex items-center justify-center text-lg font-medium\">37\r</div><div class=\"bg-background border border-input rounded-full w-12 h-12 flex items-center justify-center text-lg font-medium\">38\r</div><div class=\"bg-background border border-input rounded-full w-12 h-12 flex items-center justify-center text-lg font-medium\">39\r</div><div class=\"bg-background border border-input rounded-full w-12 h-12 flex items-center justify-center text-lg font-medium\">40\r</div><div class=\"bg-background border border-input rounded-full w-12 h-12 flex items-center justify-center text-lg font-medium\">41\r</div><div class=\"bg-background border border-input rounded-full w-12 h-12 flex items-center justify-center text-lg font-medium\">42\r</div><div class=\"bg-background border border-input rounded-full w-12 h-12 flex items-center justify-center text-lg font-medium\">43\r</div><div class=\"bg-background border border-input rounded-full w-12 h-12 flex items-center justify-center text-lg font-medium\">44\r</div><div class=\"bg-background border border-input rounded-full w-12 h-12 flex items-center justify-center text-lg font-medium\">45\r</div><div class=\"bg-background border border-input rounded-full w-12 h-12 flex items-center justify-center text-lg font-medium\">46\r</div><div class=\"bg-background border border-input rounded-full w-12 h-12 flex items-center justify-center text-lg font-medium\">47\r</div><div class=\"bg-background border border-input rounded-full w-12 h-12 flex items-center justify-center text-lg font-medium\">48\r</div><div class=\"bg-background border border-input rounded-full w-12 h-12 flex items-center justify-center text-lg font-medium\">49\r</div><div class=\"bg-background border border-input rounded-full w-12 h-12 flex items-center justify-center text-lg font-medium\">50\r</div></div></div></section><section class=\"w-full py-12 md:py-24 lg:py-32\"><div class=\"container px-4 md:px-6\"><div class=\"flex flex-col items-center justify-center space-y-4 text-center\"><div class=\"space-y-2\"><div class=\"inline-block rounded-lg bg-muted px-3 py-1 text-sm\">Prêmio da Rifa</div><h2 class=\"text-3xl font-bold tracking-tighter sm:text-5xl\">Ganhe um Carro Zero Km</h2><p class=\"max-w-[900px] text-muted-foreground md:text-xl/relaxed lg:text-base/relaxed xl:text-xl/relaxed\">O grande prêmio da nossa rifa é um carro zero quilômetro. Não perca a chance de levar esse incrível\r prêmio para casa!\r</p></div><img src=\"/placeholder.svg\" width=\"550\" height=\"310\" alt=\"Carro Zero Km\" class=\"mx-auto aspect-video overflow-hidden rounded-xl object-cover object-center sm:w-full\"></div></div></section><section class=\"w-full py-12 md:py-24 lg:py-32 border-t\"><div class=\"container grid items-center justify-center gap-4 px-4 text-center md:px-6\"><div class=\"space-y-3\"><h2 class=\"text-3xl font-bold tracking-tighter md:text-4xl/tight\">Participe Agora</h2><p class=\"mx-auto max-w-[600px] text-muted-foreground md:text-xl/relaxed lg:text-base/relaxed xl:text-xl/relaxed\">Não perca a sua chance de concorrer a prêmios incríveis. Compre seus números agora mesmo!\r</p></div><div class=\"mx-auto w-full max-w-sm space-y-2\"><form class=\"flex gap-2\"><input class=\"flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50 max-w-lg flex-1\" type=\"text\" placeholder=\"Digite seu nome\"> <button class=\"inline-flex items-center justify-center whitespace-nowrap rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 bg-primary text-primary-foreground hover:bg-primary/90 h-10 px-4 py-2\" type=\"submit\">Comprar Números\r</button></form><p class=\"text-xs text-muted-foreground\">Ao comprar seus números, você concorda com os")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(" ")
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal\frontend\index.templ`, Line: 289, Col: 75}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" <a class=\"underline underline-offset-2\" href=\"#\">Termos e Condições\r</a></p></div></div></section></main><footer class=\"flex flex-col gap-2 sm:flex-row py-6 w-full shrink-0 items-center px-4 md:px-6 border-t\"><p class=\"text-xs text-muted-foreground\">© 2024 Rifa de Números. Todos os direitos reservados.</p><nav class=\"sm:ml-auto flex gap-4 sm:gap-6\"><a class=\"text-xs hover:underline underline-offset-4\" href=\"#\">Termos de Serviço\r</a> <a class=\"text-xs hover:underline underline-offset-4\" href=\"#\">Privacidade\r</a></nav></footer></div></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
