package owl

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"strings"
	"sync"
	"time"
)

var owl *Owl

func init() {
	owl = New()
}

// Owl is a lib for get configure value from etcd.
type Owl struct {
	key      string
	value    string
	config   map[string]interface{}
	filename string
	filepath []string
	client   *clientv3.Client
	lock     sync.RWMutex
}

// New returns an initialized Owl instance.
func New() *Owl {
	return &Owl{}
}

// SetRemoteConfig set configure for the etcd. The
// client include etcd url
func SetRemoteConfig(config clientv3.Config) { owl.SetRemoteConfig(config) }
func (o *Owl) SetRemoteConfig(config clientv3.Config) {
	client, err := clientv3.New(config)
	if err != nil {
		client = nil
	}
	o.client = client
}

// SetRemoteAddr set url for the etcd.
func SetRemoteAddr(addr []string) { owl.SetRemoteAddr(addr) }
func (o *Owl) SetRemoteAddr(addr []string) {
	conf := clientv3.Config{
		Endpoints:        addr,
		AutoSyncInterval: 0,
		DialTimeout:      5 * time.Second,
	}
	client, err := clientv3.New(conf)
	if err != nil {
		client = nil
	}
	o.client = client
}

// GetRemoteKeys get keys from etcd by prefix
func GetRemoteKeys(prefix string) ([]string, error) { return owl.GetRemoteKeys(prefix) }
func (o *Owl) GetRemoteKeys(prefix string) ([]string, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	resp, err := o.client.Get(ctx, prefix, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}
	if resp == nil {
		return nil, nil
	}
	var keys []string
	for _, v := range resp.Kvs {
		keys = append(keys, string(v.Key))
	}
	return keys, nil
}

// GetRemote get config content from etcd by key
func GetRemote(key string) (string, error) { return owl.GetRemote(key) }
func (o *Owl) GetRemote(key string) (string, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	resp, err := o.client.Get(ctx, key)
	if err != nil {
		return "", err
	}
	var value string

	for _, v := range resp.Kvs {
		value = string(v.Value)
	}

	return value, nil
}

// PutRemote value into etcd.
func PutRemote(key, value string) error { return owl.PutRemote(key, value) }
func (o *Owl) PutRemote(key, value string) error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	_, err := o.client.Put(ctx, key, value)
	if err != nil {
		return err
	}
	return nil
}

// Watcher watch key's value in etcd
func Watcher(key string, c chan string) { owl.Watcher(key, c) }
func (o *Owl) Watcher(key string, c chan string) {
	rch := o.client.Watch(context.Background(), key)
	for resp := range rch {
		for _, ev := range resp.Events {
			switch ev.Type {
			case mvccpb.PUT:
				c <- string(ev.Kv.Value)
			case mvccpb.DELETE:
				c <- ""
			default:
			}
		}
	}
}

func SetConfName(name string) { owl.SetConfName(name) }
func (o *Owl) SetConfName(name string) {
	o.lock.Lock()
	defer o.lock.Unlock()
	o.filename = name
}

func AddConfPath(path string) { owl.AddConfPath(path) }
func (o *Owl) AddConfPath(path string) {
	o.lock.Lock()
	defer o.lock.Unlock()
	o.filepath = append(o.filepath, path)
}

func ReadConf() error { return owl.ReadConf() }
func (o *Owl) ReadConf() error {
	if o.filename == "" {
		return errors.WithStack(errors.New("config name not set"))
	}
	if o.filepath == nil {
		return errors.WithStack(errors.New("config path not set"))
	}

	file, err := o.findConfigFile()
	if err != nil {
		return err
	}

	content, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(content, &o.config)
	if err != nil {
		return err
	}
	return nil
}

func (o *Owl) findConfigFile() (string, error) {
	for _, v := range o.filepath {
		exist, _ := exists(v + o.filename)
		if exist {
			return v + o.filename, nil
		}
	}
	return "", errors.New("file is not exist")
}

func ReadInConf(content []byte) error { return owl.ReadInConf(content) }
func (o *Owl) ReadInConf(content []byte) error {
	err := yaml.Unmarshal(content, &o.config)
	if err != nil {
		return err
	}
	return nil
}

func Get(key string) interface{} { return owl.Get(key) }
func (o *Owl) Get(key string) interface{} {
	keys := strings.Split(key, ".")
	return o.find(o.config, keys)
}

//func GetString(key string) string                              { return owl.GetString(key) }
//func (o *Owl) GetString(key string) string                     { return "" }
//func GetStringMap(key string) map[string]interface{}           { return owl.GetStringMap(key) }
//func (o *Owl) GetStringMap(key string) map[string]interface{}  { return nil }
//func GetStringMapString(key string) map[string]string          { return owl.GetStringMapString(key) }
//func (o *Owl) GetStringMapString(key string) map[string]string { return nil }
//func GetStringSlice(key string) []string                       { return owl.GetStringSlice(key) }
//func (o *Owl) GetStringSlice(key string) []string              { return nil }
//func GetInt(key string) int                                    { return owl.GetInt(key) }
//func (o *Owl) GetInt(key string) int                           { return 0 }
//func GetIntSlice(key string) []int                             { return owl.GetIntSlice(key) }
//func (o *Owl) GetIntSlice(key string) []int                    { return nil }
//func GetUint(key string) uint                                  { return owl.GetUint(key) }
//func (o *Owl) GetUint(key string) uint                         { return 0 }
//func GetFloat64(key string) float64                            { return owl.GetFloat64(key) }
//func (o *Owl) GetFloat64(key string) float64                   { return 0 }
//func GetBool(key string) bool                                  { return owl.GetBool(key) }
//func (o *Owl) GetBool(key string) bool                         { return true }
//func GetTime(key string) time.Time                             { return owl.GetTime(key) }
//func (o *Owl) GetTime(key string) time.Time                    { return time.Time{} }
//func GeteDuration(key string) time.Duration                    { return owl.GeteDuration(key) }
//func (o *Owl) GeteDuration(key string) time.Duration           { return 0 }
func GetAll() map[string]interface{}          { return owl.GetAll() }
func (o *Owl) GetAll() map[string]interface{} { return o.config }

func (o *Owl) find(source map[string]interface{}, path []string) interface{} {
	if len(path) == 0 {
		return source
	}
	next, ok := source[path[0]]
	if ok {
		if len(path) == 1 {
			return next
		}
		switch source[path[0]].(type) {
		case map[interface{}]interface{}:
			return o.find(cast.ToStringMap(source[path[0]]), path[1:])
		case map[string]interface{}:
			return o.find(source[path[0]].(map[string]interface{}), path[1:])
		default:
			return nil
		}
	}
	return nil
}
