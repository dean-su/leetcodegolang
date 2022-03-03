/*
Medium

There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1. You are given an array prerequisites where prerequisites[i] = [ai, bi] indicates that you must take course bi first if you want to take course ai.

For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.
Return true if you can finish all courses. Otherwise, return false.

 

Example 1:

Input: numCourses = 2, prerequisites = [[1,0]]
Output: true
Explanation: There are a total of 2 courses to take. 
To take course 1 you should have finished course 0. So it is possible.
Example 2:

Input: numCourses = 2, prerequisites = [[1,0],[0,1]]
Output: false
Explanation: There are a total of 2 courses to take. 
To take course 1 you should have finished course 0, and to take course 0 you should also have finished course 1. So it is impossible.
*/

func canFinish(numCourses int, prerequisites [][]int) bool {
    if len(prerequisites) == 0 {return true}
    
    inDegree := make([]int, numCourses)
    graph := make([][]int, numCourses)
    for _, v := range prerequisites{
        inDegree[v[0]]++
        graph[v[1]] = append(graph[v[1]], v[0])
    }
    
    q := make([]int, 0)
    for i := 0; i < len(inDegree); i++ {
        if inDegree[i] == 0 {
            q = append(q, i)
            inDegree[i]--
        }
        
    }
    if len(q) == 0 {return false}
    c := 0
    for len(q) != 0 {
        curr := q[0]
        q = q[1:]
        c++
        for _, v := range graph[curr] {
            inDegree[v]--
            if inDegree[v] == 0 {
                q = append(q, v)
            }
        }
    }
    if c == numCourses {return true}
    return false
}