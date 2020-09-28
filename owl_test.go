package owl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var exampleYaml = `
name: test
addr: :8080
test:
  test01: test01
  test02: test02
`

var exampleByteYaml = []byte(exampleYaml)

var emptyByteYaml = []byte("")

func resetOwl() {
	owl.filename = ""
	owl.filepath = nil
	owl.value = ""
	owl.key = ""
	owl.config = nil
	owl.client = nil
}

func TestNew(t *testing.T) {
	owlTest := New()
	assert.NotNil(t, owlTest)
	assert.NotSame(t, owlTest, owl)
}

func TestSetRemoteAddr(t *testing.T) {
	SetRemoteAddr([]string{"localhost:2379"})
	assert.NotNil(t, owl.client)
}
func TestOwl_SetRemoteAddr(t *testing.T) {
	tests := []struct {
		name string
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestSetRemoteConfig(t *testing.T) {
	tests := []struct {
		name string
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
func TestOwl_SetRemoteConfig(t *testing.T) {
	tests := []struct {
		name string
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestGetRemoteKeys(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

		})
	}
}
func TestOwl_GetRemoteKeys(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestGetRemote(t *testing.T) {
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
func TestOwl_GetRemote(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestPutRemote(t *testing.T) {
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
func TestOwl_PutRemote(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestWatcher(t *testing.T) {
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
func TestOwl_Watcher(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestAddConfPath(t *testing.T) {
	AddConfPath("1")
	AddConfPath("2")
	assert.Equal(t, 2, len(owl.filepath))
	assert.Equal(t, "1", owl.filepath[0])
	assert.Equal(t, "2", owl.filepath[1])
}
func TestOwl_AddConfPath(t *testing.T) {
	o := New()
	o.AddConfPath("1")
	o.AddConfPath("2")
	assert.Equal(t, 2, len(o.filepath))
	assert.Equal(t, "1", o.filepath[0])
	assert.Equal(t, "2", o.filepath[1])
}

func TestSetConfName(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  string
	}{
		{"a", "test.yaml", "test.yaml"},
		{"b", "test1.yaml", "test1.yaml"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resetOwl()
			SetConfName(tt.value)
			assert.NotEmpty(t, owl.filename)
			assert.Equal(t, owl.filename, tt.want)
		})
	}
}
func TestOwl_SetConfName(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  string
	}{
		{"a", "test.yaml", "test.yaml"},
		{"b", "test1.yaml", "test1.yaml"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := New()
			o.SetConfName(tt.value)
			assert.NotEmpty(t, o.filename)
			assert.Equal(t, o.filename, tt.want)
		})
	}
}

func TestReadConf(t *testing.T) {
	AddConfPath("./mock/")
	SetConfName("test.yaml")
	err := ReadConf()
	assert.Nil(t, err)
	SetConfName("test1.yaml")
	err = ReadConf()
	assert.Error(t, err)
}
func TestOwl_ReadConf(t *testing.T) {
	o := New()
	o.AddConfPath("./mock/")
	o.SetConfName("test.yaml")
	err := o.ReadConf()
	assert.Nil(t, err)
	o.SetConfName("test1.yaml")
	err = o.ReadConf()
	assert.Error(t, err)
}

func TestReadInConf(t *testing.T) {
	type args struct {
		content []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "", args: args{content: exampleByteYaml}, wantErr: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ReadInConf(tt.args.content); (err != nil) != tt.wantErr {
				t.Errorf("ReadInConf() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
func TestOwl_ReadInConf(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name    string
		content []byte
		wantErr bool
	}{
		{"", exampleByteYaml, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := New()
			err := o.ReadInConf(tt.content)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestGet(t *testing.T) {
	tests := []struct {
		name string
		args string
		want interface{}
	}{
		{name: "name", args: "name", want: "test"},
		{name: "test01", args: "test.test01", want: "test01"},
	}
	resetOwl()
	SetConfName("test.yaml")
	AddConfPath("./mock/")
	err := ReadConf()
	assert.Nil(t, err)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, Get(tt.args), tt.want)
		})
	}
}
func TestOwl_Get(t *testing.T) {
	tests := []struct {
		name string
		args string
		want interface{}
	}{
		{name: "name", args: "name", want: "test"},
		{name: "test01", args: "test.test01", want: "test01"},
	}
	o := New()
	o.SetConfName("test.yaml")
	o.AddConfPath("./mock/")
	err := o.ReadConf()
	assert.Nil(t, err)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, o.Get(tt.args), tt.want)
		})
	}
}
func TestGetAll(t *testing.T) {
	tests := []struct {
		name    string
		content []byte
		wantNil bool
	}{
		{"1", exampleByteYaml, false},
		{"2", emptyByteYaml, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resetOwl()
			err := ReadInConf(tt.content)
			assert.Nil(t, err)
			if tt.wantNil {
				assert.Nil(t, GetAll())
			} else {
				assert.NotNil(t, GetAll())
			}
		})
	}
}
func TestOwl_GetAll(t *testing.T) {
	tests := []struct {
		name    string
		content []byte
		wantNil bool
	}{
		{"1", exampleByteYaml, false},
		{"2", emptyByteYaml, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := New()
			err := o.ReadInConf(tt.content)
			assert.Nil(t, err)
			if tt.wantNil {
				assert.Nil(t, o.GetAll())
			} else {
				assert.NotNil(t, o.GetAll())
			}
		})
	}
}
func TestOwl_findConfigFile(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		filename string
		want     string
		wantErr  bool
	}{
		{"a", "./mock/", "test.yaml", "./mock/test.yaml", false},
		{"b", "./mock/", "test1.yaml", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := New()
			o.filepath = []string{tt.value}
			o.filename = tt.filename
			file, err := o.findConfigFile()
			assert.Equal(t, file, tt.want)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
