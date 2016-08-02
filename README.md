
# go-graph - Graphs from shell STDIN.

Some parts and ideas borrowed from rtop-vis!

## what
Print nice graphs from bash STDIN.
[example.png](https://s32.postimg.org/5x7eocy9h/example.png)

## usage
For example:

`while true ; do awk 'BEGIN{srand();print int(rand()\*(63000-2000))+2000 }' ;
sleep 1 ; done | ./main --port 7071 --title Testing`

## contribute

Pull requests welcome.

## known bugs
See [issues page](https://github.com/Adar/go-graph/issues)

## License
MIT license, see LICENSE.txt for details.
