package test

import "github.com/javinc/playgo/goryo/resources/test"

func Find(o *test.Options) []test.Model {
    rows, _ := TestResource.Find(*o)

    return rows
}
