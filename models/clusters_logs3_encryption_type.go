/*
 * Databricks
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package models

type ClustersLogS3EncryptionType string

// List of ClustersAwsAvailability
const (
	SSES3               ClustersLogS3EncryptionType = "sse-s3"
	SSEKMS          ClustersLogS3EncryptionType = "sse-k"
)
