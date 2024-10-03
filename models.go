package main

type Constructor struct {
	Name string `json:"team_name"`
	Base string `json:"base"`
}

type ConstructorDetails struct {
	Base               string `json:"base"`
	TeamPrincipal      string `json:"team_principal"`
	TechnicalChief     string `json:"technical_chief"`
	Chasis             string `json:"chasis"`
	EngineSupplier     string `json:"engine_supplier"`
	WorldChampionships string `json:"world_championships"`
	Wins               string `json:"wins"`
	PolePositions      string `json:"pole_positions"`
	FastestLaps        string `json:"fastest_laps"`
	FirstTeamEntry     string `json:"first_team_entry"`
	Drivers            []Driver
}

type Driver struct {
	Name   string `json:"name"`
	Number string `json:"number"`
}

type Constructors []Constructor

type ConstructorFullInfo struct {
	TeamName           string             `json:"team_name"`
	ConstructorDetails ConstructorDetails `json:"details"`
}

type ConstructorsFullInfo []ConstructorFullInfo
