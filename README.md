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
   



*******************************************************************************************************************************
T.O.C (Table of Contents for this Go web server/service)

Dockerfile :
   This Dockerfile automatically pulls the correct dependencies for the Go server (any third-party libs) for image creation (images are stored in a Docker Registry) and then turned into running instances on any target OS (containers). 
   
   This Dockerfile automatically pulls in the latest Golang runtime (1.10+), the Gorilla Toolkit (Go) mux router (http://www.gorillatoolkit.org/pkg/mux). This Dockerfile can get extended to add additionaly third-party dependencies by issueing added
   dependency requirements to the internally selected OS image's package manager through the Docker RUN. The OS image for this Go server
   is the Alpine Linux (lightweight Linux distro) suited for production environments and is lighter than Ubuntu images.
   
   The goal is to start a container (self-reliant running image of the Go server independent of any guest OS it runs on) and to acheive this goal, from this Dockerfile an image has to get created.
   
   To create a Docker image for this Go server and prepare for its deployment you issue a 'docker build -t <image-name>:[tag] .
   
   The -t flag or for full unabbreviated form --tag list (docker build --tag list <image-name>:[tag] .)
   associates the 'name:tag' pairing where in conventional practice the 'name' part of the pairing is 
   'organization-name/appservice-name:<version-tag>. **IMPORTANT ** : The '.' period at the end is declaring the image be build from the context (root project directory from where the Dockerfile lives and all the contents at or under this context level). 
   
   To build the Go server for this project follow this :
   
   sudo docker build -t circadence/go-poc-server:1.0 . 
   
   This will take anywhere from a few minutes to longer dependending on the size of the image. The output of the layering/building process is visible to the command shell as you watch it build. A succesful message results after image is built.
   Dockerfiles work from existing images you layer upon previous layers (Dockerfile RUN command on a line is one layer, followed by another Dockerfile RUN is another) so the image sizes can be large (> 500MB or a close to a GB) for even small web server/web service applications. Dockerfiles can include execution instructions to slim down the image size. To slim down a Docker image you can Google Docker image size reduction for best practices. This current Dockefile references the Alpine Linux distro. Adding dev packages to this Dockerfile use the Alpine Linux distro package manager 'apk' through RUN apk add --update or --no-cache <package-dependency> where if Ubuntu/Debian OS Linux distro was chose the equivalent Dockerfile instructions would be RUN apt-get intall <package-dependency>.
   
  
  After the Docker image has been built you run the following to see the image in the local Docker image repository. In Docker, once installed on your host OS (on your dev machine) the Docker engine has a Docker CLI (the client process), a local Docker registry where all your Docker images are listed, and a Docker daemon (a server to execute the Docker CLI instructions you pass it).  
  
  To see the built image you cannot just issue a ls -l or ls -al command (if on Linux which is preferred). You have to issue the following to see the named image you built with the following :
  
  sudo docker image ls 
  
  
  This will show a column headered row with the three key fields you should know.
  
  1. REPOSITORY
  2. TAG
  3. IMAGE ID 
  
  The successfully built image will be shown as follows :
  
  REPOSITORY                        TAG         IMAGE ID 
  
  circadence/go-poc-server          1.0         <generated-hash-val>
   
   
  
  To run the docker image (turn into a Docker container)
  
  sudo docker run -p 8080:8080 --name=poc-server circadence/go-poc-server:1.0
  
  Here the --name flag gives your running Docker container a friendly Docker name when you look for running containers with the 
  sudo docker ps -a command.
  
  The value passed to the --name flag is arbitrary but chose a intuitive name.
  
  
  For this PoC the requirement is to copy up the locally built image to Cloud Providers (Azure or AWS). 
  
 One method is to publish your built image (not the container) to a Docker Registry (Arficatory, Dockerhub, ...) but an alternate
 route is to essentially tar up the Docker image, scp the tarred image to a Cloud server instance, ssh into the Cloud server and untar the image and the run the loaded image on that server instance.
 
 Steps to do this :
 
 1. sudo docker save -o poc-server.tar circadence/go-poc-server:1.0
 2. scp -i /path-to/<pair.ppk> poc-server.tar cloud-host@public-cloud-ip:~  (where cloud-host is ec2-user or azure)
 3. ssh -i /path-to/<pair.ppk> cloud-host@public-cloud-ip
 4. sudo docker load -i /path-to/poc-server.tar
 5. sudo docker run -p 8080:8080 --name=poc-server circadence/go-poc-server:1.0
 
 After step #5 the go server will be running inside the running container. You start up a browser and point it to 
 http://public-cloud-ip:8080/static/index.html
 
 
 The app is running correctly.
 
 
 To shutdown the server issue the following:
 
 CTRL-C as this gracefully shuts down the server.
 
 Follow this with:
 
 sudo docker ps <container-name> where <container-name> is the friendly name you issued during the sudo docker run command in this case
 poc-server.
 
 
 
 ** T.O.C Continued **
 
 data.go :
 
   This is the in-memory data structure (an Exercise of type Golang struct) and a declared Golang slice (dynamic array of Exercises) from which the repo.go Repo* functions reference on behalf of the server_handler.go web service request handlers. This approach separates the underlying storage from the service functions. If the direction changes to point to a SQL database or NoSQL database the repo.go functions do the lookup logic into that chosen production database and only the repo.go source file requires recompilation.
   
 server-logger.go :
 
   This serves as a logger of which request was routed to and which request handler is currently processing.
   
 server.go :
 
   This the main executable entrypoint to the go server. It declares the port on which the server will run and has QoS event-handling
   functionality to handle SIGEVT (CTRL-C) shutdown events to gracefully shutdown the server. 
   
   **IMPORTANT ** As of now the current server DOES NOT include any kind of onion layering/wrapping middleware (auth middleware with JWT) or any other kind of pre-functional handler wrapping which is trivial to add with Golang web request middleware packages such as Negroni.
   
  server_handlers.go :
  
   This is where the RESTful web servivce request handlers are defined. To service a new web request you add code in two places.
   The first is routes.go (where the declared routes are configured) and here in server_handlers.go to define the Golang functions to fulfill the routes defined in routes.go. All of the controller or handler funcs here accept a http.ResponseWriter and http.Request parameters to satisfy the polymorphic contract of being web service handling functions.
   
   router.go :
   
      This is where the core of the service lives. It delcares the route handling logic and indirectly dispatches all requests to the server_handlers.go handler funcs.
   
   routes.go :
   
      This is where the routes (path routing) is declared. To add a new service request handler you add a new route following the route
      structure defined in here and follow that with the server_handlers.go file you have to add the new service function handler to match the route declared here.
      
  error.go :
  
   This is JSON encodable struct that defines the error code for the service handlers.
   
   
   
 vendor/  :
 
 The vendor directory is a package-version controlled directory automatically generated with the dep init command that includes a manifest file (Gopkg.toml) and a Gopkg.lock file. To add a new dependency you issue a dep ensure add <go-package-repo/pkg>
 This command syncs up the vendor directory with the new dependency and the Gopkg.toml file).
 
 
 To git pull down this entire Git source code you follow your git clone <https://repo-name/app> with a dep ensure command as long as you have Go dep installed on your dev machine.
 
 
 
 
 
 
 
 
 
 
 
 
 
  
   
   
