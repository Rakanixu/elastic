/*
This client is intended for testing purposes. It consumes its API and asserts responses.
Bear in mind that there is only client for API so far, we are happy if the API behaves as expected, so
we can derive that the service itself is working as intended.

Some use cases may need just the service.
For completeness, both API and srv may have its own client.

In the microservices world, interfaces are less likely to change than the internal implementation.
The effort / value relationship this kind of testing offers is bigger than unit testing.
As the complexity of a service implementation increase, unit testing will turn necessary.

Our approach, 20% effort gives back 80% of the value.
*/

package main

import (
	"log"
	"net/http"
	"strings"
	"testing"
)

func main() {

}

type TestHelper struct {
	input        []int
	expectations []int
}

/*func TestCalculateFactorial(t *testing.T) {
	helper := TestHelper{
		input:        []int{1, 3, 7},
		expectations: []int{1, 6, 5040},
	}

	for index, value := range helper.input {
		result := CalculateFactorial(value)

		if result != helper.expectations[index] {
			t.Fatalf("Expected %d, got %d", helper.expectations[index], result)
		}
	}

}*/

func TestCreate(t *testing.T) {
	json := `{
	    "index":"tests",
	    "type": "test",
	    "id": "id",
	    "data":  {
		"att1": "value1",
		"att2": false,
		"innerobj": {
		    "att1": 46,
		    "att2": 1.1
		}
	    }
	}`

	reader := strings.NewReader(json)

	// TODO: configure domain dinamically
	req, err := http.NewRequest("POST", "http://localhost:8080/elasticsearch/create", reader)
	if err != nil {
		// ?? try x3 times again? Test failed? server not there yet? Bug on http, NOO?? fuckkkkkkk
		// Error generating newRequest, I would assume bad data on my test...
	}
	res, err := http.DefaultClient.Do(req)
	//res, err := http.DefaultClient.Post("POST", "http://localhost:8080/elasticsearch/create", reader)
	if err != nil {
		// ?? try x3 times again? Test failed? server not there yet? Bug on http, NOO?? fuckkkkkkk
	}

	if res == nil {
		//Something wrong, test failed
	}
	log.Println("-----------------------------")

	log.Println(res.Body)
	log.Println(res.Status)
	log.Println(res.Request)
}
