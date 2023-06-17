package kasm

type Image struct {
	Image_id                     string                        `json:"image_id"`
	Cores                        int                           `json:"cores"`
	Description                  string                        `json:"description"`
	Docker_registry              string                        `json:"docker_registry"`
	Docker_token                 string                        `json:"docker_token"`
	Docker_user                  string                        `json:"docker_user"`
	Enabled                      bool                          `json:"enabled"`
	Friendly_name                string                        `json:"friendly_name"`
	Hash                         string                        `json:"hash"`
	Memory                       int64                         `json:"memory"`
	Name                         string                        `json:"name"`
	X_res                        int                           `json:"x_res"`
	Y_res                        int                           `json:"y_res"`
	ImageAttributes              []imagesImageAttributesModel  `json:"imageAttributes"`
	Restrict_to_network          bool                          `json:"restrict_to_network"`
	Restrict_network_names       []string                      `json:"restrict_network_names"`
	Restrict_to_server           bool                          `json:"restrict_to_server"`
	Server_id                    string                        `json:"server_id"`
	Volume_mappings              map[string]volumeMappingModel `json:"volume_mappings"`
	Run_config                   map[string]any                `json:"run_config"`
	Image_src                    string                        `json:"image_src"`
	Available                    bool                          `json:"available"`
	Exec_config                  map[string]any                `json:"exec_config"`
	Restrict_to_zone             bool                          `json:"restrict_to_zone"`
	Zone_id                      string                        `json:"zone_id"`
	Zone_name                    string                        `json:"zone_name"`
	Persistent_profile_path      string                        `json:"persistent_profile_path"`
	Filter_policy_id             string                        `json:"filter_policy_id"`
	Filter_policy_name           string                        `json:"filter_policy_name"`
	Filter_policy_force_disabled bool                          `json:"filter_policy_force_disabled"`
	Session_time_limit           int64                         `json:"session_time_limit"`
	Categories                   []string                      `json:"categories"`
	Default_category             string                        `json:"default_category"`
	Allow_network_selection      bool                          `json:"allow_network_selection"`
	Require_gpu                  bool                          `json:"require_gpu"`
	Gpu_count                    int                           `json:"gpu_count"`
	Hidden                       bool                          `json:"hidden"`
	Notes                        string                        `json:"notes"`
	Image_type                   string                        `json:"image_type"`
	Server_pool_id               string                        `json:"server_pool_id"`
	Link_url                     string                        `json:"link_url"`
	Cpu_allocation_method        string                        `json:"cpu_allocation_method"`
	Uncompressed_size_mb         int                           `json:"uncompressed_size_mb"`
	Server                       map[string]any                `json:"server"`
	Server_pool                  map[string]any                `json:"server_pool"`
}

type imagesImageAttributesModel struct {
	Image_id string `json:"image_id"`
	Attr_id  string `json:"attr_id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Value    string `json:"value"`
}

type volumeMappingModel struct {
	Bind       string `json:"bind"`
	Mode       string `json:"mode"`
	Uid        int    `json:"uid"`
	Gid        int    `json:"gid"`
	Required   bool   `json:"required"`
	Skip_check bool   `json:"skip_check"`
}
