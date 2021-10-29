package usecase

// @Created 29/10/2021
// @Updated
func SequenceExists(data []int, input []int) bool {

	if len(input) > len(data) {
		return false
	}

	prevSequenceIndex := 0
	currSequenceIndex := 0
	
	for inputIndex, in := range input {

		for dataIndex, d := range data {
			if in == d {					
				currSequenceIndex = dataIndex

				if(inputIndex > 0){
					if(prevSequenceIndex + 1) != currSequenceIndex {
						return false
					}	
				}

				if(inputIndex == len(input)-1) {
					return true
				}

				break
			}
		}

		prevSequenceIndex = currSequenceIndex
	}

	return false
}