/*
 * Databricks
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package models

type GroupsListParentsRequest struct {
	UserName string `json:"user_name,omitempty"`

	GroupName string `json:"group_name,omitempty"`

	ParentName string `json:"parent_name,omitempty"`
}
