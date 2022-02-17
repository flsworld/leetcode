def is_palindrome(s):
    for i in range(len(s)):
        if not (s[i] == s[::-1][i]):
            return False
    return True


def longest_palindrome(s):
    candidates = []
    for i in range(len(s)):
        st = s[i:]
        while st:
            if is_palindrome(st):
                candidates.append(st)
                break
            else:
                st = st[:-1]
    return sorted(candidates, key=lambda x: len(x), reverse=True)[0]


def longestPalindrome(s):
    ans = ""
    while len(s) > len(ans):
        for i in range(len(s) - 1, -1, -1):
            if (
                s[0] == s[i]
                and s[0 : i + 1] == s[i::-1]
                and len(s[0 : i + 1]) > len(ans)
            ):
                ans = s[0 : i + 1]
        s = s[1:]
    return ans


if __name__ == "__main__":
    sample = "babad"
    r = longest_palindrome(sample)
    print(r)

    longestPalindrome(sample)