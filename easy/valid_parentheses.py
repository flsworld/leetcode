def is_valid(s):
    hashmap = {")": "(", "]": "[", "}": "{"}
    stack = []
    for st in s:
        if st in ["(", "[", "{"]:
            stack.append(st)
        elif stack:
            item = stack.pop()
            if item != hashmap[st]:
                return False
        else:
            return False

    if not stack:
        return True
    else:
        return False


if __name__ == "__main__":
    sample = "()[]{}"
    is_valid(sample)
