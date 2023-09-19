package models

type Model struct {
	ID              string           `json:"id"`
	Name            string           `json:"name"`
	Desc            string           `json:"desc"`
	DeviceResources []DeviceResource `json:"device_resources"`
	DeviceCommands  []DeviceCommand  `json:"device_commands"`
}

type DeviceResource struct {
	Name         string `json:"name"`
	Desc         string `json:"desc"`
	ReadWrite    string `json:"read_write"`
	ValueType    string `json:"value_type"`
	Units        string `json:"units"`
	DefaultValue string `json:"default_value"`
	Minimum      string `json:"minimum"`
	Maximum      string `json:"maximum"`
}

type DeviceCommand struct {
	Name               string              `json:"name"`
	Desc               string              `json:"desc"`
	ReadWrite          string              `json:"read_write"`
	ResourceOperations []ResourceOperation `json:"resource_operations"`
}

type ResourceOperation struct {
	ResourceName string `json:"resource_name"`
	Value        string `json:"value"`
}
