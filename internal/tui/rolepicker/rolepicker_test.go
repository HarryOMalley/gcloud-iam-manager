package rolepicker

import (
	"testing"

	"gcloud-iam-manager/internal/iam"
)

func TestNew(t *testing.T) {
	// A flat list of roles to be converted into a tree
	roles := []*iam.Role{
		{Name: "roles/viewer"},
		{Name: "roles/editor"},
		{Name: "roles/owner"},
		{Name: "roles/iam.roleAdmin"},
		{Name: "roles/iam.serviceAccountAdmin"},
	}

	m := New(roles)

	// Assert the tree structure
	if len(m.root.children) != 1 {
		t.Fatalf("Expected root to have 1 child ('roles'), got %d", len(m.root.children))
	}

	rolesNode := m.root.children[0]
	if rolesNode.name != "roles" {
		t.Errorf("Expected first child to be 'roles', got %q", rolesNode.name)
	}

	// Check for viewer, editor, owner, and the 'iam' parent node
	if len(rolesNode.children) != 4 {
		t.Fatalf("Expected 'roles' node to have 4 children, got %d", len(rolesNode.children))
	}

	// Check for the nested 'iam' node
	iamNode := rolesNode.findChild("iam")
	if iamNode == nil {
		t.Fatal("Expected to find child node 'iam'")
	}

	if len(iamNode.children) != 2 {
		t.Errorf("Expected 'iam' node to have 2 children, got %d", len(iamNode.children))
	}

	// Check for a leaf node
	saAdminNode := iamNode.findChild("serviceAccountAdmin")
	if saAdminNode == nil {
		t.Fatal("Expected to find child node 'serviceAccountAdmin'")
	}
	if saAdminNode.role == nil {
		t.Fatal("Expected leaf node to have a non-nil role")
	}
	if saAdminNode.role.Name != "roles/iam.serviceAccountAdmin" {
		t.Errorf("Node has incorrect role name: %s", saAdminNode.role.Name)
	}
}
