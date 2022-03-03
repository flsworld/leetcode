ayaah = {
    0: 5,
    1: 2,
    2: 4,
    3: 6,
    4: 3,
    5: 7,
}


def is_adjacent(p: int, q: int):
    return q - p <= 1


def solution(A: list):
    hashmap = {}
    for x, y in zip(range(1, len(A) - 1), A[1:-1]):
        hashmap[x] = y

    sorted_cost = sorted(hashmap.items(), key=lambda x: x[1])
    min_cost = sorted_cost[-1][1] + sorted_cost[-2][1]
    started = False
    for x in range(len(hashmap)):
        pivot, pivot_cost = sorted_cost[x]
        for position, position_cost in sorted_cost:
            if is_adjacent(pivot, position):
                continue
            cost = position_cost + pivot_cost
            if started and cost > min_cost:
                break
            if cost < min_cost:
                min_cost = cost
            break
        started = True

    return min_cost


if __name__ == "__main__":
    sample = [5, 2, 4, 6, 3, 7]
    res = solution(sample)
    print(res)
