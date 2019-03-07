


// # this file contains the three various mitigation techniques
// # that can be used by the defense layer, all traffic from the 
// # diagnosis layer comes into the mitigation layer which 
// # mitigates and creates flow rules etc.



function static_mitigation() {
	// # add rules to controller to drop the attack traffic
	// # tags traffic as attack or legitimate
}

	





function dynamic_mitigation(traffic_stats) {

	for i,j in ingress_locations.items():
		
		if(j["current_load"] > j["peak_load"])
			j["peak_load"] = current_load

		peak_load = j["peak_load"]

		if(j["current_load"] < j["min_load"])
			j["min_load"] = current_load

		min_load = j["min_load"]

		j["capacity"] = random.uniform(min_load,max_load)


	// # set flow rules for new capacity in the controller
	// # create packet rules for attack packers which is essentially the static mitigation

}
def adaptive_mitigation(traffic_stats):


	



