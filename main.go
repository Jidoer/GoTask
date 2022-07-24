package main

import (
	"GoTask/worker"
	"fmt"
	"github.com/kataras/iris/v12"
	"time"
)

func main() {

	start := time.Now()
	/*
	for i :=1;i<10;i++ {
		w := worker.NewWokerPool(1000).Run()
		a := &worker.RunShell{"ls -l /"}
		w.PutJob(a)
	}
	cost := time.Since(start)
	fmt.Printf("cost=[%s]",cost)
*?

	/*
	for i:=0; i<100000000; i++  {
		a:= &worker.PrintNum{Num: i}
		w.PutJob(a)
	}
	//w.Stop()
	cost := time.Since(start)
	fmt.Printf("cost=[%s]",cost)
*/

	app := iris.New()
	tmpl := iris.Django("./Template", ".html")
	tmpl.Reload(false)
	tmpl.AddFunc("greet", func(s string) string {
		return "Greetings, " + s + "!"
	})
	app.RegisterView(tmpl)

	/************************************************************************
	 * ///////////////////////////////////////////////////////////////////////
	 *                JomeTaskManager V 1.0 Powered By FlyKO
	 * ///////////////////////////////////////////////////////////////////////
	 ************************************************************************/


	app.Get("/api/{run:string}", func(context iris.Context) {
		w := worker.NewWokerPool(1000).Run()
		a := &worker.Runmain{"cd / && ls -l"}
		//RunShell{"ls -l /"}
		w.PutJob(a)
		//w.Stop()


		cost := time.Since(start)
		fmt.Printf("cost=[%s]",cost)

		context.JSON(iris.Map{"success": true, "error_message": "target ok"})

	})




/*
	app.Get("/{template:string}", func(ctx iris.Context) {
		t := ctx.Params().Get("template")
		app.Logger().Info(t)
		ctx.HTML("!")
		for i :=1;i<10;i++ {
			w := worker.NewWokerPool(1000).Run()
			a := &worker.RunShell{"echo START && sleep 10"}
			w.PutJob(a)
		}
		cost := time.Since(start)
		fmt.Printf("cost=[%s]",cost)
	}
 */

	//app.Favicon("./static/favicon.ico")
	//INDEX Page STATIC
	app.HandleDir("/js", "./static/js")
	app.Run(iris.Addr(":6066"))

}
