## once i've written the main func, i have to init a go module:
    + that would be:
        $ go mod init github.com/reponame

## to hot reload and run the app while development, I'll use cosmtrek/air ("air"). Done like this:
    + import it: 
        $ go get github.com/cosmtrek/air

    + run the app using:
        $ go run github.com/cosmtrek/air
    
    this will start up the webserver and watch all the files in proj dir.

## once we've run air, we should see somehting on localhost:3000

