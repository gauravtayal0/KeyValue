package main

import (
	"github.com/gauravtayal0/KeyValue/internal"
)

func main()  {
	internal.NewDB("memory")

	internal.DBInstance.Set("Delhi", internal.Value{"pollution_level": "very high"})
	internal.DBInstance.Set("jakarta", internal.Value{"latitude": -6.0, "longitude": 106})
	internal.DBInstance.Set("bangalore", internal.Value{"pollution_level": "very high", "latitude": 1.0, "longitude": 106})
	internal.DBInstance.Set("crocin", internal.Value{"category": []string{"cold", "flu"}})

	internal.DBInstance.Get("Delhi")
	internal.DBInstance.Get("Delhis")

	internal.DBInstance.Delete("Delhi")
	internal.DBInstance.Get("Delhi")


	internal.DBInstance.GetForColumnValue("pollution_level", "very high")
}
