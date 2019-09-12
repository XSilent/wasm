# wasm
Example code for Go, targeting Wasm

# install
go get -u github.com/shurcooL/goexec

# export path in .bash_profile
export PATH=$PATH:/usr/local/go/bin:~/go/bin/

# make changes affective in .bash_profile without logoff/logon
source ~/.bash_profile

# serve files with
goexec "http.ListenAndServe(\`:8080\`, http.FileServer(http.Dir(\`.\`)))"
