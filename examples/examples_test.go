package examples

import (
	"fmt"
	"testing"

	"github.com/stephennancekivell/go-future/future"
	"github.com/stretchr/testify/assert"
)

func Test_run_two_parallel_results(t *testing.T) {
	assert := assert.New(t)
	// oneF := future.New(func() string {
	// 	return "hello"
	// })

	fmt.Println("start")
	assert.Equal(1, 1, "should be same")

	twoF := future.New(func() string {
		return "world"
	})

	// one := oneF.Wait()
	// two := twoF.Get()

	assert.Equal(twoF.Get(), "world", "should be same")

	// one, hello := future.Par2(
	// 	func() int {
	// 		return 1
	// 	}, func() tuple.T2[string, int] {
	// 		return tuple.New2("hello", 1)
	// 	},
	// )

	// fmt.Printf("one: %v two: %v three:%v\n", one, two, hello)
}

func Test_run_more_complex(t *testing.T) {

	getRecipe := func() string { return "recipe" }
	getWater := func() string { return "water" }
	getOtherIngredients := func(r string) string { return "water" }

	shouldPreheat := func(r string) bool { return true }

	preheatOvern := func() {}

	type Ingredients struct {
		a string
		b string
	}
	type otherIngredients = string

	cook := func(s Ingredients) {}

	recipeFuture := future.New(getRecipe)

	ingredientsFuture := future.New(func() Ingredients {
		// we always need water, so get that before the receipe is ready
		waterFuture := future.New(getWater)
		recipe := recipeFuture.Get()

		// now that that recipe is ready we can get the other ingredients.
		otherIngredientsFuture := future.New(func() otherIngredients { // TODO making a future just to wait on it. Should come up with something more convincing.
			return getOtherIngredients(recipe)
		})
		return Ingredients{
			waterFuture.Get(),
			otherIngredientsFuture.Get(),
		}
	})

	recipe := recipeFuture.Get()

	// we might need to preheat the oven
	if shouldPreheat(recipe) {
		// we can do this in parallel to getting the ingredients
		preheatOvern()
	}

	ingredients := ingredientsFuture.Get()

	cook(ingredients)

	var vegetableFutures []future.Future[string]
	choppedVeges := future.Sequence(vegetableFutures).Get()

	fmt.Printf("chopped %v", choppedVeges)

}
