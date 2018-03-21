Prerequisite developer steps are to install Golang.

1. INSTALL Golang :
   https://golang.org/ and https://golang.org/dl/ and select the proper target OS for install
 
2. Configure Golang $GOROOT and $GOPATH environment variables
   In order to develop Go code in the correct environment after you install Golang
   
   Setting the GOROOT first:
   
   For Linux users:
   The standard is to download the .tar.gz file tar -xvf go.x.x.tar.gz and sudo mv the extracted go directory to /usr/local
   Next edit the .bashrc or .profile script and EXPORT GOROOT /usr/local/go
   Next in the .bashrc or .profile script add the $GOROOT/bin to the $PATH
   Next source the file and enter :
   
   > go version
   
   You should see version info of the go runtime.
   
   
   
   For Windows users:
   The standard is to download and install to root of the C:\ drive (C:\Go). 
   Next add a new envriorment variable for GOROOT and set is value to C:\Go
   Next add %GOROOT\bin% to the %PATH% 
   Next open new Windows Shell and enter :
   
   > go version
   
    You should see version info of the go runtime.
   
   
   For detailed info on installation visit https://golang.org/doc/install
   
   
   
   Configuring GOPATH
   
   There is no forced prerequisite to the path convention as long as you can intuitively know where you're Go code will reside.
   A unofficial standard location is to create a directory path ending with x/y/z/go or x/y/z/gowork where x/y/z are a path to your $HOME
   or %HOME% directory where you will have project code in. Next after you create your 'go' or 'gowork'directory you will create three directories
   all at the exact tree level : mkdir pkg, mkdir src and mkdir pkg.
   
   > go/
         pkg/
         src/
         bin/
   
   
 You CD into src/ and this is where you will create a directory where your root file structure lives.
 The typical convention is to use organization/product : circadence/trainer and on github.com the git project will look like
 github.com/circadence/trainer.
 
 With you add packages for third-party dependencies. Since the proper way is to package control your vendored depedencies for reasons
 you don't want to pull down the code from github.com only to have reconstiute the missing dependency compile errors. The next stage is using
 the Go Dep (dep) CLI.


3. Golang has a package management control service like NodeJS NPM. There are a top two (Glide and Dep). 
   https://github.com/golang/dep
   
   Install Go dep (at a Linux or Windows Shell) :
   
   > go get -u github.com/golang/dep/cmd/dep
   
   
   Dep is now installed.
   
   After you Git clone an Golang project that has a vendored directory (a directory that was created with dep originally) you will see
   a *.lock file and .toml file where the latter is project depedency manifest file. The dep CLI adds your required dependency meta data in this
   manifest .toml file.
   
   After you git clone the Golang project you just enter :
   
   > dep ensure
   
   This will automatically pull in the third party dependency packages into your own copy of the vendored directory and after a go build (Go compile command)
   and Go install (to distributd the compiled resources (.a static libs and .exe) into your $GOPATH/pkg and $GOPATH/bin directories. 
   
   ** IMPORTANT **
   
   Git uses a go get <repo>/project CLI command to pull down third party packages it puts into the global $GOPATH/pkg directory however for 
   keeping your project dependencies insulated to your project you have to use the Go dep service to add any future requiring depedencies as it will
   keep its vendored directory and manifest in sync.
   
   > $ dep ensure -add github.com/pkg/errors  # Add the third-party github.com/pkg/errors to our vendored directory and sync the manifest
   
   
   Go dep CLI details https://golang.github.io/dep/ and https://golang.github.io/dep/docs/daily-dep.html
   
