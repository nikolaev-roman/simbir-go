package repository

const (
	searchToRentWithType = `
	SELECT *
	FROM transports
	WHERE ACOS(
		SIN(radians(latitude)) * SIN(radians(?)) + 
		COS(radians(latitude)) * COS(radians(?)) * 
		COS(radians(longitude) - radians(?))
	) * 6380 < ?
		and transport_type = ?
`
	searchToRentWithoutType = `
SELECT *
FROM transports
WHERE ACOS(
	SIN(radians(latitude)) * SIN(radians(?)) + 
	COS(radians(latitude)) * COS(radians(?)) * 
	COS(radians(longitude) - radians(?))
) * 6380 < ?
`
)
