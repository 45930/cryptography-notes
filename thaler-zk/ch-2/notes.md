## Reed-Solomon Fingerprinting

This algorithm is suitable for cooperative parties which trust each other but want to minimize communication to each other - e.g. see if 2 files are equal without sending the entire contents of one file to the other party.  Instead we can just the fingerprint and the hash function used to calculate the fingerprint, chosen at random from a family of hash functions.

*   Pick a prime much larger than the size of your input

*   Convert your input to integers (e.g. ASCII encoding would be integers 1...128)

    *   Input becomes a vector of length n such as \[1, 12, 3, 5, 19]

    *   Call this vector a

    *   Think of a as the coefficients of a polynomial of degree n-1

*   Define the **Field** F<sub>p</sub> as the set of integers modulo p

*   We then call the family of hash functions *H* the set of all functions which evaluate the output of our polynomial for the different values of r E F<sub>p</sub>

*   If Alice and Bob want to evaluate if 2 files are the same, Alice selects r at random, evaluates H<sub>r</sub>(a) `(evaluates the polynomial with coefficients a at x=r)`, then sends r and H<sub>r</sub>(a) to Bob who can use r to compute the same on his file and compare the result.

*   The algorithm has an error term of 1 - ((n-1) / p) | given p > n^2, we can say the error is < (1 - 1/n)

EX/

Alphabet: A = 1; B = 2; C = 3;

Message: CAB

a: \[3, 1, 2] (n=3)

p = 11

r (chosen at random) = 3

Evaluation: 3(3)^2 + 1(3) + 2 = 27 + 3 + 2 = 32 % 11 = 10 = H<sub>3</sub>(a)

Different message: BAC

b: \[2, 1,  3]

Evaluation: 2(3)^2 + 1(3) + 3 = 18 + 3 + 3 = 24 % 11 = 2 = H<sub>3</sub>(b)


## Freivald's Algorithm

One expression of Reed-Solomon is Freivald's algortihm for determining if the product of 2 nXn matrices (A, B) is equal to a claimed solution (C).

Let r be an element of F<sub>p</sub>, then let x be a vector \[r^0, r^1, r^2, ..., r^n-1], then check if A \* Bx == Cx

The intuition for this is that Cx is essentially the Reed Solomon fingerprint of C at point r in F<sub>p</sub>.  If A \* B is indeed equal to C, then we can say with high probability that the reed solomon fingerprint will also be equal, and will only be equal if the equality of the matrices holds.

### Runtime

This algorithm runs in O(n^2) given that an nxn matrix multiplied by a vector x is done in O(n) and we do it 3 times - Cx, Bx `(call this w)`, and Aw.

### Completeness and Soundness

We know that this works because the operation of evaluating Cx = A . Bx is the same as saying C is a polynomial evaluated at r and A . Bx is a polynomial evaluated at r and the properties of Reed-Solomon can be applied.

### Distance Amplifying Property

Take example vectors \[2, 1, 1] and \[2, 1, 0].  The vectors differ on 1 value.  If we evaluate the reed-solomon encoding on these vectors over F\<sub>11\</sub> then we can expect the full fingerprint to differ on about 80% of the values.  The formula is `1 - (n-1)/p`.  1 - (3-1)/11 = 1 - 2/11 \~ 80%.  The intuition for why there are more differences in the encodings is that the encodings are much larger than the inputs.  The input is size n (3) and the encoding is size p >> n (11).  Thus, differences are amplified in the output.

In this example, if the encoding evaluated at any randomly chosen r E F\<sub>p\</sub> is equal for 2 input vectors of size 3, then we can be 80% sure they are the same input.

EX/

a: \[2, 1, 1]

b: \[2, 1, 0]

enc(a): \[2, 4, 8, 3, 0, 10, 0, 3, 8, 0, 2]

enc(b): \[2, 3, 4, 5, 6, 7, 8, 9, 10, 0, 1]

same output: 2 / 11 times, as predicted by the expression (n-1)/p


## Lagrange Interpolation

In space n-1, q<sub>a</sub> is the polynomial of degree at least n-1 which is equal to a<sub>i+1</sub> for \[a<sub>1</sub>...a<sub>n</sub>].  q can also be called the univariate low degree extension of a.

q is constructed using the lagrange basis polynomials.  For each value i = {0...n-1}, the lagrange basis polynomial of i equals 1 at i and 0 elsewhere ([wiki](https://en.wikipedia.org/wiki/Lagrange_polynomial) ).  Thus q is the sum of all a<sub>i+1</sub>\*LBP(i) over the space n-1.

The vector a which is the evaluation of the polynomial q at each point is an alternative way of describing the polynomial as compared with a vector of the coeffecients of the polynomial. 