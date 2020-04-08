package models

type InstanceprofileInfo struct {
	InstanceProfileArn string `json:"instance_profile_arn,omitempty"`
	IsMetaInstanceProfile bool `json:"is_meta_instance_profile,omitempty"`
}