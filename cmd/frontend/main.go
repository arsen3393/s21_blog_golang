package main

import (
	"Go_Day06/config"
	"fmt"
	fileserver "github.com/ffss92/fileserver"
	"log"
	"net/http"
	"os"
)

func main() {
	cfg := config.GetConfig()

	static := os.DirFS(cfg.ClientConf.StaticFilesPath)
	stylesPath := fmt.Sprintf("%s/styles/", cfg.ClientConf.StaticFilesPath)
	staticStyles := os.DirFS(stylesPath)

	scriptsPath := fmt.Sprintf("%s/scripts/", cfg.ClientConf.StaticFilesPath)
	staticScripts := os.DirFS(scriptsPath)

	mux := http.NewServeMux()
	mux.Handle("/", fileserver.ServeSPA(static, "index.html"))
	mux.Handle("/admin", fileserver.ServeSPA(static, "admin.html"))
	mux.Handle("/loginscript", fileserver.ServeSPA(staticScripts, "login.js"))
	mux.Handle("/styles/", http.StripPrefix("/styles/", fileserver.ServeFS(staticStyles)))
	mux.Handle("/scripts/", http.StripPrefix("/scripts/", fileserver.ServeFS(staticScripts)))
	mux.Handle("/logo", fileserver.ServeSPA(static, "amazing_logo.png"))

	port := fmt.Sprintf(":%d", cfg.ClientConf.Port)
	fmt.Printf("Listening and serving HTTP on %s\n", port)
	err := http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatal(err)
	}
}
