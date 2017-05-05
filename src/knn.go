package domino

/*
	k-Nearest Neighbors
	------------------------------
	The kNN algorithm can be used for regression or classification problems.

	In classification problems, aggregate attribute proximity

	Pseudocode (Classification):
		kNN (ClassifiedDataSet, Sample) {
			1. Traverse every data point in a 2D or 3D space
			2. Generate distance from sample
			3. Select the k closest samples via distance metric (see distance.go)
			4. Elect winner from labels, typically via vote
			5. In ties, implement tie-breaking rubric
		}

	Pseudocode (Regression):

*/

func NewKNN() Learner {
	var l Learner
	l.Type = "KNN"
	return l
}

func (l *Learner) KNN() {

}
