package cmd

import (
	"bytes"
	"context"
	"encoding/gob"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/urfave/cli"
	"go.etcd.io/etcd/clientv3"

	"github.com/coreos/etcd/mvcc/mvccpb"
)

func restoreCmd() cli.Command {
	return cli.Command{
		Name:   "restore",
		Usage:  "restore K/V from file",
		Action: restoreAction,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:     "address, a",
				Usage:    "etcd address",
				Value:    defaultEtcdAddress,
				Required: false,
			},
			cli.StringFlag{
				Name:     "file, f",
				Usage:    "restore from `FILE`",
				Required: true,
			},
			cli.BoolFlag{
				Name:     "silent, s",
				Usage:    "verbose mode",
				Required: false,
			},
			cli.StringFlag{
				Name:     "user, u",
				Usage:    "user name for etcd",
				Required: false,
			},
			cli.StringFlag{
				Name:     "password, pw",
				Usage:    "password for etcd",
				Required: false,
			},
		},
	}
}

func restoreAction(c *cli.Context) error {
	address := c.String("address")
	if len(address) == 0 {
		return errors.New("address shouldn't be empty")
	}

	file := c.String("file")
	if len(file) == 0 {
		return errors.New("file shouldn't be empty")
	}

	silent := c.Bool("silent")

	user := c.String("user")
	password := c.String("password")

	return restore(address, file, user, password, !silent)
}

func restore(addr, filename, user, password string, print bool) error {
	dd, err := readDumpData(filename)
	if err != nil {
		return err
	}

	cfg := clientv3.Config{
		Endpoints:   []string{addr},
		DialTimeout: 5 * time.Second,
	}

	if len(user) != 0 && len(password) != 0 {
		cfg.Username = user
		cfg.Password = password
	}

	cli, err := clientv3.New(cfg)
	if err != nil {
		return err
	}
	defer cli.Close()

	ctx := context.Background()

	for _, kvB := range dd {
		var kv mvccpb.KeyValue
		if err := kv.Unmarshal(kvB); err != nil {
			return err
		}

		pCtx, kCancel := context.WithTimeout(ctx, 5*time.Second)
		_, err = cli.Put(pCtx, string(kv.Key), string(kv.Value))
		kCancel()
		if err != nil {
			return err
		}

		if print {
			fmt.Println(string(kv.Key))
		}
	}

	return nil
}

func readDumpData(filename string) (dumpData, error) {
	b, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	dd := make(dumpData, 0)

	var buffer bytes.Buffer
	buffer.Write(b)

	dec := gob.NewDecoder(&buffer)
	err = dec.Decode(&dd)
	if err != nil {
		return nil, err
	}

	return dd, nil
}
