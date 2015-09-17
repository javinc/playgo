package test

import "github.com/javinc/playgo/goryo/resources/test"

func Find(o *test.Options) test.Model {
    return test.Model {
        Name: o.Filters.Name + " OOO",
        Description: o.Filters.Description,
    }
}
