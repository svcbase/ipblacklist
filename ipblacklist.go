package ipblacklist

var ipstart []uint32
var ipterminus []uint32

func init() {
	ipstart = make([]uint32, 0, 10)
	ipterminus = make([]uint32, 0, 10)
}

func AddSegment(start, terminus uint32) {
	ipstart = append(ipstart, start)
	ipterminus = append(ipterminus, terminus)
}

func Inblacklist(ip uint32) (flag bool) { //need optimize algorithm
	flag = false
	if ip == 2130706433 { //127.0.0.1 - localhost
		return
	}
	for i, start := range ipstart {
		terminus := ipterminus[i]
		if ip >= start && ip <= terminus {
			flag = true
			break
		}
	}
	return
}
