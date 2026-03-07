#!/usr/bin/env python3


def factorial(x):
    if x == 0:
        return 1
    y = x
    for i in range(x - 1, 0, -1):
        y *= i
    return y


xs = [
    10**6,
    60 * 10**6,
    3600 * 10**6,
    24 * 3600 * 10**6,
    30 * 24 * 3600 * 10**6,
    365 * 24 * 3600 * 10**6,
    100 * 365 * 24 * 3600 * 10**6,
]

if __name__ == "__main__":
    cache = {}
    xs[-1]
    for i in range(2, 100):
        y = factorial(i)
        cache[i] = y
        if y > xs[-1]:
            break

    left = sorted(list(cache.items()))[:-1]
    right = sorted(list(cache.items()))[1:]
    pairs = [(l, r) for (l, r) in zip(left, right)]

    for x in xs:
        for (lk, lv), (rk, rv) in pairs:
            if x >= lv and x < rv:
                print(f"{x} in [{lv},{rv}] for {lk}!..{rk}!")
                break
