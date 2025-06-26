def longestCommonPrefix(strs) -> str:
    l = []
    for i in zip(*strs):
        if (len(set(i))) > 1:
            break
        l.append(i[0])
    return "".join(l)


if __name__ == "__main__":
    strings = ["flower", "flow", "flight"]
    longestCommonPrefix(strings)
