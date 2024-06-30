package MyFunc

import "fmt"

func InitializeDefinition(weaponCategory *WeaponCategory) {
	weaponType := map[uint8]string{
		1: "mele",
		2: "shield",
		3: "staff",
		4: "chime",
	}
	scalingVal := map[string]float32{
		"S": 3,
		"A": 2,
		"B": 1.5,
		"C": 1.25,
		"D": 1,
		"E": 0,
		"-": 0,
	}
	weaponCategory.weaponType = weaponType
	weaponCategory.scalingVal = scalingVal
}

type Inventory struct {
	weapons []WeaponValue
}

func (v WeaponValue) New(droppedWeapon *WeaponValue) {

	droppedWeapon.name = ""
	droppedWeapon.weapon_type = "mele"
	droppedWeapon.damage = 22
	droppedWeapon.damage_reduction = 0
	droppedWeapon.damage_type = "raw"
	droppedWeapon.faith_scaling = "E"
	droppedWeapon.int_scaling = "E"
	droppedWeapon.dex_scaling = "B"
	droppedWeapon.str_scaling = "C"

}

func (i Inventory) InitializeInv(Player_inv *Inventory) {
	var droppedWeapon WeaponValue
	droppedWeapon.New(&droppedWeapon)
	Player_inv.weapons = append(Player_inv.weapons, droppedWeapon)
	fmt.Println(Player_inv.weapons)

}

func InitializeData(Player_inv *Inventory, weaponCategory *WeaponCategory) {
	InitializeDefinition(weaponCategory)
	Player_inv.InitializeInv(Player_inv)
}
