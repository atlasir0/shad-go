package hogwarts

func GetCourseList(prerequisites map[string][]string) []string {
	inDegree := make(map[string]int)

	adjList := make(map[string][]string)

	for course, prereqs := range prerequisites {
		if _, exists := inDegree[course]; !exists {
			inDegree[course] = 0
		}
		for _, prereq := range prereqs {
			if _, exists := inDegree[prereq]; !exists {
				inDegree[prereq] = 0
			}
			inDegree[course]++
			adjList[prereq] = append(adjList[prereq], course)
		}
	}

	// Queue for BFS
	queue := []string{}
	for course, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, course)
		}
	}

	result := []string{}

	for len(queue) > 0 {
		course := queue[0]
		queue = queue[1:]
		result = append(result, course)

		for _, nextCourse := range adjList[course] {
			inDegree[nextCourse]--
			if inDegree[nextCourse] == 0 {
				queue = append(queue, nextCourse)
			}
		}
	}

	if len(result) != len(inDegree) {
		panic("cyclic dependency detected")
	}

	return result
}
