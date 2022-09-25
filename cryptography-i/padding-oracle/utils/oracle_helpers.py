def guess_helper(ciphertext, index, guess):
    cipher_bytes = bytearray.fromhex(ciphertext)
    index_byte = cipher_bytes[index]
    cipher_bytes[len(cipher_bytes) - 1 - index] = index_byte ^ guess ^ 1
    return cipher_bytes.hex()
