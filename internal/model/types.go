package model

type Environment struct {
	Name             string `yaml:"name"`
	MonthlyBudgetNOK int    `yaml:"monthlyBudgetNOK"`
}

type System struct {
	Name             string        `yaml:"name"`
	Owner            string        `yaml:"owner"`
	MonthlyBudgetNOK int           `yaml:"monthlyBudgetNOK"`
	Environments     []Environment `yaml:"environments"`
	Apps             []App         `yaml:"apps"`
}

type Postgres struct {
	Port int `yaml:"port"`
}

type App struct {
	Name      string   `yaml:"name"`
	Port      int      `yaml:"port"`
	Postgres  Postgres `yaml:"postgres"`
	Directory string   `yaml:"directory"`
}
