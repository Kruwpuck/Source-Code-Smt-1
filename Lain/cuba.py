def a(n):
    if n == 0: return 0
    if n == 1: return 4
    else: return 6*a(n-1)-5*a(n-2)
print(a(20))