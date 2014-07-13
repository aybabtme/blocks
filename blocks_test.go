package blocks_test

import (
	"fmt"
	"github.com/aybabtme/blocks"
)

func ExampleBlock() {
	top := blocks.Block(
		true, true,
		false, false,
	)
	bottom := blocks.Block(
		false, false,
		true, true,
	)
	fmt.Printf("the top %c and the bottom %c.\n", top, bottom)

	left := blocks.Block(
		true, false,
		true, false,
	)

	right := blocks.Block(
		false, true,
		false, true,
	)
	fmt.Printf("the left %c and the right %c.\n", left, right)

	topleft := blocks.Block(
		true, false,
		false, false,
	)
	topright := blocks.Block(
		false, true,
		false, false,
	)
	fmt.Printf("the topleft %c and the topright %c.\n", topleft, topright)

	botleft := blocks.Block(
		false, false,
		true, false,
	)
	botright := blocks.Block(
		false, false,
		false, true,
	)
	fmt.Printf("the botleft %c and the botright %c.\n", botleft, botright)

	// Output:
	// the top ▀ and the bottom ▄.
	// the left ▌ and the right ▐.
	// the topleft ▘ and the topright ▝.
	// the botleft ▖ and the botright ▗.
}

func ExampleBlocks() {

	bitmap := [][]bool{
		{true, true, true, true, true, true, true, true, true, true, true, true, true},
		{true, false, false, false, false, false, false, false, false, false, false, false, true},
		{true, false, false, true, true, false, false, true, true, false, false, false, true},
		{true, false, false, true, true, false, false, true, true, false, false, false, true},
		{true, false, false, false, false, false, false, false, false, false, false, false, true},
		{true, false, false, false, false, false, true, false, false, false, false, false, true},
		{true, false, false, false, false, true, true, false, false, false, false, false, true},
		{true, false, false, false, false, false, false, false, false, false, false, false, true},
		{true, false, true, false, false, false, false, false, false, true, false, false, true},
		{true, false, false, true, false, false, false, false, true, false, false, false, true},
		{true, false, false, false, true, true, true, true, false, false, false, false, true},
		{true, false, false, false, false, false, false, false, false, false, false, false, true},
		{true, true, true, true, true, true, true, true, true, true, true, true, true},
	}
	fmt.Println("Happy now?")
	for _, str := range blocks.Blocks(bitmap) {
		fmt.Printf("%s\n", string(str))
	}
	// Output:

	// Happy now?
	// ▛▀▀▀▀▀▌
	// ▌▐▌▐▌ ▌
	// ▌  ▖  ▌
	// ▌ ▝▘  ▌
	// ▌▚  ▞ ▌
	// ▌ ▀▀  ▌
	// ▀▀▀▀▀▀▘
}
