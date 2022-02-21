class Node:
    def __init__(self, value, next_=None):
        self.value = value
        self.next = next_


a = Node("A")
b = Node("B")
c = Node("C")
d = Node("D")

a.next = b
b.next = c
c.next = d


def get_node_value_rec(head, index):
    if not head.next:
        return None
    if index == 0:
        return head.value
    return get_node_value_rec(head.next, index - 1)


def get_node_value(head, index):
    current = head
    count = 0
    while current:
        if count == index:
            return current.value
        count += 1
        current = current.next


if __name__ == "__main__":
    r = get_node_value(a, 2)
    assert r == "C"
