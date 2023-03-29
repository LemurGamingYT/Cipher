from time import perf_counter


def code():
    i = 0
    while i < 50000000:
        i += 1


st = perf_counter()

code()

et = perf_counter()

print("Time to Execute: " + str(round((et - st) * 1000, 4)) + " ms")
