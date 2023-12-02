Adding a new year command 
`cobra-cli add-command "YEAR"`

if the YEAR is a number 2015, 2016 ect. then you will need to update the file name and the command name because Golang doesn't like then numbers 

Then create a folder in the `cmd/` dir for that year, this is where the problems will go each day. 
Add the new directory to the `main.go` file like this `_ "github.com/rjprice04/advent_of_code/cmd/NEW_YEAR"`. 


Adding new day command 
`cobra-cli add-command dayXX -p YEAR`

then move it to the corret year dirctory `mv cmd/dayXX.go cmd/YEAR/dayXX.go


