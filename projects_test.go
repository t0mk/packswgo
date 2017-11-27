package packswgo

import (
	"log"
	"testing"
)

func TestAccProject(t *testing.T) {
	skipUnlessAcceptanceTestsAllowed(t)

	c := setup(t)
	defer projectTeardown(c)

	rs := testProjectPrefix + randString8()
	pcr := ProjectCreateInput{Name: rs}
	p, resp, err := c.CreateProject(pcr)
	log.Println(string(resp.Payload))

	if err != nil {
		t.Fatal(err)
	}
	err = CheckForAppError(resp)
	if err != nil {
		t.Fatal(err)
	}
	if p.Name != rs {
		t.Fatalf("Expected new project name to be %s, not %s", rs, p.Name)
	}
	rs = testProjectPrefix + randString8()
	pur := ProjectUpdateInput{Name: rs}
	p, _, err = c.UpdateProject(p.Id, pur)
	if err != nil {
		t.Fatal(err)
	}
	if p.Name != rs {
		t.Fatalf("Expected the name of the updated project to be %s, not %s", rs, p.Name)
	}
	gotProject, _, err := c.FindProjectById(p.Id, "")
	if err != nil {
		t.Fatal(err)
	}
	if gotProject.Name != rs {
		t.Fatalf("Expected the name of the GOT project to be %s, not %s", rs, gotProject.Name)
	}
	_, err = c.DeleteProject(p.Id)
	if err != nil {
		t.Fatal(err)
	}

}
