package MyFunc

type WeaponValue struct {
	name             string
	weapon_type      string
	scalingVal       string
	damage           uint8
	damage_reduction uint8
	damage_type      string
	faith_scaling    string
	int_scaling      string
	dex_scaling      string
	str_scaling      string
}

type WeaponCategory struct {
	weaponType map[uint8]string
	scalingVal map[string]float32
}
