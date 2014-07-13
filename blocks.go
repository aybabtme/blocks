// Package blocks returns unicode blocks matching 2x2 boolean matrices.
// This is useful to draw blocks on a terminal output, as it produces
// twice the resolution of using simple █ blocks.
package blocks

// Hex | D C B A | Block
// ----|---------|------
//  0  | 0 0 0 0 | ' '
//  1  | 0 0 0 1 | '▘'
//  2  | 0 0 1 0 | '▝'
//  3  | 0 0 1 1 | '▀'
//  4  | 0 1 0 0 | '▖'
//  5  | 0 1 0 1 | '▌'
//  6  | 0 1 1 0 | '▞'
//  7  | 0 1 1 1 | '▛'
//  8  | 1 0 0 0 | '▗'
//  9  | 1 0 0 1 | '▚'
//  a  | 1 0 1 0 | '▐'
//  b  | 1 0 1 1 | '▜'
//  c  | 1 1 0 0 | '▄'
//  d  | 1 1 0 1 | '▙'
//  e  | 1 1 1 0 | '▟'
//  f  | 1 1 1 1 | '█'
var blocks = [16]rune{
	' ',
	'▘',
	'▝',
	'▀',
	'▖',
	'▌',
	'▞',
	'▛',
	'▗',
	'▚',
	'▐',
	'▜',
	'▄',
	'▙',
	'▟',
	'█',
}

// Block returns a unicode character that
// represents a 2 by 2 grid, where the parts
// of the grid are set according to this
// chart:
//
//     -----
//     |A|B|
//     -----
//     |C|D|
//     -----
//
// For instance, b=true, c=true gives:
//
//      ▞
//
func Block(a, b, c, d bool) rune {
	i := 0x00
	if a {
		i |= 0x01
	}
	if b {
		i |= 0x02
	}
	if c {
		i |= 0x04
	}
	if d {
		i |= 0x08
	}
	return blocks[i]
}

// Blocks transforms an m by n matrix of booleans to a m/2 by n/2 matrix
// of runes, where true entries are represented by a filled block, false
// entries by nothing.
func Blocks(in [][]bool) [][]rune {
	var (
		h = len(in)
		w = len(in[0])

		oddH = len(in)%2 != 0
		oddW = len(in[0])%2 != 0
	)
	if oddH {
		h++
	}
	if oddW {
		w++
	}
	out := make([][]rune, h/2)
	for i := range out {
		out[i] = make([]rune, w/2)
	}

	// read the input 2x2
	for i := 0; i < h; i += 2 {
		for j := 0; j < w; j += 2 {
			var a, b, c, d bool

			a = in[i][j]
			if !oddW || j < len(in[0])-1 {
				b = in[i][j+1]
			}
			if !oddH || i < len(in)-1 {
				c = in[i+1][j]
			}

			if (oddH && i+1 == len(in)) || (oddW && j+1 == len(in[0])) {
				// could demorgan this thing, too tired to think about it
			} else {
				d = in[i+1][j+1]
			}

			out[i/2][j/2] = Block(a, b, c, d)
		}
	}
	return out
}
