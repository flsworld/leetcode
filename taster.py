from collections import Counter, defaultdict


def main(word_list):
    single_products = set()
    res = defaultdict(list)
    for item in word_list:
        if "-" not in item:
            single_products.add(item)
            res[1].append(item)
        else:
            if not any(w for w in single_products if w in item):
                single_products.add(item)
                res[1].append(item)
            else:
                print(single_products)
                cpt = 0
                for x in single_products:
                    if x in item:
                        cpt += 1
                res[cpt].append(item)
    return res


if __name__ == "__main__":
    sample = [
        "bread-orange-soap",
        "paper-chocolate",
        "chocolate",
        "hot-dog-milk",
        "hot-dog",
        "finger-nail-bread",
        "bread",
        "orange",
        "tee-shirt-bread",
        "soap",
        "paper",
        "tee-shirt",
        "milk",
        "finger-nail",
        "orange-bread-hot-dog-milk",
        "tee-shirt-hot-dog",
        "bread",
    ]

    s = sorted(sample, key=lambda x: Counter(x)["-"])

    r = main(s)
    print(r)
