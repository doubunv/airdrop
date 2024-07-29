package dto

type ListenerContract struct {
	Chain        string
	Name         string
	ContractAddr string
	Events       []string
	AbiJSON      string
}
