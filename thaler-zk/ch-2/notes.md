## Reed-Solomon Fingerprinting

This algorithm is suitable for cooperative parties which trust each other but want to minimize communication to each other - e.g. see if 2 files are equal without sending the entire contents of one file to the other party.  Instead we can just the fingerprint and the hash function used to calculate the fingerprint, chosen at random from a family of hash functions.

*   Pick a prime much larger than the size of your input

*   Convert your input to integers (e.g. ASCII encoding would be integers 1...128)

    *   Input becomes a vector of length n such as \[1, 12, 3, 5, 19]

    *   Call this vector a

    *   Think of a~~n~~ as the coefficients of a polynomial of degree n-1

*   Define the **Field** F~~p~~ as the set of integers modulo p

*   We then call the family of hash functions *H* the set of all functions which evaluate the output of our polynomial for the different values of r E F~~p~~

*   If Alice and Bob want to evaluate if 2 files are the same, Alice selects r at random, evaluates H~~r~~(a) `(evaluates the polynomial with coefficients a at x=r)`, then sends r and H~~r~~(a) to Bob who can use r to compute the same on his file and compare the result.

*   The algorithm has an error term of 1 - ((n-1) / p) | given p > n^2, we can say the error is < (1 - 1/n)

EX/

Alphabet: A = 1; B = 2; C = 3;

Message: CAB

a: \[3, 1, 2] (n=3)

p = 11

r (chosen at random) = 3

Evaluation: 3(3)^2 + 1(3) + 2 = 27 + 3 + 2 = 32 % 11 = 10 = H~~3~~(a)

Different message: BAC

b: \[2, 1,  3]

Evaluation: 2(3)^2 + 1(3) + 3 = 18 + 3 + 3 = 24 % 11 = 2 = H~~3~~(b)

