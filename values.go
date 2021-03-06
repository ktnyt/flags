package flags

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// BoolValue represents a boolean argument value.
type BoolValue bool

// NewBoolValue creates a new BoolValue.
func NewBoolValue(init bool) *BoolValue {
	p := new(bool)
	*p = init
	return (*BoolValue)(p)
}

// Set will set attempt to convert the given string to a value.
func (p *BoolValue) Set(s string) error {
	v, err := strconv.ParseBool(s)
	if err != nil {
		return fmt.Errorf("`%s` cannot be interpreted as %T", s, v)
	}
	*p = BoolValue(v)
	return nil
}

// String satisfies the fmt.Stringer interface.
func (p BoolValue) String() string {
	return strconv.FormatBool(bool(p))
}

// IntValue represents a integer argument value.
type IntValue int

// NewIntValue creates a new IntValue.
func NewIntValue(init int) *IntValue {
	p := new(int)
	*p = init
	return (*IntValue)(p)
}

// Set will set attempt to convert the given string to a value.
func (p *IntValue) Set(s string) error {
	v, err := strconv.Atoi(s)
	if err != nil {
		return fmt.Errorf("`%s` cannot be interpreted as %T", s, v)
	}
	*p = IntValue(v)
	return nil
}

// String satisfies the fmt.Stringer interface.
func (p IntValue) String() string {
	return strconv.Itoa(int(p))
}

// FloatValue represents a float argument value.
type FloatValue float64

// NewFloatValue creates a new FloatValue.
func NewFloatValue(init float64) *FloatValue {
	p := new(float64)
	*p = init
	return (*FloatValue)(p)
}

// Set will set attempt to convert the given string to a value.
func (p *FloatValue) Set(s string) error {
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return fmt.Errorf("`%s` cannot be interpreted as %T", s, v)
	}
	*p = FloatValue(v)
	return nil
}

// String satisfies the fmt.Stringer interface.
func (p FloatValue) String() string {
	return strconv.FormatFloat(float64(p), 'g', -1, 64)
}

// StringValue represents a string argument value.
type StringValue string

// NewStringValue creates a new StringValue.
func NewStringValue(init string) *StringValue {
	p := new(string)
	*p = init
	return (*StringValue)(p)
}

// Set will set attempt to convert the given string to a value.
func (p *StringValue) Set(s string) error {
	*p = StringValue(s)
	return nil
}

// String satisfies the fmt.Stringer interface.
func (p StringValue) String() string {
	return string(p)
}

// OpenValue represents a file argument value for opening.
type OpenValue os.File

// NewOpenValue creates a new OpenValue.
func NewOpenValue(init *os.File) *OpenValue {
	p := new(os.File)
	if init != nil {
		*p = *init
	}
	return (*OpenValue)(p)
}

// Set will set attempt to convert the given string to a value.
func (p *OpenValue) Set(s string) error {
	f, err := os.Open(s)
	if err != nil {
		return err
	}
	*p = OpenValue(*f)
	return nil
}

// String satisfies the fmt.Stringer interface.
func (p *OpenValue) String() string {
	return (*os.File)(p).Name()
}

// CreateValue represents a file argument value for creating.
type CreateValue os.File

// NewCreateValue creates a new CreateValue.
func NewCreateValue(init *os.File) *CreateValue {
	p := new(os.File)
	if init != nil {
		*p = *init
	}
	return (*CreateValue)(p)
}

// Set will set attempt to convert the given string to a value.
func (p *CreateValue) Set(s string) error {
	f, err := os.Create(s)
	if err != nil {
		return err
	}
	*p = CreateValue(*f)
	return nil
}

// String satisfies the fmt.Stringer interface.
func (p *CreateValue) String() string {
	return (*os.File)(p).Name()
}

// StringSliceValue represents a variable number string argument value.
type StringSliceValue []string

// NewStringSliceValue creates a new StringSliceValue.
func NewStringSliceValue(init []string) *StringSliceValue {
	p := new([]string)
	*p = init
	return (*StringSliceValue)(p)
}

// Len will return the length of the slice value.
func (v StringSliceValue) Len() int { return len(v) }

// Set will set attempt to convert and append the given string to the slice.
func (p *StringSliceValue) Set(s string) error {
	ss := []string(*p)
	ss = append(ss, s)
	*p = StringSliceValue(ss)
	return nil
}

// String satisfies the fmt.Stringer interface.
func (p StringSliceValue) String() string {
	return fmt.Sprintf("[%s]", strings.Join([]string(p), ", "))
}

// OpenSliceValue represents a variable number open argument value.
type OpenSliceValue []*os.File

// NewOpenSliceValue creates a new OpenSliceValue.
func NewOpenSliceValue(init []*os.File) *OpenSliceValue {
	p := new([]*os.File)
	*p = init
	return (*OpenSliceValue)(p)
}

// Len will return the length of the slice value.
func (v OpenSliceValue) Len() int { return len(v) }

// Set will set attempt to convert and append the given string to the slice.
func (p *OpenSliceValue) Set(s string) error {
	ff := []*os.File(*p)
	f, err := os.Open(s)
	if err != nil {
		return err
	}
	ff = append(ff, f)
	*p = OpenSliceValue(ff)
	return nil
}

// String satisfies the fmt.Stringer interface.
func (v OpenSliceValue) String() string {
	ss := make([]string, len(v))
	for i, f := range v {
		ss[i] = f.Name()
	}
	return fmt.Sprintf("[%s]", strings.Join(ss, ", "))
}
