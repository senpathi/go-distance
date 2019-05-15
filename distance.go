package go_distance

import "math"


//distanceFromPythagoras function returns the distance (in meters) between two points of
//     a given longitude and latitude relatively accurately through the pythagoras formula with accuracy for small distances
//
//https://www.movable-type.co.uk/scripts/latlong.html

func distanceFromPythagoras(lat1, lon1, lat2, lon2 float64) float64{

	var d_ew, d_ns, d_lu, r float64
	d_ew = (toRad(lon2 - lon1)) * math.Cos((toRad(lat1 + lat2))/2)
	d_ns = toRad(lat2 - lat1)
	d_lu = math.Sqrt(d_ew * d_ew + d_ns * d_ns)

	r = 6371000 //Mean radius of planet Earth

	d_mi := r * d_lu

	return d_mi
}


//distanceFromHSin function returns the distance (in meters) between two points of
//     a given longitude and latitude relatively accurately (using a spherical
//     approximation of the Earth) through the Haversin Distance Formula for
//     great arc distance on a sphere with accuracy for small distances
//
// http://en.wikipedia.org/wiki/Haversine_formula
// https://gist.github.com/cdipaolo/d3f8db3848278b49db68

func distanceFromHaversine(lat1, lon1, lat2, lon2 float64) float64{

	var la1, lo1, la2, lo2, r float64
	la1 = toRad(lat1)
	lo1 = toRad(lon1)
	la2 = toRad(lat2)
	lo2 = toRad(lon2)

	r = 6378100 // Earth radius in METERS (nominal "zero tide" equatorial)

	h := math.Pow(math.Sin((la2-la1)/2), 2) + math.Cos(la1)*math.Cos(la2)*math.Pow(math.Sin((lo2-lo1)/2), 2)

	return 2 * r * math.Asin(math.Sqrt(h))
}


//distanceFromCosine function returns the distance (in meters) between two points of
//     a given longitude and latitude relatively accurately Spherical Law of Cosines with accuracy for small distances
//
// https://www.movable-type.co.uk/scripts/latlong.html

func distanceFromCosine(lat1, lon1, lat2, lon2 float64) float64{

	var la1, la2, diffLon, r float64
	la1 = toRad(lat1)
	la2 = toRad(lat2)
	diffLon = toRad(lon2-lon1)

	r = 6371000 //Mean radius of planet Earth

	d := math.Acos( math.Sin(la1)*math.Sin(la2) + math.Cos(la1)*math.Cos(la2) * math.Cos(diffLon) ) * r

	return d
}

func toRad(a float64) float64{
	return a * math.Pi / 180
}
