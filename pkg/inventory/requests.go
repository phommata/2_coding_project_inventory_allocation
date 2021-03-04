package inventory

import "2_coding_project_inventory_allocation/pkg/entities"

type InventoryAllocationRequest struct {
	Header int `json:"Header"`
	Lines  []entities.Product `json:"Lines"`
}