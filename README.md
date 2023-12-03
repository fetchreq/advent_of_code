# Advent of Code

## Getting stared
First intall [Go](https://go.dev/doc/install), if you don't have it already.
Then create a `.env` file `cp .env.example .env` and add your session token from the AoC website as the value. You should be able to see it in the network requests in dev tools

Running a days problems 

`go run main.go YEAR dayXX`

this will download the file from AoC for that year and day if you don't already have it loacally in the `input/YEAR/` directory

## Adding New Commands
To create new commands you will need the [cobra-cli tool](https://github.com/spf13/cobra-cli/blob/main/README.md)
`go install github.com/spf13/cobra-cli@latest`

### Add Year Commands

1. Adding a new year command 
`cobra-cli add-command "YEAR"`

2. Create a folder in the `cmd/` dir for that year 
`mkdir cmd/YEAR` 
    * this is where the problems will go each day. 
3. Add the new directory to the `main.go` file 
    * `_ "github.com/rjprice04/advent_of_code/cmd/NEW_YEAR"`. 
4. Create a folder in the input diretory `mkdir input/YEAR` 
    * The year will be used to find the file in the `ReadInput` function

### Add Day Commands

1. Adding new day command 
`cobra-cli add-command dayXX -p YEAR`

2. Move it to the correct year dirctory 
`mv cmd/dayXX.go cmd/YEAR/dayXX.go


