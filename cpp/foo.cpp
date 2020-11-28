#include <stdio.h>
#include <math.h>
#include <stdlib.h>
#include <x86intrin.h>

void AddTwoInts(int a, int b, int* out) {
    *out = a + b;
}

void SqrtAvx256(float* in, float* out) {
    __m256 v1 = _mm256_load_ps(in);
    __m256 v2 = _mm256_sqrt_ps(v1);
    _mm256_storeu_ps(out, v2);
}
