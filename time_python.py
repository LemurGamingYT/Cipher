from time import perf_counter


def code():
    pass


st = perf_counter()

code()

et = perf_counter()

print("Time to Execute: " + str(round((et - st) * 1000, 4)) + " ms")
