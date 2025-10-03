package rolepicker

import (
	"strings"

	"gcloud-iam-manager/internal/iam"
	tea "github.com/charmbracelet/bubbletea"
)

// node represents a single entry in the role tree.
type node struct {
	name     string
	parent   *node
	children []*node
	role     *iam.Role // nil if not a final role
}

// Model represents the state of the role picker component.
type Model struct {
	root         *node               // The root of the role tree
	cursor       int                 // Index of the currently focused node in the visibleNodes list
	selected     map[string]struct{} // Set of selected role names
	expanded     map[*node]bool      // Set of expanded nodes
	visibleNodes []*node             // A flat list of currently visible nodes
}

// New creates a new role picker model from a flat list of roles.
func New(roles []*iam.Role) Model {
	root := &node{name: "root"}
		// The first part is always "roles", so we create that node first.
		roleroot := &node{name: "roles", parent: root}
		root.children = append(root.children, roleroot)
	
		for _, role := range roles {
			parts := strings.Split(role.Name, "/")
			if len(parts) < 2 {
				continue
			}
	
			currentNode := roleroot // Always start from the 'roles' node
			path := parts[1]
	
			for _, subPart := range strings.Split(path, ".") {
				child := currentNode.findChild(subPart)
				if child == nil {
					child = &node{name: subPart, parent: currentNode}
					currentNode.children = append(currentNode.children, child)
				}
				currentNode = child
			}
			currentNode.role = role // Attach the full role to the final leaf node
		}
	m := Model{
		root:     root,
		selected: make(map[string]struct{}),
		expanded: make(map[*node]bool),
	}
	m.visibleNodes = m.getVisibleNodes()
	return m
}

// getVisibleNodes recursively builds a flat list of nodes that are currently visible.
func (m *Model) getVisibleNodes() []*node {
	var nodes []*node
	var recurse func(*node)
	recurse = func(n *node) {
		for _, child := range n.children {
			nodes = append(nodes, child)
			if m.expanded[child] {
				recurse(child)
			}
		}
	}

	// Start recursion from the main 'roles' node.
	if len(m.root.children) > 0 {
		recurse(m.root.children[0])
	}
	return nodes
}

// IsSelected checks if a role name is in the selected set.
func (m *Model) IsSelected(roleName string) bool {
	_, ok := m.selected[roleName]
	return ok
}

// findChild searches for a child node by name.
func (n *node) findChild(name string) *node {
	for _, child := range n.children {
		if child.name == name {
			return child
		}
	}
	return nil
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		currentNode := m.visibleNodes[m.cursor]
		switch msg.String() {
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.visibleNodes)-1 {
				m.cursor++
			}
		case "right", "l":
			if len(currentNode.children) > 0 {
				m.expanded[currentNode] = true
				m.visibleNodes = m.getVisibleNodes()
			}
		case "left", "h":
			if len(currentNode.children) > 0 {
				delete(m.expanded, currentNode)
			} else if currentNode.parent != nil {
				delete(m.expanded, currentNode.parent)
			}
			m.visibleNodes = m.getVisibleNodes()

		case " ": // spacebar
			if currentNode.role != nil {
				if m.IsSelected(currentNode.role.Name) {
					delete(m.selected, currentNode.role.Name)
				} else {
					m.selected[currentNode.role.Name] = struct{}{}
				}
			}
		}
	}
	return m, nil
}

func (m Model) View() string {
	var b strings.Builder
	b.WriteString("Select roles to assign (space to select, enter to confirm):\n\n")

	var renderNode func(*node, int)
	renderNode = func(n *node, depth int) {
		// Render the current node
		isCursorOnNode := m.visibleNodes[m.cursor] == n
		
		cursor := "  "
		if isCursorOnNode {
			cursor = "> "
		}
		b.WriteString(cursor)

		b.WriteString(strings.Repeat("  ", depth))

		if len(n.children) > 0 {
			if m.expanded[n] {
				b.WriteString("[-] ")
			} else {
				b.WriteString("[+] ")
			}
		} else {
			selection := "[ ] "
			if n.role != nil && m.IsSelected(n.role.Name) {
				selection = "[x] "
			}
			b.WriteString(selection)
		}

		b.WriteString(n.name + "\n")

		// If expanded, render children
		if m.expanded[n] {
			for _, child := range n.children {
				renderNode(child, depth+1)
			}
		}
	}

	// Find the subset of the tree to render
	// This is a simplified approach; a real implementation would use a viewport
	for _, node := range m.root.children {
		// We only render the children of the true root, which is our main 'roles' node
		if m.expanded[node] || node.name == "roles" { // Start with roles expanded
			for _, child := range node.children {
				renderNode(child, 0)
			}
		}
	}

	return b.String()
}
