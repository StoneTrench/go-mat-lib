package path

import (
	"errors"
	"math"
	"slices"
)

func remove[T any](slice []T, i int) []T {
	if i < 0 || i >= len(slice) {
		return slice
	}

	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func process_path[T any, G any](came_from map[string]T, curr T, goal G, get_id func(goal G, s T) string) []T {
	path := []T{}
	for {
		path = append(path, curr)
		next, exists := came_from[get_id(goal, curr)]
		if !exists {
			break
		}
		curr = next
	}
	slices.Reverse(path)
	return path
}

func GenericPathfinder[T any, G any, V float32 | float64 | int32 | int64](
	start T,
	goal G,
	max_iterations int,
	get_id func(goal G, s T) string,
	is_reached func(goal G, s T) bool,
	get_neighbours func(goal G, s T) []T,
	get_score_d func(goal G, s T, n T) V,
	get_score_h func(goal G, s T) V,
) (path []T, err error) {
	INFTY := V(math.Inf(+1))

	_start_id := get_id(goal, start)
	_start_f_score := get_score_h(goal, start)

	open_set_contains := map[string]bool{
		_start_id: true,
	}
	open_set_data := []T{start}
	open_set_id := []string{_start_id}

	came_from := make(map[string]T)

	g_score := map[string]V{
		_start_id: 0,
	}
	f_score := map[string]V{
		_start_id: _start_f_score,
	}

	lowest_h := _start_f_score
	lowest_h_state := start

	for len(open_set_data) > 0 && max_iterations > 0 {
		max_iterations--
		var curr T
		curr_fscore := INFTY
		var curr_index int = -1
		var curr_id string

		for i, data := range open_set_data {
			id := open_set_id[i]
			v_fscore, exists := f_score[id]
			if !exists {
				v_fscore = INFTY
			}

			if v_fscore <= curr_fscore {
				curr = data
				curr_fscore = v_fscore
				curr_index = i
				curr_id = id
			}
		}

		if curr_index == -1 {
			break
		}

		if is_reached(goal, curr) {
			return process_path(came_from, curr, goal, get_id), nil
		}

		open_set_data = remove(open_set_data, curr_index)
		open_set_id = remove(open_set_id, curr_index)
		open_set_contains[curr_id] = false

		curr_gscore := g_score[curr_id]
		for _, neigh := range get_neighbours(goal, curr) {
			neigh_id := get_id(goal, neigh)

			tentative_g_score := curr_gscore + get_score_d(goal, curr, neigh)

			neigh_gscore := INFTY
			if g, exists := g_score[neigh_id]; exists {
				neigh_gscore = g
			}

			if tentative_g_score < neigh_gscore {
				came_from[neigh_id] = curr
				g_score[neigh_id] = tentative_g_score

				neigh_hscore := get_score_h(goal, neigh)
				neigh_fscore := tentative_g_score + neigh_hscore

				f_score[neigh_id] = neigh_fscore
				if !open_set_contains[neigh_id] {
					open_set_data = append(open_set_data, neigh)
					open_set_id = append(open_set_id, neigh_id)
					open_set_contains[neigh_id] = true

					if lowest_h >= neigh_hscore {
						lowest_h = neigh_hscore
						lowest_h_state = neigh
					}
				}
			}
		}
	}

	return process_path(came_from, lowest_h_state, goal, get_id), errors.New("partial path found")
}
