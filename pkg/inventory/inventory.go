// SPDX-License-Identifier: BSD-3-Clause

package inventory

import "github.com/google/uuid"

type Location struct {
	Building string
	Room     string
	Rack     string
	Slot     string
}

type Machine struct {
	ID           uuid.UUID
	Location     Location
	Hostname     string
	IP           string
	SerialNumber string
	ProductName  string
	Manfuacturer string
	AssetTag     string
}

type Component struct {
	ID           uuid.UUID
	SerialNumber string
	ProductName  string
	Manufacturer string
	AssetTag     string
}

type Inventory struct {
	Machines   map[uuid.UUID]Machine
	Components map[uuid.UUID]Component
}

func (i *Inventory) AddMachine(m Machine) {
	i.Machines[m.ID] = m
}

func (i *Inventory) RemoveMachine(id uuid.UUID) {
	delete(i.Machines, id)
}

func (i *Inventory) AddComponent(c Component) {
	i.Components[c.ID] = c
}

func (i *Inventory) RemoveComponent(id uuid.UUID) {
	delete(i.Components, id)
}

func (i *Inventory) GetMachine(id uuid.UUID) (Machine, bool) {
	m, ok := i.Machines[id]
	return m, ok
}

func (i *Inventory) GetComponent(id uuid.UUID) (Component, bool) {
	c, ok := i.Components[id]
	return c, ok
}
