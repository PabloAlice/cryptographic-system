# cryptographic-system
Academical cryptographic system

### PRESENT ALGORITHM
light cryptographic system


### STEPS

 - first at all: [install golang](https://golang.org/doc/install)

 - go get -u github.com/kardianos/govendor

 - clone this repo inside ${GOPATH}/src/github/yourUserName/
 this is because we are using https://github.com/kardianos/govendor

 - govendor sync, (if you have dependencies troubles, check govendor docs)

 - for the next step, you should have node.js lts installed.
 - at root project: cd client && npm i && npm run build
 - go back to the root project.
 - "gin main.go" should start the app in localhost:3000
 in case, you cannot run it with gin command (check https://github.com/codegangsta/gin)
otherwise you can use go run main.go
