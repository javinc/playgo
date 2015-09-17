package test

func Find(o *Options) Model {
    

    return Model {
        Name: o.Filters.Name + " OOO",
        Description: o.Filters.Description,
    }
}
