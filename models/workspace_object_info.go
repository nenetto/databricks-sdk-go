/*
 * Databricks
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package models

type WorkspaceObjectInfo struct {
	ObjectType *WorkspaceObjectType `json:"object_type,omitempty"`

	Path string `json:"path,omitempty"`

	Language *WorkspaceLanguage `json:"language,omitempty"`
}
