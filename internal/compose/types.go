package compose

import composetypes "github.com/compose-spec/compose-go/v2/types"

type composableApp interface {
	name() string
	addEnvVars(env map[string]string)
	emitEnvVars(app composableApp)
	serviceConfig() composetypes.ServiceConfig
}

func newContext(startingPort int) composeContext {
	return &composeContextImpl{
		port: startingPort,
		m:    map[string]composableApp{},
	}
}

type composeContext interface {
	nextPort() int
	addApp(name string, app composableApp)
	apps() map[string]composableApp
}

type composeContextImpl struct {
	port int
	m    map[string]composableApp
}

func (c *composeContextImpl) apps() map[string]composableApp {
	return c.m
}

func (c *composeContextImpl) addApp(name string, app composableApp) {
	c.m[name] = app
}

func (c *composeContextImpl) nextPort() int {
	p := c.port
	c.port++
	return p
}
