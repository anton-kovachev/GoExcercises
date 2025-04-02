/*
  Write the following interface, structs and methods as defined below.
  
      Item: an interface that defines a single method:
      
        
          GetPrice() float64: a method that returns a float
          representing the price of an item.
        
      
    
    
      Gravel: a struct that implements the
      Item interface and has the following fields:
      
        
          lbs: a float64 value representing the
          quantity of gravel.
        
        
          pricePerLb: a float64 value representing the
          price per lb of gravel.
        
        
          grade: a string value representing the grade
          of the gravel. The possible values for this are "fine",
          "medium" and "coarse".
    
      Shovel: a struct that implements the
      Item interface and has the following fields:
      
        
          size: a string value representing the size
          of the shovel. The possible values are "S",
          "M" and "L".
        
        
          price: the price of a shovel of the specified
          size.
        
      
    After writing the types and methods specified above complete the
    CalculateOrderPrice function. This function accepts a slice of
    Item's (representing a users order) and returns the total price
    of the order (i.e., the sum of the price of all the items).
*/

package main

type Item interface {
    GetPrice() float64
}

type Gravel struct {
    lbs float64
    pricePerLb float64
    grade string
}

func (g *Gravel) GetPrice() float64 {
    return g.lbs * g.pricePerLb
}

type Shovel struct {
    size string
    price float64
}

func (s *Shovel) GetPrice() float64 {
    return s.price
}

func CalculateOrderPrice(order []Item) float64 {
	total := 0.0
    for _, item := range order {
        total += item.GetPrice()
    }

    return total
}
