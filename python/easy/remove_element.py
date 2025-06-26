from typing import Optional


def remove_element(nums: list[Optional[int]], val: int) -> int:
    count = 0
    for i in range(len(nums)):
        if nums[i] == val:
            nums[i] = None
            count += 1

    nums.sort(key=lambda x: (x is None, x))
    nums = nums[:-count]
    return count, nums


def removeElement(nums: list[int], val: int) -> int:
    count = 0
    imax = len(nums)
    i = 0
    while i < imax:
        if nums[i] == val:
            count += 1
            nums[i] = nums[imax - 1]
            imax -= 1
        else:
            i += 1
    return len(nums) - count


def removeElement(nums: list[int], val: int) -> int:
    c = nums.count(val)
    for i in range(0, c):
        nums.remove(val)
    return len(nums) - c


def removeElement(self, nums, val):
    i = 0
    for j in range(len(nums)):
        if nums[j] != val:
            nums[i] = nums[j]
            i += 1
    return i


def removeElement(nums, val):
    i: int = 0
    n = len(nums)
    while i < n:
        if nums[i] == val:
            nums[i] = nums[n - 1]
            # reduce array size by one
            n -= 1
        else:
            i += 1
    return n


if __name__ == "__main__":
    sample = [0, 1, 2, 2, 3, 0, 4, 2]
    res = remove_element(sample, 2)
    print(res)
    assert res[1] == [0, 0, 1, 3, 4]

    sample = [3, 2, 2, 3]
    res = remove_element(sample, 3)
    print(res)
    assert res[1] == [2, 2]
