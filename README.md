# mazeCli
A cli for generating maze image files.

A maze generator which works by first generating
a random path through the plane and 
then adds different branches to the maze
which themselves also have branches.
It saves the maze as a png image.

Usage:
  mazeCli [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  gen         Generates the maze and outputs it into a png file.
  help        Help about any command

Flags:
  -h, --help   help for mazeCli
  
 # Generate a maze
 
It generates the maze for the given arguments and saves it
to the specified output file.

Usage:
  mazeCli gen [flags]

Flags:
  -h, --help         help for gen
  -o, --out string   Sets the output file. (default "maze.png")
  -s, --seed int     Sets the seed of the maze. (default time in ns)
  -z, --size int     Sets the size of the maze. (default 10)
