package go_distance

import "math"


//https://stackoverflow.com/questions/1664799/calculating-distance-between-two-points-using-pythagorean-theorem?answertab=votes#tab-top

func distanceFromPythagoras(lat1, lon1, lat2, lon2 float64) float64{
	var d_ew, d_ns, d_lu, r float64
	d_ew = (toRad(lon2 - lon1)) * math.Cos((toRad(lat1))/2)
	d_ns = toRad(lat2 - lat1)
	d_lu = math.Sqrt(d_ew * d_ew + d_ns * d_ns)

	r = 6371000 //Mean radius of planet Earth

	d_mi := r * d_lu

	return d_mi
}


func hsin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}

// Distance function returns the distance (in meters) between two points of
//     a given longitude and latitude relatively accurately (using a spherical
//     approximation of the Earth) through the Haversin Distance Formula for
//     great arc distance on a sphere with accuracy for small distances
//
// point coordinates are supplied in degrees and converted into rad. in the func
//
// distance returned is METERS!!!!!!
// http://en.wikipedia.org/wiki/Haversine_formula

func distanceFromHSin(lat1, lon1, lat2, lon2 float64) float64{
	// convert to radians
	// must cast radius as float to multiply later
	var la1, lo1, la2, lo2, r float64
	la1 = toRad(lat1)
	lo1 = toRad(lon1)
	la2 = toRad(lat2)
	lo2 = toRad(lon2)

	r = 6378100 // Earth radius in METERS (nominal "zero tide" equatorial)

	// calculate
	h := hsin(la2-la1) + math.Cos(la1)*math.Cos(la2)*hsin(lo2-lo1)

	return 2 * r * math.Asin(math.Sqrt(h))
}

func toRad(a float64) float64{
	return a * math.Pi / 180
}
