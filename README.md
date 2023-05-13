# mazeCli
A cli for generating maze image files.

A maze generator which works by first generating<br />
a random path through the plane and <br />
then adds different branches to the maze<br />
which themselves also have branches.<br />
It saves the maze as a txt file.<br />

Usage:<br />
&nbsp;&nbsp;mazeCli [command]<br />
<br />
Available Commands:<br />
&nbsp;&nbsp;completion - Generate the autocompletion script for the specified shell<br />
&nbsp;&nbsp;gen -        Generates the maze and outputs it into a txt file.<br />
&nbsp;&nbsp;help -       Help about any command<br />
<br />
Flags:<br />
&nbsp;&nbsp;-h, --help -  help for mazeCli<br />
&nbsp;&nbsp;<br />
 # Generate a maze
 <br />
It generates the maze for the given arguments and saves it<br />
to the specified output file.<br />
<br />
Usage:<br />
&nbsp;&nbsp;mazeCli gen [flags]<br />
<br />
Flags:<br />
&nbsp;&nbsp;-h, --help -        help for gen<br />
&nbsp;&nbsp;-o, --out string -  Sets the output file. (default "maze.txt")<br />
&nbsp;&nbsp;-s, --seed int -    Sets the seed of the maze. (default time in ns)<br />
&nbsp;&nbsp;-z, --size int -    Sets the size of the maze. (default 10)<br />
