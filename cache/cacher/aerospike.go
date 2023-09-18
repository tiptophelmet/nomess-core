package cacher

import (
	"fmt"
	"time"

	"github.com/aerospike/aerospike-client-go/v6"
	aerospikeurl "github.com/tiptophelmet/aerospike-url"
	"github.com/tiptophelmet/nomess-core/v4/errs"
	"github.com/tiptophelmet/nomess-core/v4/logger"
)

type AerospikeCacher struct {
	client    *aerospike.Client
	namespace string
}

func InitAerospikeCacher() *AerospikeCacher {
	cacher := &AerospikeCacher{}
	return cacher
}

func (ac *AerospikeCacher) Connect(url string) error {
	clientFactory, err := aerospikeurl.Parse(url)
	if err != nil {
		logger.Error("failed to connect to aerospike: %v", err.Error())
		return errs.ErrCacheStoreConnectionFailed
	}

	ac.client, err = clientFactory.BuildClient()
	if err != nil {
		logger.Error("failed to connect to aerospike: %v", err.Error())
		return errs.ErrCacheStoreConnectionFailed
	}

	ac.namespace = clientFactory.GetNamespace()

	return nil
}

func (ac *AerospikeCacher) Set(key string, val []byte, exp time.Duration) error {
	setKey := fmt.Sprintf("set.%v", key)
	aeroKey, err := aerospike.NewKey(ac.namespace, setKey, key)
	if err != nil {
		return err
	}

	bin := aerospike.NewBin(key, val)

	policy := aerospike.NewWritePolicy(0, uint32(exp.Seconds()))

	err = ac.client.PutBins(policy, aeroKey, bin)
	if err != nil {
		return err
	}

	return nil
}

func (ac *AerospikeCacher) Has(key string) (bool, error) {
	setKey := fmt.Sprintf("set.%v", key)
	aeroKey, err := aerospike.NewKey(ac.namespace, setKey, key)
	if err != nil {
		return false, err
	}

	return ac.client.Exists(nil, aeroKey)
}

func (ac *AerospikeCacher) Get(key string) ([]byte, error) {
	setKey := fmt.Sprintf("set.%v", key)
	aeroKey, err := aerospike.NewKey(ac.namespace, setKey, key)
	if err != nil {
		return nil, err
	}

	record, err := ac.client.Get(nil, aeroKey)
	if err != nil {
		return nil, err
	}

	return record.Bins[key].([]byte), nil
}

func (ac *AerospikeCacher) Expire(key string, exp time.Duration) error {
	policy := aerospike.NewWritePolicy(0, uint32(exp.Seconds()))
	policy.RecordExistsAction = aerospike.UPDATE_ONLY

	setKey := fmt.Sprintf("set.%v", key)
	aeroKey, err := aerospike.NewKey(ac.namespace, setKey, key)
	if err != nil {
		return err
	}

	return ac.client.Touch(policy, aeroKey)
}

func (ac *AerospikeCacher) ExpireTime(key string) (time.Duration, error) {
	setKey := fmt.Sprintf("set.%v", key)
	aeroKey, err := aerospike.NewKey(ac.namespace, setKey, key)
	if err != nil {
		return -1, err
	}

	record, err := ac.client.GetHeader(nil, aeroKey)
	if err != nil {
		return -1, err
	}

	return time.Duration(record.Expiration), nil
}

func (ac *AerospikeCacher) Delete(key string) (bool, error) {
	setKey := fmt.Sprintf("set.%v", key)
	aeroKey, err := aerospike.NewKey(ac.namespace, setKey, key)
	if err != nil {
		return false, err
	}

	deletePolicy := aerospike.NewWritePolicy(0, 0)
	deletePolicy.DurableDelete = true

	return ac.client.Delete(deletePolicy, aeroKey)
}

func (ac *AerospikeCacher) Flush() error {
	return ac.client.Truncate(nil, ac.namespace, "", nil)
}
