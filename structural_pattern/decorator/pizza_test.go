package decorator

import (
	"testing"
)

func TestPizza(t *testing.T) {

	veggiePizza := &veggeMania{}

	//Add cheese topping
	veggiePizzaWithCheese := &cheeseTopping{
		pizza: veggiePizza,
	}

	//Add tomato topping
	veggiePizzaWithCheeseAndTomato := &tomatoTopping{
		pizza: veggiePizzaWithCheese,
	}

	t.Logf("Price of veggieMania pizza with tomato and cheese topping is %d\n", veggiePizzaWithCheeseAndTomato.getPrice())

	peppyPaneerPizza := &peppyPaneer{}
	// Add tomato topping
	peppyPaneerPizzaWithTomato := &tomatoTopping{pizza: peppyPaneerPizza}

	//Add cheese topping
	peppyPaneerPizzaWithTomatoAndCheese := &cheeseTopping{
		pizza: peppyPaneerPizzaWithTomato,
	}

	t.Logf("Price of peppyPaneer with tomato and cheese topping is %d\n", peppyPaneerPizzaWithTomatoAndCheese.getPrice())

}
