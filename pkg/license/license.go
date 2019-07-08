package license

// License ...
type License struct {
	Data []string
}

// AppendLine ...
func (l *License) AppendLine(line string) {
	l.Data = append(l.Data, line)
}

// Lines ...
func (l *License) Lines() int {
	return len(l.Data)
}
