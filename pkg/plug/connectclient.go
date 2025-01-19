package plug

import (
	"fmt"
	"net/rpc"
	"time"

	"github.com/enuesaa/teatime/pkg/repository"
	"github.com/enuesaa/teatime/pkg/repository/db"
)

// should implement ProviderInterface
type ConnectClient struct {
	client *rpc.Client
	repos repository.Repos
}

func (cc *ConnectClient) logErr(err error) {
	doc := db.Log{
		Created: time.Now().Format(time.RFC3339),
		Message: fmt.Sprintf("Error: %s", err.Error()),
	}
	query := cc.repos.DB.Logs()
	query.Create(doc)
}

func (cc *ConnectClient) OnStartup() error {
	var resp error
	// do not catch err becasuse this does not return err. see ConnectServer.
	cc.client.Call("Plugin.OnStartup", new(interface{}), &resp)
	if resp != nil {
		cc.logErr(resp)
	}
	return resp
}

func (cc *ConnectClient) OnShutdown() error {
	var resp error
	// do not catch err becasuse this does not return err. see ConnectServer.
	cc.client.Call("Plugin.OnShutdown", new(interface{}), &resp)
	if resp != nil {
		cc.logErr(resp)
	}
	return resp
}

func (cc *ConnectClient) Info() (Info, error) {
	var resp ConnectResult[Info]
	if err := cc.client.Call("Plugin.Info", new(interface{}), &resp); err != nil {
		cc.logErr(err)
		return resp.Data, err
	}
	if resp.Err != nil {
		cc.logErr(resp.Err)
	}
	return resp.Data, resp.Err
}

func (cc *ConnectClient) List(props ListProps) ([]db.Tea, error) {
	var resp ConnectResult[[]db.Tea]
	if err := cc.client.Call("Plugin.List", props, &resp); err != nil {
		cc.logErr(err)
		return resp.Data, err
	}
	if resp.Err != nil {
		cc.logErr(resp.Err)
	}
	return resp.Data, resp.Err
}

func (cc *ConnectClient) Get(props GetProps) (db.Tea, error) {
	var resp ConnectResult[db.Tea]
	if err := cc.client.Call("Plugin.Get", props, &resp); err != nil {
		cc.logErr(err)
		return resp.Data, err
	}
	if resp.Err != nil {
		cc.logErr(resp.Err)
	}
	return resp.Data, resp.Err
}

func (cc *ConnectClient) Create(props CreateProps) (string, error) {
	var resp ConnectResult[string]
	if err := cc.client.Call("Plugin.Create", props, &resp); err != nil {
		cc.logErr(err)
		return resp.Data, err
	}
	if resp.Err != nil {
		cc.logErr(resp.Err)
	}
	return resp.Data, resp.Err
}

func (cc *ConnectClient) Update(props UpdateProps) (string, error) {
	var resp ConnectResult[string]
	if err := cc.client.Call("Plugin.Update", props, &resp); err != nil {
		cc.logErr(err)
		return resp.Data, err
	}
	if resp.Err != nil {
		cc.logErr(resp.Err)
	}
	return resp.Data, resp.Err
}

func (cc *ConnectClient) Delete(props DeleteProps) (bool, error) {
	var resp ConnectResult[bool]
	if err := cc.client.Call("Plugin.Delete", props, &resp); err != nil {
		cc.logErr(err)
		return resp.Data, err
	}
	if resp.Err != nil {
		cc.logErr(resp.Err)
	}
	return resp.Data, resp.Err
}
