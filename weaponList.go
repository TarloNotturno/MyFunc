package MyFunc

func ListOfRawWeapon() []WeaponValue {
	var init_weapons []WeaponValue
	var current_weapon WeaponValue
	current_weapon.name = "Fire Knight's Greatsword"
	current_weapon.weapon_type = "mele"
	current_weapon.damage = 147
	current_weapon.damage_reduction = 0
	current_weapon.damage_type = "raw"
	current_weapon.faith_scaling = "-"
	current_weapon.int_scaling = "-"
	current_weapon.dex_scaling = "D"
	current_weapon.str_scaling = "D"
	init_weapons = append(init_weapons, current_weapon)
	current_weapon.name = "Ancient Meteoric Ore Greatsword"
	current_weapon.weapon_type = "mele"
	current_weapon.damage = 138
	current_weapon.damage_reduction = 0
	current_weapon.damage_type = "raw"
	current_weapon.faith_scaling = "-"
	current_weapon.int_scaling = "-"
	current_weapon.dex_scaling = "E"
	current_weapon.str_scaling = "D"
	init_weapons = append(init_weapons, current_weapon)

	current_weapon.name = "Greatsword"
	current_weapon.weapon_type = "mele"
	current_weapon.scalingVal = "D"
	current_weapon.damage = 164
	current_weapon.damage_reduction = 0
	current_weapon.damage_type = "raw"
	current_weapon.faith_scaling = "-"
	current_weapon.int_scaling = "-"
	current_weapon.dex_scaling = "E"
	current_weapon.str_scaling = "C"
	init_weapons = append(init_weapons, current_weapon)
	return init_weapons

}
