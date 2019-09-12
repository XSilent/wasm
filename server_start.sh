echo Listening on 8080
goexec "http.ListenAndServe(\`:8080\`, http.FileServer(http.Dir(\`.\`)))"
