# you can write to stdout for debugging purposes, e.g.
# print("this is a debug message")

def is_in_range(value: int):
    return 0 <= value < 2**20


def push(stack: list, element: int):
    if is_in_range(element):
        stack.append(element)
    return stack


def dup(stack: list):
    temp = stack.copy()
    last = temp.pop()
    stack.append(last)
    return stack


def add(stack: list):
    if len(stack) <= 1:
        return None
    last = stack.pop()
    penultimate = stack.pop()
    sum_ = int(penultimate) + int(last)
    if not is_in_range(sum_):
        return None
    stack.append(str(sum_))
    return stack


def substract(stack: list):
    if len(stack) <= 1:
        return None
    last = stack.pop()
    penultimate = stack.pop()
    diff = int(last) - int(penultimate)
    if not is_in_range(diff):
        return None
    stack.append(str(diff))
    return stack


def solution(s):
    s = s.split()
    stack = []
    for element in s:
        if element == "+":
            stack = add(stack)
        elif element == "-":
            stack = substract(stack)
        elif element == "DUP":
            dup(stack)
        elif element == "POP":
            stack.pop()
        else:
            push(stack, int(element))

    if stack is None:
        return -1
    topmost = stack.pop()
    return int(topmost)


if __name__ == "__main__":
    sample = "1048575 DUP +"
    s = sample.split()
    print(solution(s))
