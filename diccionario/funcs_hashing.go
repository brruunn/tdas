package diccionario

// Hashing de suma o modular
func hashingSuma(clave []byte, largo int) int {
	var suma uint32
	for _, c := range clave {
		suma += uint32(c)
	}
	return int(suma % uint32(largo))
}

// Hashing de FNV-1a-64
func hashingFNV(clave []byte, largo int) int {
	var h uint64 = 14695981039346656037
	for _, c := range clave {
		h *= 1099511628211
		h ^= uint64(c)
	}
	return int(h % uint64(largo))
}

// Hashing de Jenkins One-at-a-time
func hashJenkins(clave []byte, largo int) int {
	var h uint32
	for i := range clave {
		h += uint32(clave[i])
		h += (h << 10)
		h ^= (h >> 6)
	}
	h += (h << 3)
	h ^= (h >> 11)
	h += (h << 15)
	return int(h % uint32(largo))
}
