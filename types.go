package main

type Inputs struct {
	Name       string `json:"name"`
	Type       string `json:"type"`
	Components string `json:"components"`
	Indexed    bool   `json:"indexed"`
}

type Outputs struct {
	InternalType string `json:"internalType"`
	Name         string `json:"name"`
	Type         string `json:"type"`
}

type Abi struct {
	Inputs          []Inputs  `json:"inputs"`
	Name            string    `json:"name"`
	Outputs         []Outputs `json:"outputs"`
	StateMutability string    `json:"stateMutability"`
	Type            string    `json:"type"`
}
