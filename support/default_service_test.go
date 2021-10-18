package support

import (
	"testing"

	"github.com/bitwormhole/starter-cli/filters"
)

func TestDefaultService(t *testing.T) {

	service := &DefaultSerivce{}

	service.AddFilter(500, &filters.ContextFilter{Service: service})
	service.AddFilter(400, &filters.MultilineSupportFilter{})
	service.AddFilter(300, &filters.HandlerFinderFilter{})
	service.AddFilter(200, &filters.ExecutorFilter{})
	service.AddFilter(0, &filters.NopFilter{})

	//	service.RegisterHandler("help", &handlers.Help{})

	// client := service.GetClient(nil)
	// err := client.ExecuteScript("help hello world")
	// if err != nil {
	// 	t.Error(err)
	// }
}
