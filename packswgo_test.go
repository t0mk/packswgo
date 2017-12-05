package packswgo

import (
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"testing"
	"time"
)

const (
	packetTokenEnvVar = "PACKET_AUTH_TOKEN"
	packngoAccTestVar = "PACKSWGO_TEST_ACTUAL_API"
	testProjectPrefix = "PACKSWGO_TEST_DELME_2d767456_"
	testFacilityVar   = "PACKSWGO_TEST_FACILITY"
)

func testFacility() string {
	envFac := os.Getenv(testFacilityVar)
	if envFac != "" {
		return envFac
	}
	return "ewr1"
}

func randString8() string {
	n := 8
	rand.Seed(time.Now().UnixNano())
	letterRunes := []rune("acdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func setupWithProject(t *testing.T) (*Client, string, func()) {
	c := setup(t)
	p, _, err := c.CreateProject(ProjectCreateInput{Name: testProjectPrefix + randString8()})
	if err != nil {
		t.Fatal(err)
	}

	return c, p.Id, func() {
		_, err := c.DeleteProject(p.Id)
		if err != nil {
			panic(fmt.Errorf("while deleting %s: %s", p, err))
		}
	}

}

func skipUnlessAcceptanceTestsAllowed(t *testing.T) {
	if os.Getenv(packngoAccTestVar) == "" {
		t.Skipf("%s is not set", packngoAccTestVar)
	}
}

func setup(t *testing.T) *Client {
	c, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}
	return c
}

func projectTeardown(c *Client) {
	ps, _, err := c.FindProjects("", 1, 10)
	if err != nil {
		panic(fmt.Errorf("while teardown: %s", err))
	}
	for _, p := range ps.Projects {
		if strings.HasPrefix(p.Name, testProjectPrefix) {
			_, err := c.DeleteProject(p.Id)
			if err != nil {
				panic(fmt.Errorf("while deleting %s: %s", p, err))
			}
		}
	}
}

func TestAccInvalidCredentials(t *testing.T) {
	skipUnlessAcceptanceTestsAllowed(t)
	c, err := NewClientWithToken("wrongApiToken")
	if err != nil {
		t.Fatal(err)
	}
	_, r, err := c.FindProjects("", 1, 10)
	if err != nil {
		t.Fatal(err)
	}
	expectedErr := string(r.Payload)
	matched, err := regexp.MatchString(".*Invalid.*", expectedErr)
	if err != nil {
		t.Fatalf("Err while matching err string from response err %s: %s", expectedErr, err)
	}
	if r.StatusCode != 401 {
		t.Fatalf("Expected 401 as response code, got: %d", r.StatusCode)
	}

	if !matched {
		t.Fatalf("Unexpected error string: %s", expectedErr)
	}

}
