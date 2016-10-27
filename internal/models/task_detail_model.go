package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*TaskDetailModel task detail model

swagger:model TaskDetailModel
*/
type TaskDetailModel struct {

	/* cluster arn

	Required: true
	*/
	ClusterArn *string `json:"clusterArn"`

	/* container instance arn

	Required: true
	*/
	ContainerInstanceArn *string `json:"containerInstanceArn"`

	/* containers

	Required: true
	*/
	Containers []*TaskDetailContainerModel `json:"containers"`

	/* created at

	Required: true
	*/
	CreatedAt *string `json:"createdAt"`

	/* desired status

	Required: true
	*/
	DesiredStatus *string `json:"desiredStatus"`

	/* last status

	Required: true
	*/
	LastStatus *string `json:"lastStatus"`

	/* overrides

	Required: true
	*/
	Overrides *TaskDetailOverridesModel `json:"overrides"`

	/* started at
	 */
	StartedAt string `json:"startedAt,omitempty"`

	/* started by
	 */
	StartedBy string `json:"startedBy,omitempty"`

	/* stopped at
	 */
	StoppedAt string `json:"stoppedAt,omitempty"`

	/* stopped reason
	 */
	StoppedReason string `json:"stoppedReason,omitempty"`

	/* task arn

	Required: true
	*/
	TaskArn *string `json:"taskArn"`

	/* task definition arn

	Required: true
	*/
	TaskDefinitionArn *string `json:"taskDefinitionArn"`

	/* updated at

	Required: true
	*/
	UpdatedAt *string `json:"updatedAt"`

	/* version

	Required: true
	*/
	Version *int32 `json:"version"`
}

// Validate validates this task detail model
func (m *TaskDetailModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterArn(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateContainerInstanceArn(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateContainers(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDesiredStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLastStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOverrides(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTaskArn(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTaskDefinitionArn(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaskDetailModel) validateClusterArn(formats strfmt.Registry) error {

	if err := validate.Required("clusterArn", "body", m.ClusterArn); err != nil {
		return err
	}

	return nil
}

func (m *TaskDetailModel) validateContainerInstanceArn(formats strfmt.Registry) error {

	if err := validate.Required("containerInstanceArn", "body", m.ContainerInstanceArn); err != nil {
		return err
	}

	return nil
}

func (m *TaskDetailModel) validateContainers(formats strfmt.Registry) error {

	if err := validate.Required("containers", "body", m.Containers); err != nil {
		return err
	}

	for i := 0; i < len(m.Containers); i++ {

		if swag.IsZero(m.Containers[i]) { // not required
			continue
		}

		if m.Containers[i] != nil {

			if err := m.Containers[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *TaskDetailModel) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("createdAt", "body", m.CreatedAt); err != nil {
		return err
	}

	return nil
}

func (m *TaskDetailModel) validateDesiredStatus(formats strfmt.Registry) error {

	if err := validate.Required("desiredStatus", "body", m.DesiredStatus); err != nil {
		return err
	}

	return nil
}

func (m *TaskDetailModel) validateLastStatus(formats strfmt.Registry) error {

	if err := validate.Required("lastStatus", "body", m.LastStatus); err != nil {
		return err
	}

	return nil
}

func (m *TaskDetailModel) validateOverrides(formats strfmt.Registry) error {

	if err := validate.Required("overrides", "body", m.Overrides); err != nil {
		return err
	}

	if m.Overrides != nil {

		if err := m.Overrides.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *TaskDetailModel) validateTaskArn(formats strfmt.Registry) error {

	if err := validate.Required("taskArn", "body", m.TaskArn); err != nil {
		return err
	}

	return nil
}

func (m *TaskDetailModel) validateTaskDefinitionArn(formats strfmt.Registry) error {

	if err := validate.Required("taskDefinitionArn", "body", m.TaskDefinitionArn); err != nil {
		return err
	}

	return nil
}

func (m *TaskDetailModel) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updatedAt", "body", m.UpdatedAt); err != nil {
		return err
	}

	return nil
}

func (m *TaskDetailModel) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}
