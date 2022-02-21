from typing import List, Any


def execute(lists: List[List[Any]], proc) -> List[List[Any]]:
    out = []
    if not isinstance(lists, list):
        raise ValueError("Can't sort it")
    for list_ in lists:
        if not isinstance(list_, list):
            raise ValueError("Can't sort it")
        types = []
        for element in list_:
            if types and type(element) not in types:
                raise ValueError("Wrong type")
            types.append(type(element))
        out.append(proc(list_))
    return out


def custom_max(list1: List):
    if not list1:
        return None
    return max(list1)


def max_lists(lists: List[List[Any]]) -> List[Any]:
    return execute(lists, custom_max)


def sort_lists(lists: List[List[Any]]) -> List[Any]:
    return execute(lists, sorted)


if __name__ == "__main__":
    input = [[3, 1, 2], ["a", "c", "b"], []]
    expected = [[1, 2, 3], ["a", "b", "c"], [], "a"]
    expected = [3, "c", None]
    r = sort_lists(input)
    r = max_lists(input)
    print(r)
