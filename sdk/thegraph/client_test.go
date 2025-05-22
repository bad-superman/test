package thegraph

import "testing"

func TestQueryIndexers(t *testing.T) {
	client := NewClient("ae757456b31a0feaf48338e1a91402eb")
	result := &QueryIndexersResponse{}
	err := client.QueryIndexers("ANk8smWo9Y1FCBY6EBU5mG2ArWzYJ1iex7iQb6SCB41X", 10, 0, result)
	if err != nil {
		t.Fatal(err)
	}
}
