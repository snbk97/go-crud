## CONFIGURE
set target build in make file `GOOS` and `GOARCH`

either run `go mod tidy` or `make ready`

## RUN
To run the project run for developement purpose
command:  `make watch`

## BUILD
To run the project run for developement purpose
command:  `make build`

## FOLDER STRUCTURE
- builds should be in `/bin` folder
- backend logic should be in `/internal`
- common code, configs, constants should be in `/common`
- front end app should be in `/web`